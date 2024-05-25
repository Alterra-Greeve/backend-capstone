package database

import (
	DataAdmin "backendgreeve/features/admin/data"
	ImpactCategoryData "backendgreeve/features/impactcategory/data"
	Product "backendgreeve/features/product/data"
	DataUsers "backendgreeve/features/users/data"
	WebhookData "backendgreeve/features/webhook/data"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	db.AutoMigrate(&DataUsers.User{})
	db.AutoMigrate(&DataUsers.VerifyOTP{})
	db.AutoMigrate(&WebhookData.PaymentNotification{})
	db.AutoMigrate(&DataAdmin.Admin{})
	db.AutoMigrate(&ImpactCategoryData.ImpactCategory{})
	db.AutoMigrate(&Product.Product{})
	db.AutoMigrate(&Product.ProductImage{})
	db.AutoMigrate(&Product.ProductImpactCategory{})
	return nil
}
