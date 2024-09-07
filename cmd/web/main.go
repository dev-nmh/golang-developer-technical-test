package main

import (
	"github/golang-developer-technical-test/internal/config"
	"github/golang-developer-technical-test/internal/delivery/http/controller"
	"github/golang-developer-technical-test/internal/delivery/http/middleware"
	"github/golang-developer-technical-test/internal/delivery/http/route"
	"github/golang-developer-technical-test/internal/repository"
	"github/golang-developer-technical-test/internal/usecase"
)

func main() {
	config.InitCache()
	viperConfig := config.NewViper()
	log := config.NewLogger(viperConfig)
	DB := config.NewDatabase(viperConfig, log)
	validator := config.NewValidator(viperConfig)
	coudinary := config.NewCloudinary(viperConfig)

	App := config.NewEcho(viperConfig, log, validator)

	userRepository := repository.NewUserRepository(log)

	CloudinaryUploader := repository.NewCloudinaryUploader(coudinary, viperConfig.GetString("cdn.cloudinary.upload_folder"))
	userUseCase := usecase.NewUserUseCase(DB, log, validator, userRepository, CloudinaryUploader)
	userController := controller.NewUserController(log, userUseCase)
	middleware := middleware.NewMiddleware(viperConfig)

	routeConfig := route.RouteConfig{
		App:            App,
		UserController: userController,
		Middleware:     middleware,
	}

	routeConfig.Setup()
	App.Logger.Info(App.Start(":" + viperConfig.GetString("web.port")))
}
