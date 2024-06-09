package database

import (
	DataAdmin "backendgreeve/features/admin/data"
	Cart "backendgreeve/features/cart/data"
	Challenge "backendgreeve/features/challenges/data"
	Forums "backendgreeve/features/forums/data"
	ImpactCategoryData "backendgreeve/features/impactcategory/data"
	Product "backendgreeve/features/product/data"
	Transaction "backendgreeve/features/transaction/data"
	DataUsers "backendgreeve/features/users/data"
	Voucher "backendgreeve/features/voucher/data"
	WebhookData "backendgreeve/features/webhook/data"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	db.AutoMigrate(&DataUsers.User{})
	db.AutoMigrate(&DataUsers.VerifyOTP{})
	db.AutoMigrate(&WebhookData.PaymentNotification{})
	db.AutoMigrate(&DataAdmin.Admin{})
	db.AutoMigrate(&Voucher.Voucher{})
	db.AutoMigrate(&ImpactCategoryData.ImpactCategory{})
	db.AutoMigrate(&Forums.Forum{})
	db.AutoMigrate(&Forums.MessageForum{})
	db.AutoMigrate(&Product.Product{})
	db.AutoMigrate(&Product.ProductImage{})
	db.AutoMigrate(&Product.ProductImpactCategory{})
	db.AutoMigrate(&Product.ProductLog{})
	db.AutoMigrate(&Challenge.Challenge{})
	db.AutoMigrate(&Challenge.ChallengeLog{})
	db.AutoMigrate(&Challenge.ChallengeImpactCategory{})
	db.AutoMigrate(&Challenge.ChallengeConfirmation{})
	db.AutoMigrate(&Challenge.ChallengeConfirmationImage{})
	db.AutoMigrate(&Cart.Cart{})
	db.AutoMigrate(&Transaction.Transaction{})
	db.AutoMigrate(&Transaction.TransactionItem{})
	return nil
}
