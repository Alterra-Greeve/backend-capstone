package seeds

import (
	"backendgreeve/features/product/data"

	"gorm.io/gorm"
)

func CreateProduct(db *gorm.DB, product data.Product) error {
	product.Stock = 500
	return db.Where("id = ?", product.ID).FirstOrCreate(&product).Error
}
