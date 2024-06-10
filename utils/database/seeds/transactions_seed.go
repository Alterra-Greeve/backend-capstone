package seeds

import (
	transaction "backendgreeve/features/transaction/data"

	"gorm.io/gorm"
)

func CreateTransaction(db *gorm.DB, transaction transaction.Transaction) error {
	return db.Where("id = ?", transaction.ID).FirstOrCreate(&transaction).Error
}
