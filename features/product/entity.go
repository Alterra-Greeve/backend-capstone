package product

import "github.com/labstack/echo/v4"

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
}

type ProductHandlerInterface interface {
	Get() echo.HandlerFunc
	GetById() echo.HandlerFunc
	GetByCategoryID() echo.HandlerFunc
	GetByName() echo.HandlerFunc
	Create() echo.HandlerFunc
	Update() echo.HandlerFunc
	Delete() echo.HandlerFunc
}

type ProductServiceInterface interface {
	Get() ([]Product, error)
	GetById(id string) (Product, error)
	GetByCategoryID(categoryID string) ([]Product, error)
	GetByName(name string) ([]Product, error)
	Create(product Product) error
	Update(product Product) error
	Delete(product Product) error
}

type ProductDataInterface interface {
	Get() ([]Product, error)
	GetById(id string) (Product, error)
	GetByCategoryID(categoryID string) ([]Product, error)
	GetByName(name string) ([]Product, error)
	Create(product Product) error
	Update(product Product) error
	Delete(product Product) error
}
