package config

import (
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type IApp interface {
	Setup(AppConfig)
}
type AppConfig struct {
	DB         *gorm.DB
	EchoServer *echo.Echo
	Log        *logrus.Logger
	Config     *viper.Viper
}

func NewAPP(DB *gorm.DB, EchoServer *echo.Echo, Log *logrus.Logger, Config *viper.Viper) AppConfig {
	return AppConfig{
		DB:         DB,
		EchoServer: EchoServer,
		Log:        Log,
		Config:     Config,
	}
}

func (AppConfig) Setup() {

}
