package handler

type TransactionResponse struct {
	ID      string `json:"id"`
	Amount  int    `json:"amount"`
	SnapURL string `json:"snap_url"`
}
