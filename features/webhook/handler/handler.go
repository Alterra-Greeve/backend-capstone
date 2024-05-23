package handler

import (
	"backendgreeve/features/webhook"

	"github.com/labstack/echo/v4"
)

type WebhookRequest struct {
	s webhook.MidtransNotificationService
}

func New(s webhook.MidtransNotificationService) webhook.MidtransNotificationHandler {
	return &WebhookRequest{
		s: s,
	}
}

func (h *WebhookRequest) HandleNotification() echo.HandlerFunc {
	return func(c echo.Context) error {
		var notification webhook.PaymentNotification
		err := c.Bind(&notification)
		if err != nil {
			return echo.NewHTTPError(400, err.Error())
		}
		err = h.s.HandleNotification(notification)
		if err != nil {
			return echo.NewHTTPError(500, err.Error())
		}
		return c.JSON(200, map[string]string{
			"message": "success",
		})
	}
}
