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
	ImageURL         string
	DateStart        time.Time
	DateEnd          time.Time
	ImpactCategories []ChallengeImpactCategory
}

type ChallengeImpactCategory struct {
	ID               string
	ChallengeID      string
	ImpactCategoryID string
	ImpactCategory   ImpactCategory
}

type ImpactCategory struct {
	ID          string
	Name        string
	ImpactPoint int
	IconURL     string
}

type ChallengeConfirmation struct {
	ID        string
	UserID    string
	Status    string
	Challenge Challenge
}

type ChallengeConfirmationUpdate struct {
	ID    string
	Image []string
}
type ChallengeHandlerInterface interface {
	GetAllForUser() echo.HandlerFunc
	GetByID() echo.HandlerFunc
	Swipe() echo.HandlerFunc
	GetUserParticipate() echo.HandlerFunc
	GetAllForAdmin() echo.HandlerFunc
	Create() echo.HandlerFunc
	Update() echo.HandlerFunc
	Delete() echo.HandlerFunc
}

type ChallengeServiceInterface interface {
	GetAllForUser(userId string) ([]Challenge, error)
	GetByID(challengeId string) (Challenge, error)
	Swipe(userId string, challengeId string, challengeType string) error
	GetUserParticipate(userId string, status string) ([]ChallengeConfirmation, error)
	GetChallengeParticipant(challengeId string) (int, error)

	GetAllForAdmin(page int) ([]Challenge, int, error)
	Create(challenge Challenge) error
	Update(challenge Challenge) error
	Delete(challengeId string) error
}

type ChallengeDataInterface interface {
	GetAllForAdmin(page int) ([]Challenge, int, error)
	GetAllForUser(userId string) ([]Challenge, error)
	GetChallengeParticipant(challengeId string) (int, error)
	GetByID(challengeId string) (Challenge, error)
	Swipe(userId string, challengeId string, challengeType string) error
	AddToLogs(userId string, challengeId string, status string) error
	IsUserParticipate(userId string, challengeId string) (error, bool)
	GetUserParticipate(userId string, status string) ([]ChallengeConfirmation, error)
	Create(Challenge) error
	Update(Challenge) error
	Delete(string) error
}
