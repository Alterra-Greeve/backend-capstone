package data

import (
	"backendgreeve/constant"
	"backendgreeve/features/challenges"
	"log"
	"math"
	"time"

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
	var challenges []challenges.Challenge
	subQuery := cd.DB.Model(&ChallengeLog{}).Where("user_id = ?", userId).Select("challenge_id")
	tx := cd.DB.Model(&Challenge{}).Preload("ImpactCategories.ImpactCategory").Where("id NOT IN (?)", subQuery).Find(&challenges)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return challenges, nil
}

func (cd *ChallengeData) GetChallengeParticipant(challengeId string) (int, error) {
	var count int64
	tx := cd.DB.Model(&ChallengeLog{}).Where("challenge_id = ? AND status = ?", challengeId, "Diterima").Count(&count)
	if tx.Error != nil {
		return 0, tx.Error
	}
	return int(count), nil
}
func (cd *ChallengeData) GetByID(challengeId string) (challenges.Challenge, error) {
	// Kode Anda di sini
	var challenge challenges.Challenge
	tx := cd.DB.Model(&Challenge{}).Preload("ImpactCategories.ImpactCategory").Find(&challenge, "id = ?", challengeId)
	if tx.Error != nil {
		return challenges.Challenge{}, tx.Error
	}
	log.Println(challenge)
	return challenge, nil
}

func (cd *ChallengeData) GetUserParticipate(userId string, status string) ([]challenges.ChallengeConfirmation, error) {
	var challengeConfirmations []ChallengeConfirmation

	// Memuat challenge confirmation dengan challenge dan impact categories terkait
	tx := cd.DB.Model(&ChallengeConfirmation{}).
		Preload("Challenge").
		Preload("Challenge.ChallengeImpactCategories").
		Preload("Challenge.ChallengeImpactCategories.ImpactCategory").
		Find(&challengeConfirmations)
	if tx.Error != nil {
		return nil, tx.Error
	}

	// Konversi dari []ChallengeConfirmation ke challenges.ChallengeConfirmation
	var challengeData []challenges.ChallengeConfirmation
	for _, confirmation := range challengeConfirmations {
		// Membuat slice untuk menyimpan impact categories
		var impactCategories []challenges.ChallengeImpactCategory
		for _, category := range confirmation.Challenge.ChallengeImpactCategories {
			impactCategories = append(impactCategories, challenges.ChallengeImpactCategory{
				ID:               category.ID,
				ChallengeID:      category.ChallengeID,
				ImpactCategoryID: category.ImpactCategoryID,
				ImpactCategory: challenges.ImpactCategory{
					ID:          category.ImpactCategory.ID,
					Name:        category.ImpactCategory.Name,
					ImpactPoint: category.ImpactCategory.ImpactPoint,
					IconURL:     category.ImpactCategory.IconURL,
				},
			})
		}

		// Memasukkan data ke dalam slice challengeData
		challengeData = append(challengeData, challenges.ChallengeConfirmation{
			ID:     confirmation.ID,
			UserID: confirmation.UserID,
			Status: confirmation.Status,
			Challenge: challenges.Challenge{
				ID:               confirmation.Challenge.ID,
				Title:            confirmation.Challenge.Title,
				Difficulty:       confirmation.Challenge.Difficulty,
				Description:      confirmation.Challenge.Description,
				Exp:              confirmation.Challenge.Exp,
				Coin:             confirmation.Challenge.Coin,
				ImageURL:         confirmation.Challenge.ImageURL,
				DateStart:        confirmation.Challenge.DateStart,
				DateEnd:          confirmation.Challenge.DateEnd,
				ImpactCategories: impactCategories,
			},
		})
	}

	return challengeData, nil
}

func (cd *ChallengeData) Swipe(userId string, challengeId string, challengeType string) error {
	err, isUserParticipate := cd.IsUserParticipate(userId, challengeId)
	if err != nil {
		return err
	}
	if isUserParticipate {
		return constant.ErrUserAlreadyParticipate
	}
	tx := cd.DB.Begin()
	if tx.Error != nil {
		return tx.Error
	}
	err = cd.AddToLogs(userId, challengeId, challengeType)
	if err != nil {
		tx.Rollback()
		return constant.ErrInsertChallengeLog
	}
	switch {
	case challengeType == "Diterima":
		err := cd.ParticipateChallenge(userId, challengeId)
		if err != nil {
			tx.Rollback()
			return err
		}
		tx.Commit()
		return nil
	case challengeType == "Ditolak":
		tx.Commit()
		return nil
	default:
		tx.Rollback()
		return constant.ErrChallengeLogType
	}
}

func (cd *ChallengeData) AddToLogs(userId string, challengeId string, challengeType string) error {
	userLog := ChallengeLog{
		ID:          uuid.New().String(),
		UserID:      userId,
		ChallengeID: challengeId,
		Status:      challengeType,
	}

	tx := cd.DB.Create(&userLog)

	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

func (cd *ChallengeData) Create(challenge challenges.Challenge) error {
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
	tx := cd.DB.Model(&Challenge{}).Where("id = ?", challenge.ID).Updates(&challenge)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (cd *ChallengeData) Delete(challengeId string) error {
	tx := cd.DB.Delete(&Challenge{}, "id = ?", challengeId)

	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

func (cd *ChallengeData) ParticipateChallenge(userId string, challengeId string) error {
	challengeConfrimation := ChallengeConfirmation{
		ID:          uuid.New().String(),
		ChallengeID: challengeId,
		UserID:      userId,
		Status:      "Pending",
	}
	tx := cd.DB.Create(&challengeConfrimation)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (cd *ChallengeData) UploadParticipateProof(userId string, challengeConfirmationId string, images []string) error {
	tx := cd.DB.Begin()
	if tx.Error != nil {
		return tx.Error
	}
	err := cd.DeleteImageProof(challengeConfirmationId)
	if err != nil {
		tx.Rollback()
		return err
	}
	err = cd.DB.Model(&ChallengeConfirmation{}).Where("id = ?", challengeConfirmationId).Update("date_upload", time.Now()).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	for _, img := range images {
		uploadParticipateProof := ChallengeConfirmationImage{
			ID:                      uuid.New().String(),
			ChallengeConfirmationID: challengeConfirmationId,
			UserID:                  userId,
			ImageURL:                img,
		}
		tx := cd.DB.Create(&uploadParticipateProof)
		if tx.Error != nil {
			tx.Rollback()
			return tx.Error
		}
	}
	tx.Commit()
	return nil
}
func (cd *ChallengeData) IsUserParticipate(userId string, challengeId string) (error, bool) {
	var count int64
	tx := cd.DB.Model(&ChallengeLog{}).Where("user_id = ? AND challenge_id = ?", userId, challengeId).Count(&count)
	if tx.Error != nil {
		return tx.Error, false
	}
	if count > 0 {
		return nil, true
	}
	return nil, false
}

func (cd *ChallengeData) DeleteImageProof(challengeConfirmationId string) error {
	tx := cd.DB.Delete(&ChallengeConfirmationImage{}, "challenge_confirmation_id = ?", challengeConfirmationId)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}
