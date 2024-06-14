package data

import (
	"backendgreeve/features/challenges/data"
	"backendgreeve/features/orders"

	"gorm.io/gorm"
)

type OrdersData struct {
	DB *gorm.DB
}

func New(db *gorm.DB) orders.OrdersDataInterface {
	return &OrdersData{
		DB: db,
	}
}

func (d *OrdersData) GetOrdersProduct() ([]orders.OrdersProduct, error) {
	var ordersProducts []orders.OrdersProduct

	rows, err := d.DB.Table("transactions").
		Select(`transactions.user_id, users.username, users.email,
			products.id AS product_id, products.name AS product_name, products.coin AS product_coin,
			transaction_items.quantity, 
			transactions.created_at, transactions.updated_at`).
		Joins("JOIN users ON users.id = transactions.user_id").
		Joins("JOIN transaction_items ON transaction_items.transaction_id = transactions.id").
		Joins("JOIN products ON products.id = transaction_items.product_id").Order("transactions.created_at DESC").
		Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var ordersProduct orders.OrdersProduct
		if err := rows.Scan(&ordersProduct.UserID, &ordersProduct.Username, &ordersProduct.Email,
			&ordersProduct.ProductID, &ordersProduct.ProductName, &ordersProduct.Coin,
			&ordersProduct.Qty,
			&ordersProduct.CreatedAt, &ordersProduct.UpdatedAt); err != nil {
			return nil, err
		}

		var impactCategories []orders.ImpactCategory
		if err := d.DB.Table("product_impact_categories").
			Select("impact_categories.id, impact_categories.name, impact_categories.impact_point").
			Joins("JOIN impact_categories ON impact_categories.id = product_impact_categories.impact_category_id").
			Where("product_impact_categories.product_id = ?", ordersProduct.ProductID).
			Scan(&impactCategories).Error; err != nil {
			return nil, err
		}

		var impactPointTotal int
		for _, category := range impactCategories {
			impactPointTotal += category.ImpactPoint
		}

		ordersProduct.ImpactCategories = impactCategories
		ordersProduct.ImpactPointTotal = impactPointTotal
		totalCoin, err := d.GetTotalCoin()
		if err != nil {
			return nil, err
		}
		ordersProduct.TotalCoin = totalCoin
		ordersProducts = append(ordersProducts, ordersProduct)
	}

	return ordersProducts, nil
}

func (d *OrdersData) GetTotalCoin() (int, error) {
	var totalCoin int
	tx := d.DB.Table("transactions").
		Select("SUM(products.coin * transaction_items.quantity)").
		Joins("JOIN transaction_items ON transaction_items.transaction_id = transactions.id").
		Joins("JOIN products ON products.id = transaction_items.product_id").
		Row()
	if tx.Err() != nil {
		return 0, tx.Err()
	}
	tx.Scan(&totalCoin)
	return totalCoin, nil
}

// Challenge
func (cd *OrdersData) GetOrdersChallenge() ([]orders.ChallengeConfirmation, error) {
	var challengeConfirmations []data.ChallengeConfirmation
	tx := cd.DB.Preload("User").
		Preload("Challenge").
		Preload("Challenge.ChallengeImpactCategories").
		Preload("Challenge.ChallengeImpactCategories.ImpactCategory").
		Find(&challengeConfirmations)
	if tx.Error != nil {
		return nil, tx.Error
	}

	var challengeData []orders.ChallengeConfirmation
	for _, cc := range challengeConfirmations {
		var impactCategories []orders.ImpactCategory
		var impactPointTotal int
		for _, cic := range cc.Challenge.ChallengeImpactCategories {
			impactCategories = append(impactCategories, orders.ImpactCategory{
				ID:   cic.ImpactCategory.ID,
				Name: cic.ImpactCategory.Name,
			})
			impactPointTotal += cic.ImpactCategory.ImpactPoint
		}

		challengeData = append(challengeData, orders.ChallengeConfirmation{
			ID:               cc.ID,
			Username:         cc.User.Username,
			Status:           cc.Status,
			ImpactCategories: impactCategories,
			ImpactPointTotal: impactPointTotal,
			Challenge: orders.Challenge{
				ID:                        cc.Challenge.ID,
				Title:                     cc.Challenge.Title,
				Description:               cc.Challenge.Description,
				Exp:                       cc.Challenge.Exp,
				Coin:                      cc.Challenge.Coin,
				DateStart:                 cc.Challenge.DateStart,
				DateEnd:                   cc.Challenge.DateEnd,
				Difficulty:                cc.Challenge.Difficulty,
				ImageURL:                  cc.Challenge.ImageURL,
				ChallengeImpactCategories: []orders.ChallengeImpactCategory{},
			},
			CreatedAt: cc.CreatedAt,
			UpdatedAt: cc.UpdatedAt,
		})
	}

	return challengeData, nil
}
