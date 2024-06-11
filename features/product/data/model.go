package data

import (
	impactcategoryentity "backendgreeve/features/impactcategory"
	impactcategory "backendgreeve/features/impactcategory/data"
	product "backendgreeve/features/product"
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

type ProductReccomendation struct {
	*gorm.Model
	ID               string                        `gorm:"primary_key;type:varchar(50);not null;column:id"`
	UserID           string                        `gorm:"type:varchar(50);not null;column:user_id"`
	ImpactCategoryID string                        `gorm:"type:varchar(50);not null;column:impact_category_id"`
	User             user.User                     `gorm:"foreignKey:UserID;references:ID"`
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

func (product *ProductLog) TableName() string {
	return "product_logs"
}

func (pd *Product) ToEntity() product.Product {
	impactCategories := make([]product.ProductImpactCategory, len(pd.ImpactCategories))
	for i, impactCategory := range pd.ImpactCategories {
		impactCategories[i] = product.ProductImpactCategory{
			ID:               impactCategory.ImpactCategoryID,
			ImpactCategoryID: impactCategory.ImpactCategoryID,
			ProductID:        impactCategory.ProductID,
			ImpactCategory: impactcategoryentity.ImpactCategory{
				Name:        impactCategory.ImpactCategory.Name,
				ImpactPoint: impactCategory.ImpactCategory.ImpactPoint,
				IconURL:     impactCategory.ImpactCategory.IconURL,
			},
		}
	}
	images := make([]product.ProductImage, len(pd.Images))
	for i, image := range pd.Images {
		images[i] = product.ProductImage{
			ID:       image.ID,
			ImageURL: image.ImageURL,
			Position: image.Position,
		}
	}
	return product.Product{
		ID:               pd.ID,
		Name:             pd.Name,
		Description:      pd.Description,
		Price:            pd.Price,
		Coin:             pd.Coin,
		Stock:            pd.Stock,
		CreatedAt:        pd.CreatedAt,
		UpdatedAt:        pd.UpdatedAt,
		Images:           images,
		ImpactCategories: impactCategories,
	}
}
