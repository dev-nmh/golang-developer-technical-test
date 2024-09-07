package converter

import (
	"github/golang-developer-technical-test/internal/entity"
	"github/golang-developer-technical-test/internal/model"
)

func UserToResponse(user *entity.MsUser) *model.UserResponseDetail {
	return &model.UserResponseDetail{
		PkMsUser: user.PkMsUser,

		UserData: model.UserData{
			Nik:        user.Nik,
			FullName:   user.FullName,
			LegalName:  user.LegalName,
			BirthPlace: user.BirthPlace,
			BirthDate:  user.BirthDate,
			Salary:     user.Salary,
		},
		UserAssetRecord: model.UserAssetRecord{
			ImageKtp:    user.ImageKtp,
			ImageSelfie: user.ImageSelfie,
		},
	}
}
