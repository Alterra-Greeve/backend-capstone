package database

import (
	DataAdmin "backendgreeve/features/admin/data"
	Forums "backendgreeve/features/forums/data"
	ImpactCategoryData "backendgreeve/features/impactcategory/data"
	DataUsers "backendgreeve/features/users/data"
	WebhookData "backendgreeve/features/webhook/data"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	db.AutoMigrate(&DataUsers.User{})
	db.AutoMigrate(&DataUsers.VerifyOTP{})
	db.AutoMigrate(&WebhookData.PaymentNotification{})
	db.AutoMigrate(&DataAdmin.Admin{})
	db.AutoMigrate(&ImpactCategoryData.ImpactCategory{})
	db.AutoMigrate(&Forums.Forum{})
	db.AutoMigrate(&Forums.MessageForum{})
	return nil
}
