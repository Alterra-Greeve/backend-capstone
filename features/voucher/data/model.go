package data

import (
	"time"

	"gorm.io/gorm"
)

type Voucher struct {
	*gorm.Model
	ID          string    `gorm:"primary_key;type:varchar(50);not null;column:id"`
	Name        string    `gorm:"type:varchar(255);not null;column:name"`
	Code        string    `gorm:"type:varchar(255);not null;column:code"`
	Discount    string    `gorm:"type:varchar(255);not null;column:discount"`
	Description string    `gorm:"type:varchar(255);not null;column:description"`
	ExpiredAt   time.Time `gorm:"type:datetime;not null;column:expired_at"`
}

func (voucher *Voucher) TableName() string {
	return "vouchers"
}
