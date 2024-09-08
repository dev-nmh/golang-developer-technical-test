package model

import (
	"time"

	"github.com/google/uuid"
)

/*
pk_tr_loan_header
fk_ms_user
fk_ms_payment_status
fk_ms_item_type
contract_number
asset_name
is_active
created_by
created_at
updated_by
updated_at
*/

type RequestLoan struct {
	FkMsUser        uuid.UUID `json:"user_id,omitempty" validate:"required"`
	CreateBy        uuid.UUID `json:"created_by,omitempty" validate:"required"`
	FkMsItemType    uuid.UUID `json:"item_type_id" validate:"required"`
	FkMsSource      string    `json:"source_id" validate:"required"`
	TenorId         string    `json:"tenor_id" validate:"required"`
	ContractNumber  string    `json:"contract_number" validate:"required"`
	AssetName       string    `json:"asset_name" validate:"required"`
	OtrAmount       float64   `json:"otr_amount" validate:"required"`
	TransactionDate time.Time `json:"transaction_date" validate:"required"`
	DueDate         time.Time `json:"due_date"`
}
