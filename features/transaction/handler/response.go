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
