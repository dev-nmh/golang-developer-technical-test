package entity

import "github.com/google/uuid"

type MapUserTenor struct {
	PkMapUserTenor uuid.UUID `gorm:"column:pk_map_user_tenor;primaryKey"`
	FkMsUser       uuid.UUID `gorm:"column:fk_ms_user"`
	FkMsTenor      string    `gorm:"column:fk_ms_tenor"`
	Amount         float64   `gorm:"column:amount"`
	Stamp
}

func (e *MapUserTenor) TableName() string {
	return "map_user_tenor"
}
