package usecase

import (
	"context"
	"errors"
	"github/golang-developer-technical-test/internal/config"
	"github/golang-developer-technical-test/internal/constant"
	"github/golang-developer-technical-test/internal/entity"
	"github/golang-developer-technical-test/internal/model"
	"github/golang-developer-technical-test/internal/model/converter"
	"github/golang-developer-technical-test/internal/repository"
	"github/golang-developer-technical-test/internal/util"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type AccountUseCase struct {
	db                *gorm.DB
	log               *logrus.Logger
	validate          *validator.Validate
	config            *viper.Viper
	accountRepository *repository.AccountRepository
	userRepository    *repository.UserRepository
	jwtGenerator      *config.JwtGenerator
}

func NewAccountUseCase(db *gorm.DB, logger *logrus.Logger, validate *validator.Validate, config *viper.Viper, userRepository *repository.UserRepository, accountRepository *repository.AccountRepository, jwtGenerator *config.JwtGenerator) *AccountUseCase {
	return &AccountUseCase{
		db:                db,
		log:               logger,
		validate:          validate,
		config:            config,
		accountRepository: accountRepository,
		userRepository:    userRepository,
		jwtGenerator:      jwtGenerator,
	}
}

func (acc *AccountUseCase) Register(ctx context.Context, request *model.AccoutRequest) (*model.AccountResponse, error) {
	tx := acc.db.WithContext(ctx)
	total, err := acc.accountRepository.CountByWhere(tx, map[string]interface{}{
		"email": request.Email,
	})
	if err != nil {
		acc.log.Warnf("Failed to generate token: %+v", err)
		return nil, echo.ErrInternalServerError
	}
	if total > 0 {
		return nil, echo.ErrConflict
	}

	salt := uuid.New().String()
	if hashPassword, err := util.HashPassword(acc.config.GetString("app.app_key") + request.Password + salt); err != nil {
		return nil, echo.ErrInternalServerError
	} else {
		request.Password = hashPassword
	}
	accountId, err := uuid.NewV7()
	if err != nil {
		return nil, echo.ErrInternalServerError
	}
	create := &entity.MsAccount{
		PkMsAccount:  accountId,
		FkMsRole:     constant.USER_ROLES_USER_INT,
		Password:     request.Password,
		PasswordSalt: salt,
		Email:        request.Email,
		Stamp: entity.Stamp{
			CreatedBy: accountId.String(),
			UpdatedBy: accountId.String(),
		},
	}
	startTransaction := tx.Begin()
	defer startTransaction.Rollback()
	if err := acc.accountRepository.Create(startTransaction, create); err != nil {
		acc.log.Warnf("Error On Created Entity: %+v", err)
		return nil, echo.ErrInternalServerError
	}

	token, err := acc.jwtGenerator.CreateAccessTokenUser(&model.UserBaseClaims{
		BaseClaims: model.BaseClaims{
			ID:     create.PkMsAccount,
			Email:  create.Email,
			RoleId: strconv.Itoa(create.FkMsRole),
		},
		UserID: uuid.Nil,
	})

	if err != nil {
		acc.log.Warnf("Failed to generate token: %+v", err)
		startTransaction.Rollback()
		return nil, echo.ErrInternalServerError
	}
	if err := startTransaction.Commit().Error; err != nil {
		startTransaction.Rollback()
		acc.log.Warnf("Failed to commit transaction: %+v", err)
		return nil, echo.ErrInternalServerError
	}
	respose := converter.AccountToResponse(*create, &uuid.Nil, token, nil)
	return respose, nil
}

func (acc *AccountUseCase) Verify(ctx context.Context, request *model.AccoutRequest) (interface{}, error) {
	tx := acc.db.WithContext(ctx)
	var record entity.MsAccount
	err := acc.accountRepository.FindByWhere(tx, &record, map[string]interface{}{
		"email": request.Email,
	})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, echo.ErrNotFound

		}
		return nil, echo.ErrInternalServerError
	}
	if (record == entity.MsAccount{}) {
		return nil, echo.ErrNotFound
	}
	isTrue := util.CheckPasswordHash(acc.config.GetString("app.app_key")+request.Password+record.PasswordSalt, record.Password)
	if !isTrue {
		return nil, echo.ErrForbidden
	}
	if record.FkMsRole != constant.USER_ROLES_USER_INT {
		token, err := acc.jwtGenerator.CreateAccessTokenAdmin(&model.BaseClaims{
			ID:     record.PkMsAccount,
			Email:  record.Email,
			RoleId: strconv.Itoa(record.FkMsRole),
		})
		if err != nil {
			acc.log.Warnf("Failed to generate token: %+v", err)
			return nil, echo.ErrInternalServerError
		}
		respose := converter.AccountToResponse(record, nil, token, nil)
		return respose, nil
	}

	var user entity.MsUser
	err = acc.userRepository.FindByWhere(acc.db, &user, map[string]interface{}{
		"fk_ms_account": record.PkMsAccount,
	})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			user.PkMsUser = uuid.Nil
		} else {
			acc.log.Warnf("Failed to get user record: %+v", err)
			return nil, echo.ErrInternalServerError
		}
	}
	token, err := acc.jwtGenerator.CreateAccessTokenUser(&model.UserBaseClaims{
		BaseClaims: model.BaseClaims{
			ID:     record.PkMsAccount,
			Email:  record.Email,
			RoleId: strconv.Itoa(record.FkMsRole),
		},
		UserID: user.PkMsUser,
	})

	if err != nil {
		acc.log.Warnf("Failed to generate token: %+v", err)
		return nil, echo.ErrInternalServerError
	}
	respose := converter.AccountToResponse(record, &user.PkMsUser, token, nil)
	return respose, nil
}
