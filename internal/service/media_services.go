package services

import (
	"github/golang-developer-technical-test/internal/model"
	"github/golang-developer-technical-test/internal/util"

	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
)

var (
	validate = validator.New()
)

type mediaUpload interface {
	FileUpload(file model.File) (string, error)
	RemoteUpload(url model.Url) (string, error)
}

type media struct {
	config *viper.Viper
}

func NewMediaUpload() mediaUpload {
	return &media{}
}

func (m *media) FileUpload(file model.File) (string, error) {
	//validate
	err := validate.Struct(file)
	if err != nil {
		return "", err
	}

	//upload
	uploadUrl, err := util.ImageUploadCoudinaryHelper(m.config, file.File)
	if err != nil {
		return "", err
	}
	return uploadUrl, nil
}

func (m *media) RemoteUpload(url model.Url) (string, error) {
	//validate
	err := validate.Struct(url)
	if err != nil {
		return "", err
	}

	//upload
	uploadUrl, errUrl := util.ImageUploadCoudinaryHelper(m.config, url.Url)
	if errUrl != nil {
		return "", err
	}
	return uploadUrl, nil
}
