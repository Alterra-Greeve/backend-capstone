package orders

import (
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
	ID          string
	Name        string
	ImpactPoint int
}

// Challenge
type ChallengeConfirmation struct {
	ID               string
	Username         string
	Status           string
	Challenge        Challenge
	ImpactCategories []ImpactCategory
	ImpactPointTotal int
	CreatedAt        time.Time
	UpdatedAt        time.Time
}

type Challenge struct {
	ID                        string
	Title                     string
	Difficulty                string
	Description               string
	Exp                       int
	Coin                      int
	ImageURL                  string
	DateStart                 time.Time
	DateEnd                   time.Time
	ChallengeImpactCategories []ChallengeImpactCategory
}

type ChallengeImpactCategory struct {
	ID               string
	ChallengeID      string
	ImpactCategoryID string
	ImpactCategory   ImpactCategory
}

//	type User struct {
//		ID       string
//		Username string
//		Email    string
//	}
type OrdersHandlerInterface interface {
	GetOrdersProduct() echo.HandlerFunc
	GetOrdersChallenge() echo.HandlerFunc
}

type OrdersServiceInterface interface {
	GetOrdersProduct() ([]OrdersProduct, error)
	GetOrdersChallenge() ([]ChallengeConfirmation, error)
}

type OrdersDataInterface interface {
	GetOrdersProduct() ([]OrdersProduct, error)
	// GetTotalImpactPoint() (int, error)
	GetTotalCoin() (int, error)

	// Challenge
	GetOrdersChallenge() ([]ChallengeConfirmation, error)
}
