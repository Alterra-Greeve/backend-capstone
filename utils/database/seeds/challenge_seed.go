package seeds

import (
	"backendgreeve/features/challenges/data"

	"gorm.io/gorm"
)

func CreateChallenge(db *gorm.DB, challenge data.Challenge) error {
	return db.Where("id = ?", challenge.ID).FirstOrCreate(&challenge).Error
}

func CreateChallengeConfirmation(db *gorm.DB, challengeConfirmation data.ChallengeConfirmation) error {
	return db.Where("id = ?", challengeConfirmation.ID).FirstOrCreate(&challengeConfirmation).Error
}

func CreateChallengeLog(db *gorm.DB, challengeLog data.ChallengeLog) error {
	return db.Where("id = ?", challengeLog.ID).FirstOrCreate(&challengeLog).Error
}
