package model

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
)

type BaseClaims struct {
	Email  string
	ID     uuid.UUID
	RoleId string
}
type UserBaseClaims struct {
	BaseClaims
	UserID uuid.UUID
}
type JwtBaseClaims struct {
	BaseClaims
	jwt.StandardClaims
}

type JwtUserClaims struct {
	UserBaseClaims
	jwt.StandardClaims
}

type JwtCustomRefreshClaims struct {
	ID string `json:"id"`
	jwt.StandardClaims
}
