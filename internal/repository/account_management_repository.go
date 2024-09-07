package repository

import (
	"github/golang-developer-technical-test/internal/entity"

	"github.com/sirupsen/logrus"
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
