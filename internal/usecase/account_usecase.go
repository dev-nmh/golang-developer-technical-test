package usecase

import (
	"github/golang-developer-technical-test/internal/repository"

	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type AccountUseCase struct {
	DB                *gorm.DB
	Log               *logrus.Logger
	Validate          *validator.Validate
	config            *viper.Viper
	accountRepository *repository.AccountRepository
}

func NewAccountUseCase(db *gorm.DB, logger *logrus.Logger, validate *validator.Validate, config *viper.Viper, accountRepository *repository.AccountRepository) *AccountUseCase {
	return &AccountUseCase{
		DB:                db,
		Log:               logger,
		Validate:          validate,
		config:            config,
		accountRepository: accountRepository,
	}
}

func (acc *AccountUseCase) RegisterUserRequest() {

}

func (acc *AccountUseCase) Verify() {

}
