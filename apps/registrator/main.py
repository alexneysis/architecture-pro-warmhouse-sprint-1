
from fastapi import FastAPI, HTTPException, Path, Request, status
from database import lifespan
from models import MessageResponse, SensorCreate, Sensor
import random

app = FastAPI(
    title="Sensor Service",
    version="1.0.0",
    description="Async FastAPI microservice for sensors",
    lifespan=lifespan,
)

def get_random_temperature() -> float:
    return random.random() * 100


@app.get("/health")
async def healthcheck() -> dict[str, str]:
    return {"status": "ok"}


@app.get("/api/v1/sensors", response_model=list[Sensor], summary="Get all sensors")
async def get_sensors(request: Request) -> list[Sensor]:
    sensors = await request.app.state.db.get_sensors()
    for sensor in sensors:
        sensor.value = get_random_temperature()

    return sensors


@app.get(
    "/api/v1/sensors/{sensor_id}",
    response_model=Sensor,
    summary="Get sensor by id",
)
async def get_sensor_by_id(
    request: Request,
    sensor_id: int = Path(..., ge=1, description="Sensor ID"),
) -> Sensor:
    sensor = await request.app.state.db.get_sensor_by_id(sensor_id)
    if sensor is None:
        raise HTTPException(
            status_code=status.HTTP_404_NOT_FOUND,
            detail="Sensor not found",
        )

    sensor.value = get_random_temperature()

    return sensor


@app.post(
    "/api/v1/sensors",
    response_model=Sensor,
    status_code=status.HTTP_201_CREATED,
    summary="Create sensor",
)
async def create_sensor(request: Request, payload: SensorCreate) -> Sensor:
    return await request.app.state.db.create_sensor(payload)


@app.delete(
    "/api/v1/sensors/{sensor_id}",
    response_model=MessageResponse,
    summary="Delete sensor",
)
async def delete_sensor(
    request: Request,
    sensor_id: int = Path(..., ge=1, description="Sensor ID"),
) -> MessageResponse:
    deleted = await request.app.state.db.delete_sensor(sensor_id)
    if not deleted:
        raise HTTPException(
            status_code=status.HTTP_404_NOT_FOUND,
            detail="Sensor not found",
        )
    return MessageResponse(message="Sensor deleted successfully")