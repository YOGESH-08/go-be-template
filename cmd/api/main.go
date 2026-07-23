package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/CodeChefVIT/go-backend-template/pkg/logging"
	"github.com/CodeChefVIT/go-backend-template/pkg/middlewares"
	"github.com/CodeChefVIT/go-backend-template/pkg/router"
	"github.com/CodeChefVIT/go-backend-template/pkg/utils"
	"github.com/labstack/echo/v4"
	emiddleware "github.com/labstack/echo/v4/middleware"
)

func main() {
	// Initialize logger
	logging.InitLogger()

	// Load configuration
	utils.LoadConfig()

	// Initialize DB pool
	utils.InitDB()
	defer utils.CloseDB()

	// Initialize Redis
	utils.InitRedis()
	defer utils.CloseRedis()

	// Initialize Echo instance
	e := echo.New()

	// Register request validator
	e.Validator = utils.NewValidator()

	// Setup middlewares
	e.Use(emiddleware.Recover())
	e.Use(middlewares.Logger)
	e.Use(emiddleware.Secure())
	e.Use(emiddleware.RateLimiter(emiddleware.NewRateLimiterMemoryStore(20)))

	// Configure CORS in non-production environments
	if len(utils.Config.CORSOrigins) > 0 && utils.Config.Env != "production" {
		e.Use(emiddleware.CORSWithConfig(emiddleware.CORSConfig{
			AllowOrigins:     utils.Config.CORSOrigins,
			AllowMethods:     []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
			AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
			AllowCredentials: true,
		}))
		logging.Infof("CORS enabled for origins: %v", utils.Config.CORSOrigins)
	}

	// Register routes
	router.RegisterRoutes(e)

	// Start server in goroutine to allow graceful shutdown
	serverErrCh := make(chan error, 1)
	go func() {
		logging.Infof("Starting HTTP server on port %s", utils.Config.Port)
		if err := e.Start(":" + utils.Config.Port); err != nil && err != http.ErrServerClosed {
			serverErrCh <- err
		}
	}()

	// Graceful shutdown setup
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)

	select {
	case sig := <-sigCh:
		logging.Infof("Received shutdown signal: %s", sig)
	case err := <-serverErrCh:
		logging.Fatalf("HTTP server error: %v", err)
	}

	// Shutdown Echo context with 10s timeout
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := e.Shutdown(shutdownCtx); err != nil {
		logging.Errorf("Failed to gracefully shutdown HTTP server: %v", err)
	} else {
		logging.Infof("HTTP server shut down successfully")
	}

	logging.Infof("Shutdown process completed")
}
