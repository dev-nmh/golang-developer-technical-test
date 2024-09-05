package entity

type MsSource struct {
	PkMsSource int `gorm:"column:pk_ms_source;primaryKey"`
	DefaultEntityMasterColumn
	Stamp
}

func (e *MsSource) TableName() string {
	return "ms_source"
}
