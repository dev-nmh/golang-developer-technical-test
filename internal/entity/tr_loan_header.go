package entity

import (
	"time"

	"github.com/google/uuid"
)

type TrLoanHeader struct {
	PkTrLoanHeader    uuid.UUID `gorm:"column:pk_tr_loan_header;primaryKey"`
	FkMsUser          uuid.UUID `gorm:"column:fk_ms_user"`
	FkMsPaymentStatus int       `gorm:"column:fk_ms_payment_status"`
	FkMsItemType      uuid.UUID `gorm:"column:fk_ms_item_type"`
	ContractNumber    float64   `gorm:"column:contract_number"`
	AssetName         time.Time `gorm:"column:asset_name"`
	Stamp
}

func (e *TrLoanHeader) TableName() string {
	return "tr_loan_header"
}
