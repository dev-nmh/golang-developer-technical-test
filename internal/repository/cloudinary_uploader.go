package repository

import (
	"context"
	"log"
	"mime/multipart"

	"braces.dev/errtrace"
	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/google/uuid"
)

type CloudinaryUploader struct {
	folderName string
	cld        *cloudinary.Cloudinary
}

func NewCloudinaryUploader(coudinary *cloudinary.Cloudinary, folderName string) *CloudinaryUploader {
	return &CloudinaryUploader{
		folderName: folderName,
		cld:        coudinary,
	}
}

func (cu *CloudinaryUploader) Upload(ctx context.Context, file interface{}, fileName string) (*uploader.UploadResult, error) {
	uploadParams := uploader.UploadParams{
		Folder:   cu.folderName,
		PublicID: fileName,
	}
	log.Println("=========================Start Uploading Image=========================")
	result, err := cu.cld.Upload.Upload(ctx, file, uploadParams)
	if err != nil {
		return nil, errtrace.Wrap(err)
	}
	log.Println("Upload result:", result)
	log.Println("Secure URL:", result.SecureURL)
	log.Println("Display Name:", result.DisplayName)
	log.Println("Format:", result.Format)
	log.Println("Public ID:", result.PublicID)
	log.Println("Type:", result.Type)

	if result.SecureURL == "" {
		log.Println("Warning: Secure URL is empty.")
	}
	log.Println("=========================End Uploading Image=========================")

	return result, nil
}
func (cu *CloudinaryUploader) UploadFromMultipartHeader(file *multipart.FileHeader) (string, string, error) {
	ctx := context.Background()
	id, err := uuid.NewV7()
	if err != nil {
		return "", "", errtrace.Wrap(err)
	}
	result, errCloudinary := cu.Upload(ctx, file, id.String())
	if errCloudinary != nil {
		return "", "", errtrace.Wrap(err)
	}
	urlHttp := result.SecureURL
	fileName := result.DisplayName + "." + result.Format

	return urlHttp, fileName, nil
}

func (cu *CloudinaryUploader) Delete(ctx context.Context, public_id string) error {

	_, err := cu.cld.Upload.Destroy(ctx, uploader.DestroyParams{PublicID: public_id})

	return errtrace.Wrap(err)
}
