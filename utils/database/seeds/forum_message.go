package seeds

import (
	forumMessage "backendgreeve/features/forums/data"

	"gorm.io/gorm"
)

func CreateForumMessage(db *gorm.DB, message forumMessage.MessageForum) error {
	return db.Where("id = ?", message.ID).FirstOrCreate(&message).Error
}
