package main

import (
	"github/golang-developer-technical-test/internal/config"
	"github/golang-developer-technical-test/internal/delivery/http/controller"
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

	EchoContext := config.NewEcho(viperConfig, log, validator)

	userRepository := repository.NewUserRepository(log)
	log.Println(viperConfig)
	CloudinaryUploader := repository.NewCloudinaryUploader(coudinary, viperConfig.GetString("cdn.cloudinary.upload_folder"))

	userUseCase := usecase.NewUserUseCase(DB, log, validator, userRepository, CloudinaryUploader)
	userController := controller.NewUseController(log, userUseCase)

	routeConfig := route.RouteConfig{
		App:            EchoContext,
		UserController: userController,
	}

	routeConfig.Setup()
	EchoContext.Logger.Info(EchoContext.Start(":" + viperConfig.GetString("web.port")))
}
