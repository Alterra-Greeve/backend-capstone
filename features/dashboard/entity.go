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
	TotalMembership           int
}

type MonthlyImpact struct {
	Date        string
	ImpactPoint []ImpactPoint
}

type ImpactPoint struct {
	Name       string
	TotalPoint int
}

type UserCoin struct {
	ID   string `json:"id"`
	Type string `json:"type"`
	Coin int    `json:"coin"`
}

type DashboardHandlerInterface interface {
	GetDashboard() echo.HandlerFunc
	GetUserCoin() echo.HandlerFunc
}

type DashboardServiceInterface interface {
	GetDashboard() (Dashboard, error)
	GetMonthlyImpact() ([]MonthlyImpact, error)
	GetUserCoin(userID string) ([]UserCoin, error)
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
	GetTotalMembership() (int, error)

	GetTransactionCoinEarned(userId string) ([]UserCoin, error)
	GetChallengeCoinEarned(userId string) ([]UserCoin, error)
}
