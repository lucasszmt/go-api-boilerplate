package handler

import (
	"context"
	"database/sql"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

type HealthHandler struct {
	db *sql.DB
}

func NewHealthHandler(db *sql.DB) *HealthHandler {
	return &HealthHandler{db: db}
}

type HealthResponse struct {
	Status string `json:"status"`
}

type ReadyResponse struct {
	Status   string            `json:"status"`
	Services map[string]string `json:"services"`
}

// Health returns a simple health check response
func (h *HealthHandler) Health(c echo.Context) error {
	return c.JSON(http.StatusOK, HealthResponse{
		Status: "ok",
	})
}

// Ready checks if all dependencies are available
func (h *HealthHandler) Ready(c echo.Context) error {
	services := make(map[string]string)
	status := http.StatusOK

	// Check database connection
	ctx, cancel := context.WithTimeout(c.Request().Context(), 2*time.Second)
	defer cancel()

	if err := h.db.PingContext(ctx); err != nil {
		services["database"] = "unhealthy"
		status = http.StatusServiceUnavailable
	} else {
		services["database"] = "healthy"
	}

	response := ReadyResponse{
		Status:   "ok",
		Services: services,
	}

	if status != http.StatusOK {
		response.Status = "degraded"
	}

	return c.JSON(status, response)
}
