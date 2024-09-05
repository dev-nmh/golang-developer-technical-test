package entity

type MsPaymentStatus struct {
	PkMsPaymentStatus int `gorm:"column:pk_ms_payment_status;primaryKey"`
	DefaultEntityMasterColumn
	Stamp
}

func (e *MsPaymentStatus) TableName() string {
	return "ms_payment_status"
}
