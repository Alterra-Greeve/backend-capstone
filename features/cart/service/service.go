package service

import (
	"backendgreeve/constant"
	"backendgreeve/features/cart"
)

type CartService struct {
	d cart.CartDataInterface
}

func New(d cart.CartDataInterface) cart.CartServiceInterface {
	return &CartService{
		d: d,
	}
}

func (cs *CartService) Create(cart cart.NewCart) error {
	isCartExist, err := cs.d.IsCartExist(cart.UserID, cart.ProductID)
	if err != nil {
		return err
	}

	if isCartExist {
		return cs.d.InsertIncrement(cart.UserID, cart.ProductID)
	}

	return cs.d.Create(cart)
}

func (cs *CartService) Update(cart cart.UpdateCart) error {
	if cart.Type != "increment" && cart.Type != "decrement" && cart.Type != "qty" {
		return constant.ErrFieldType
	}
	// if cart.Type == "increment" || cart.Type == "decrement" && cart.Quantity != 0 {
	// 	return constant.ErrFieldChoiceOneType
	// }

	cartQty, err := cs.d.GetCartQty(cart.UserID, cart.ProductID)
	if err != nil {
		return err
	}

	if cartQty == 1 && cart.Type == "decrement" {
		return cs.d.Delete(cart.UserID, cart.ProductID)
	}

	if cart.Type == "increment" {
		return cs.d.InsertIncrement(cart.UserID, cart.ProductID)
	} else if cart.Type == "decrement" {
		return cs.d.InsertDecrement(cart.UserID, cart.ProductID)
	} else if cart.Type == "qty" {
		return cs.d.InsertByQuantity(cart.UserID, cart.ProductID, cart.Quantity)
	}
	return constant.ErrFieldType
}

func (cs *CartService) Delete(userId string, productId string) error {
	return cs.d.Delete(userId, productId)
}

func (cs *CartService) Get(userId string) (cart.Cart, error) {
	return cs.d.Get(userId)
}
