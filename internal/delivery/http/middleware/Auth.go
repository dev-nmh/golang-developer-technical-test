package middleware

import (
	"github/golang-developer-technical-test/internal/model"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
)

type Middleware struct {
	config *viper.Viper
}

func NewMiddleware(cfg *viper.Viper) *Middleware {
	return &Middleware{config: cfg}
}
func (m Middleware) AuthApiKey(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		key := m.config.GetString("app.api_key")
		apiKey := c.Request().Header.Get("X-API-KEY")
		if apiKey != key { // Replace "secret" with your actual API key logic.
			var response model.JSONResponse
			response.Data = http.StatusText(http.StatusUnauthorized)
			response.StatusCode = http.StatusUnauthorized
			return c.JSON(http.StatusUnauthorized, response)
		}
		return next(c)
	}
}
