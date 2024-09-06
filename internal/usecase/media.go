package usecase

type MediaUploader interface {
	Upload(file []byte, fileName string) (string, error)
	Delete(publicID string) error
}

type MediaUseCase struct {
	uploader MediaUploader
}

func NewMediaUseCase(uploader MediaUploader) *MediaUseCase {
	return &MediaUseCase{uploader: uploader}
}

func (uc *MediaUseCase) UploadMedia(file []byte, fileName string) (string, error) {
	return uc.uploader.Upload(file, fileName)
}

func (uc *MediaUseCase) DeleteMedia(publicID string) error {
	return uc.uploader.Delete(publicID)
}
