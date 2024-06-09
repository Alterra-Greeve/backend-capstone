package service

import (
	"backendgreeve/config"
	"backendgreeve/constant"
	"backendgreeve/features/transaction"

	"github.com/google/uuid"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
)

type TransactionService struct {
	d      transaction.TransactionDataInterface
	config config.MidtransConfig
}

func New(d transaction.TransactionDataInterface, config config.MidtransConfig) transaction.TransactionServiceInterface {
	return &TransactionService{
		d:      d,
		config: config,
	}
}
func (s *TransactionService) GetUserTransaction(userId string) ([]transaction.Transaction, error) {
	// Implement logic here
	transaction, err := s.d.GetUserTransaction(userId)
	if err != nil {
		return nil, err
	}
	return transaction, nil
}

func (s *TransactionService) GetTransactionByID(transactionId string) (transaction.Transaction, error) {
	// Implement logic here
	transactions, err := s.d.GetTransactionByID(transactionId)
	if err != nil {
		return transaction.Transaction{}, err
	}
	return transactions, nil
}

func (s *TransactionService) CreateTransaction(Transaction transaction.CreateTransaction) error {
	// Implement logic here
	var transactionData transaction.Transaction
	if transactionData.Address == "" {
		return constant.ErrAddressEmpty
	}
	
	transactionData.Address = Transaction.Address
	transactionData.ID = uuid.New().String()
	transactionData.UserID = Transaction.UserID
	transactionData.Status = "pending"
	transactionData.VoucherID = "asd123"
	transactionData.Total = 100000
	req := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  transactionData.ID,
			GrossAmt: int64(transactionData.Total),
		},
	}
	var client snap.Client
	client.New(s.config.ServerKey, midtrans.Sandbox)
	snapResponse, err := client.CreateTransaction(req)
	if err != nil {
		return err
	}
	transactionData.SnapURL = snapResponse.RedirectURL

	error := s.d.CreateTransactions(transactionData)
	if error != nil {
		return error
	}
	return nil
}

func (s *TransactionService) UpdateTransaction(transaction transaction.UpdateTransaction) error {
	// Implement logic here
	return nil
}

func (s *TransactionService) DeleteTransaction(transactionId string) error {
	// Implement logic here
	return nil
}
