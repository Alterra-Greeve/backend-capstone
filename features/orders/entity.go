package orders

import (
	"backendgreeve/features/challenges"
	"time"

	"github.com/labstack/echo/v4"
)

type OrdersProduct struct {
	UserID           string
	Username         string
	Email            string
	ImpactCategories []ImpactCategory
	ProductID        string
	ProductName      string
	Qty              int
	Coin             int
	ImpactPointTotal int
	TotalCoin        int
	CreatedAt        time.Time
	UpdatedAt        time.Time
}

type ImpactCategory struct {
	ID   string
	Name string
}

// Challenge
type OrderChallengeConfirmation struct {
	ID               string
	Username         string
	Status           string
	Challenge        challenges.Challenge
	ImpactCategories []ImpactCategory
	ImpactPointTotal int
	CreatedAt        time.Time
	UpdatedAt        time.Time
}
type Challenge struct {
	ID          string
	Title       string
	Difficulty  string
	Description string
	Exp         int
	Coin        int
	ImageURL    string
	DateStart   time.Time
	DateEnd     time.Time
}

type OrdersHandlerInterface interface {
	GetOrdersProduct() echo.HandlerFunc
	// GetOrdersChallenge() echo.HandlerFunc
}

type OrdersServiceInterface interface {
	GetOrdersProduct() ([]OrdersProduct, error)
	// GetOrdersChallenge() ([]OrderChallengeConfirmation, error)
}

type OrdersDataInterface interface {
	GetOrdersProduct() ([]OrdersProduct, error)
	GetTotalImpactPoint() (int, error)
	GetTotalCoin() (int, error)

	// Challenge
	// GetOrdersChallenge() ([]OrderChallengeConfirmation, error)
}
