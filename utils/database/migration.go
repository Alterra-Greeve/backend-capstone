package database

import (
	DataAdmin "backendgreeve/features/admin/data"
	DataUsers "backendgreeve/features/users/data"
	WebhookData "backendgreeve/features/webhook/data"
	ImpactCategoryData "backendgreeve/features/impactcategory/data"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	db.AutoMigrate(&DataUsers.User{})
	db.AutoMigrate(&DataUsers.VerifyOTP{})
	db.AutoMigrate(&WebhookData.PaymentNotification{})
	db.AutoMigrate(&DataAdmin.Admin{})
	db.AutoMigrate(&ImpactCategoryData.ImpactCategory{})
	return nil
}
