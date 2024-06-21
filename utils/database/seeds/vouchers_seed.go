package seeds

import (
	vouchers "backendgreeve/features/voucher/data"

	"gorm.io/gorm"
)

func CreateVoucher(db *gorm.DB, voucher vouchers.Voucher) error {
	return db.Where("id = ?", voucher.ID).FirstOrCreate(&voucher).Error
}
