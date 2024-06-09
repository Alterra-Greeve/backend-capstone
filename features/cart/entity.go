package cart

import (
	"backendgreeve/features/product"
	"backendgreeve/features/users"

	"github.com/labstack/echo/v4"
)

type Cart struct {
	User  users.User
	Items []CartItems
}

type NewCart struct {
	UserID    string
	ProductID string
}

type CartItems struct {
	Quantity int
	Product  product.Product
}

type UpdateCart struct {
	UserID    string
	ProductID string
	Quantity  int
	Type      string
}

type CartHandlerInterface interface {
	Create() echo.HandlerFunc
	Update() echo.HandlerFunc
	Delete() echo.HandlerFunc
	Get() echo.HandlerFunc
}

type CartServiceInterface interface {
	Create(cart NewCart) error
	Update(cart UpdateCart) error
	Delete(userId string, productId string) error
	Get(userId string) (Cart, error)
}

type CartDataInterface interface {
	Create(cart NewCart) error
	Update(cart UpdateCart) error
	Delete(userId string, productId string) error
	Get(userId string) (Cart, error)
	GetCartDetail(userId string, productId string) (Cart, error)
	IsCartExist(userId string, productId string) (bool, error)
	GetCartQty(userId string, productId string) (int, error)
	InsertIncrement(userId string, productId string) error
	InsertDecrement(userId string, productId string) error
	InsertByQuantity(userId string, productId string, quantity int) error
}
