package handler

import "time"

type VoucherRequest struct {
	Name        string    `json:"name"`
	Code        string    `json:"code"`
	Discount    string    `json:"discount"`
	Description string    `json:"description"`
	ExpiredAt   time.Time `json:"expired_at"`
}
