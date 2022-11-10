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
					logger.LogInfo("echo", "%d | %s %s <%v>", v.Status, v.Method, v.URI, v.Latency)
					return nil
				}
				logger.LogError("echo", "%d | %s %s <%s> (%s)", v.Status, v.Method, v.URI, v.Latency, v.Error)
				return v.Error
			},
			LogError:   true,
			LogLatency: true,
			LogMethod:  true,
			LogStatus:  true,
			LogURI:     true,
		},
	)
}
