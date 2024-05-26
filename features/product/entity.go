package product

import (
	"backendgreeve/features/impactcategory"

	"github.com/labstack/echo/v4"
)

type Product struct {
	ID               string
	Name             string
	Description      string
	Price            float64
	Coin             int
	Images           []ProductImage
	ImpactCategories []ProductImpactCategory
}

type ProductImage struct {
	ID        string
	ProductID string
	ImageURL  string
	Position  int
}

type ProductImpactCategory struct {
	ID               string
	ProductID        string
	ImpactCategoryID string
	ImpactCategory   impactcategory.ImpactCategory
}

type ImpactCategory struct {
	ID          string
	Name        string
	ImpactPoint int
}

type ProductHandlerInterface interface {
	Get() echo.HandlerFunc
	GetById() echo.HandlerFunc
	GetByCategory() echo.HandlerFunc
	GetByName() echo.HandlerFunc
	Create() echo.HandlerFunc
	Update() echo.HandlerFunc
	Delete() echo.HandlerFunc
}

type ProductServiceInterface interface {
	Get() ([]Product, error)
	GetById(id string) (Product, error)
	GetByPage(page int) ([]Product, int, error)
	GetByCategory(category string, page int) ([]Product, int, error)
	GetByName(name string, page int) ([]Product, int, error)
	Create(product Product) error
	Update(product Product) error
	Delete(productId string) error
}

type ProductDataInterface interface {
	Get() ([]Product, error)
	GetByPage(page int) ([]Product, int, error)
	GetById(id string) (Product, error)
	GetByCategory(category string, page int) ([]Product, int, error)
	GetByName(name string, page int) ([]Product, int, error)
	Create(product Product) error
	Update(product Product) error
	Delete(productId string) error
}
