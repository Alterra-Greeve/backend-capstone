package data

import (
	"gorm.io/gorm"
)

type PaymentNotification struct {
	*gorm.Model
	ID                string `gorm:"primary_key;type:varchar(255);not null;column:id"`
	OrderID           string `gorm:"type:varchar(255);not null;column:order_id"`
	TransactionTime   string `gorm:"column:transaction_time"`
	TransactionStatus string `gorm:"column:transaction_status"`
	TransactionID     string `gorm:"column:transaction_id"`
	SignatureKey      string `gorm:"column:signature_key"`
	PaymentType       string `gorm:"column:payment_type"`
	MerchantID        string `gorm:"column:merchant_id"`
	GrossAmount       string `gorm:"column:gross_amount"`
	FraudStatus       string `gorm:"column:fraud_status"`
	Currency          string `gorm:"column:currency"`
	SettlementTime    string `gorm:"column:settlement_time"`
}

func (PaymentNotification) TableName() string {
	return "payment_notifications"
}
