package middleware

import (
	"log/slog"
	"time"

	"github.com/labstack/echo/v4"
)

func Logger(logger *slog.Logger) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			start := time.Now()

			err := next(c)

			req := c.Request()
			res := c.Response()

			attrs := []slog.Attr{
				slog.String("method", req.Method),
				slog.String("path", req.URL.Path),
				slog.Int("status", res.Status),
				slog.Duration("latency", time.Since(start)),
				slog.String("ip", c.RealIP()),
				slog.String("user_agent", req.UserAgent()),
			}

			if requestID, ok := c.Get("request_id").(string); ok {
				attrs = append(attrs, slog.String("request_id", requestID))
			}

			if err != nil {
				attrs = append(attrs, slog.String("error", err.Error()))
				logger.LogAttrs(req.Context(), slog.LevelError, "request failed", attrs...)
			} else {
				logger.LogAttrs(req.Context(), slog.LevelInfo, "request completed", attrs...)
			}

			return err
		}
	}
}
