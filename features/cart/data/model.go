package data

import (
	"backendgreeve/features/cart"
	"backendgreeve/features/impactcategory"
	product "backendgreeve/features/product"
	productData "backendgreeve/features/product/data"
	user "backendgreeve/features/users"
	userData "backendgreeve/features/users/data"

	"gorm.io/gorm"
)

type Cart struct {
	*gorm.Model
	ID                    string                              `gorm:"primary_key;type:varchar(50);column:id"`
	UserID                string                              `gorm:"not null;type:varchar(50);column:user_id"`
	ProductID             string                              `gorm:"not null;type:varchar(50);column:product_id"`
	Quantity              int                                 `gorm:"not null;column:quantity"`
	Product               productData.Product                 `gorm:"foreignKey:ProductID;references:ID"`
	ProductImage          []productData.ProductImage          `gorm:"foreignKey:ProductID;references:ID"`
	ProductImpactCategory []productData.ProductImpactCategory `gorm:"foreignKey:ProductID;references:ID"`
	User                  userData.User                       `gorm:"foreignKey:UserID;references:ID"`
}

func (c *Cart) TableName() string {
	return "carts"
}

func (c *Cart) CastDataToCartEntity() cart.Cart {
	var images []product.ProductImage
	var impactCategories []product.ProductImpactCategory

	for _, img := range c.Product.Images {
		images = append(images, product.ProductImage{
			ID:        img.ID,
			ProductID: img.ProductID,
			ImageURL:  img.ImageURL,
			Position:  img.Position,
		})
	}

	for _, impact := range c.Product.ImpactCategories {
		impactCategories = append(impactCategories, product.ProductImpactCategory{
			ID:               impact.ID,
			ProductID:        impact.ProductID,
			ImpactCategoryID: impact.ImpactCategoryID,
			ImpactCategory: impactcategory.ImpactCategory{
				ID:          impact.ImpactCategory.ID,
				Name:        impact.ImpactCategory.Name,
				ImpactPoint: impact.ImpactCategory.ImpactPoint,
				IconURL:     impact.ImpactCategory.IconURL,
			},
		})
	}

	return cart.Cart{
		User: user.User{
			ID:       c.UserID,
			Username: c.User.Username,
			Email:    c.User.Email,
			Address:  c.User.Address,
			Phone:    c.User.Phone,
		},
		Items: []cart.CartItems{
			{
				Quantity: c.Quantity,
				Product: product.Product{
					ID:               c.Product.ID,
					Name:             c.Product.Name,
					Description:      c.Product.Description,
					Price:            c.Product.Price,
					Coin:             c.Product.Coin,
					Stock:            c.Product.Stock,
					CreatedAt:        c.Product.CreatedAt,
					UpdatedAt:        c.Product.UpdatedAt,
					Images:           images,
					ImpactCategories: impactCategories,
				},
			},
		},
	}
}
