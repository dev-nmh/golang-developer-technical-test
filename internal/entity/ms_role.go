package entity

type MsRole struct {
	PkMsRole int `gorm:"column:pk_ms_role;primaryKey"`
	DefaultEntityMasterColumn
	Stamp
}

func (e *MsRole) TableName() string {
	return "ms_role"
}
