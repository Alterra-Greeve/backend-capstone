package data

import (
	product "backendgreeve/features/product/data"
	user "backendgreeve/features/users/data"

	"gorm.io/gorm"
)

type Cart struct {
	*gorm.Model
	ID                    string                          `gorm:"primary_key;type:varchar(50);column:id"`
	UserID                string                          `gorm:"not null;type:varchar(50);column:user_id"`
	ProductID             string                          `gorm:"not null;type:varchar(50);column:product_id"`
	Quantity              int                             `gorm:"not null;column:quantity"`
	Product               product.Product                 `gorm:"foreignKey:ProductID;references:ID"`
	ProductImage          []product.ProductImage          `gorm:"foreignKey:ProductID;references:ID"`
	ProductImpactCategory []product.ProductImpactCategory `gorm:"foreignKey:ProductID;references:ID"`
	User                  user.User                       `gorm:"foreignKey:UserID;references:ID"`
}

func (c *Cart) TableName() string {
	return "carts"
}
