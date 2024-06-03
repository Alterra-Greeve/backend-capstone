package challenges

import (
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
	DateStart        time.Time
	DateEnd          time.Time
	Images           []ChallengeImage
	ImpactCategories []ChallengeImpactCategory
}

type ChallengeImage struct {
	ID          string
	ChallengeID string
	ImageURL    string
	Position    int
}

type ChallengeImpactCategory struct {
	ID               string
	ChallengeID      string
	ImpactCategoryID string
}

type ChallengeHandlerInterface interface {
	GetAllChallengeForUser() echo.HandlerFunc
	GetUserChallenge() echo.HandlerFunc
	GetChallengeByID() echo.HandlerFunc
	SwipeChallenge() echo.HandlerFunc

	GetAllChallengeForAdmin() echo.HandlerFunc
	CreateChallenge() echo.HandlerFunc
	UpdateChallenge() echo.HandlerFunc
	DeleteChallenge() echo.HandlerFunc
}

type ChallengeServiceInterface interface {
	GetAllChallengeForUser() ([]Challenge, int, error)
	GetUserChallenge(userId string) ([]Challenge, int, error)
	GetChallengeByID(challengeId string) (Challenge, error)
	SwipeChallenge(userId string, challengeId string) (Challenge, error)

	CreateChallenge(challenge Challenge) error
	UpdateChallenge(challenge Challenge) error
	DeleteChallenge(challengeId string) error
}

type ChallengeDataInterface interface {
	CreateChallenge(Challenge) (Challenge, error)
	UpdateChallenge(Challenge) (Challenge, error)
	DeleteChallenge(string) error
}
