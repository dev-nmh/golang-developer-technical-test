package repository

import (
	"github/golang-developer-technical-test/internal/entity"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
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

func (ur UserRepository) GetApprovalStatusById(db *gorm.DB, pkMsUser uuid.UUID) (int, error) {
	var approvalStatus int
	if err := db.Model(entity.MsUser{
		PkMsUser: pkMsUser,
	}).Scan(approvalStatus).Error; err != nil {
		return 0, err
	}
	return approvalStatus, nil

}
func (ur UserRepository) UpdateApprovalStatusById(db *gorm.DB, approvalStaus int, pkMsUser uuid.UUID) error {
	return db.Model(entity.MsUser{
		PkMsUser: pkMsUser,
	}).Update("fk_ms_approval_status", approvalStaus).Error
}
