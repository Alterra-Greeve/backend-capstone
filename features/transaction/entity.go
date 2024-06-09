package transaction

import (
	"backendgreeve/features/cart"

	"github.com/labstack/echo/v4"
)

type Transaction struct {
	ID            string
	UserID        string
	VoucherID     string
	Address       string
	Description   string
	Status        string
	Total         float64
	Coin          int
	PaymentMethod string
	SnapURL       string
}

type CreateTransaction struct {
	UserID      string
	VoucherCode string
	Address     string
	Description string
	SnapURL     string
}

type UpdateTransaction struct {
	ID     string
	Status string
}

type TransactionItems struct {
	ID            string
	TransactionID string
	ProductID     string
	Qty           int
}

type TransactionHandlerInterface interface {
	GetUserTransaction() echo.HandlerFunc
	CreateTransaction() echo.HandlerFunc
	UpdateTransaction() echo.HandlerFunc
	DeleteTransaction() echo.HandlerFunc
}

type TransactionServiceInterface interface {
	GetUserTransaction(userId string) ([]Transaction, error)
	GetTransactionByID(transactionId string) (Transaction, error)
	CreateTransaction(transaction CreateTransaction) error
	UpdateTransaction(transaction UpdateTransaction) error
	DeleteTransaction(transactionId string) error
}

type TransactionDataInterface interface {
	GetUserTransaction(userId string) ([]Transaction, error)
	GetTransactionByID(transactionId string) (Transaction, error)
	CreateTransactions(transaction Transaction) error
	UpdateTransaction(transaction Transaction) error
	DeleteTransaction(transactionId string) error
	GetVoucherID(voucherCode string) (string, error)
	GetShoppingCartItems(userId string) ([]cart.Cart, error)
	AddTransactionItems(transactionItems []TransactionItems) error
	GetTotalPrice(userId string) (float64, error)
	GetTotalPriceWithDiscount(userId string, voucherId string) (float64, error)
	GetTotalPriceWithCoin(userId string) (float64, error)
	GetTotalPriceWithDiscountAndCoin(userId string, voucherId string) (float64, error)
}
