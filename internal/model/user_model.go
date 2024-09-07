package model

import (
	"mime/multipart"
	"time"

	"github.com/google/uuid"
)

type RegisterUserRequest struct {
	UserData
	UserAssetRegister
}
type UserData struct {
	FkMsAccount        uuid.UUID `form:"-" json:"account_id"`
	FkMsApprovalStatus int       `form:"-" json:"approval_status"`
	Nik                string    `form:"NIK" json:"NIK" validate:"required,numeric,len=16"`
	FullName           string    `form:"full_name" json:"full_name" validate:"required,max=100"`
	LegalName          string    `form:"legal_name" json:"legal_name" validate:"required,max=60"`
	BirthPlace         string    `form:"birth_place" json:"birth_place" validate:"required,max=50"`
	BirthDate          time.Time `form:"birth_date" json:"birth_date" validate:"required"`
	Salary             int       `form:"salary" json:"salary" validate:"required"`
}
type UserAssetRegister struct {
	ImageKtp    *multipart.FileHeader `form:"image_ktp" json:"image_ktp" validate:"required"`
	ImageSelfie *multipart.FileHeader `form:"image_selfie" json:"image_selfie" validate:"required"`
}

type UserAssetRecord struct {
	ImageKtp    string `json:"image_ktp"`
	ImageSelfie string `json:"image_selfie"`
}

type UserResponseDetail struct {
	PkMsUser uuid.UUID `json:"user_id"`
	UserData
	UserAssetRecord
}
