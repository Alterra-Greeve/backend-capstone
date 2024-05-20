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
	return s.d.HandleNotification(notification)
}
