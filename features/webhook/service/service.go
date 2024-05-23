package service

import "backendgreeve/features/webhook"

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
	if transactionStatus == "capture" {
		if fraudStatus == "accept" {
			// TODO set transaction status on your database to 'success'
			// and response with 200 OK
		}
	} else if transactionStatus == "settlement" {
		// TODO set transaction status on your database to 'success'
		// and response with 200 OK
	} else if transactionStatus == "cancel" || transactionStatus == "deny" || transactionStatus == "expire" {
		// TODO set transaction status on your database to 'failure'
		// and response with 200 OK
	} else if transactionStatus == "pending" {
		// TODO set transaction status on your database to 'pending' / waiting payment
		// and response with 200 OK
	}
	return s.d.HandleNotification(notification)
}
