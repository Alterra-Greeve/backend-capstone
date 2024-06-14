package data

import (
	"backendgreeve/features/challenges"
	impactcategory "backendgreeve/features/impactcategory/data"
	user "backendgreeve/features/users/data"
	"database/sql"
	"time"

	"gorm.io/gorm"
)

type Challenge struct {
	*gorm.Model
	ID                        string                    `gorm:"primaryKey;type:varchar(50);not null;column:id"`
	Title                     string                    `gorm:"type:varchar(255);not null;column:title"`
	Difficulty                string                    `gorm:"type:varchar(255);not null;column:difficulty"`
	Description               string                    `gorm:"type:varchar(255);not null;column:description"`
	Exp                       int                       `gorm:"type:int;not null;column:exp"`
	Coin                      int                       `gorm:"type:int;not null;column:coin"`
	ImageURL                  string                    `gorm:"type:varchar(255);not null;column:image_url"`
	DateStart                 time.Time                 `gorm:"type:datetime;not null;column:date_start"`
	DateEnd                   time.Time                 `gorm:"type:datetime;not null;column:date_end"`
	ChallengeImpactCategories []ChallengeImpactCategory `gorm:"foreignKey:ChallengeID"`
}

type ChallengeImpactCategory struct {
	*gorm.Model
	ID               string                        `gorm:"primaryKey;type:varchar(50);not null;column:id"`
	ChallengeID      string                        `gorm:"type:varchar(50);not null;column:challenge_id"`
	ImpactCategoryID string                        `gorm:"type:varchar(50);not null;column:impact_category_id"`
	Challenge        Challenge                     `gorm:"foreignKey:ChallengeID;references:ID"`
	ImpactCategory   impactcategory.ImpactCategory `gorm:"foreignKey:ImpactCategoryID;references:ID"`
}

type ChallengeConfirmation struct {
	*gorm.Model
	ID                          string                       `gorm:"primaryKey;type:varchar(50);not null;column:id"`
	ChallengeID                 string                       `gorm:"type:varchar(50);not null;column:challenge_id"`
	UserID                      string                       `gorm:"type:varchar(50);not null;column:user_id"`
	Description                 string                       `gorm:"type:varchar(255);column:description"`
	DateUpload                  sql.NullTime                 `gorm:"type:datetime;column:date_upload"`
	Status                      string                       `gorm:"type:varchar(50);not null;column:status"`
	Challenge                   Challenge                    `gorm:"foreignKey:ChallengeID;references:ID;association_foreignkey:ID"`
	User                        user.User                    `gorm:"foreignKey:UserID;references:ID"`
	ChallengeConfirmationImages []ChallengeConfirmationImage `gorm:"foreignKey:ChallengeConfirmationID"`
	CreatedAt                   time.Time                    `gorm:"column:created_at"`
	UpdatedAt                   time.Time                    `gorm:"column:updated_at"`
}

type ChallengeConfirmationImage struct {
	*gorm.Model
	ID                      string                `gorm:"primaryKey;type:varchar(50);not null;column:id"`
	ImageURL                string                `gorm:"type:varchar(255);not null;column:image_url"`
	ChallengeConfirmationID string                `gorm:"type:varchar(50);not null;column:challenge_confirmation_id"`
	ChallengeConfirmation   ChallengeConfirmation `gorm:"foreignKey:ChallengeConfirmationID;references:ID"`
}

type ChallengeLog struct {
	*gorm.Model
	ID          string    `gorm:"primaryKey;type:varchar(50);not null;column:id"`
	UserID      string    `gorm:"type:varchar(50);not null;column:user_id"`
	ChallengeID string    `gorm:"type:varchar(50);not null;column:challenge_id"`
	Status      string    `gorm:"type:varchar(50);not null;column:status"`
	User        user.User `gorm:"foreignKey:UserID;references:ID"`
	Challenge   Challenge `gorm:"foreignKey:ChallengeID;references:ID"`
}

func (*Challenge) TableName() string {
	return "challenges"
}

func (*ChallengeImpactCategory) TableName() string {
	return "challenge_impact_categories"
}

func (*ChallengeConfirmation) TableName() string {
	return "challenge_confirmations"
}

func (*ChallengeConfirmationImage) TableName() string {
	return "challenge_confirmation_image"
}

func (*ChallengeLog) TableName() string {
	return "challenge_logs"
}

func (cd *ChallengeConfirmation) ConvertToArrayEntity(challengeConfirmations []ChallengeConfirmation) []challenges.ChallengeConfirmation {
	var challengeData []challenges.ChallengeConfirmation
	for _, confirmation := range challengeConfirmations {
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
		var images []challenges.ChallengeConfirmationImage
		for _, image := range confirmation.ChallengeConfirmationImages {
			images = append(images, challenges.ChallengeConfirmationImage{
				ID:       image.ID,
				ImageURL: image.ImageURL,
			})
		}
		challengeData = append(challengeData, challenges.ChallengeConfirmation{
			ID:                          confirmation.ID,
			UserID:                      confirmation.UserID,
			Status:                      confirmation.Status,
			ChallengeConfirmationImages: images,
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
	return challengeData
}
