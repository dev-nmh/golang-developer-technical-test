package repository

import (
	"github/golang-developer-technical-test/internal/entity"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type UserTenorRepository struct {
	Repository[entity.MapUserTenor]
	Log *logrus.Logger
}

func NewUserTenorRepository(log *logrus.Logger) *UserTenorRepository {
	return &UserTenorRepository{
		Log: log,
	}
}
func (utr UserTenorRepository) DeleteUserTenorByUserId(db *gorm.DB, userId uuid.UUID) error {
	return db.Model(entity.MapUserTenor{}).Where("fk_ms_user", userId).Delete(nil).Error
}
func (utr UserTenorRepository) CreateListUserTenor(db *gorm.DB, mapUserTenor []entity.MapUserTenor) error {
	return db.Create(mapUserTenor).Error
}
