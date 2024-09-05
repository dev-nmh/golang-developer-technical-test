package entity

type MsTenor struct {
	PkMsTenor int `gorm:"column:pk_ms_tenor;primaryKey"`
	DefaultEntityMasterColumn
	Stamp
}

func (e *MsTenor) TableName() string {
	return "ms_tenor"
}
