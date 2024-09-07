package model

type MsAccount struct {
	Email    string `validate:"required,max=100"`
	Password string `validate:"required,max=14"`
}
