package service

import (
	transaction "backendgreeve/features/transaction/data"
	"backendgreeve/features/webhook"
)

type WebhookService struct {
	d webhook.MidtransNotificationData
}

func New(data webhook.MidtransNotificationData) webhook.MidtransNotificationService {
	return &WebhookService{
		d: data,
	}
}

func (s *WebhookService) HandleNotification(notification webhook.PaymentNotification) error {
	// orderId := notification.OrderID
	transactionStatus := notification.TransactionStatus
	fraudStatus := notification.FraudStatus
	transactionData := transaction.Transaction{
		ID:            notification.OrderID,
		PaymentMethod: notification.PaymentType,
	}
	if transactionStatus == "capture" {
		if fraudStatus == "accept" {
			// TODO set transaction status on your database to 'success'
			// and response with 200 OK
			transactionData.Status = "Berhasil"
		}
	} else if transactionStatus == "settlement" {
		// TODO set transaction status on your database to 'success'
		// and response with 200 OK
		transactionData.Status = "Berhasil"
	} else if transactionStatus == "cancel" || transactionStatus == "deny" || transactionStatus == "expire" {
		// TODO set transaction status on your database to 'failure'
		// and response with 200 OK
		transactionData.Status = "Gagal"
	} else if transactionStatus == "pending" {
		// TODO set transaction status on your database to 'pending' / waiting payment
		// and response with 200 OK
		transactionData.Status = "Gagal"
	}
	return s.d.HandleNotification(notification, transactionData)
}
