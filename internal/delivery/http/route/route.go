package route

import (
	"github/golang-developer-technical-test/internal/constant"
	"github/golang-developer-technical-test/internal/delivery/http/controller"
	"github/golang-developer-technical-test/internal/delivery/http/middleware"

	"github.com/labstack/echo/v4"
)

type RouteConfig struct {
	App               *echo.Echo
	UserController    *controller.UserController
	AccountController *controller.AccountController
	Middleware        *middleware.Middleware
}

func (c *RouteConfig) Setup() {
	c.SetupApiKeyAuthRoute()
	c.SetupAuthJwt()
	c.SetupUserAuth()
	c.SetupAdminAuth()

}

func (c *RouteConfig) SetupApiKeyAuthRoute() {
	private := c.App.Group(constant.PREFIX_API + "/public")
	private.Use(c.Middleware.AuthApiKey)
	private.POST("/auth", c.AccountController.Auth)
	private.POST("/register", c.AccountController.Register)

}
func (c *RouteConfig) SetupAuthJwt() {

}
func (c *RouteConfig) SetupAdminAuth() {
	private := c.App.Group(constant.PREFIX_API + "/admin")
	private.Use(c.Middleware.AuthAdminJWT)
	// private.POST("/user", c.UserController.CreateProfile)
}

func (c *RouteConfig) SetupUserAuth() {
	private := c.App.Group(constant.PREFIX_API)
	private.Use(c.Middleware.AuthUserJWT)
	private.POST("/user", c.UserController.CreateProfile)
}
