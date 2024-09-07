package entity

import "github.com/google/uuid"

type MsAccount struct {
	PkMsAccount  uuid.UUID `gorm:"column:pk_ms_account;primaryKey"`
	FkMsRole     int       `gorm:"column:fk_ms_role; required"`
	Email        string    `gorm:"column:email; required"`
	Password     string    `gorm:"column:password; required"`
	PasswordSalt string    `gorm:"column:password_salt; required"`
	Stamp
}

func (e *MsAccount) TableName() string {
	return "ms_account"
}
