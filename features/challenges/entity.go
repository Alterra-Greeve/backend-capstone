package challenges

import "time"

type Challenge struct {
	ID          string
	Title       string
	Difficulty  string
	Description string
	Exp         int
	Coin        int
	DateStart   time.Time
	DateEnd     time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type ChallengeImage struct {
	ID          string
	ChallengeID string
	ImageURL    string
	Position    int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type ChallengeImpactCategory struct {
	ID               string
	ChallengeID      string
	ImpactCategoryID string
	CreatedAt        time.Time
	UpdatedAt        time.Time
}

type ChallengeLog struct {
	ID          string
	UserID      string
	ChallengeID string
	Status      string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type ChallengeHandlerInterface interface {
	CreateChallenge(Challenge) (Challenge, error)
	UpdateChallenge(Challenge) (Challenge, error)
	DeleteChallenge(string) error
}

type ChallengeServiceInterface interface {
	FindChallengeByTitle(string) (Challenge, error)
	FindChallengeByTitleAndUserID(string, string) (Challenge, error)
	FindChallengeByTitleAndUserIDAndStatus(string, string, string) (Challenge, error)
}

type ChallengeDataInterface interface {
	CreateChallenge(Challenge) (Challenge, error)
	UpdateChallenge(Challenge) (Challenge, error)
	DeleteChallenge(string) error
}

