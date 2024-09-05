package entity

import (
	"time"

	"github.com/google/uuid"
)

type TrLoanBilling struct {
	PkTrLoanBilling   uuid.UUID `gorm:"column:pk_tr_loan_billing;primaryKey"`
	FkTrLoanHeader    uuid.UUID `gorm:"column:fk_tr_loan_header"`
	FkMsBillingStatus int       `gorm:"column:fk_ms_billing_status"`
	SortOrder         int       `gorm:"column:sort_order"`
	PayoffBalance     float64   `gorm:"column:payoff_balance"`
	ExpiredDate       time.Time `gorm:"column:expired_date"`
	Stamp
}

func (e *TrLoanBilling) TableName() string {
	return "tr_loan_billing"
}
