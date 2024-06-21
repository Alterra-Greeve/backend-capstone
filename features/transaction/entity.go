package transaction

import (
	"backendgreeve/features/product"
	"backendgreeve/features/users"
	"time"

	"github.com/labstack/echo/v4"
)

type Transaction struct {
	ID            string
	UserID        string
	VoucherID     string
	Address       string
	Status        string
	Total         float64
	Coin          int
	PaymentMethod string
	SnapURL       string
}

type CreateTransaction struct {
	UserID      string
	UsingCoin   bool
	VoucherCode string
}

type UpdateTransaction struct {
	ID            string
	Status        string
	PaymentMethod string
}

type TransactionItems struct {
	ID            string
	TransactionID string
	ProductID     string
	Qty           int
	Product       product.Product
}

type TransactionData struct {
	ID               string
	VoucherID        string
	Status           string
	Total            float64
	Coin             int
	SnapURL          string
	User             users.User
	TransactionItems []TransactionItems
	CreatedAt        time.Time
	UpdatedAt        time.Time
}

type TransactionHandlerInterface interface {
	GetUserTransaction() echo.HandlerFunc
	CreateTransaction() echo.HandlerFunc
	UpdateTransaction() echo.HandlerFunc
	DeleteTransaction() echo.HandlerFunc

	//
	GetAllTransaction() echo.HandlerFunc
}

type TransactionServiceInterface interface {
	GetUserTransaction(userId string) ([]TransactionData, error)
	GetTransactionByID(transactionId string) (TransactionData, error)
	CreateTransaction(transaction CreateTransaction) (Transaction, error)
	UpdateTransaction(transaction UpdateTransaction) error
	DeleteTransaction(transactionId string) error

	//
	GetAllTransaction() ([]TransactionData, error)
}

type TransactionDataInterface interface {
	GetUserTransaction(userId string) ([]TransactionData, error)
	GetTransactionByID(transactionId string) (TransactionData, error)
	CreateTransactions(transaction Transaction) error
	UpdateTransaction(transaction Transaction) error
	DeleteTransaction(transactionId string) error
	GetUserAddress(userId string) (string, error)
	GetUserCoin(userId string) (int, error)
	DecreaseUserCoin(userId string, coin int, total float64) (float64, int, error)
	GetVoucherID(voucherCode string) (string, error)
	AddTransactionItems(userId string, transactionId string) error
	GetTotalPrice(userId string) (float64, error)
	GetTotalPriceWithDiscount(total float64, voucherId string) (float64, error)
	GetTotalPriceWithCoin(userId string) (float64, error)
	GetTotalPriceWithDiscountAndCoin(userId string, voucherId string) (float64, error)

	//
	GetAllTransaction() ([]TransactionData, error)
}
