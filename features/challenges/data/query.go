package data

import (
	"backendgreeve/constant"
	"backendgreeve/features/challenges"
	"math"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ChallengeData struct {
	DB *gorm.DB
}

func New(db *gorm.DB) challenges.ChallengeDataInterface {
	return &ChallengeData{
		DB: db,
	}
}

func (cd *ChallengeData) GetAllForAdmin(page int) ([]challenges.Challenge, int, error) {
	// Perbaikan Kode Anda di sini
	var challenges []challenges.Challenge

	var totalChallenges int64
	countTx := cd.DB.Model(&Challenge{}).Count(&totalChallenges)
	if countTx.Error != nil {
		return nil, 0, countTx.Error
	}

	challengePerPage := 20
	totalPages := int(math.Ceil(float64(totalChallenges) / float64(challengePerPage)))

	tx := cd.DB.Model(&Challenge{}).Preload("ImpactCategories.ImpactCategory").
		Offset((page - 1) * challengePerPage).Limit(challengePerPage).Find(&challenges)
	if tx.Error != nil {
		return nil, 0, tx.Error
	}
	return challenges, totalPages, nil
}

func (cd *ChallengeData) GetAllForUser(userId string) ([]challenges.Challenge, error) {
	// Kode Anda di sini
	var challenges []challenges.Challenge
	subQuery := cd.DB.Model(&ChallengeLog{}).Where("user_id = ?", userId).Select("challenge_id")
	tx := cd.DB.Model(&Challenge{}).Preload("ImpactCategories.ImpactCategory").Where("id NOT IN (?)", subQuery).Find(&challenges)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return challenges, nil
}

func (cd *ChallengeData) GetByID(challengeId string) (challenges.Challenge, error) {
	// Kode Anda di sini
	var challenge challenges.Challenge
	tx := cd.DB.Model(&Challenge{}).Preload("ImpactCategories.ImpactCategory").Find(&challenge, "id = ?", challengeId)
	if tx.Error != nil {
		return challenges.Challenge{}, tx.Error
	}
	return challenge, nil
}

func (cd *ChallengeData) Swipe(userId string, challengeId string, challengeType string) error {
	isUserParticipate := cd.IsUserParticipate(userId, challengeId)
	if isUserParticipate {
		return constant.ErrUserAlreadyParticipate
	}

	tx := cd.DB.Create(&ChallengeLog{
		ID:          uuid.New().String(),
		UserID:      userId,
		ChallengeID: challengeId,
		Status:      challengeType,
	})

	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (cd *ChallengeData) AddToLogs(userId string, challengeId string, status string) error {
	// Kode Anda di sini
	return nil
}

func (cd *ChallengeData) Create(challenge challenges.Challenge) error {
	// Kode Anda di sini
	challengeQuery := Challenge{
		ID:          uuid.New().String(),
		Title:       challenge.Title,
		Description: challenge.Description,
		Exp:         challenge.Exp,
		Coin:        challenge.Coin,
		DateStart:   challenge.DateStart,
		DateEnd:     challenge.DateEnd,
		Difficulty:  challenge.Difficulty,
		ImageURL:    challenge.ImageURL,
	}

	for _, category := range challenge.ImpactCategories {
		challengeQuery.ChallengeImpactCategories = append(challengeQuery.ChallengeImpactCategories, ChallengeImpactCategory{
			ID:               uuid.New().String(),
			ChallengeID:      challengeQuery.ID,
			ImpactCategoryID: category.ImpactCategoryID,
		})
	}
	tx := cd.DB.Create(&challengeQuery)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (cd *ChallengeData) Update(challenge challenges.Challenge) error {
	// Kode Anda di sini
	return nil
}

func (cd *ChallengeData) Delete(challengeId string) error {
	// Kode Anda di sini
	return nil
}

func (cd *ChallengeData) IsUserParticipate(userId string, challengeId string) bool {
	var count int64
	cd.DB.Model(&ChallengeLog{}).Where("user_id = ? AND challenge_id = ?", userId, challengeId).Count(&count)
	if count > 0 {
		return true
	}
	return false
}
