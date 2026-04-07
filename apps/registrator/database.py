
import os
from collections.abc import Sequence
from contextlib import asynccontextmanager
from typing import Any

import asyncpg
from fastapi import FastAPI

from models import Sensor, SensorCreate


class Database:
    def __init__(self, dsn: str) -> None:
        self._dsn = dsn
        self.pool: asyncpg.Pool | None = None

    async def connect(self) -> None:
        self.pool = await asyncpg.create_pool(dsn=self._dsn, min_size=1, max_size=10)

    async def disconnect(self) -> None:
        if self.pool is not None:
            await self.pool.close()
            self.pool = None

    async def init_schema(self) -> None:
        assert self.pool is not None, "Database pool is not initialized"
        query = """
        CREATE TABLE IF NOT EXISTS sensors (
            id BIGSERIAL PRIMARY KEY,
            name VARCHAR(255) NOT NULL,
            type VARCHAR(64) NOT NULL,
            location VARCHAR(255),
            value DOUBLE PRECISION NOT NULL DEFAULT 0,
            unit VARCHAR(32),
            status VARCHAR(64) NOT NULL,
            last_updated TIMESTAMPTZ NOT NULL DEFAULT NOW(),
            created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
        );
        CREATE INDEX IF NOT EXISTS idx_sensors_type ON sensors(type);
        CREATE INDEX IF NOT EXISTS idx_sensors_location ON sensors(location);
        """
        async with self.pool.acquire() as conn:
            await conn.execute(query)

    async def get_sensors(self) -> list[Sensor]:
        assert self.pool is not None, "Database pool is not initialized"
        query = """
        SELECT id, name, type, location, value, unit, status, last_updated, created_at
        FROM sensors
        ORDER BY id;
        """
        async with self.pool.acquire() as conn:
            rows = await conn.fetch(query)
        return [self._row_to_sensor(row) for row in rows]

    async def get_sensor_by_id(self, sensor_id: int) -> Sensor | None:
        assert self.pool is not None, "Database pool is not initialized"
        query = """
        SELECT id, name, type, location, value, unit, status, last_updated, created_at
        FROM sensors
        WHERE id = $1;
        """
        async with self.pool.acquire() as conn:
            row = await conn.fetchrow(query, sensor_id)
        return self._row_to_sensor(row) if row else None

    async def create_sensor(self, payload: SensorCreate) -> Sensor:
        assert self.pool is not None, "Database pool is not initialized"
        query = """
        INSERT INTO sensors (name, type, location, value, unit, status)
        VALUES ($1, $2, $3, $4, $5, $6)
        RETURNING id, name, type, location, value, unit, status, last_updated, created_at;
        """
        async with self.pool.acquire() as conn:
            row = await conn.fetchrow(
                query,
                payload.name,
                payload.type.value,
                payload.location,
                payload.value,
                payload.unit,
                payload.status.value,
            )
        assert row is not None
        return self._row_to_sensor(row)

    async def delete_sensor(self, sensor_id: int) -> bool:
        assert self.pool is not None, "Database pool is not initialized"
        query = "DELETE FROM sensors WHERE id = $1;"
        async with self.pool.acquire() as conn:
            result = await conn.execute(query, sensor_id)
        return result.endswith("1")

    @staticmethod
    def _row_to_sensor(row: asyncpg.Record | dict[str, Any]) -> Sensor:
        return Sensor.model_validate(dict(row))


def get_database_url() -> str:
    return os.getenv(
        "DATABASE_URL",
        "postgresql://postgres:postgres@localhost:5432/smarthome",
    )


def create_database() -> Database:
    return Database(get_database_url())


@asynccontextmanager
async def lifespan(app: FastAPI):
    db = create_database()
    await db.connect()
    await db.init_schema()
    app.state.db = db
    try:
        yield
    finally:
        await db.disconnect()