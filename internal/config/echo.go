package config

import (
	"github/golang-developer-technical-test/internal/util"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func NewEcho(config *viper.Viper, log *logrus.Logger, validator *validator.Validate) *echo.Echo {
	e := echo.New()
	allowOrigin := strings.Split(config.GetString("app.allow-origins"), ",")
	e.Use(util.RequestLogger(log))
	e.Use(middleware.Recover())
	e.Use(middleware.Secure())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     allowOrigin, // Allow requests from any origin.
		AllowCredentials: true,        // Allow including credentials in requests.
		AllowMethods: []string{ // Allow specified HTTP methods.
			http.MethodGet,
			http.MethodHead,
			http.MethodPut,
			http.MethodPatch,
			http.MethodPost,
			http.MethodDelete,
		},
	}))
	// Set the custom error handler
	e.HTTPErrorHandler = util.ErrorHandler(log)
	e.Validator = &Validator{
		Validator: validator,
	}
	// Set the application name and other configurations
	e.Debug = config.GetBool("debug")
	return e
}
