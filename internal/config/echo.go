package config

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func NewEcho(config *viper.Viper, log *logrus.Logger) *echo.Echo {
	e := echo.New()

	e.Use(RequestLogger(log))
	e.Use(middleware.Recover())

	// Set the custom error handler
	e.HTTPErrorHandler = ErrorHandler(log)

	// Set the application name and other configurations
	e.Debug = config.GetBool("debug")

	return e
}

func RequestLogger(log *logrus.Logger) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			start := time.Now()

			log.WithFields(logrus.Fields{
				"method": c.Request().Method,
				"uri":    c.Request().RequestURI,
				"ip":     c.RealIP(),
			}).Info("Received request")

			err := next(c)

			log.WithFields(logrus.Fields{
				"method":  c.Request().Method,
				"uri":     c.Request().RequestURI,
				"status":  c.Response().Status,
				"latency": time.Since(start).String(),
			}).Info("Completed request")

			return err
		}
	}
}

func ErrorHandler(log *logrus.Logger) echo.HTTPErrorHandler {
	return func(err error, c echo.Context) {
		if he, ok := err.(*echo.HTTPError); ok {
			log.WithFields(logrus.Fields{
				"error":  he.Message,
				"status": he.Code,
				"path":   c.Request().URL.Path,
			}).Error("HTTP error occurred")
			c.JSON(he.Code, he)
		} else {
			log.WithFields(logrus.Fields{
				"error": err.Error(),
				"path":  c.Request().URL.Path,
			}).Error("Unexpected error occurred")
			c.JSON(http.StatusInternalServerError, map[string]string{"message": "Internal Server Error"})
		}
	}
}
