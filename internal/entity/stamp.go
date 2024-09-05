package entity

import "time"

type DefaultEntityMasterColumn struct {
	Title       string `gorm:"column:title"`
	Description string `gorm:"column:description"`
}

type Stamp struct {
	IsActive  bool      `gorm:"column:is_active;default:1"`
	CreatedBy string    `gorm:"column:created_by;<-:create;"`
	CreatedAt time.Time `gorm:"column:created_at;<-:create;autoCreateTime:true;default:CURRENT_TIMESTAMP()"`
	UpdatedBy string    `gorm:"column:updated_by;"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoUpdateTime;default:CURRENT_TIMESTAMP()"`
}
