package repository

import (
	"github/golang-developer-technical-test/internal/entity"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type AccountRepository struct {
	Repository[entity.MsAccount]
	Log *logrus.Logger
}

func NewAccountRepository(log *logrus.Logger) *AccountRepository {
	return &AccountRepository{
		Log: log,
	}
}

func (r *UserRepository) FindByUserName(db *gorm.DB, account *entity.MsAccount, email string) error {
	return db.Where("Email = ?", email).First(account).Error
}
