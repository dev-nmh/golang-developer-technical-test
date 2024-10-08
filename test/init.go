package test

import (
	"github/golang-developer-technical-test/internal/config"
	"github/golang-developer-technical-test/internal/delivery/http/controller"
	"github/golang-developer-technical-test/internal/delivery/http/middleware"
	"github/golang-developer-technical-test/internal/delivery/http/route"
	"github/golang-developer-technical-test/internal/repository"
	"github/golang-developer-technical-test/internal/usecase"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

var db *gorm.DB
var viperConfig *viper.Viper
var log *logrus.Logger
var coudinary *cloudinary.Cloudinary
var App *echo.Echo
var userController *controller.UserController
var accountController *controller.AccountController
var accountRepository *repository.AccountRepository
var accountUseCase *usecase.AccountUseCase

func init() {
	config.InitCache()
	viperConfig = config.NewViper()
	log = config.NewLogger(viperConfig)
	db = config.NewDatabase(viperConfig, log)
	validator := config.NewValidator(viperConfig)
	coudinary = config.NewCloudinary(viperConfig)
	jwtGenerator := config.NewJwtGenerator(viperConfig)
	App = config.NewEcho(viperConfig, log, validator)
	log.Println(jwtGenerator)
	CloudinaryUploader := repository.NewCloudinaryUploader(coudinary, viperConfig.GetString("cdn.cloudinary.upload_folder"))
	userRepository := repository.NewUserRepository(log)
	accountRepository = repository.NewAccountRepository(log)
	mapUserTenorRepository := repository.NewUserTenorRepository(log)
	transcationLoanRepository := repository.NewTranscationLoanRepository(log)
	tenorLoanRepository := repository.NewTenorRepository(log)
	sourceRepository := repository.NewSourceRepository(log)

	userUseCase := usecase.NewUserUseCase(db, log, validator, userRepository, CloudinaryUploader)
	accountUseCase = usecase.NewAccountUseCase(db, log, validator, viperConfig, userRepository, accountRepository, jwtGenerator)

	userController = controller.NewUserController(log, userUseCase)
	accountController = controller.NewAccountController(log, accountUseCase)

	loanUseCase := usecase.NewLoanUseCase(db, log, validator, userRepository, mapUserTenorRepository)
	loanController := controller.NewLoanController(log, loanUseCase)
	middleware := middleware.NewMiddleware(viperConfig)

	transcationLoanUseCase := usecase.NewTranscationLoanUseCase(db, log, validator, userRepository, mapUserTenorRepository, transcationLoanRepository, tenorLoanRepository, sourceRepository)
	transcationLoanController := controller.NewTranscationLoanController(log, transcationLoanUseCase)

	routeConfig := route.RouteConfig{
		App:                       App,
		UserController:            userController,
		AccountController:         accountController,
		LoanController:            loanController,
		Middleware:                middleware,
		TranscationLoanController: transcationLoanController,
	}

	routeConfig.Setup()
}
