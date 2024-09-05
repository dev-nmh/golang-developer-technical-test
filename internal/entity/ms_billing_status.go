package entity

type MsBillingStatus struct {
	PkMsBillingStatus int `gorm:"column:pk_ms_billing_status;primaryKey"`
	DefaultEntityMasterColumn
	Stamp
}

func (e *MsBillingStatus) TableName() string {
	return "ms_billing_status"
}
