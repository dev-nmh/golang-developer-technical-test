package entity

import (
	"time"

	"github.com/google/uuid"
)

type MsUser struct {
	PkMsUser           uuid.UUID `gorm:"column:pk_ms_user;primaryKey"`
	FkMsAccount        uuid.UUID `gorm:"column:fk_ms_account;"`
	FkMsApprovalStatus int       `gorm:"column:fk_ms_approval_status;"`
	Nik                string    `gorm:"column:NIK"`
	FullName           string    `gorm:"column:full_name"`
	LegalName          string    `gorm:"column:legal_name"`
	BirthPlace         string    `gorm:"column:birth_place"`
	BirthDate          time.Time `gorm:"column:birth_date"`
	Salary             int       `gorm:"column:salary"`
	ImageKtp           string    `gorm:"column:image_ktp"`
	ImageSelfie        string    `gorm:"column:image_selfie"`
	Stamp
}

func (e *MsUser) TableName() string {
	return "ms_user"
}
