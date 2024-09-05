package repository

import (
	"github/golang-developer-technical-test/internal/entity"

	"github.com/sirupsen/logrus"
)

type UserRepository struct {
	Repository[entity.MsUser]
	Log *logrus.Logger
}

func NewUserRepository(log *logrus.Logger) *UserRepository {
	return &UserRepository{
		Log: log,
	}
}
