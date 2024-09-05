package entity

import "github.com/google/uuid"

type MsItemType struct {
	PkMsItemType uuid.UUID `gorm:"column:pk_ms_item_type;primaryKey"`
	DefaultEntityMasterColumn
	Stamp
}

func (e *MsItemType) TableName() string {
	return "ms_item_type"
}
