package data

import (
	"backendgreeve/features/dashboard"
	"backendgreeve/features/product"
	productData "backendgreeve/features/product/data"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type DashboardData struct {
	DB *gorm.DB
}

func New(db *gorm.DB) dashboard.DashboardDataInterface {
	return &DashboardData{
		DB: db,
	}
}

func (d *DashboardData) GetDashboard() (dashboard.Dashboard, error) {
	totalProduct, err := d.GetTotalProduct()
	if err != nil {
		return dashboard.Dashboard{}, err
	}
	totalNewProductThisMonth, err := d.GetTotalNewProductThisMonth()
	if err != nil {
		return dashboard.Dashboard{}, err
	}
	totalNewProductPercentage, err := d.GetTotalNewProductPercentage()
	if err != nil {
		return dashboard.Dashboard{}, err
	}
	totalUser, err := d.GetTotalUser()
	if err != nil {
		return dashboard.Dashboard{}, err
	}
	newProduct, err := d.GetNewProduct()
	if err != nil {
		return dashboard.Dashboard{}, err
	}
	return dashboard.Dashboard{
		TotalProduct:              totalProduct,
		TotalNewProductThisMonth:  totalNewProductThisMonth,
		TotalNewProductPercentage: totalNewProductPercentage,
		TotalUser:                 totalUser,
		NewProduct:                newProduct,
	}, nil
}

func (d *DashboardData) GetTotalProduct() (int, error) {
	var totalProduct int64
	tx := d.DB.Table("products").Count(&totalProduct)
	if tx.Error != nil {
		return 0, tx.Error
	}
	return int(totalProduct), nil
}

func (d *DashboardData) GetTotalProductPercentage() (string, error) {
	var totalProduct int64
	var totalProductThisMonth int64
	var totalProductPercentage float64

	txAll := d.DB.Table("products").Count(&totalProduct)
	if txAll.Error != nil {
		return "0%", txAll.Error
	}

	txThisMonth := d.DB.Table("products").Where("MONTH(created_at) = MONTH(NOW())").Count(&totalProductThisMonth)
	if txThisMonth.Error != nil {
		return "0%", txThisMonth.Error
	}

	if totalProduct > 0 {
		totalProductPercentage = float64(totalProductThisMonth) / float64(totalProduct) * 100
	} else {
		totalProductPercentage = 100
	}

	return fmt.Sprintf("%.2f%%", totalProductPercentage), nil
}

func (d *DashboardData) GetTotalNewProductThisMonth() (int, error) {
	var totalNewProductThisMonth int64
	tx := d.DB.Table("products").
		Where("MONTH(created_at) = MONTH(NOW())").
		Count(&totalNewProductThisMonth)
	if tx.Error != nil {
		return 0, tx.Error
	}
	return int(totalNewProductThisMonth), nil
}

func (d *DashboardData) GetTotalNewProductPercentage() (string, error) {
	var totalProductLastMonth int64
	var totalProductThisMonth int64
	var totalProductPercentage float64

	txLastMonth := d.DB.Table("products").Where("MONTH(created_at) = MONTH(NOW()) - 1").Count(&totalProductLastMonth)
	if txLastMonth.Error != nil {
		return "0%", txLastMonth.Error
	}

	txThisMonth := d.DB.Table("products").Where("MONTH(created_at) = MONTH(NOW())").Count(&totalProductThisMonth)
	if txThisMonth.Error != nil {
		return "0%", txThisMonth.Error
	}

	if totalProductLastMonth > 0 {
		totalProductPercentage = float64(totalProductThisMonth) / float64(totalProductLastMonth) * 100
	} else {
		totalProductPercentage = 100
	}

	return fmt.Sprintf("%.2f%%", totalProductPercentage), nil
}

func (d *DashboardData) GetNewProduct() ([]product.Product, error) {
	var productDB []productData.Product
	var top3Products []product.Product
	tx := d.DB.Model(&productData.Product{}).
		Preload("Images").
		Preload("ImpactCategories.ImpactCategory").
		Order("created_at DESC").
		Limit(3).
		Find(&productDB)
	if tx.Error != nil {
		return nil, tx.Error
	}
	for _, product := range productDB {
		top3Products = append(top3Products, product.ToEntity())
	}
	return top3Products, nil
}

func (d *DashboardData) GetTotalUser() (int, error) {
	var totalUser int64
	tx := d.DB.Table("users").Count(&totalUser)
	if tx.Error != nil {
		return 0, tx.Error
	}
	return int(totalUser), nil
}

func (d *DashboardData) GetNewUserPercentage() (string, error) {
	// Implement logic here
	var totalUser int64
	var totalUserThisMonth int64
	var totalUserPercentage float64
	txAll := d.DB.Table("users").Count(&totalUser)
	if txAll.Error != nil {
		return "0%", txAll.Error
	}
	txThisMonth := d.DB.Table("users").Where("MONTH(created_at) = MONTH(NOW())").Count(&totalUserThisMonth)
	if txThisMonth.Error != nil {
		return "0%", txThisMonth.Error
	}
	return fmt.Sprintf("%.2f%%", totalUserPercentage), nil
}

func (d *DashboardData) GetMonthlyImpactProduct(month string) ([]dashboard.ImpactPoint, error) {
	var impactPoints []dashboard.ImpactPoint

	monthTime, err := time.Parse("2006-01", month)
	if err != nil {
		return nil, fmt.Errorf("invalid month format: %w", err)
	}

	startDate := monthTime
	endDate := monthTime.AddDate(0, 1, 0)

	err = d.DB.Table("transaction_items").
		Select("impact_categories.name as name, SUM(impact_categories.impact_point * transaction_items.quantity) as total_point").
		Joins("JOIN products ON products.id = transaction_items.product_id").
		Joins("JOIN product_impact_categories ON product_impact_categories.product_id = products.id").
		Joins("JOIN impact_categories ON impact_categories.id = product_impact_categories.impact_category_id").
		Where("transaction_items.created_at BETWEEN ? AND ?", startDate, endDate).
		Group("impact_categories.name").
		Scan(&impactPoints).Error

	if err != nil {
		return nil, err
	}

	return impactPoints, nil
}

func (d *DashboardData) GetMonthlyImpactChallenge(month string) ([]dashboard.ImpactPoint, error) {
	var impactPoints []dashboard.ImpactPoint

	monthTime, err := time.Parse("2006-01", month)
	if err != nil {
		return nil, fmt.Errorf("invalid month format: %w", err)
	}

	startDate := monthTime
	endDate := monthTime.AddDate(0, 1, 0)

	err = d.DB.Table("challenge_confirmations").
		Select("impact_categories.name as name, SUM(impact_categories.impact_point) as total_point").
		Joins("JOIN challenges ON challenges.id = challenge_confirmations.challenge_id").
		Joins("JOIN challenge_impact_categories ON challenge_impact_categories.challenge_id = challenges.id").
		Joins("JOIN impact_categories ON impact_categories.id = challenge_impact_categories.impact_category_id").
		Where("challenge_confirmations.created_at BETWEEN ? AND ?", startDate, endDate).
		Where("challenge_confirmations.status = ?", "Berhasil").
		Group("impact_categories.name").
		Scan(&impactPoints).Error

	if err != nil {
		return nil, err
	}

	return impactPoints, nil
}
