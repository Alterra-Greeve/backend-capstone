package handler

import (
	"backendgreeve/features/product"
)

type ProductResponse struct {
	ID          string                  `json:"product_id"`
	Name        string                  `json:"name"`
	Description string                  `json:"description"`
	Price       float64                 `json:"price"`
	Coin        int                     `json:"coin"`
	Stock       int                     `json:"stock"`
	CreatedAt   string                  `json:"created_at"`
	UpdatedAt   string                  `json:"updated_at"`
	Category    []ProductImpactCategory `json:"category"`
	Images      []ProductImage          `json:"images"`
}

type ProductAdminResponse struct {
	ID          string   `json:"product_id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Price       float64  `json:"price"`
	Coin        int      `json:"coin"`
	Stock       int      `json:"stock"`
	CreatedAt   string   `json:"created_at"`
	UpdatedAt   string   `json:"updated_at"`
	Category    []string `json:"category"`
	Images      []string `json:"image_url"`
}
type ProductImpactCategory struct {
	ImpactCategory ImpactCategory `json:"impact_category"`
}

type ProductImage struct {
	ImageURL string `json:"image_url"`
	Position int    `json:"position"`
}

type ImpactCategory struct {
	Name        string `json:"name"`
	ImpactPoint int    `json:"impact_point"`
	IconURL     string `json:"icon_url"`
}

type MetadataResponse struct {
	TotalPage int `json:"total_page"`
	Page      int `json:"current_page"`
}

func (p ProductResponse) ToResponse(product product.Product) ProductResponse {
	images := make([]ProductImage, len(product.Images))
	for i, image := range product.Images {
		images[i] = ProductImage{
			ImageURL: image.ImageURL,
			Position: image.Position,
		}
	}
	impactCategories := make([]ProductImpactCategory, len(product.ImpactCategories))
	for i, impactCategory := range product.ImpactCategories {
		impactCategories[i] = ProductImpactCategory{
			ImpactCategory: ImpactCategory{
				Name:        impactCategory.ImpactCategory.Name,
				ImpactPoint: impactCategory.ImpactCategory.ImpactPoint,
				IconURL:     impactCategory.ImpactCategory.IconURL,
			},
		}
	}
	return ProductResponse{
		ID:          product.ID,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		Coin:        product.Coin,
		Stock:       product.Stock,
		CreatedAt:   product.CreatedAt.Format("02-01-2006"),
		UpdatedAt:   product.UpdatedAt.Format("02-01-2006"),
		Images:      images,
		Category:    impactCategories,
	}
}

func (p ProductAdminResponse) ToResponse(product product.Product) ProductAdminResponse {
	images := make([]string, len(product.Images))
	for i, image := range product.Images {
		images[i] = image.ImageURL
	}

	category := make([]string, len(product.ImpactCategories))
	for i, impactCategory := range product.ImpactCategories {
		category[i] = impactCategory.ImpactCategory.ID
	}
	return ProductAdminResponse{
		ID:          product.ID,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		Coin:        product.Coin,
		Stock:       product.Stock,
		CreatedAt:   product.CreatedAt.Format("02-01-2006"),
		UpdatedAt:   product.UpdatedAt.Format("02-01-2006"),
		Images:      images,
		Category:    category,
	}
}
