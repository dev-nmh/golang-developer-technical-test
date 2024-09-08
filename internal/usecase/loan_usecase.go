package usecase

import (
	"context"
	"errors"
	"github/golang-developer-technical-test/internal/constant"
	"github/golang-developer-technical-test/internal/entity"
	"github/golang-developer-technical-test/internal/model"
	"github/golang-developer-technical-test/internal/model/converter"
	"github/golang-developer-technical-test/internal/repository"
	"time"

	"braces.dev/errtrace"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type LoanUseCase struct {
	db                  *gorm.DB
	log                 *logrus.Logger
	validate            *validator.Validate
	userRepository      *repository.UserRepository
	userTenorRepository *repository.UserTenorRepository
}

func NewLoanUseCase(db *gorm.DB, log *logrus.Logger, validate *validator.Validate, userRepository *repository.UserRepository, userTenorRepository *repository.UserTenorRepository) *LoanUseCase {
	return &LoanUseCase{
		db:                  db,
		log:                 log,
		validate:            validate,
		userRepository:      userRepository,
		userTenorRepository: userTenorRepository,
	}
}

func (loan LoanUseCase) ApprovalUser(ctx context.Context, req *model.UserApproval) (interface{}, error) {
	tx := loan.db.WithContext(ctx)
	var recordUser entity.MsUser
	err := loan.userRepository.FindByWhere(tx, &recordUser, map[string]interface{}{
		"pk_ms_user": req.UserId,
	})
	if err != nil {
		loan.log.Warnf("Failed to generate token: %+v", err)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, echo.ErrNotFound
		}
		return nil, echo.ErrInternalServerError

	}

	if recordUser.FkMsApprovalStatus == req.ApprovalId {
		return nil, echo.ErrConflict
	}
	strTrx := tx.Begin()
	defer strTrx.Rollback()
	recordUser.FkMsApprovalStatus = req.ApprovalId
	recordUser.UpdatedAt = time.Now()
	recordUser.UpdatedBy = req.AdminId.String()
	if err := loan.userRepository.Update(tx, &recordUser); err != nil {
		loan.log.Warnf("Failed to generate token: %+v", err)
		return nil, errtrace.Wrap(err)
	}
	if recordUser.FkMsApprovalStatus == constant.APPROVAL_STATUS_REJECTED {
		if err := strTrx.Commit().Error; err != nil {
			return nil, errtrace.Wrap(err)
		}
		req.UserTenor = []model.UserLimitTenor{}
		return req, nil
	}

	entitiesUserTenor, err := converter.UserTenorsToEntities(req.UserTenor, req.UserId, req.AdminId, req.AdminId)
	if err != nil {
		return nil, errtrace.Wrap(err)
	}
	if err := loan.userTenorRepository.DeleteUserTenorByUserId(strTrx, req.UserId); err != nil {
		return nil, errtrace.Wrap(err)
	}

	if err := loan.userTenorRepository.CreateListUserTenor(strTrx, entitiesUserTenor); err != nil {
		return nil, errtrace.Wrap(err)
	}

	if err := strTrx.Commit().Error; err != nil {
		return nil, errtrace.Wrap(err)
	}
	return req, nil
}
