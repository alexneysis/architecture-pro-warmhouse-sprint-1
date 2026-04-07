from __future__ import annotations

from datetime import datetime
from enum import StrEnum

from pydantic import BaseModel, Field, ConfigDict


class SensorType(StrEnum):
    temperature = "temperature"
    humidity = "humidity"
    motion = "motion"
    light = "light"
    smoke = "smoke"
    generic = "generic"


class SensorStatus(StrEnum):
    active = "active"
    inactive = "inactive"
    warning = "warning"
    error = "error"


class SensorCreate(BaseModel):
    name: str = Field(min_length=1, max_length=255)
    type: SensorType
    location: str | None = Field(default=None, max_length=255)
    value: float = 0
    unit: str | None = Field(default=None, max_length=32)
    status: SensorStatus = SensorStatus.active


class Sensor(BaseModel):
    model_config = ConfigDict(from_attributes=True)

    id: int
    name: str
    type: SensorType
    location: str | None
    value: float
    unit: str | None
    status: SensorStatus
    last_updated: datetime
    created_at: datetime


class MessageResponse(BaseModel):
    message: str