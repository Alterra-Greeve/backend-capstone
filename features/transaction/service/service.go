package service

import (
	"backendgreeve/config"
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
func (s *TransactionService) GetUserTransaction(userId string) ([]transaction.TransactionData, error) {
	// Implement logic here
	transaction, err := s.d.GetUserTransaction(userId)
	if err != nil {
		return nil, err
	}

	return transaction, nil
}

func (s *TransactionService) GetTransactionByID(transactionId string) (transaction.TransactionData, error) {
	// Implement logic here
	transactions, err := s.d.GetTransactionByID(transactionId)
	if err != nil {
		return transaction.TransactionData{}, err
	}
	return transactions, nil
}

func (s *TransactionService) CreateTransaction(Transaction transaction.CreateTransaction) (transaction.Transaction, error) {
	var transactionData transaction.Transaction

	transactionData.ID = uuid.New().String()
	transactionData.UserID = Transaction.UserID
	transactionData.Status = "Pending"
	address, err := s.d.GetUserAddress(Transaction.UserID)
	if err != nil {
		return transaction.Transaction{}, err
	}
	// if address == "" {
	// 	return transaction.Transaction{}, constant.ErrAddressEmpty
	// }

	transactionData.Address = address

	total, err := s.d.GetTotalPrice(Transaction.UserID)
	if err != nil {
		return transaction.Transaction{}, err
	}

	transactionData.Total = total
	if Transaction.VoucherCode != "" {
		voucher, err := s.d.GetVoucherID(Transaction.VoucherCode)
		if err != nil {
			return transaction.Transaction{}, err
		}
		transactionData.VoucherID = voucher
	} else {
		transactionData.VoucherID = ""
	}

	req := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  transactionData.ID,
			GrossAmt: int64(transactionData.Total),
		},
	}

	if Transaction.UsingCoin {
		coin, err := s.d.GetUserCoin(Transaction.UserID)
		if err != nil {
			return transaction.Transaction{}, err
		}
		newTotal, usedCoin, err := s.d.DecreaseUserCoin(Transaction.UserID, coin, transactionData.Total)
		if err != nil {
			return transaction.Transaction{}, err
		}
		transactionData.Coin = usedCoin
		transactionData.Total = newTotal
	}

	var client snap.Client
	client.New(s.config.ServerKey, midtrans.Sandbox)
	snapResponse, err := client.CreateTransaction(req)
	if snapResponse == nil {
		return transaction.Transaction{}, err
	}

	transactionData.SnapURL = snapResponse.RedirectURL

	error := s.d.CreateTransactions(transactionData)
	if error != nil {
		return transaction.Transaction{}, error
	}
	err = s.d.AddTransactionItems(Transaction.UserID, transactionData.ID)
	if err != nil {
		return transaction.Transaction{}, err
	}
	return transactionData, nil
}

func (s *TransactionService) UpdateTransaction(transaction transaction.UpdateTransaction) error {
	// Implement logic here
	return nil
}

func (s *TransactionService) DeleteTransaction(transactionId string) error {
	// Implement logic here
	return nil
}
