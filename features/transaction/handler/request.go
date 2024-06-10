package handler

type TransactionRequest struct {
	VoucherCode string `json:"voucher_code"`
	UsingCoin   bool   `json:"using_coin"`
}
