package test

import (
	"github/golang-developer-technical-test/internal/config"
	"github/golang-developer-technical-test/internal/delivery/http/controller"
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
var EchoContext *echo.Echo
var UserController *controller.UserController

func init() {
	config.InitCache()
	viperConfig = config.NewViper()
	log = config.NewLogger(viperConfig)
	db = config.NewDatabase(viperConfig, log)
	validator := config.NewValidator(viperConfig)
	coudinary = config.NewCloudinary(viperConfig)

	EchoContext = config.NewEcho(viperConfig, log, validator)

	userRepository := repository.NewUserRepository(log)
	CloudinaryUploader := repository.NewCloudinaryUploader(coudinary, viperConfig.GetString("cdn.cloudinary.upload_folder"))

	userUseCase := usecase.NewUserUseCase(db, log, validator, userRepository, CloudinaryUploader)
	UserController = controller.NewUseController(log, userUseCase)

	routeConfig := route.RouteConfig{
		App:            EchoContext,
		UserController: UserController,
	}

	routeConfig.Setup()
}
