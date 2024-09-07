package converter

import (
	"github/golang-developer-technical-test/internal/entity"
	"github/golang-developer-technical-test/internal/model"

	"github.com/google/uuid"
)

func AccountToResponse(account entity.MsAccount, userId *uuid.UUID, token string, refreshToken *string) *model.AccountResponse {
	return &model.AccountResponse{
		AccountId:    account.PkMsAccount,
		UserId:       userId,
		RoleId:       account.FkMsRole,
		AccessToken:  token,
		RefreshToken: refreshToken,
	}
}
