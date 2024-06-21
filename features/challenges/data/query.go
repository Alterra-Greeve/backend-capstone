package data

import (
	"backendgreeve/constant"
	"backendgreeve/features/challenges"
	impactcategory "backendgreeve/features/impactcategory/data"
	user "backendgreeve/features/users/data"
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
	var challengesList []challenges.Challenge

	// Hitung total challenge
	var totalChallenges int64
	countTx := cd.DB.Model(&challenges.Challenge{}).Count(&totalChallenges)
	if countTx.Error != nil {
		return nil, 0, countTx.Error
	}

	challengePerPage := 20
	totalPages := int(math.Ceil(float64(totalChallenges) / float64(challengePerPage)))

	// Ambil challenges dengan pagination tanpa JOIN
	tx := cd.DB.
		Table("challenges").
		Offset((page - 1) * challengePerPage).
		Scan(&challengesList)

	if tx.Error != nil {
		return nil, 0, tx.Error
	}

	// Ambil kategori dampak untuk setiap challenge
	for i := range challengesList {
		var impactCategories []challenges.ChallengeImpactCategory
		rows, err := cd.DB.Raw(`
			SELECT cic.id, cic.challenge_id, cic.impact_category_id, ic.id, ic.name, ic.impact_point, ic.icon_url
			FROM challenge_impact_categories cic
			JOIN impact_categories ic ON cic.impact_category_id = ic.id
			WHERE cic.challenge_id = ? AND cic.deleted_at IS NULL
		`, challengesList[i].ID).Rows()
		if err != nil {
			return nil, 0, err
		}
		defer rows.Close()

		for rows.Next() {
			var cic challenges.ChallengeImpactCategory
			var ic impactcategory.ImpactCategory
			if err := rows.Scan(&cic.ID, &cic.ChallengeID, &cic.ImpactCategoryID, &ic.ID, &ic.Name, &ic.ImpactPoint, &ic.IconURL); err != nil {
				return nil, 0, err
			}
			cic.ImpactCategory = challenges.ImpactCategory{
				ID:          ic.ID,
				Name:        ic.Name,
				ImpactPoint: ic.ImpactPoint,
				IconURL:     ic.IconURL,
			}
			impactCategories = append(impactCategories, cic)
		}
		challengesList[i].ImpactCategories = impactCategories
	}

	return challengesList, totalPages, nil
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
	var challenge challenges.Challenge

	err := cd.DB.Table("challenges").
		Select("challenges.*, cic.id AS cic_id, cic.challenge_id, cic.impact_category_id, ic.id AS ic_id, ic.name, ic.impact_point, ic.icon_url").
		Joins("LEFT JOIN challenge_impact_categories AS cic ON cic.challenge_id = challenges.id AND cic.deleted_at IS NULL").
		Joins("LEFT JOIN impact_categories AS ic ON cic.impact_category_id = ic.id").
		Where("challenges.id = ?", challengeId).
		Scan(&challenge).Error

	if err != nil {
		return challenges.Challenge{}, err
	}

	// Manually map the results
	challenge.ImpactCategories = []challenges.ChallengeImpactCategory{}
	rows, err := cd.DB.Raw(`
		SELECT cic.id, cic.challenge_id, cic.impact_category_id, ic.id, ic.name, ic.impact_point, ic.icon_url
		FROM challenge_impact_categories cic
		JOIN impact_categories ic ON cic.impact_category_id = ic.id
		WHERE cic.challenge_id = ? AND cic.deleted_at IS NULL
	`, challengeId).Rows()
	if err != nil {
		return challenges.Challenge{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var cic ChallengeImpactCategory
		var ic impactcategory.ImpactCategory
		if err := rows.Scan(&cic.ID, &cic.ChallengeID, &cic.ImpactCategoryID, &ic.ID, &ic.Name, &ic.ImpactPoint, &ic.IconURL); err != nil {
			return challenges.Challenge{}, err
		}
		cic.ImpactCategory = ic
		challenge.ImpactCategories = append(challenge.ImpactCategories, challenges.ChallengeImpactCategory{
			ID: cic.ID,
			ImpactCategory: challenges.ImpactCategory{
				ID:          ic.ID,
				Name:        ic.Name,
				ImpactPoint: ic.ImpactPoint,
				IconURL:     ic.IconURL,
			},
		})
	}
	log.Println(challenge)
	if challenge.ID == "" {
		return challenges.Challenge{}, constant.ErrChallengeNotFound
	}
	return challenge, nil
}

func (cd *ChallengeData) GetUserParticipate(userId string, status string) ([]challenges.ChallengeConfirmation, error) {
	var challengeConfirmations []ChallengeConfirmation

	tx := cd.DB.Model(&ChallengeConfirmation{}).
		Preload("Challenge").
		Preload("ChallengeConfirmationImages").
		Preload("Challenge.ChallengeImpactCategories").
		Preload("Challenge.ChallengeImpactCategories.ImpactCategory").
		Where("user_id = ? AND status = ?", userId, status).
		Find(&challengeConfirmations)
	if tx.Error != nil {
		return nil, tx.Error
	}

	challengeData := new(ChallengeConfirmation).ConvertToArrayEntity(challengeConfirmations)

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
	tx := cd.DB.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	err := cd.DB.Model(&ChallengeImpactCategory{}).Where("challenge_id = ?", challenge.ID).Delete(&ChallengeImpactCategory{}).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	err = cd.DB.Model(&Challenge{}).Where("id = ?", challenge.ID).Updates(&challenge).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	var challengeImpactCategories []ChallengeImpactCategory
	for _, category := range challenge.ImpactCategories {
		challengeImpactCategories = append(challengeImpactCategories, ChallengeImpactCategory{
			ID:               uuid.New().String(),
			ChallengeID:      challenge.ID,
			ImpactCategoryID: category.ImpactCategoryID,
		})
	}
	err = cd.DB.Create(&challengeImpactCategories).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
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

func (cd *ChallengeData) UploadParticipateProof(challengeConfirmationId string, images []string) error {
	tx := cd.DB.Begin()
	if tx.Error != nil {
		return tx.Error
	}
	err := cd.DeleteImageProof(challengeConfirmationId)
	if err != nil {
		tx.Rollback()
		return err
	}
	err = cd.DB.Model(&ChallengeConfirmation{}).Where("id = ?", challengeConfirmationId).Update("date_upload", time.Now()).Update("status", "Diterima").Error
	if err != nil {
		tx.Rollback()
		return err
	}

	for _, img := range images {
		uploadParticipateProof := ChallengeConfirmationImage{
			ID:                      uuid.New().String(),
			ChallengeConfirmationID: challengeConfirmationId,
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

func (cd *ChallengeData) GetChallengeForUserByID(challengeConfirmationId string) (challenges.ChallengeConfirmation, error) {
	var challengeConfirmations ChallengeConfirmation
	tx := cd.DB.Model(&ChallengeConfirmation{}).
		Preload("User").
		Preload("Challenge").
		Preload("Challenge.ChallengeImpactCategories").
		Preload("Challenge.ChallengeImpactCategories.ImpactCategory").
		Preload("ChallengeConfirmationImages").
		Where("id = ?", challengeConfirmationId).
		Find(&challengeConfirmations)
	if tx.Error != nil {
		return challenges.ChallengeConfirmation{}, tx.Error
	}
	var challengeData challenges.ChallengeConfirmation
	challengeData.ID = challengeConfirmations.ID
	challengeData.UserID = challengeConfirmations.UserID
	challengeData.Status = challengeConfirmations.Status
	challengeData.CreatedAt = challengeConfirmations.CreatedAt
	challengeData.UpdatedAt = challengeConfirmations.UpdatedAt
	challengeData.Challenge = challenges.Challenge{
		ID:          challengeConfirmations.Challenge.ID,
		Title:       challengeConfirmations.Challenge.Title,
		Description: challengeConfirmations.Challenge.Description,
		Exp:         challengeConfirmations.Challenge.Exp,
		Coin:        challengeConfirmations.Challenge.Coin,
		DateStart:   challengeConfirmations.Challenge.DateStart,
		DateEnd:     challengeConfirmations.Challenge.DateEnd,
		Difficulty:  challengeConfirmations.Challenge.Difficulty,
		ImageURL:    challengeConfirmations.Challenge.ImageURL,
	}
	for _, category := range challengeConfirmations.Challenge.ChallengeImpactCategories {
		challengeData.Challenge.ImpactCategories = append(challengeData.Challenge.ImpactCategories, challenges.ChallengeImpactCategory{
			ID: category.ID,
			ImpactCategory: challenges.ImpactCategory{
				ID:          category.ImpactCategory.ID,
				Name:        category.ImpactCategory.Name,
				ImpactPoint: category.ImpactCategory.ImpactPoint,
				IconURL:     category.ImpactCategory.IconURL,
			},
		})
	}
	for _, image := range challengeConfirmations.ChallengeConfirmationImages {
		challengeData.ChallengeConfirmationImages = append(challengeData.ChallengeConfirmationImages, challenges.ChallengeConfirmationImage{
			ID:       image.ID,
			ImageURL: image.ImageURL,
		})
	}

	return challengeData, nil
}

func (cd *ChallengeData) EditChallengeForUserByID(challengeConfirmationId string, images []string) error {
	tx := cd.DB.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	err := cd.DeleteImageProof(challengeConfirmationId)
	if err != nil {
		tx.Rollback()
		return err
	}

	err = cd.DB.Model(&ChallengeConfirmation{}).Where("id = ?", challengeConfirmationId).Update("date_upload", time.Now()).Update("status", "Diterima").Error
	if err != nil {
		tx.Rollback()
		return err
	}
	err = cd.InsertCoinAndExpUser(challengeConfirmationId)
	if err != nil {
		tx.Rollback()
		return err
	}
	for _, img := range images {
		uploadParticipateProof := ChallengeConfirmationImage{
			ID:                      uuid.New().String(),
			ChallengeConfirmationID: challengeConfirmationId,
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

func (cd *ChallengeData) InsertCoinAndExpUser(challengeConfirmationId string) error {
	var challengeConfirmation ChallengeConfirmation
	tx := cd.DB.Begin()
	if tx.Error != nil {
		return tx.Error
	}
	err := cd.DB.Model(&ChallengeConfirmation{}).Where("id = ?", challengeConfirmationId).Find(&challengeConfirmation).Error
	if err != nil {
		return err
	}

	var challengeData Challenge
	err = cd.DB.Model(&Challenge{}).Where("id = ?", challengeConfirmation.ChallengeID).Find(&challengeData).Error
	if err != nil {
		return err
	}

	var userData user.User
	err = cd.DB.Model(&user.User{}).Where("id = ?", challengeConfirmation.UserID).Find(&userData).Error
	if err != nil {
		return err
	}

	var userDataUpdate user.User
	userDataUpdate.Coin = userData.Coin + challengeData.Coin
	userDataUpdate.Exp = userData.Exp + challengeData.Exp

	err = cd.DB.Model(&user.User{}).Where("id = ?", challengeConfirmation.UserID).Updates(&userDataUpdate).Error
	if err != nil {
		return err
	}
	tx.Commit()
	return nil
}
