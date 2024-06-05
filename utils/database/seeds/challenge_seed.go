package seeds

import (
	"backendgreeve/features/challenges/data"

	"gorm.io/gorm"
)

func CreateChallenge(db *gorm.DB, challenge data.Challenge) error {
	return db.Where("id = ?", challenge.ID).FirstOrCreate(&challenge).Error
}
