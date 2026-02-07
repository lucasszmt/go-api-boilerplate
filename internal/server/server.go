package server

import (
	"context"
	"database/sql"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/labstack/echo/v4"

	"go-api-boilerplate/internal/config"
	"go-api-boilerplate/internal/handler"
	"go-api-boilerplate/internal/middleware"
)

type Server struct {
	echo   *echo.Echo
	config *config.Config
	db     *sql.DB
	logger *slog.Logger
}

func New(cfg *config.Config, db *sql.DB, logger *slog.Logger) *Server {
	e := echo.New()

	s := &Server{
		echo:   e,
		config: cfg,
		db:     db,
		logger: logger,
	}

	s.setupMiddleware()
	s.setupRoutes()

	return s
}

func (s *Server) setupMiddleware() {
	s.echo.Use(middleware.RequestID())
	s.echo.Use(middleware.Logger(s.logger))
	s.echo.Use(middleware.Recover(s.logger))
	s.echo.Use(middleware.CORS())
}

func (s *Server) setupRoutes() {
	// Health check
	h := handler.NewHealthHandler(s.db)
	s.echo.GET("/health", h.Health)
	s.echo.GET("/health/ready", h.Ready)

	// API v1 routes
	v1 := s.echo.Group("/api/v1")
	_ = v1 // Add your routes here
}

func (s *Server) Start() error {
	addr := fmt.Sprintf("%s:%d", s.config.Server.Host, s.config.Server.Port)

	// Configure server timeouts
	s.echo.Server.ReadTimeout = s.config.Server.ReadTimeout
	s.echo.Server.WriteTimeout = s.config.Server.WriteTimeout

	// Graceful shutdown
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	// Start server in goroutine
	go func() {
		s.logger.Info("starting server", "address", addr)
		if err := s.echo.Start(addr); err != nil && err != http.ErrServerClosed {
			s.logger.Error("server error", "error", err)
			os.Exit(1)
		}
	}()

	// Wait for interrupt signal
	<-ctx.Done()

	s.logger.Info("shutting down server")

	// Shutdown with timeout
	shutdownCtx, cancel := context.WithTimeout(context.Background(), s.config.Server.ShutdownTimeout)
	defer cancel()

	if err := s.echo.Shutdown(shutdownCtx); err != nil {
		return fmt.Errorf("server shutdown error: %w", err)
	}

	s.logger.Info("server stopped gracefully")
	return nil
}
