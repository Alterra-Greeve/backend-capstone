package data

import (
	"gorm.io/gorm"
)

type ImpactCategory struct {
	*gorm.Model
	ID          string `gorm:"primary_key;type:varchar(50);not null;column:id"`
	Name        string `gorm:"type:varchar(255);not null;column:name"`
	ImpactPoint int    `gorm:"type:int;not null;column:impact_point"`
	Description string `gorm:"type:varchar(255);not null;column:description"`
	ImageURL    string `gorm:"type:varchar(255);not null;column:image_url"`
	IconURL     string `gorm:"type:varchar(255);not null;column:icon_url"`
}

func (ImpactCategory) TableName() string {
	return "impact_categories"
}
