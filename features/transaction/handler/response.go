package handler

import "backendgreeve/features/transaction"

type TransactionResponse struct {
	ID      string `json:"id"`
	Amount  int    `json:"amount"`
	SnapURL string `json:"snap_url"`
}

type TransactionUserResponse struct {
	ID           string  `json:"id"`
	Total        float64 `json:"total"`
	Status       string  `json:"status"`
	SnapURL      string  `json:"snap_url"`
	ProductName  string  `json:"product_name"`
	ProductImage string  `json:"product_image"`
}

func (t TransactionUserResponse) FromEntity(transaction transaction.TransactionData) TransactionUserResponse {
	response := TransactionUserResponse{}
	response.ID = transaction.ID
	response.Total = transaction.Total
	response.Status = transaction.Status
	response.SnapURL = transaction.SnapURL
	response.ProductName = transaction.TransactionItems[0].Product.Name
	response.ProductImage = transaction.TransactionItems[0].Product.Images[0].ImageURL
	return response
}

type TransactionAllUserResponses struct {
	ID           string  `json:"id"`
	User         string  `json:"username"`
	Total        float64 `json:"total_transaction"`
	Status       string  `json:"status"`
	SnapURL      string  `json:"snap_url"`
	ProductName  string  `json:"product_name"`
	ProductImage string  `json:"product_image"`
	CreatedAt    string  `json:"created_at"`
	UpdatedAt    string  `json:"updated_at"`
}

func (t *TransactionAllUserResponses) FromEntity(transaction transaction.TransactionData) TransactionAllUserResponses {
	response := TransactionAllUserResponses{}
	response.ID = transaction.ID
	response.User = transaction.User.Name
	response.Total = transaction.Total
	response.Status = transaction.Status
	response.SnapURL = transaction.SnapURL
	if len(transaction.TransactionItems) > 0 {
		response.ProductName = transaction.TransactionItems[0].Product.Name
		if len(transaction.TransactionItems[0].Product.Images) > 0 {
			response.ProductImage = transaction.TransactionItems[0].Product.Images[0].ImageURL
		}
	}
	response.CreatedAt = transaction.CreatedAt.Format("02/01/2006 15:04")
	response.UpdatedAt = transaction.UpdatedAt.Format("02/01/2006 15:04")
	return response
}
