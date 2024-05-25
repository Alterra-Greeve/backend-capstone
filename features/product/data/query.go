package data

import (
	"backendgreeve/features/product"

	"gorm.io/gorm"
)

type ProductData struct {
	*gorm.DB
}

func New(db *gorm.DB) product.ProductDataInterface {
	return &ProductData{
		DB: db,
	}
}

func (pd *ProductData) Get() ([]product.Product, error) {
	var products []product.Product
	tx := pd.DB.Preload("Images").Preload("ImpactCategories").Find(&products)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return products, nil
}

func (pd *ProductData) GetById(id string) (product.Product, error) {
	var products product.Product
	tx := pd.DB.Preload("Images").Preload("ImpactCategories").Find(&products, "id = ?", id)
	if tx.Error != nil {
		return products, tx.Error
	}
	return products, nil
}

func (pd *ProductData) GetByCategoryID(categoryID string) ([]product.Product, error) {
	var products []product.Product
	tx := pd.DB.Find(&products, "impact_category_id LIKE ?", "%"+categoryID+"%")
	if tx.Error != nil {
		return nil, tx.Error
	}
	return products, nil
}

func (pd *ProductData) GetByName(name string) ([]product.Product, error) {
	var products []product.Product
	tx := pd.DB.Find(&products, "name LIKE ?", "%"+name+"%")
	if tx.Error != nil {
		return nil, tx.Error
	}
	return products, nil
}

func (pd *ProductData) Create(product product.Product) error {
	tx := pd.DB.Create(product)
	return tx.Error
}

func (pd *ProductData) Update(product product.Product) error {
	tx := pd.DB.Model(product).Where("id = ?", product.ID).Updates(product)
	return tx.Error
}

func (pd *ProductData) Delete(product product.Product) error {
	tx := pd.DB.Delete(product)
	return tx.Error
}
