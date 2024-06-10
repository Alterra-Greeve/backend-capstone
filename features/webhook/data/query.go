package data

import (
	transactionEntity "backendgreeve/features/transaction"
	transactionData "backendgreeve/features/transaction/data"
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

func (w *WebhookData) HandleNotification(notification webhook.PaymentNotification, transaction transactionData.Transaction) error {
	transactionUpdate := transactionEntity.UpdateTransaction{
		ID:            transaction.ID,
		Status:        transaction.Status,
		PaymentMethod: transaction.PaymentMethod,
	}

	err := w.DB.Model(&transactionData.Transaction{}).Where("id = ?", transaction.ID).Updates(&transactionUpdate).Error
	if err != nil {
		return err
	}
	return nil
}
