package util

import (
	"context"
	"time"

	"github.com/cloudinary/cloudinary-go"
	"github.com/cloudinary/cloudinary-go/api/uploader"
	"github.com/spf13/viper"
)

func ImageUploadCoudinaryHelper(config *viper.Viper, input interface{}) (string, error) {
	cloudName := config.GetString("cdn.cloudinary.cloud_name")
	apiKey := config.GetString("cdn.cloudinary.api_key")
	apiSecret := config.GetString("cdn.cloudinary.api_secret")
	uploadFolder := config.GetString("cdn.cloudinary.upload_folder")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	//create cloudinary instance
	cld, err := cloudinary.NewFromParams(cloudName, apiKey, apiSecret)
	if err != nil {
		return "", err
	}

	//upload file
	uploadParam, err := cld.Upload.Upload(ctx, input, uploader.UploadParams{Folder: uploadFolder})
	if err != nil {
		return "", err
	}
	return uploadParam.SecureURL, nil
}
