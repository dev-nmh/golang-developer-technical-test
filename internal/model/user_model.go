package model

import (
	"mime/multipart"
	"time"

	"github.com/google/uuid"
)

type RegisterUserRequest struct {
	Nik         string                `form:"NIK" json:"NIK"`
	FullName    string                `form:"full_name" json:"full_name"`
	LegalName   string                `form:"legal_name" json:"legal_name"`
	BirthPlace  string                `form:"birth_place" json:"birth_place"`
	BirthDate   time.Time             `form:"birth_date" json:"birth_date"`
	Salary      string                `form:"salary" json:"salary"`
	ImageKtp    *multipart.FileHeader `form:"image_ktp" json:"image_ktp"`
	ImageSelfie *multipart.FileHeader `form:"image_selfie" json:"image_selfie"`
}

type UserResponse struct {
	PkMsUser uuid.UUID `json:"user_id"`
	RegisterUserRequest
}
