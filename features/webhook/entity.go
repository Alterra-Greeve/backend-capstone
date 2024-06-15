package webhook

import (
	"github.com/labstack/echo/v4"
	transaction "backendgreeve/features/transaction/data"
)

type PaymentNotification struct {
	TransactionTime   string `json:"transaction_time"`
	TransactionStatus string `json:"transaction_status"`
	TransactionID     string `json:"transaction_id"`
	SignatureKey      string `json:"signature_key"`
	PaymentType       string `json:"payment_type"`
	OrderID           string `json:"order_id"`
	MerchantID        string `json:"merchant_id"`
	GrossAmount       string `json:"gross_amount"`
	FraudStatus       string `json:"fraud_status"`
	Currency          string `json:"currency"`
	SettlementTime    string `json:"settlement_time,omitempty"`
}

type MidtransNotificationHandler interface {
	HandleNotification() echo.HandlerFunc
}

type MidtransNotificationService interface {
	HandleNotification(notification PaymentNotification) error
}

type MidtransNotificationData interface {
	HandleNotification(notification PaymentNotification, transaction transaction.Transaction) error
	InsertUserCoin(transactionID string) error
}
