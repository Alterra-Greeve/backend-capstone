package data

import (
	impactcategory "backendgreeve/features/impactcategory/data"

	"gorm.io/gorm"
)

type Product struct {
	*gorm.Model
	ID               string                  `gorm:"primary_key;type:varchar(50);not null;column:id"`
	Name             string                  `gorm:"type:varchar(255);not null;column:name"`
	Description      string                  `gorm:"type:varchar(255);column:description"`
	Price            float64                 `gorm:"type:int;not null;column:price"`
	Coin             int                     `gorm:"type:int;not null;column:coin"`
	Images           []ProductImage          `gorm:"foreignKey:ProductID"`
	ImpactCategories []ProductImpactCategory `gorm:"foreignKey:ProductID"`
}

type ProductImage struct {
	*gorm.Model
	ID        string  `gorm:"primary_key;type:varchar(50);not null;column:id"`
	ProductID string  `gorm:"type:varchar(50);not null;column:product_id"`
	ImageURL  string  `gorm:"type:varchar(255);not null;column:image_url"`
	Position  int     `gorm:"type:int;not null;column:position"`
	Product   Product `gorm:"foreignKey:ProductID;references:ID"`
}

type ProductImpactCategory struct {
	*gorm.Model
	ID               string                        `gorm:"primary_key;type:varchar(50);not null;column:id"`
	ProductID        string                        `gorm:"type:varchar(50);not null;column:product_id"`
	ImpactCategoryID string                        `gorm:"type:varchar(50);not null;column:impact_category_id"`
	Product          Product                       `gorm:"foreignKey:ProductID;references:ID"`
	ImpactCategory   impactcategory.ImpactCategory `gorm:"foreignKey:ImpactCategoryID;references:ID"`
}

func (product *Product) TableName() string {
	return "products"
}

func (product *ProductImage) TableName() string {
	return "product_images"
}

func (product *ProductImpactCategory) TableName() string {
	return "product_impact_categories"
}
