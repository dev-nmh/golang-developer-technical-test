package main

import (
	"github/golang-developer-technical-test/internal/config"
	"github/golang-developer-technical-test/internal/delivery/http/controller"
	"github/golang-developer-technical-test/internal/delivery/http/route"
)

func main() {
	config.InitCache()
	viperConfig := config.NewViper()
	log := config.NewLogger(viperConfig)
	// db := config.NewDatabase(viperConfig, log)
	validator := config.NewValidator(viperConfig)

	EchoContext := config.NewEcho(viperConfig, log, validator)
	userController := controller.NewUseController(log)

	routeConfig := route.RouteConfig{
		App:            EchoContext,
		UserController: userController,
	}

	routeConfig.Setup()
	EchoContext.Logger.Info(EchoContext.Start(":" + viperConfig.GetString("web.port")))
}
