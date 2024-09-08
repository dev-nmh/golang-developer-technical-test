package usecase

import (
	"context"
	"errors"
	"github/golang-developer-technical-test/internal/constant"
	"github/golang-developer-technical-test/internal/entity"
	"github/golang-developer-technical-test/internal/model"
	"github/golang-developer-technical-test/internal/repository"
	"github/golang-developer-technical-test/internal/util"
	"net/http"
	"time"

	"braces.dev/errtrace"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type TranscationLoanUseCase struct {
	db                        *gorm.DB
	log                       *logrus.Logger
	validate                  *validator.Validate
	userRepository            *repository.UserRepository
	userTenorRepository       *repository.UserTenorRepository
	transcationLoanRepository *repository.TranscationLoanRepository
	tenorLoanRepository       *repository.TenorRepository
	sourceRepository          *repository.SourceRepository
}

func NewTranscationLoanUseCase(db *gorm.DB, log *logrus.Logger, validate *validator.Validate, userRepository *repository.UserRepository, userTenorRepository *repository.UserTenorRepository, transcationLoanRepository *repository.TranscationLoanRepository, tenorLoanRepository *repository.TenorRepository, sourceRepository *repository.SourceRepository) *TranscationLoanUseCase {
	return &TranscationLoanUseCase{
		db:                        db,
		log:                       log,
		validate:                  validate,
		userRepository:            userRepository,
		userTenorRepository:       userTenorRepository,
		transcationLoanRepository: transcationLoanRepository,
		tenorLoanRepository:       tenorLoanRepository,
		sourceRepository:          sourceRepository,
	}
}

func (c TranscationLoanUseCase) CreateLoan(ctx context.Context, req *model.RequestLoan) (interface{}, int, error) {
	tx := c.db.WithContext(ctx).Begin()
	defer tx.Rollback()
	c.log.Println("Masuk Pak eko -1")
	if approvalStatus, err := c.userRepository.GetApprovalStatusById(tx, req.FkMsUser); err != nil {
		c.log.Warnf("Failed to get approvalStatus: %+v", err)
		return nil, http.StatusInternalServerError, errtrace.Wrap(err)
	} else if approvalStatus != constant.APPROVAL_STATUS_APPROVED {
		return nil, http.StatusForbidden, errtrace.Wrap(echo.ErrForbidden)
	}

	var tenorLimit entity.MsTenor
	var userLimitTenor entity.MapUserTenor
	var source entity.MsSource
	c.log.Println("Masuk Pak eko 0")

	var total *float64
	if err := c.userTenorRepository.FindByWhere(tx, &userLimitTenor, map[string]interface{}{
		"fk_ms_user":  req.FkMsUser,
		"fk_ms_tenor": req.TenorId,
	}); err != nil {
		return nil, http.StatusNotFound, errtrace.Wrap(err)
	}
	c.log.Println("Masuk Pak eko 1")
	if err := c.tenorLoanRepository.FindByWhere(tx, &tenorLimit, map[string]interface{}{
		"pk_ms_tenor": req.TenorId,
	}); err != nil {
		return nil, http.StatusNotFound, errtrace.Wrap(err)
	}
	c.log.Println("Masuk Pak eko 2")
	if err := c.sourceRepository.FindByWhere(tx, &source, map[string]interface{}{
		"pk_ms_source": req.FkMsSource,
	}); err != nil {
		return nil, http.StatusNotFound, errtrace.Wrap(err)
	}
	c.log.Println("Masuk Pak eko 3")
	total, err := c.transcationLoanRepository.GetTotalNotPaidLoanByUserTenorId(tx, req.FkMsUser)
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, http.StatusInternalServerError, errtrace.Wrap(err)
		}
	}
	if total != nil {
		if req.OtrAmount > *total {
			return nil, http.StatusNotAcceptable, errtrace.Wrap(err)

		}
	}
	headerId, _ := uuid.NewV7()
	now := time.Now()
	stamp := entity.Stamp{
		CreatedBy: req.FkMsUser.String(),
		UpdatedAt: now,
		UpdatedBy: req.FkMsUser.String(),
		CreatedAt: now,
	}
	loanHeaderEntity := entity.TrLoanHeader{
		PkTrLoanHeader:    headerId,
		FkMsUser:          req.FkMsUser,
		FkMsPaymentStatus: constant.PAYMENT_STATUS_ACTIVE,
		FkMsItemType:      req.FkMsItemType,
		ContractNumber:    req.ContractNumber,
		AssetName:         req.AssetName,
		Stamp:             stamp,
	}
	loanDetailEntity := entity.TrLoanDetail{
		PkTrLoanDetail:  uuid.New(),
		FkTrLoanHeader:  headerId,
		FkMsSource:      req.FkMsSource,
		FkMapUserTenor:  userLimitTenor.PkMapUserTenor,
		OtrAmount:       req.OtrAmount,
		LoanBalance:     req.OtrAmount + util.GenerateBasicInterest(req.OtrAmount, tenorLimit.InterestRatePercent, float64(tenorLimit.TenorMonths), source.AdminFee),
		TransactionDate: now,
		DueDate:         now.Local().AddDate(0, tenorLimit.TenorMonths, 0),
		Stamp:           stamp,
	}
	if err := c.transcationLoanRepository.Header.Create(tx, &loanHeaderEntity); err != nil {
		return nil, http.StatusOK, errtrace.Wrap(err)
	}
	if err := c.transcationLoanRepository.Detail.Create(tx, &loanDetailEntity); err != nil {
		return nil, http.StatusOK, errtrace.Wrap(err)
	}
	if err := tx.Commit().Error; err != nil {
		return nil, http.StatusInternalServerError, errtrace.Wrap(err)
	}
	// c.userTenorRepository.FindByWhere(tx, UserLimitTenor)
	return req, http.StatusOK, nil
}
