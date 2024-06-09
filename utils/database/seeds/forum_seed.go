package seeds

import (
	forum "backendgreeve/features/forums/data"

	"gorm.io/gorm"
)

func CreateForum(db *gorm.DB, forums forum.Forum) error {
	return db.Where("id = ?", forums.ID).FirstOrCreate(&forums).Error
}
