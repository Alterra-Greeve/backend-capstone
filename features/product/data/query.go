package data

import (
	"backendgreeve/constant"
	"backendgreeve/features/product"
	"log"

	"github.com/google/uuid"
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
	tx := pd.DB.Preload("Images").Preload("ImpactCategories.ImpactCategory").Find(&products).Where("deleted_at IS NULL")
	if tx.Error != nil {
		return nil, tx.Error
	}
	return products, nil
}

func (pd *ProductData) GetByPage(page int) ([]product.Product, int, error) {
	var products []product.Product

	var totalProducts int64
	countTx := pd.DB.Model(&Product{}).Count(&totalProducts)
	if countTx.Error != nil {
		return nil, 0, constant.ErrProductEmpty
	}

	productsPerPage := 20
	totalPages := int((totalProducts + int64(productsPerPage) - 1) / int64(productsPerPage))

	tx := pd.DB.Model(&Product{}).Preload("Images").Preload("ImpactCategories.ImpactCategory").
		Offset((page - 1) * productsPerPage).Limit(productsPerPage).Find(&products)
	if tx.Error != nil {
		return nil, 0, constant.ErrGetProduct
	}
	if tx.RowsAffected == 0 {
		return nil, 0, constant.ErrProductNotFound
	}
	return products, totalPages, nil
}

func (pd *ProductData) GetByPageAdmin(page int) ([]product.Product, int, error) {
	var products []product.Product

	var totalProducts int64
	countTx := pd.DB.Model(&Product{}).Count(&totalProducts)
	if countTx.Error != nil {
		return nil, 0, constant.ErrProductEmpty
	}

	productsPerPage := 20
	totalPages := int((totalProducts + int64(productsPerPage) - 1) / int64(productsPerPage))

	tx := pd.DB.Model(&Product{}).Preload("Images").Preload("ImpactCategories.ImpactCategory").
		Offset((page - 1) * productsPerPage).Find(&products)
	if tx.Error != nil {
		return nil, 0, constant.ErrGetProduct
	}
	if tx.RowsAffected == 0 {
		return nil, 0, constant.ErrProductNotFound
	}
	return products, totalPages, nil
}

func (pd *ProductData) GetById(id string) (product.Product, error) {
	var products product.Product

	tx := pd.DB.Model(&Product{}).Preload("Images").Preload("ImpactCategories.ImpactCategory").
		Find(&products, "id = ?", id)
	if tx.Error != nil {
		return products, constant.ErrGetProduct
	}
	return products, nil
}

func (pd *ProductData) GetByIdUser(id string, userId string) (product.Product, error) {
	var products product.Product

	tx := pd.DB.Model(&Product{}).Preload("Images").Preload("ImpactCategories.ImpactCategory").
		Find(&products, "id = ?", id)
	if tx.Error != nil {
		return products, constant.ErrGetProduct
	}
	err := pd.InsertToProductLog(userId, id)
	if err != nil {
		return products, err
	}
	return products, nil
}

func (pd *ProductData) InsertToProductLog(userId string, productId string) error {
	productLogs := ProductLog{
		ID:        uuid.New().String(),
		UserID:    userId,
		ProductID: productId,
	}
	tx := pd.DB.Create(&productLogs)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}
func (pd *ProductData) GetByCategory(categoryName string, page int) ([]product.Product, int, error) {
	var products []product.Product
	var totalProducts int64

	countTx := pd.DB.Model(&Product{}).
		Joins("JOIN product_impact_categories ON product_impact_categories.product_id = products.id").
		Joins("JOIN impact_categories ON impact_categories.id = product_impact_categories.impact_category_id").
		Where("impact_categories.name = ?", categoryName).
		Count(&totalProducts)
	if countTx.Error != nil {
		return nil, 0, constant.ErrProductEmpty
	}

	productsPerPage := 20
	totalPages := int((totalProducts + int64(productsPerPage) - 1) / int64(productsPerPage))

	tx := pd.DB.Model(&Product{}).
		Joins("JOIN product_impact_categories ON product_impact_categories.product_id = products.id").
		Joins("JOIN impact_categories ON impact_categories.id = product_impact_categories.impact_category_id").
		Where("impact_categories.name = ?", categoryName).
		Preload("Images").
		Preload("ImpactCategories.ImpactCategory").
		Offset((page - 1) * productsPerPage).Limit(productsPerPage).
		Find(&products)
	if tx.Error != nil {
		return nil, 0, constant.ErrGetProduct
	}
	if len(products) == 0 {
		return nil, 0, constant.ErrProductNotFound
	}
	return products, totalPages, nil
}

func (pd *ProductData) GetByName(name string, page int) ([]product.Product, int, error) {
	var products []product.Product
	var totalProducts int64

	countTx := pd.DB.Model(&product.Product{}).Where("name LIKE ?", "%"+name+"%").Count(&totalProducts)
	if countTx.Error != nil {
		return nil, 0, constant.ErrProductEmpty
	}

	productsPerPage := 20
	totalPages := int((totalProducts + int64(productsPerPage) - 1) / int64(productsPerPage))

	tx := pd.DB.Preload("Images").Preload("ImpactCategories.ImpactCategory").Where("name LIKE ?", "%"+name+"%").Offset((page - 1) * productsPerPage).Limit(productsPerPage).Find(&products)
	if tx.Error != nil {
		return nil, 0, constant.ErrGetProduct
	}
	if len(products) == 0 {
		return nil, 0, constant.ErrProductNotFound
	}
	return products, totalPages, nil
}

