package model

import "github.com/google/uuid"

type AccoutRequest struct {
	Email    string `validate:"required,max=100" "json:"email"`
	Password string `validate:"required,max=14" "json:"password"`
}

type AccountResponse struct {
	AccountId    uuid.UUID  `json:"account_id"`
	UserId       *uuid.UUID `json:"user_id,omitempty"`
	RoleId       int        `json:"role_id"`
	AccessToken  string     `json:"access_token"`
	RefreshToken *string    `json:"refresh_token,omitempty"`
}
