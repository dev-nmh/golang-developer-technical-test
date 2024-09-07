package entity

type MsApprovalStatus struct {
	PkMsApprovalStatus int `gorm:"column:pk_ms_approval_status;primaryKey"`
	DefaultEntityMasterColumn
	Stamp
}

func (e *MsApprovalStatus) TableName() string {
	return "ms_approval_status"
}
