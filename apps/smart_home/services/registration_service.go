package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// RegistrationService proxies sensor CRUD requests to the registration API
type RegistrationService struct {
	BaseURL    string
	HTTPClient *http.Client
}

// NewRegistrationService creates a new RegistrationService
func NewRegistrationService(baseURL string) *RegistrationService {
	return &RegistrationService{
		BaseURL: baseURL,
		HTTPClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

// GetSensors fetches all sensors from the registration API
func (s *RegistrationService) GetSensors() (json.RawMessage, int, error) {
	return s.doRequest(http.MethodGet, "/api/v1/sensors", nil)
}

// GetSensorByID fetches a single sensor by ID from the registration API
func (s *RegistrationService) GetSensorByID(id int) (json.RawMessage, int, error) {
	return s.doRequest(http.MethodGet, fmt.Sprintf("/api/v1/sensors/%d", id), nil)
}

// CreateSensor creates a new sensor via the registration API
func (s *RegistrationService) CreateSensor(body []byte) (json.RawMessage, int, error) {
	return s.doRequest(http.MethodPost, "/api/v1/sensors", body)
}

// DeleteSensor deletes a sensor by ID via the registration API
func (s *RegistrationService) DeleteSensor(id int) (json.RawMessage, int, error) {
	return s.doRequest(http.MethodDelete, fmt.Sprintf("/api/v1/sensors/%d", id), nil)
}

// doRequest performs an HTTP request and returns the raw response body, status code, and any error
func (s *RegistrationService) doRequest(method, path string, body []byte) (json.RawMessage, int, error) {
	url := s.BaseURL + path

	var reqBody io.Reader
	if body != nil {
		reqBody = bytes.NewReader(body)
	}

	req, err := http.NewRequest(method, url, reqBody)
	if err != nil {
		return nil, 0, fmt.Errorf("error creating request: %w", err)
	}
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	resp, err := s.HTTPClient.Do(req)
	if err != nil {
		return nil, 0, fmt.Errorf("error sending request to registration API: %w", err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, resp.StatusCode, fmt.Errorf("error reading response body: %w", err)
	}

	return json.RawMessage(data), resp.StatusCode, nil
}

