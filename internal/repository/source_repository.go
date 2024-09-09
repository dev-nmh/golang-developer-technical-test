package repository

import (
	"github/golang-developer-technical-test/internal/entity"

	"github.com/sirupsen/logrus"
)

type SourceRepository struct {
	Repository[entity.MsSource]
	Log *logrus.Logger
}

func NewSourceRepository(log *logrus.Logger) *SourceRepository {
	return &SourceRepository{
		Log: log,
	}
}
