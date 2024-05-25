package data

import (
	"time"

	"gorm.io/gorm"
)

type ImpactCategory struct {
	*gorm.Model
	ID          string          `gorm:"primary_key;type:varchar(50);not null;column:id"`
	Name        string          `gorm:"type:varchar(255);not null;column:name"`
	ImpactPoint int             `gorm:"type:int;not null;column:impact_point"`
	IconURL     string          `gorm:"type:varchar(255);not null;column:icon_url"`
	CreatedAt   time.Time       `gorm:"column:created_at"`
	UpdatedAt   time.Time       `gorm:"column:updated_at"`
	DeletedAt   *gorm.DeletedAt `gorm:"column:deleted_at"`
}

func (ImpactCategory) TableName() string {
	return "impact_categories"
}
