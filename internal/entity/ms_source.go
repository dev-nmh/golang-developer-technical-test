package entity

type MsSource struct {
	PkMsSource int `gorm:"column:pk_ms_source;primaryKey"`
	DefaultEntityMasterColumn
	AdminFee float64 `gorm:"column:admin_fee"`
	Stamp
}

func (e *MsSource) TableName() string {
	return "ms_source"
}
