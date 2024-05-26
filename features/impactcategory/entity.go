package impactcategory

import (
	"time"

	"github.com/labstack/echo/v4"
)

type ImpactCategory struct {
	ID          string
	Name        string
	ImpactPoint int
	IconURL     string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type ImpactCategoryUpdate struct {
	ID          string
	Name        string
	ImpactPoint int
	IconURL     string
	UpdatedAt   time.Time
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
