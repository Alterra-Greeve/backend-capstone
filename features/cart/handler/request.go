package handler

type CreateCartRequest struct {
	ProductID string `json:"product_id"`
}
type UpdateCartRequest struct {
	ProductID string `json:"product_id"`
	Type      string `json:"type"`
	Quantity  int    `json:"qty"`
}
