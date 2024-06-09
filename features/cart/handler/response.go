package handler

import (
	product "backendgreeve/features/product/handler"
)

type CartResponse struct {
	User  User        `json:"user"`
	Items []CartItems `json:"items"`
}

type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Address  string `json:"address"`
}

type CartItems struct {
	Quantity int                     `json:"quantity"`
	Product  product.ProductResponse `json:"product"`
}
