package route

import (
	"github/golang-developer-technical-test/internal/delivery/http/controller"

	"github.com/labstack/echo/v4"
)

type RouteConfig struct {
	App            *echo.Echo
	UserController *controller.UserController
}

func (c *RouteConfig) Setup() {
	c.SetupPublic()
	// c.SetupApiKetAuthRoute()
}

func (c *RouteConfig) SetupPublic() {
	c.App.POST("/api/user", c.UserController.Register)
}

// func (c *RouteConfig) SetupApiKetAuthRoute() {

// }
