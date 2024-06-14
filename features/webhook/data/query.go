package data

import (
	productData "backendgreeve/features/product/data"
	transactionEntity "backendgreeve/features/transaction"
	transactionData "backendgreeve/features/transaction/data"
	userData "backendgreeve/features/users/data"
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
	tx := w.DB.Begin()

	err := w.DB.Model(&transactionData.Transaction{}).Where("id = ?", transaction.ID).Updates(&transactionUpdate).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	err = w.InsertUserCoin(transaction.ID)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func (w *WebhookData) InsertUserCoin(transactionId string) error {
	var transaction transactionData.Transaction
	err := w.DB.Where("id = ?", transactionId).First(&transaction).Error
	if err != nil {
		return err
	}
	var user userData.User
	err = w.DB.Where("id = ?", transaction.UserID).First(&user).Error
	if err != nil {
		return err
	}
	var transactionItem []transactionData.TransactionItem
	err = w.DB.Where("transaction_id = ?", transaction.ID).Find(&transactionItem).Error
	if err != nil {
		return err
	}
	var product []productData.Product
	var totalCoinxQty int
	for i, item := range transactionItem {
		err = w.DB.Where("id = ?", item.ProductID).First(&product).Error
		totalCoinxQty += product[i].Coin * item.Quantity
		if err != nil {
			return err
		}
	}
	userUpdate := userData.User{
		Coin: user.Coin + totalCoinxQty,
	}

	err = w.DB.Model(&userData.User{}).Where("id = ?", user.ID).Updates(&userUpdate).Error
	if err != nil {
		return err
	}
	return nil
}
