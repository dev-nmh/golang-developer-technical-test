package repository

import (
	"context"
	"io"
	"log"
	"mime/multipart"
	"net/url"
	"path"
	"strings"

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

func (cu *CloudinaryUploader) UploadFromBytes(ctx context.Context, file []byte, fileName string) (string, error) {
	uploadParams := uploader.UploadParams{
		Folder:   cu.folderName,
		PublicID: fileName,
	}
	log.Println("=========================Start Uploading Image=========================")
	log.Println(cu.cld.Config)
	result, err := cu.cld.Upload.Upload(ctx, file, uploadParams)
	if err != nil {
		return "", errtrace.Wrap(err)
	}
	log.Println("Upload result:", result)
	log.Println("Secure URL:", result.SecureURL)
	if result.SecureURL == "" {
		log.Println("Warning: Secure URL is empty.")
	}
	log.Println("=========================End Uploading Image=========================")

	return result.SecureURL, nil
}
func (cu *CloudinaryUploader) UploadFromMultipartHeader(file *multipart.FileHeader) (string, error) {
	ctx := context.Background()
	formFile, err := file.Open()

	if err != nil {
		return "", errtrace.Wrap(err)
	}
	defer formFile.Close()

	fileBytes, err := io.ReadAll(formFile)
	if err != nil {
		return "", errtrace.Wrap(err)
	}
	id, err := uuid.NewV7()
	if err != nil {
		return "", errtrace.Wrap(err)
	}
	urlFile, errCloudinary := cu.UploadFromBytes(ctx, fileBytes, id.String())
	if errCloudinary != nil {
		return "", errtrace.Wrap(err)
	}

	return urlFile, nil
}

func (cu *CloudinaryUploader) getPublicIDFromURL(imageURL string) (string, error) {

	parsedURL, err := url.Parse(imageURL)
	if err != nil {
		return "", errtrace.Wrap(err)
	}

	fileName := path.Base(parsedURL.Path)

	// Remove the file extension (e.g., .jpg, .png)
	publicID := strings.TrimSuffix(fileName, path.Ext(fileName))

	return publicID, nil
}
func (cu *CloudinaryUploader) Delete(ctx context.Context, imageURL string) error {

	publicID, err := cu.getPublicIDFromURL(imageURL)
	if err != nil {
		return errtrace.Wrap(err)
	}

	_, err = cu.cld.Upload.Destroy(ctx, uploader.DestroyParams{PublicID: publicID})
	return errtrace.Wrap(err)
}
