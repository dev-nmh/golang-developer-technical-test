package route

import (
	"github/golang-developer-technical-test/internal/config"
	"github/golang-developer-technical-test/internal/delivery/http/controller"
	"github/golang-developer-technical-test/internal/delivery/http/middleware"

	"github.com/labstack/echo/v4"
)

type RouteConfig struct {
	App            *echo.Echo
	UserController *controller.UserController
	Middleware     *middleware.Middleware
}

func (c *RouteConfig) Setup() {

	c.SetupPublic()
	c.SetupApiKeyAuthRoute()
	c.SetupJWTAuth()
}

func (c *RouteConfig) SetupPublic() {

	// c.App.POST("/api/user", c.UserController.Register)
}

func (c *RouteConfig) SetupApiKeyAuthRoute() {
	private := c.App.Group(config.PREFIX_API)
	private.Use(c.Middleware.AuthApiKey)
	private.POST("/user", c.UserController.Register)
}

func (c *RouteConfig) SetupJWTAuth() {
	private := c.App.Group(config.PREFIX_API + "/admin")
	private.Use(c.Middleware.AuthApiKey)
	private.POST("/user", c.UserController.Register)
}
