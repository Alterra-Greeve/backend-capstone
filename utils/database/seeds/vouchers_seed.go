package seeds

import (
	voucher "backendgreeve/features/Voucher/data"
	"time"

	"gorm.io/gorm"
)

func CreateVoucher(db *gorm.DB, id, name, code, discount, description string, expiryDate time.Time) error {
	voucherRecord := voucher.Voucher{
		ID:          id,
		Name:        name,
		Code:        code,
		Discount:    discount,
		Description: description,
		ExpiredAt:   expiryDate,
	}

	return db.Where("id = ?", id).FirstOrCreate(&voucherRecord).Error
}
