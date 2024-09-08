package test

import (
	"github/golang-developer-technical-test/internal/model"

	"github.com/google/uuid"
)

var accessToken string
var account uuid.UUID
var tokenUser string
var tokenAdmin string
var userId uuid.UUID

func CreateAccountUser() model.AccoutRequest {
	return model.AccoutRequest{Email: "naufal@mail.com", Password: "user"}
}
func CreateAccountAdmin() model.AccoutRequest {
	return model.AccoutRequest{Email: "admin@mail.com", Password: "admin"}
}
