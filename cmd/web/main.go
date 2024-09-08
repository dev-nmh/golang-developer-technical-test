package main

import (
	"github/golang-developer-technical-test/internal/config"
	"github/golang-developer-technical-test/internal/delivery/http/controller"
	"github/golang-developer-technical-test/internal/delivery/http/middleware"
	"github/golang-developer-technical-test/internal/delivery/http/route"
	"github/golang-developer-technical-test/internal/repository"
	"github/golang-developer-technical-test/internal/usecase"

	"gorm.io/gorm"
)

func main() {
	config.InitCache()
	viperConfig := config.NewViper()
	log := config.NewLogger(viperConfig)
	db := config.NewDatabase(viperConfig, log)
	validator := config.NewValidator(viperConfig)
	coudinary := config.NewCloudinary(viperConfig)
	jwtGenerator := config.NewJwtGenerator(viperConfig)
	App := config.NewEcho(viperConfig, log, validator)

	userRepository := repository.NewUserRepository(log)
	accountRepository := repository.NewAccountRepository(log)
	mapUserTenorRepository := repository.NewUserTenorRepository(log)
	transcationLoanRepository := repository.NewTranscationLoanRepository(log)
	tenorLoanRepository := repository.NewTenorRepository(log)
	sourceRepository := repository.NewSourceRepository(log)

	cloudinaryUploader := repository.NewCloudinaryUploader(coudinary, viperConfig.GetString("cdn.cloudinary.upload_folder"))

	userUseCase := usecase.NewUserUseCase(db, log, validator, userRepository, cloudinaryUploader)
	userController := controller.NewUserController(log, userUseCase)

	accountUseCase := usecase.NewAccountUseCase(db, log, validator, viperConfig, userRepository, accountRepository, jwtGenerator)
	accountController := controller.NewAccountController(log, accountUseCase)

	loanUseCase := usecase.NewLoanUseCase(db, log, validator, userRepository, mapUserTenorRepository)
	loanController := controller.NewLoanController(log, loanUseCase)

	transcationLoanUseCase := usecase.NewTranscationLoanUseCase(db, log, validator, userRepository, mapUserTenorRepository, transcationLoanRepository, tenorLoanRepository, sourceRepository)
	transcationLoanController := controller.NewTranscationLoanController(log, transcationLoanUseCase)

	middleware := middleware.NewMiddleware(viperConfig)

	routeConfig := route.RouteConfig{
		App:                       App,
		UserController:            userController,
		AccountController:         accountController,
		LoanController:            loanController,
		Middleware:                middleware,
		TranscationLoanController: transcationLoanController,
	}

	routeConfig.Setup()
	App.Logger.Info(App.Start(":" + viperConfig.GetString("app.port")))

	connections := []gorm.DB{*db}
	for _, connection := range connections {
		if db, err := connection.DB(); err != nil {
			log.Println("Error close databse", err.Error())
		} else {
			defer db.Close()
		}
	}
}
