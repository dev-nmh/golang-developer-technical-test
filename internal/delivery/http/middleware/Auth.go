package middleware

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
)

func AuthApiKey(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		key := viper.GetString("app.api_key")
		apiKey := c.Request().Header.Get("X-API-KEY")
		if apiKey != key { // Replace "secret" with your actual API key logic.
			return c.JSON(http.StatusUnauthorized, map[string]string{"message": "Invalid API Key"})
		}
		return next(c)
	}
}
