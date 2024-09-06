package config

import (
	"mime/multipart"
	"strconv"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
)

type Validator struct {
	Validator *validator.Validate
}

func NewValidator(viper *viper.Viper) *validator.Validate {
	var validate = validator.New()
	validate.RegisterValidation("filetypeimg", FileTypeImage)
	validate.RegisterValidation("maxsize", FileSize)

	return validate
}
func (cv *Validator) Validate(i interface{}) error {
	return cv.Validator.Struct(i)
}

func FileTypeImage(fl validator.FieldLevel) bool {
	file, ok := fl.Field().Interface().(*multipart.FileHeader)
	if !ok {
		return false
	}

	// Example file type check, adjust as needed
	allowedTypes := map[string]bool{
		"jpeg": true,
		"jpg":  true,
		"png":  true,
	}

	// Get file extension
	fileExt := strings.ToLower(file.Filename[strings.LastIndex(file.Filename, ".")+1:])
	return allowedTypes[fileExt]
}

func FileSize(fl validator.FieldLevel) bool {
	file, ok := fl.Field().Interface().(*multipart.FileHeader)
	if !ok {
		return false
	}

	// Convert the size limit (in bytes) from the tag
	maxSize := fl.Param()
	maxSizeInt, err := strconv.ParseInt(maxSize, 10, 64)
	if err != nil {
		return false
	}

	// Check the file size
	if file.Size > maxSizeInt {
		return false
	}
	return true
}
