package controllers_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/CodeChefVIT/go-backend-template/internal/controllers"
	"github.com/CodeChefVIT/go-backend-template/internal/dto"
	"github.com/labstack/echo/v4"
)

func TestHealthCheck(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	err := controllers.HealthCheck(c)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Status code should be 503 Service Unavailable because DBPool and RedisClient are uninitialized (nil) in testing
	if rec.Code != http.StatusServiceUnavailable {
		t.Errorf("expected status %d, got %d", http.StatusServiceUnavailable, rec.Code)
	}

	var resp dto.SuccessResponse
	if err := json.Unmarshal(rec.Body.Bytes(), &resp); err != nil {
		t.Fatalf("failed to parse response body: %v", err)
	}

	if resp.Message != "Health check completed" {
		t.Errorf("expected message 'Health check completed', got '%s'", resp.Message)
	}
}
