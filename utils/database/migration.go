package database

import (
	DataAdmin "backendgreeve/features/admin/data"
	DataUsers "backendgreeve/features/users/data"
	WebhookData "backendgreeve/features/webhook/data"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	db.AutoMigrate(&DataUsers.User{})
	db.AutoMigrate(&DataUsers.VerifyOTP{})
	db.AutoMigrate(&WebhookData.PaymentNotification{})
	db.AutoMigrate(&DataAdmin.Admin{})
	return nil
}
