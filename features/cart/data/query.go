package data

import (
	"backendgreeve/constant"
	"backendgreeve/features/cart"

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
		return cart.Cart{}, nil
	}

	cartData.User.ID = userId
	cartData.User.Username = carts[0].User.Username
	cartData.User.Email = carts[0].User.Email
	cartData.User.Address = carts[0].User.Address
	cartData.User.Phone = carts[0].User.Phone

	for _, cartModel := range carts {
		cartEntity := cartModel.CastDataToCartEntity()
		cartData.Items = append(cartData.Items, cartEntity.Items...)
	}

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
