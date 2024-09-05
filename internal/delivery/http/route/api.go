package route

import "github.com/labstack/echo/v4"

const prefix = "/api/v1"

type RouteConfig struct {
	App *echo.Echo
}

// func Build(e *echo.Echo) {
// 	r := e.Group(prefix)

// 	r.POST("/user")
// }
