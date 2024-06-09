package data

import (
	impactcategory "backendgreeve/features/impactcategory/data"
	user "backendgreeve/features/users/data"

	"gorm.io/gorm"
)

type Product struct {
	*gorm.Model
	ID               string                  `gorm:"primary_key;type:varchar(50);not null;column:id"`
	Name             string                  `gorm:"type:varchar(255);not null;column:name;index:,class:FULLTEXT,option:WITH PARSER ngram VISIBLE"`
	Description      string                  `gorm:"type:varchar(255);column:description"`
	Price            float64                 `gorm:"type:float;not null;column:price"`
	Coin             int                     `gorm:"type:int;not null;column:coin"`
	Stock            int                     `gorm:"type:int;not null;column:stock"`
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

type ProductLog struct {
	*gorm.Model
	ID        string    `gorm:"primary_key;type:varchar(50);not null;column:id"`
	UserID    string    `gorm:"type:varchar(50);not null;column:user_id"`
	ProductID string    `gorm:"type:varchar(50);not null;column:product_id"`
	User      user.User `gorm:"foreignKey:UserID;references:ID"`
	Product   Product   `gorm:"foreignKey:ProductID;references:ID"`
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

func (product *ProductLog) TableName() string {
	return "product_logs"
}
