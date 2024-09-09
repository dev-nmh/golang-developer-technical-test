package repository

import (
	"github/golang-developer-technical-test/internal/entity"

	"github.com/sirupsen/logrus"
)

type TenorRepository struct {
	Repository[entity.MsTenor]
	Log *logrus.Logger
}

func NewTenorRepository(log *logrus.Logger) *TenorRepository {
	return &TenorRepository{
		Log: log,
	}
}
