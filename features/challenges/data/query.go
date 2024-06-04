package data

import (
	"backendgreeve/features/challenges"
	"math"

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

func (cd *ChallengeData) GetAllForUser() ([]challenges.Challenge, int, error) {
	// Kode Anda di sini
	return nil, 0, nil
}

func (cd *ChallengeData) GetByID(challengeId string) (challenges.Challenge, error) {
	// Kode Anda di sini
	return challenges.Challenge{}, nil
}

func (cd *ChallengeData) Swipe(userId string, challengeId string) error {
	// Kode Anda di sini
	return nil
}

func (cd *ChallengeData) AddToLogs(userId string, challengeId string, status string) error {
	// Kode Anda di sini
	return nil
}

func (cd *ChallengeData) Create(challenge challenges.Challenge) error {
	// Kode Anda di sini
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
