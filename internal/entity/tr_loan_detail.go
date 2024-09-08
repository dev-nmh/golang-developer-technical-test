package entity

import (
	"time"

	"github.com/google/uuid"
)

type TrLoanDetail struct {
	PkTrLoanDetail  uuid.UUID `gorm:"column:pk_tr_loan_detail;primaryKey"`
	FkTrLoanHeader  uuid.UUID `gorm:"column:fk_tr_loan_header"`
	FkMsSource      string    `gorm:"column:fk_ms_source"`
	FkMapUserTenor  uuid.UUID `gorm:"column:fk_map_user_tenor"`
	OtrAmount       float64   `gorm:"column:otr_amount"`
	LoanBalance     float64   `gorm:"column:loan_balance"`
	TransactionDate time.Time `gorm:"column:transaction_date"`
	DueDate         time.Time `gorm:"column:due_date"`

	Stamp
}

func (e *TrLoanDetail) TableName() string {
	return "tr_loan_detail"
}
