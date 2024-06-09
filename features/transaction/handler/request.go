package handler

type TransactionRequest struct {
	ProductId uint   `json:"product_id" validate:"required"`
	Quantity  uint   `json:"quantity" validate:"required"`
	Address   string `json:"address" validate:"required"`
}
