package entity

type MsTenor struct {
	PkMsTenor           string  `gorm:"column:pk_ms_tenor;primaryKey"`
	TenorMonths         int     `gorm:"column:tenor_months"`
	InterestRatePercent float64 `gorm:"column:interest_rate_percent"`
	Stamp
}

func (e *MsTenor) TableName() string {
	return "ms_tenor"
}