func (pd *ProductData) Create(product product.Product) error {
	productQuery := Product{
		ID:          product.ID,
		Name:        product.Name,
		Coin:        product.Coin,
		Price:       product.Price,
		Description: product.Description,
	}
	for _, image := range product.Images {
		productQuery.Images = append(productQuery.Images, ProductImage{
			ID:        image.ID,
			ProductID: productQuery.ID,
			ImageURL:  image.ImageURL,
			Position:  image.Position,
		})
	}

	for _, impactCategory := range product.ImpactCategories {
		productQuery.ImpactCategories = append(productQuery.ImpactCategories, ProductImpactCategory{
			ID:               impactCategory.ID,
			ProductID:        productQuery.ID,
			ImpactCategoryID: impactCategory.ImpactCategoryID,
		})
	}
	tx := pd.DB.Create(&productQuery)
	if tx.Error != nil {
		return constant.ErrCreateProduct
	}
	return nil
}

func (pd *ProductData) Update(products product.Product) error {
	productQuery := Product{
		ID:          products.ID,
		Name:        products.Name,
		Coin:        products.Coin,
		Price:       products.Price,
		Description: products.Description,
		Stock:       products.Stock,
	}

	for _, image := range products.Images {
		productQuery.Images = append(productQuery.Images, ProductImage{
			ID:        image.ID,
			ProductID: productQuery.ID,
			ImageURL:  image.ImageURL,
			Position:  image.Position,
		})
	}

	for _, impactCategory := range products.ImpactCategories {
		productQuery.ImpactCategories = append(productQuery.ImpactCategories, ProductImpactCategory{
			ID:               impactCategory.ID,
			ProductID:        productQuery.ID,
			ImpactCategoryID: impactCategory.ImpactCategoryID,
		})
	}
	tx := pd.DB.Begin()
	err := pd.DB.Where("product_id = ?", productQuery.ID).Delete(product.ProductImage{})
	if err.Error != nil {

		return tx.Error
	}
	err = pd.DB.Where("product_id = ?", productQuery.ID).Delete(product.ProductImpactCategory{})
	if err.Error != nil {
		return tx.Error
	}

	err = pd.DB.Model(&productQuery).Where("id = ?", productQuery.ID).Save(&productQuery)
	if err.Error != nil {
		return constant.ErrUpdateProduct
	}
	return tx.Commit().Error
}

func (pd *ProductData) Delete(productId string) error {
	tx := pd.DB.Begin()

	if err := tx.Where("product_id = ?", productId).Delete(&ProductImage{}); err.Error != nil {
		tx.Rollback()
		return constant.ErrProductNotFound
	} else if err.RowsAffected == 0 {
		tx.Rollback()
		return constant.ErrProductNotFound
	}

	if err := tx.Where("product_id = ?", productId).Delete(&ProductImpactCategory{}); err.Error != nil {
		tx.Rollback()
		return constant.ErrProductNotFound
	} else if err.RowsAffected == 0 {
		tx.Rollback()
		return constant.ErrProductNotFound
	}

	if err := tx.Where("id = ?", productId).Delete(&Product{}); err.Error != nil {
		tx.Rollback()
		return constant.ErrProductNotFound
	} else if err.RowsAffected == 0 {
		tx.Rollback()
		return constant.ErrProductNotFound
	}

	return tx.Commit().Error
}

func (pd *ProductData) GetRecommendation(categoryId string) ([]product.Product, error) {
	var products []product.Product

	tx := pd.DB.Model(&Product{}).
		Preload("Images").
		Preload("ImpactCategories.ImpactCategory").
		Order("RAND()").
		Joins("JOIN product_impact_categories ON products.id = product_impact_categories.product_id").
		Where("product_impact_categories.impact_category_id = ?", categoryId).
		Limit(10).Find(&products)
	if tx.Error != nil {
		log.Println(tx.Error)
		return nil, constant.ErrGetProduct
	}
	log.Println(products)
	return products, nil
}

func (pd *ProductData) GetImpactCategoryFromTransactionItems(userId string) (string, error) {
	var impactCategoryId string
	var totalQuantity int64

	row := pd.DB.Table("transaction_items").
		Select("product_impact_categories.impact_category_id, SUM(transaction_items.quantity) as total_quantity").
		Joins("JOIN products ON transaction_items.product_id = products.id").
		Joins("JOIN product_impact_categories ON products.id = product_impact_categories.product_id").
		Where("transaction_items.transaction_id IN (SELECT id FROM transactions WHERE user_id = ?)", userId).
		Group("product_impact_categories.impact_category_id").
		Order("total_quantity DESC").
		Limit(1).
		Row()

	err := row.Scan(&impactCategoryId, &totalQuantity)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return "", nil
		}
		return "", constant.ErrGetProduct
	}

	return impactCategoryId, nil
}

func (pd *ProductData) GetImpactCategoryFromProductLog(userId string) (string, error) {
	var impactCategoryId string
	var visits int64

	row := pd.DB.Table("product_logs").
		Select("product_impact_categories.impact_category_id, COUNT(*) as visits").
		Joins("JOIN products ON product_logs.product_id = products.id").
		Joins("JOIN product_impact_categories ON products.id = product_impact_categories.product_id").
		Where("user_id = ?", userId).
		Group("product_impact_categories.impact_category_id").
		Order("visits DESC").
		Limit(1).
		Row()

	err := row.Scan(&impactCategoryId, &visits)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return "", nil
		}
		return "", constant.ErrGetProduct
	}

	return impactCategoryId, nil
}

func (pd *ProductData) GetRandomRecommendation() ([]product.Product, error) {
	var products []product.Product

	tx := pd.DB.Model(&Product{}).
		Preload("Images").
		Preload("ImpactCategories.ImpactCategory").
		Order("RAND()").
		Limit(10).
		Find(&products)
	if tx.Error != nil {
		return nil, constant.ErrGetProduct
	}
	return products, nil
}
