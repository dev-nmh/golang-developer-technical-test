package repository

import (
	"github/golang-developer-technical-test/internal/constant"
	"github/golang-developer-technical-test/internal/entity"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type TranscationLoanRepository struct {
	Header  Repository[entity.TrLoanHeader]
	Detail  Repository[entity.TrLoanDetail]
	Billing Repository[entity.TrLoanBilling]

	Log *logrus.Logger
}

func NewTranscationLoanRepository(log *logrus.Logger) *TranscationLoanRepository {
	return &TranscationLoanRepository{
		Log: log,
	}
}

func (tlr TranscationLoanRepository) GetTotalNotPaidLoanByUserTenorId(db *gorm.DB, userId uuid.UUID) (float64, error) {
	var total float64
	err := db.Model(entity.TrLoanHeader{}).Where("fk_ms_payment_status in ? ", []int{
		constant.PAYMENT_STATUS_ACTIVE, constant.PAYMENT_STATUS_DEFAULTED,
	}).Where("fk_ms_user", userId).Joins("JOIN tr_loan_detail ON tr_loan_header.pk_tr_loan_header = tr_loan_detail.fk_tr_loan_header").Select("sum(tr_loan_detail.otr_amount)").Scan(total).Error
	if err != nil {
		return 0, err
	}
	return total, nil
}
