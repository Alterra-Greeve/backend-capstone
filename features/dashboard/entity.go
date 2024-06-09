package dashboard

import (
	"backendgreeve/features/product"
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

type DashboardHandlerInterface interface {
	GetDashboard() (Dashboard, error)
}

type DashboardServiceInterface interface {
	GetDashboard() (Dashboard, error)
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
}
