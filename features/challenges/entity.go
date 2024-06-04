package challenges

import (
	"backendgreeve/features/impactcategory"
	"time"

	"github.com/labstack/echo/v4"
)

type Challenge struct {
	ID               string
	Title            string
	Difficulty       string
	Description      string
	Exp              int
	Coin             int
	ImageURL         string
	DateStart        time.Time
	DateEnd          time.Time
	ImpactCategories []ChallengeImpactCategory
}

type ChallengeImpactCategory struct {
	ID               string
	ChallengeID      string
	ImpactCategoryID string
	ImpactCategory   impactcategory.ImpactCategory
}

type ImpactCategory struct {
	ID          string
	Name        string
	ImpactPoint int
}

type ChallengeHandlerInterface interface {
	GetAllForUser() echo.HandlerFunc
	GetByID() echo.HandlerFunc
	Swipe() echo.HandlerFunc

	GetAllForAdmin() echo.HandlerFunc
	Create() echo.HandlerFunc
	Update() echo.HandlerFunc
	Delete() echo.HandlerFunc
}

type ChallengeServiceInterface interface {
	GetAllForUser() ([]Challenge, int, error)
	GetByID(challengeId string) (Challenge, error)
	SwipeChallenge(userId string, challengeId string) (Challenge, error)

	GetAllForAdmin(page int) ([]Challenge, int, error)
	Create(challenge Challenge) error
	Update(challenge Challenge) error
	Delete(challengeId string) error
}

type ChallengeDataInterface interface {
	GetAllForAdmin(page int) ([]Challenge, int, error)
	GetAllForUser() ([]Challenge, int, error)
	GetByID(challengeId string) (Challenge, error)
	Swipe(userId string, challengeId string) error
	AddToLogs(userId string, challengeId string, status string) error
	Create(Challenge) error
	Update(Challenge) error
	Delete(string) error
}
