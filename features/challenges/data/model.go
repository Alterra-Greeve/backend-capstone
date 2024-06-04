package data

import (
	impactcategory "backendgreeve/features/impactcategory/data"
	"time"

	"gorm.io/gorm"
)

type Challenge struct {
	*gorm.Model
	ID                        string                    `gorm:"primary_key;type:varchar(50);not null;column:id"`
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
	ID               string                        `gorm:"primary_key;type:varchar(50);not null;column:id"`
	ChallengeID      string                        `gorm:"type:varchar(50);not null;column:challenge_id"`
	ImpactCategoryID string                        `gorm:"type:varchar(50);not null;column:impact_category_id"`
	Challenge        Challenge                     `gorm:"foreignKey:ChallengeID;references:ID"`
	ImpactCategory   impactcategory.ImpactCategory `gorm:"foreignKey:ImpactCategoryID;references:ID"`
}

type ChallengeLog struct {
	*gorm.Model
	ID          string `gorm:"primary_key;type:varchar(50);not null;column:id"`
	UserID      string `gorm:"type:varchar(50);not null;column:user_id"`
	ChallengeID string `gorm:"type:varchar(50);not null;column:challenge_id"`
	Status      string `gorm:"type:varchar(50);not null;column:status"`
}

func (*Challenge) TableName() string {
	return "challenges"
}
func (*ChallengeImpactCategory) TableName() string {
	return "challenge_impact_categories"
}
func (*ChallengeLog) TableName() string {
	return "challenge_logs"
}
