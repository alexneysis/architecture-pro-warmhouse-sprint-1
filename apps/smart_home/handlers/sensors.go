package handlers

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"

	"smarthome/db"
	"smarthome/models"
	"smarthome/services"

	"github.com/gin-gonic/gin"
)

// SensorHandler handles sensor-related requests
type SensorHandler struct {
	DB                  *db.DB
	TemperatureService  *services.TemperatureService
	RegistrationService *services.RegistrationService
}

// NewSensorHandler creates a new SensorHandler
func NewSensorHandler(db *db.DB, temperatureService *services.TemperatureService, registrationService *services.RegistrationService) *SensorHandler {
	return &SensorHandler{
		DB:                  db,
		TemperatureService:  temperatureService,
		RegistrationService: registrationService,
	}
}

// RegisterRoutes registers the sensor routes
func (h *SensorHandler) RegisterRoutes(router *gin.RouterGroup) {
	sensors := router.Group("/sensors")
	{
		sensors.GET("", h.GetSensors)
		sensors.GET("/:id", h.GetSensorByID)
		sensors.POST("", h.CreateSensor)
		sensors.PUT("/:id", h.UpdateSensor)
		sensors.DELETE("/:id", h.DeleteSensor)
		sensors.PATCH("/:id/value", h.UpdateSensorValue)
		sensors.GET("/temperature", h.GetTemperatureByLocation)
		sensors.GET("/temperature/:id", h.GetTemperatureByID)
	}
}

// GetSensors handles GET /api/v1/sensors — proxied to registration-api
func (h *SensorHandler) GetSensors(c *gin.Context) {
	data, statusCode, err := h.RegistrationService.GetSensors()
	if err != nil {
		log.Printf("Error proxying GetSensors to registration-api: %v", err)
		c.JSON(http.StatusBadGateway, gin.H{"error": fmt.Sprintf("registration API unavailable: %v", err)})
		return
	}
	c.Data(statusCode, "application/json", data)
}

// GetSensorByID handles GET /api/v1/sensors/:id — proxied to registration-api
func (h *SensorHandler) GetSensorByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid sensor ID"})
		return
	}

	data, statusCode, err := h.RegistrationService.GetSensorByID(id)
	if err != nil {
		log.Printf("Error proxying GetSensorByID to registration-api: %v", err)
		c.JSON(http.StatusBadGateway, gin.H{"error": fmt.Sprintf("registration API unavailable: %v", err)})
		return
	}
	c.Data(statusCode, "application/json", data)
}

// GetTemperatureById handles GET /api/v1/sensors/temperature/:id
func (h *SensorHandler) GetTemperatureByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid sensor ID"})
		return
	}

	// Fetch temperature data from the external API
	tempData, err := h.TemperatureService.GetTemperatureByID(strconv.Itoa(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintf("Failed to fetch temperature data: %v", err),
		})
		return
	}

	// Return the temperature data
	c.JSON(http.StatusOK, gin.H{
		"location":    tempData.Location,
		"value":       tempData.Value,
		"unit":        tempData.Unit,
		"status":      tempData.Status,
		"timestamp":   tempData.Timestamp,
		"description": tempData.Description,
	})
}

// GetTemperatureByLocation handles GET /api/v1/sensors/temperature?location=1
func (h *SensorHandler) GetTemperatureByLocation(c *gin.Context) {
	location := c.Query("location")
	if location == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Location is required"})
		return
	}

	// Fetch temperature data from the external API
	tempData, err := h.TemperatureService.GetTemperature(location)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintf("Failed to fetch temperature data: %v", err),
		})
		return
	}

	// Return the temperature data
	c.JSON(http.StatusOK, gin.H{
		"location":    tempData.Location,
		"value":       tempData.Value,
		"unit":        tempData.Unit,
		"status":      tempData.Status,
		"timestamp":   tempData.Timestamp,
		"description": tempData.Description,
	})
}

// CreateSensor handles POST /api/v1/sensors — proxied to registration-api
func (h *SensorHandler) CreateSensor(c *gin.Context) {
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to read request body"})
		return
	}

	data, statusCode, err := h.RegistrationService.CreateSensor(body)
	if err != nil {
		log.Printf("Error proxying CreateSensor to registration-api: %v", err)
		c.JSON(http.StatusBadGateway, gin.H{"error": fmt.Sprintf("registration API unavailable: %v", err)})
		return
	}
	c.Data(statusCode, "application/json", data)
}

// UpdateSensor handles PUT /api/v1/sensors/:id
func (h *SensorHandler) UpdateSensor(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid sensor ID"})
		return
	}

	var sensorUpdate models.SensorUpdate
	if err := c.ShouldBindJSON(&sensorUpdate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	sensor, err := h.DB.UpdateSensor(context.Background(), id, sensorUpdate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, sensor)
}

// DeleteSensor handles DELETE /api/v1/sensors/:id — proxied to registration-api
func (h *SensorHandler) DeleteSensor(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid sensor ID"})
		return
	}

	data, statusCode, err := h.RegistrationService.DeleteSensor(id)
	if err != nil {
		log.Printf("Error proxying DeleteSensor to registration-api: %v", err)
		c.JSON(http.StatusBadGateway, gin.H{"error": fmt.Sprintf("registration API unavailable: %v", err)})
		return
	}
	c.Data(statusCode, "application/json", data)
}

// UpdateSensorValue handles PATCH /api/v1/sensors/:id/value
func (h *SensorHandler) UpdateSensorValue(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid sensor ID"})
		return
	}

	var request struct {
		Value  float64 `json:"value" binding:"required"`
		Status string  `json:"status" binding:"required"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = h.DB.UpdateSensorValue(context.Background(), id, request.Value, request.Status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Sensor value updated successfully"})
}
