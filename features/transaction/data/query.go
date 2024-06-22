package data

import (
	cart "backendgreeve/features/cart/data"
	"backendgreeve/features/impactcategory"
	"backendgreeve/features/product"
	"backendgreeve/features/transaction"
	"backendgreeve/features/users"
	user "backendgreeve/features/users/data"
	voucher "backendgreeve/features/voucher/data"
	"strconv"
	"strings"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TransactionData struct {
	DB *gorm.DB
}

func New(db *gorm.DB) transaction.TransactionDataInterface {
	return &TransactionData{
		DB: db,
	}
}

func (td *TransactionData) CreateTransactions(transaction transaction.Transaction) error {
	transactionData := Transaction{
		ID:            transaction.ID,
		VoucherID:     transaction.VoucherID,
		UserID:        transaction.UserID,
		Total:         transaction.Total,
		Status:        transaction.Status,
		Address:       transaction.Address,
		Coin:          transaction.Coin,
		PaymentMethod: transaction.PaymentMethod,
		SnapURL:       transaction.SnapURL,
	}

	err := td.DB.Create(&transactionData).Error
	if err != nil {
		return err
	}
	return nil
}

func (td *TransactionData) GetUserTransaction(userId string) ([]transaction.TransactionData, error) {
	var transactions []Transaction

	err := td.DB.Model(&Transaction{}).
		Preload("User").
		Preload("TransactionItems").
		Preload("TransactionItems.Product").
		Preload("TransactionItems.Product.Images").
		Preload("TransactionItems.Product.ImpactCategories").
		Preload("TransactionItems.Product.ImpactCategories.ImpactCategory").
		Where("user_id = ?", userId).
		Find(&transactions).Error

	if err != nil {
		return nil, err
	}

	// Konversi hasil ke TransactionData
	var result []transaction.TransactionData
	for _, txn := range transactions {
		var txnItems []transaction.TransactionItems
		var images []product.ProductImage
		var impactCategories []product.ProductImpactCategory
		for _, item := range txn.TransactionItems {
			for _, img := range item.Product.Images {
				images = append(images, product.ProductImage{
					ID:        img.ID,
					ProductID: img.ProductID,
					ImageURL:  img.ImageURL,
					Position:  img.Position,
				})
			}
			for _, impact := range item.Product.ImpactCategories {
				impactCategories = append(impactCategories, product.ProductImpactCategory{
					ID:               impact.ID,
					ProductID:        impact.ProductID,
					ImpactCategoryID: impact.ImpactCategoryID,
					ImpactCategory: impactcategory.ImpactCategory{
						ID:          impact.ImpactCategory.ID,
						Name:        impact.ImpactCategory.Name,
						ImpactPoint: impact.ImpactCategory.ImpactPoint,
						IconURL:     impact.ImpactCategory.IconURL,
					},
				})
			}
			txnItems = append(txnItems, transaction.TransactionItems{
				ID:            item.ID,
				TransactionID: item.TransactionID,
				ProductID:     item.ProductID,
				Qty:           item.Quantity,
				Product: product.Product{
					ID:               item.Product.ID,
					Name:             item.Product.Name,
					Description:      item.Product.Description,
					Price:            item.Product.Price,
					Coin:             item.Product.Coin,
					Stock:            item.Product.Stock,
					Images:           images,
					ImpactCategories: impactCategories,
				},
			})
		}

		result = append(result, transaction.TransactionData{
			ID:        txn.ID,
			VoucherID: txn.VoucherID,
			Status:    txn.Status,
			Total:     txn.Total,
			Coin:      txn.Coin,
			SnapURL:   txn.SnapURL,
			User: users.User{
				ID:        txn.User.ID,
				Name:      txn.User.Name,
				Email:     txn.User.Email,
				Username:  txn.User.Username,
				Password:  txn.User.Password,
				Address:   txn.User.Address,
				Gender:    txn.User.Gender,
				Phone:     txn.User.Phone,
				Coin:      txn.User.Coin,
				Exp:       txn.User.Exp,
				AvatarURL: txn.User.AvatarURL,
				CreatedAt: txn.User.CreatedAt,
				UpdatedAt: txn.User.UpdatedAt,
			},
			TransactionItems: txnItems,
			CreatedAt:        txn.CreatedAt,
			UpdatedAt:        txn.UpdatedAt,
		})
	}

	return result, nil
}

func (td *TransactionData) GetTopTransactionItem(transactionId string) (transaction.TransactionItems, error) {
	// Implement logic here
	return transaction.TransactionItems{}, nil
}

func (td *TransactionData) GetTransactionByID(transactionId string) (transaction.TransactionData, error) {
	// Implement logic here
	return transaction.TransactionData{}, nil
}

func (td *TransactionData) UpdateTransaction(transaction transaction.Transaction) error {
	// Implement logic here
	return nil
}

func (td *TransactionData) DeleteTransaction(transactionId string) error {
	// Implement logic here
	return nil
}

func (td *TransactionData) GetVoucherID(voucherCode string) (string, error) {
	// Implement logic here
	var voucher voucher.Voucher
	err := td.DB.Table("vouchers").Where("code = ?", voucherCode).First(&voucher).Error
	if err != nil {
		return "", err
	}
	return voucher.ID, nil
}

func (td *TransactionData) AddTransactionItems(userId string, transactionId string) error {
	var carts []cart.Cart
	err := td.DB.Where("user_id = ?", userId).Find(&carts).Error
	if err != nil {
		return err
	}
	tx := td.DB.Begin()
	for _, cartItem := range carts {
		transactionItemId := uuid.New().String()

		transactionItem := TransactionItem{
			ID:            transactionItemId,
			TransactionID: transactionId,
			ProductID:     cartItem.ProductID,
			Quantity:      cartItem.Quantity,
		}

		err := td.DB.Create(&transactionItem).Error
		if err != nil {
			tx.Rollback()
			return err
		}

		err = td.DB.Delete(&cartItem).Error
		if err != nil {
			tx.Rollback()
			return err
		}
		tx.Commit()
	}

	return nil
}

func (td *TransactionData) GetUserAddress(userId string) (string, error) {
	var user user.User
	err := td.DB.Where("id = ?", userId).First(&user).Error
	if err != nil {
		return "", err
	}
	return user.Address, nil
}

func (td *TransactionData) GetUserCoin(userId string) (int, error) {
	var user user.User
	err := td.DB.Where("id = ?", userId).First(&user).Error
	if err != nil {
		return 0, err
	}
	return user.Coin, nil
}

func (td *TransactionData) DecreaseUserCoin(userId string, coin int, total float64) (float64, int, error) {
	var user user.User
	err := td.DB.Where("id = ?", userId).First(&user).Error
	if err != nil {
		return total, 0, err
	}

	maxCoin := int(total * 0.80)

	usedCoin := coin
	if usedCoin > maxCoin {
		usedCoin = maxCoin
	}

	if user.Coin < usedCoin {
		usedCoin = user.Coin
	}

	newTotal := total - float64(usedCoin)

	user.Coin -= usedCoin
	err = td.DB.Save(&user).Error
	if err != nil {
		return total, 0, err
	}

	return newTotal, usedCoin, nil
}
func (td *TransactionData) GetTotalPrice(userId string) (float64, error) {
	var carts []cart.Cart
	err := td.DB.Where("user_id = ?", userId).Find(&carts).Error
	if err != nil {
		return 0.0, err
	}

	var totalPrice float64
	for _, cart := range carts {
		var product product.Product
		err := td.DB.Where("id = ?", cart.ProductID).First(&product).Error
		if err != nil {
			return 0.0, err
		}
		totalPrice += float64(cart.Quantity) * product.Price
	}

	return totalPrice, nil
}

func (td *TransactionData) GetTotalPriceWithDiscount(total float64, voucherId string) (float64, error) {
	var voucher voucher.Voucher
	err := td.DB.Where("id = ?", voucherId).First(&voucher).Error
	if err != nil {
		return 0.0, err
	}

	discount := voucher.Discount

	var discountedTotal float64

	if strings.Contains(discount, "%") {
		percentValue, err := strconv.ParseFloat(strings.TrimSuffix(discount, "%"), 64)
		if err != nil {
			return 0.0, err
		}
		discountedTotal = total - (total * percentValue / 100)
	} else {
		discountValue, err := strconv.ParseFloat(discount, 64)
		if err != nil {
			return 0.0, err
		}
		discountedTotal = total - discountValue
	}

	if discountedTotal < 0 {
		discountedTotal = 0
	}

	return discountedTotal, nil
}

func (td *TransactionData) GetTotalPriceWithCoin(userId string) (float64, error) {
	// Implement logic here
	return 0.0, nil
}

func (td *TransactionData) GetTotalPriceWithDiscountAndCoin(userId string, voucherId string) (float64, error) {
	// Implement logic here
	return 0.0, nil
}

func (td *TransactionData) GetAllTransaction() ([]transaction.TransactionData, error) {
	var transactions []Transaction

	err := td.DB.Model(&Transaction{}).Preload("User").
		Preload("TransactionItems").
		Preload("TransactionItems.Product").
		Preload("TransactionItems.Product.Images").
		Preload("TransactionItems.Product.ImpactCategories").
		Preload("TransactionItems.Product.ImpactCategories.ImpactCategory").
		Order("created_at DESC").
		Find(&transactions).Error

	if err != nil {
		return nil, err
	}

	var result []transaction.TransactionData
	for _, txn := range transactions {
		var txnItems []transaction.TransactionItems
		var images []product.ProductImage
		var impactCategories []product.ProductImpactCategory
		for _, item := range txn.TransactionItems {
			for _, img := range item.Product.Images {
				images = append(images, product.ProductImage{
					ID:        img.ID,
					ProductID: img.ProductID,
					ImageURL:  img.ImageURL,
					Position:  img.Position,
				})
			}
			for _, impact := range item.Product.ImpactCategories {
				impactCategories = append(impactCategories, product.ProductImpactCategory{
					ID:               impact.ID,
					ProductID:        impact.ProductID,
					ImpactCategoryID: impact.ImpactCategoryID,
					ImpactCategory: impactcategory.ImpactCategory{
						ID:          impact.ImpactCategory.ID,
						Name:        impact.ImpactCategory.Name,
						ImpactPoint: impact.ImpactCategory.ImpactPoint,
						IconURL:     impact.ImpactCategory.IconURL,
					},
				})
			}
			txnItems = append(txnItems, transaction.TransactionItems{
				ID:            item.ID,
				TransactionID: item.TransactionID,
				ProductID:     item.ProductID,
				Qty:           item.Quantity,
				Product: product.Product{
					ID:               item.Product.ID,
					Name:             item.Product.Name,
					Description:      item.Product.Description,
					Price:            item.Product.Price,
					Coin:             item.Product.Coin,
					Stock:            item.Product.Stock,
					Images:           images,
					ImpactCategories: impactCategories,
				},
			})
		}

		result = append(result, transaction.TransactionData{
			ID:        txn.ID,
			VoucherID: txn.VoucherID,
			Status:    txn.Status,
			Total:     txn.Total,
			Coin:      txn.Coin,
			SnapURL:   txn.SnapURL,
			User: users.User{
				ID:        txn.User.ID,
				Name:      txn.User.Name,
				Email:     txn.User.Email,
				Username:  txn.User.Username,
				Password:  txn.User.Password,
				Address:   txn.User.Address,
				Gender:    txn.User.Gender,
				Phone:     txn.User.Phone,
				Coin:      txn.User.Coin,
				Exp:       txn.User.Exp,
				AvatarURL: txn.User.AvatarURL,
				CreatedAt: txn.User.CreatedAt,
				UpdatedAt: txn.User.UpdatedAt,
			},
			TransactionItems: txnItems,
			CreatedAt:        txn.CreatedAt,
			UpdatedAt:        txn.UpdatedAt,
		})
	}

	return result, nil
}
