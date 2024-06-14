package dashboard

import (
	"backendgreeve/features/product"

	"github.com/labstack/echo/v4"
)

type Dashboard struct {
	TotalProduct              int
	TotalProductPercentage    string
	TotalNewProductThisMonth  int
	TotalNewProductPercentage string
	NewProduct                []product.Product
	TotalUser                 int
	NewUserPercentage         string
}

type MonthlyImpact struct {
	Date        string
	ImpactPoint []ImpactPoint
}

type ImpactPoint struct {
	Name       string
	TotalPoint int
}

type DashboardHandlerInterface interface {
	GetDashboard() echo.HandlerFunc
}

type DashboardServiceInterface interface {
	GetDashboard() (Dashboard, error)
	GetMonthlyImpact() ([]MonthlyImpact, error)
}

type DashboardDataInterface interface {
	GetDashboard() (Dashboard, error)
	GetTotalProduct() (int, error)
	GetTotalProductPercentage() (string, error)
	GetTotalNewProductThisMonth() (int, error)
	GetTotalNewProductPercentage() (string, error)
	GetNewProduct() ([]product.Product, error)
	GetTotalUser() (int, error)
	GetNewUserPercentage() (string, error)
	GetMonthlyImpactChallenge(month string) ([]ImpactPoint, error)
	GetMonthlyImpactProduct(month string) ([]ImpactPoint, error)
}
