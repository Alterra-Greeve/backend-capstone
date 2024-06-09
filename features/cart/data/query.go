package data

import (
	"backendgreeve/constant"
	"backendgreeve/features/cart"
	"backendgreeve/features/impactcategory"
	"backendgreeve/features/product"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CartData struct {
	DB *gorm.DB
}

func New(db *gorm.DB) cart.CartDataInterface {
	return &CartData{
		DB: db,
	}
}

func (c *CartData) Create(cart cart.NewCart) error {
	newCart := Cart{
		ID:        uuid.New().String(),
		UserID:    cart.UserID,
		ProductID: cart.ProductID,
		Quantity:  1,
	}
	err := c.DB.Create(&newCart).Error
	if err != nil {
		return constant.ErrCreateCart
	}
	return nil
}

func (c *CartData) Update(cart cart.UpdateCart) error {
	err := c.DB.Model(&Cart{}).Where("user_id = ?", cart.UserID).Updates(cart).Error
	if err != nil {
		return constant.ErrUpdateCart
	}
	return nil
}

func (c *CartData) Delete(userId string, productId string) error {
	rowAffected := c.DB.Where("user_id = ? AND product_id = ?", userId, productId).Delete(&Cart{}).RowsAffected
	if rowAffected == 0 {
		return constant.ErrCartNotFound
	} else if rowAffected > 0 {
		c.DB.Model(&Cart{}).Where("user_id = ? AND product_id = ?", userId, productId).Update("quantity", 0)
		return nil
	}
	return nil

}

func (c *CartData) Get(userId string) (cart.Cart, error) {
	var carts []Cart
	var cartData cart.Cart

	// Retrieve cart data from the database
	err := c.DB.Model(&Cart{}).
		Preload("Product").
		Preload("User").
		Preload("Product.Images").
		Preload("Product.ImpactCategories").
		Preload("Product.ImpactCategories.ImpactCategory").
		Where("user_id = ?", userId).Find(&carts).Error
	if err != nil {
		return cartData, err
	}

	if len(carts) == 0 {
		// Return empty cart if no data found
		return cartData, nil
	}

	// Initialize cart user
	cartData.User.ID = userId
	cartData.User.Username = carts[0].User.Username
	cartData.User.Email = carts[0].User.Email
	cartData.User.Address = carts[0].User.Address
	cartData.User.Phone = carts[0].User.Phone
	// Populate other user fields as needed

	var cartItems []cart.CartItems
	for _, cartData := range carts {
		var images []product.ProductImage
		var impactCategories []product.ProductImpactCategory

		for _, img := range cartData.Product.Images {
			images = append(images, product.ProductImage{
				ID:        img.ID,
				ProductID: img.ProductID,
				ImageURL:  img.ImageURL,
				Position:  img.Position,
			})
		}

		for _, impact := range cartData.Product.ImpactCategories {
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

		cartItems = append(cartItems, cart.CartItems{
			Quantity: cartData.Quantity,
			Product: product.Product{
				ID:               cartData.Product.ID,
				Name:             cartData.Product.Name,
				Description:      cartData.Product.Description,
				Price:            cartData.Product.Price,
				Coin:             cartData.Product.Coin,
				Stock:            cartData.Product.Stock,
				CreatedAt:        cartData.Product.CreatedAt,
				UpdatedAt:        cartData.Product.UpdatedAt,
				Images:           images,
				ImpactCategories: impactCategories,
			},
		})
	}
	cartData.Items = cartItems
	return cartData, nil
}

func (c *CartData) GetCartDetail(userId string, productId string) (cart.Cart, error) {
	var cart cart.Cart
	err := c.DB.Where("user_id = ? AND product_id = ?", userId, productId).First(&cart).Error
	return cart, err
}

func (c *CartData) GetCartQty(userId string, productId string) (int, error) {
	var cart Cart
	err := c.DB.Where("user_id = ? AND product_id = ?", userId, productId).First(&cart).Error
	return cart.Quantity, err
}

func (c *CartData) IsCartExist(userId string, productId string) (bool, error) {
	var count int64
	tx := c.DB.Model(&Cart{}).Where("user_id = ? AND product_id = ?", userId, productId).Count(&count)
	if tx.Error != nil {
		return false, tx.Error
	}
	return count > 0, nil
}

func (c *CartData) InsertIncrement(userId string, productId string) error {
	return c.DB.Model(&Cart{}).Where("user_id = ? AND product_id = ?", userId, productId).Update("quantity", gorm.Expr("quantity + 1")).Error
}

func (c *CartData) InsertDecrement(userId string, productId string) error {
	return c.DB.Model(&Cart{}).Where("user_id = ? AND product_id = ?", userId, productId).Update("quantity", gorm.Expr("quantity - 1")).Error
}

func (c *CartData) InsertByQuantity(userId string, productId string, quantity int) error {
	return c.DB.Model(&Cart{}).Where("user_id = ? AND product_id = ?", userId, productId).Update("quantity", quantity).Error
}
