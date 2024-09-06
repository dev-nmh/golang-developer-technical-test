package usecase

import (
	"context"
	"github/golang-developer-technical-test/internal/entity"
	"github/golang-developer-technical-test/internal/model"
	"github/golang-developer-technical-test/internal/model/converter"
	"github/golang-developer-technical-test/internal/repository"
	"github/golang-developer-technical-test/internal/util"
	"sync"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type UserUseCase struct {
	DB                 *gorm.DB
	Log                *logrus.Logger
	Validate           *validator.Validate
	UserRepository     *repository.UserRepository
	CloudinaryUploader *repository.CloudinaryUploader
}

func NewUserUseCase(db *gorm.DB, logger *logrus.Logger, validate *validator.Validate, userRepository *repository.UserRepository, cloudinary *repository.CloudinaryUploader) *UserUseCase {
	return &UserUseCase{
		DB:                 db,
		Log:                logger,
		Validate:           validate,
		UserRepository:     userRepository,
		CloudinaryUploader: cloudinary,
	}
}

func (c *UserUseCase) Create(ctx context.Context, request *model.RegisterUserRequest) (*model.UserResponseDetail, error) {
	tx := c.DB.WithContext(ctx).Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	err := c.Validate.Struct(request)
	if err != nil {
		c.Log.Warnf("Invalid request body : %+v", err)
		tx.Rollback()
		return nil, echo.ErrBadRequest
	}

	total, err := c.UserRepository.CountByWhere(tx, map[string]interface{}{
		"NIK": request.Nik,
	})
	if err != nil {
		c.Log.Warnf("Failed count user from database : %+v", err)
		tx.Rollback()
		return nil, echo.ErrInternalServerError
	}

	if total > 0 {
		c.Log.Warnf("User already exists : %+v", err)
		tx.Rollback()
		return nil, echo.ErrConflict
	}

	id, err := uuid.NewV7()
	if err != nil {
		c.Log.Warnf("Failed Create UUID Id : %+v", err)
		tx.Rollback()
		return nil, echo.ErrInternalServerError
	}
	user := &entity.MsUser{
		PkMsUser:   id,
		Nik:        request.Nik,
		FullName:   request.FullName,
		LegalName:  request.LegalName,
		BirthPlace: request.BirthPlace,
		BirthDate:  request.BirthDate,
		Salary:     request.Salary,
		Stamp: entity.Stamp{
			CreatedBy: id.String(),
			UpdatedBy: id.String(),
		},
	}

	var wg sync.WaitGroup
	errChan := make(chan error, 2)
	var base64FileKtp, base64FileSelfie string

	wg.Add(2)

	go func() {
		defer wg.Done()
		base64, err := util.GetBytesFile(request.ImageKtp)
		if err != nil {
			errChan <- err
			return
		}
		base64FileKtp = base64
	}()

	go func() {
		defer wg.Done()
		// base64, err := util.GetBytesFile(request.ImageSelfie)
		base64, err := c.CloudinaryUploader.UploadFromMultipartHeader(request.ImageSelfie)
		if err != nil {
			errChan <- err
			return
		}
		base64FileSelfie = base64
	}()

	// Close errChan after all uploads are done
	go func() {
		wg.Wait()
		close(errChan)
	}()

	for err := range errChan {
		if err != nil {
			c.Log.Warnf("Failed to upload images: %+v", err)
			tx.Rollback()
			return nil, echo.ErrInternalServerError
		}
	}

	user.ImageKtp = base64FileKtp
	user.ImageSelfie = base64FileSelfie
	if err := c.UserRepository.Create(tx, user); err != nil {
		c.Log.Warnf("Failed create user to database : %+v", err)
		tx.Rollback()
		return nil, echo.ErrInternalServerError
	}
	if err := tx.Commit().Error; err != nil {
		c.Log.Warnf("Failed to commit transaction: %+v", err)
		return nil, echo.ErrInternalServerError
	}

	response := converter.UserToResponse(user)
	return response, nil
}
