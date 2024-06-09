package impactcategory

import (
	"github.com/labstack/echo/v4"
)

type ImpactCategory struct {
	ID          string
	Name        string
	ImpactPoint int
	IconURL     string
	ImageURL    string
	Description string
}

type ImpactCategoryUpdate struct {
	ID          string
	Name        string
	ImpactPoint int
	IconURL     string
}

type ImpactCategoryHandlerInterface interface {
	GetAll() echo.HandlerFunc
	GetByID() echo.HandlerFunc
	Create() echo.HandlerFunc
	Update() echo.HandlerFunc
	Delete() echo.HandlerFunc
}

type ImpactCategoryServiceInterface interface {
	GetAll() ([]ImpactCategory, error)
	GetByID(ID string) (ImpactCategory, error)
	Create(ImpactCategory) error
	Update(ImpactCategoryUpdate) error
	Delete(ImpactCategory) error
}

type ImpactCategoryDataInterface interface {
	GetAll() ([]ImpactCategory, error)
	GetByID(ID string) (ImpactCategory, error)
	Create(ImpactCategory) error
	Update(ImpactCategoryUpdate) error
	Delete(ImpactCategory) error
}
