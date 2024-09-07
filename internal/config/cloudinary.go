package config

import (
	"log"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/spf13/viper"
)

func NewCloudinary(config *viper.Viper) *cloudinary.Cloudinary {
	cld, err := cloudinary.NewFromParams(
		config.GetString("cdn.cloudinary.cloud_name"),
		config.GetString("cdn.cloudinary.api_key"),
		config.GetString("cdn.cloudinary.api_secret"),
	)
	if err != nil {
		log.Fatalf("failed to connect cloudinary: %v", err)
	}
	return cld
}
