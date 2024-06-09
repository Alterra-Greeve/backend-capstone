package data

import (
	"backendgreeve/features/cart"
	"backendgreeve/features/transaction"

	"gorm.io/gorm"
)

type TrasactionData struct {
	DB *gorm.DB
}

func New(db *gorm.DB) transaction.TransactionDataInterface {
	return &TrasactionData{
		DB: db,
	}
}

func (td *TrasactionData) CreateTransactions(transaction transaction.Transaction) error {
	transactionData := Transaction{
		ID:            transaction.ID,
		VoucherID:     transaction.VoucherID,
		UserID:        transaction.UserID,
		Total:         transaction.Total,
		Status:        transaction.Status,
		Address:       transaction.Address,
		Description:   transaction.Description,
		Coin:          transaction.Coin,
		PaymentMethod: transaction.PaymentMethod,
		SnapURL:       transaction.SnapURL,
	}

	err := td.DB.Create(&transactionData).Error
	if err != nil {
		return err
	}
	return nil
}

func (td *TrasactionData) GetUserTransaction(userId string) ([]transaction.Transaction, error) {
	// Implement logic here
	return nil, nil
}

func (td *TrasactionData) GetTransactionByID(transactionId string) (transaction.Transaction, error) {
	// Implement logic here
	return transaction.Transaction{}, nil
}

func (td *TrasactionData) UpdateTransaction(transaction transaction.Transaction) error {
	// Implement logic here
	return nil
}

func (td *TrasactionData) DeleteTransaction(transactionId string) error {
	// Implement logic here
	return nil
}

func (td *TrasactionData) GetVoucherID(voucherCode string) (string, error) {
	// Implement logic here
	return "", nil
}

func (td *TrasactionData) GetShoppingCartItems(userId string) ([]cart.Cart, error) {
	// Implement logic here
	return nil, nil
}

func (td *TrasactionData) AddTransactionItems(transactionItems []transaction.TransactionItems) error {
	// Implement logic here
	return nil
}

func (td *TrasactionData) GetTotalPrice(userId string) (float64, error) {
	// Implement logic here
	return 0.0, nil
}

func (td *TrasactionData) GetTotalPriceWithDiscount(userId string, voucherId string) (float64, error) {
	// Implement logic here
	return 0.0, nil
}

func (td *TrasactionData) GetTotalPriceWithCoin(userId string) (float64, error) {
	// Implement logic here
	return 0.0, nil
}

func (td *TrasactionData) GetTotalPriceWithDiscountAndCoin(userId string, voucherId string) (float64, error) {
	// Implement logic here
	return 0.0, nil
}
