package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/CodeChefVIT/go-backend-template/pkg/dto"
	"github.com/CodeChefVIT/go-backend-template/pkg/utils"
	"github.com/labstack/echo/v4"
)

func HealthCheck(c echo.Context) error {
	ctx, cancel := context.WithTimeout(c.Request().Context(), 2*time.Second)
	defer cancel()

	postgresStatus := "OK"
	if utils.DBPool == nil {
		postgresStatus = "DOWN"
	} else if err := utils.DBPool.Ping(ctx); err != nil {
		postgresStatus = "DOWN"
	}

	redisStatus := "OK"
	if utils.RedisClient == nil {
		redisStatus = "DOWN"
	} else if err := utils.RedisClient.Ping(ctx).Err(); err != nil {
		redisStatus = "DOWN"
	}

	status := http.StatusOK
	if postgresStatus == "DOWN" || redisStatus == "DOWN" {
		status = http.StatusServiceUnavailable
	}

	healthInfo := map[string]string{
		"postgres": postgresStatus,
		"redis":    redisStatus,
	}

	return c.JSON(status, dto.NewSuccessResponse("Health check completed", healthInfo))
}
