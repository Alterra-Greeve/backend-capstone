package database

import (
	DataUsers "backendgreeve/features/users/data"
	WebhookData "backendgreeve/features/webhook/data"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	db.AutoMigrate(&DataUsers.User{})
	db.AutoMigrate(&DataUsers.VerifyOTP{})
	db.AutoMigrate(&WebhookData.PaymentNotification{})
	return nil
}
