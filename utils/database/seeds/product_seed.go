package seeds

import (
	"backendgreeve/features/product/data"
	"gorm.io/gorm"
)


func CreateProduct(db *gorm.DB, product data.Product) error {
	return db.Where("id = ?", product.ID).FirstOrCreate(&product).Error
}
