package config

import (
	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
)

type Validator struct {
	Validator *validator.Validate
}

func NewValidator(viper *viper.Viper) *validator.Validate {
	return validator.New()
}
func (cv *Validator) Validate(i interface{}) error {
	return cv.Validator.Struct(i)
}
