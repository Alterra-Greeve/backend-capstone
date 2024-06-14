package data

import (
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
		Joins("JOIN products ON products.id = transaction_items.product_id").
		Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var ordersProduct orders.OrdersProduct
		if err := rows.Scan(&ordersProduct.UserID, &ordersProduct.Username, &ordersProduct.Email,
			&ordersProduct.ProductID, &ordersProduct.ProductName,
			&ordersProduct.Qty, &ordersProduct.Coin,
			&ordersProduct.CreatedAt, &ordersProduct.UpdatedAt); err != nil {
			return nil, err
		}

		var impactCategories []orders.ImpactCategory
		if err := d.DB.Table("product_impact_categories").
			Select("impact_categories.id, impact_categories.name").
			Joins("JOIN impact_categories ON impact_categories.id = product_impact_categories.impact_category_id").
			Where("product_impact_categories.product_id = ?", ordersProduct.ProductID).
			Scan(&impactCategories).Error; err != nil {
			return nil, err
		}
		ordersProduct.ImpactCategories = impactCategories

		ordersProducts = append(ordersProducts, ordersProduct)
	}

	for i := range ordersProducts {
		impactPointTotal, err := d.GetTotalImpactPoint()
		if err != nil {
			return nil, err
		}
		ordersProducts[i].ImpactPointTotal = impactPointTotal

		totalCoin, err := d.GetTotalCoin()
		if err != nil {
			return nil, err
		}
		ordersProducts[i].TotalCoin = totalCoin
	}

	return ordersProducts, nil
}

func (d *OrdersData) GetTotalImpactPoint() (int, error) {
	var totalImpactPoint int
	tx := d.DB.Table("impact_categories").Select("SUM(impact_point)").Row()
	if tx.Err() != nil {
		return 0, tx.Err()
	}
	tx.Scan(&totalImpactPoint)
	return totalImpactPoint, nil
}

func (d *OrdersData) GetTotalCoin() (int, error) {
	var totalProduct int64
	tx := d.DB.Table("products").Count(&totalProduct)
	if tx.Error != nil {
		return 0, tx.Error
	}
	return int(totalProduct), nil
}

// Challenge
// func (cd *OrdersData) GetOrdersChallenge() ([]orders.OrderChallengeConfirmation, error) {
// 	var challengeConfirmations []orders.OrderChallengeConfirmation
// 	tx := cd.DB.Preload("User").
// 		Preload("Challenge").
// 		Preload("Challenge.ChallengeImpactCategories").
// 		Preload("Challenge.ChallengeImpactCategories.ImpactCategory").
// 		Find(&challengeConfirmations)
// 	if tx.Error != nil {
// 		return nil, tx.Error
// 	}

// 	var challengeData []orders.OrderChallengeConfirmation
// 	for _, cc := range challengeConfirmations {
// 		var impactCategories []orders.ImpactCategory
// 		for _, cic := range cc.Challenge.ChallengeImpactCategories {
// 			impactCategories = append(impactCategories, orders.ImpactCategory{
// 				ID:   cic.ImpactCategory.ID,
// 				Name: cic.ImpactCategory.Name,
// 			})
// 		}

// 		challengeData = append(challengeData, orders.OrderChallengeConfirmation{
// 			ID: cc.ID,
// 			// Username: cc.User.Username,
// 			Status:           cc.Status,
// 			ImpactCategories: impactCategories,
// 			ImpactPointTotal: cc.ImpactPointTotal,
// 			CreatedAt:        cc.CreatedAt,
// 			UpdatedAt:        cc.UpdatedAt,
// 		})
// 	}

// 	return challengeData, nil
// }
