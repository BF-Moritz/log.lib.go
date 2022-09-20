package middleware

import (
	logger "github.com/BF-Moritz/log.lib.go"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func MakeEchoMiddleware(logger *logger.Logger) echo.MiddlewareFunc {
	return middleware.RequestLoggerWithConfig(
		middleware.RequestLoggerConfig{
			LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
				if v.Error == nil {
					logger.LogInfo("echo", "%s | %s in %v", v.Method, v.URIPath, v.Latency)
					return nil
				}
				logger.LogError("echo", "%s | %s in %s (%s)", v.Method, v.URI, v.Latency, v.Error)
				return v.Error
			},
			LogError:   true,
			LogLatency: true,
			LogMethod:  true,
			LogURIPath: true,
		},
	)
}
