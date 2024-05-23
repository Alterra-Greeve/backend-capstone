package data

import (
	"backendgreeve/features/webhook"

	"gorm.io/gorm"
)

type WebhookData struct {
	DB *gorm.DB
}

func New(db *gorm.DB) webhook.MidtransNotificationData {
	return &WebhookData{
		DB: db,
	}
}

func (w *WebhookData) HandleNotification(notification webhook.PaymentNotification) error {
	if err := w.DB.Create(&notification).Error; err != nil {
		return err
	}
	return nil
}
