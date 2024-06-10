package data

import (
	product "backendgreeve/features/product/data"
	user "backendgreeve/features/users/data"

	"gorm.io/gorm"
)

type Transaction struct {
	*gorm.Model
	ID               string            `gorm:"primary_key;type:varchar(50);not null;column:id"`
	VoucherID        string            `gorm:"type:varchar(50);column:voucher_id"`
	UserID           string            `gorm:"type:varchar(50);not null;column:user_id"`
	Address          string            `gorm:"type:varchar(255);not null;column:address"`
	Total            float64           `gorm:"type:decimal(10,2);not null;column:total"`
	Status           string            `gorm:"type:varchar(50);not null;column:status"`
	PaymentMethod    string            `gorm:"type:varchar(50);column:payment_method"`
	SnapURL          string            `gorm:"type:varchar(255);not null;column:snap_url"`
	Coin             int               `gorm:"type:int;column:coin"`
	User             user.User         `gorm:"foreignKey:UserID;references:ID"`
	TransactionItems []TransactionItem `gorm:"foreignKey:TransactionID;references:ID"`
}

type TransactionItem struct {
	*gorm.Model
	ID            string          `gorm:"primary_key;type:varchar(50);not null;column:id"`
	TransactionID string          `gorm:"type:varchar(50);not null;column:transaction_id"`
	ProductID     string          `gorm:"type:varchar(50);not null;column:product_id"`
	Quantity      int             `gorm:"type:int;not null;column:quantity"`
	Transaction   Transaction     `gorm:"foreignKey:TransactionID;references:ID"`
	Product       product.Product `gorm:"foreignKey:ProductID;references:ID"`
}

func (Transaction) TableName() string {
	return "transactions"
}
func (TransactionItem) TableName() string {
	return "transaction_items"
}
