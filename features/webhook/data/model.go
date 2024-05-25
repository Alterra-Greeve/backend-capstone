package data

import (
	"time"

	"gorm.io/gorm"
)

type PaymentNotification struct {
	*gorm.Model
	OrderID           string          `gorm:"primary_key;type:varchar(255);not null;column:order_id"`
	TransactionTime   string          `gorm:"column:transaction_time"`
	TransactionStatus string          `gorm:"column:transaction_status"`
	TransactionID     string          `gorm:"column:transaction_id"`
	SignatureKey      string          `gorm:"column:signature_key"`
	PaymentType       string          `gorm:"column:payment_type"`
	MerchantID        string          `gorm:"column:merchant_id"`
	GrossAmount       string          `gorm:"column:gross_amount"`
	FraudStatus       string          `gorm:"column:fraud_status"`
	Currency          string          `gorm:"column:currency"`
	SettlementTime    string          `gorm:"column:settlement_time"`
	CreatedAt         time.Time       `gorm:"column:created_at"`
	UpdatedAt         time.Time       `gorm:"column:updated_at"`
	DeletedAt         *gorm.DeletedAt `gorm:"column:deleted_at"`
}

func (PaymentNotification) TableName() string {
	return "payment_notifications"
}
