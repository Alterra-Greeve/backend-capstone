package handler

import "time"

type VoucherResponse struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Code        string    `json:"code"`
	Discount    string    `json:"discount"`
	Description string    `json:"description"`
	ExpiredAt   time.Time `json:"expired_at"`
}
