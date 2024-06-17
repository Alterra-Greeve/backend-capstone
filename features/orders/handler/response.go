package handler

import (
	"backendgreeve/features/orders"
)

type OrdersProductResponse struct {
	Username    string                   `json:"username"`
	Email       string                   `json:"email"`
	Coin        int                      `json:"coin"`
	Qty         int                      `json:"qty"`
	ProductName string                   `json:"product_name"`
	TotalHarga  float64                  `json:"total"`
	Helps       []ImpactCategoryResponse `json:"helps"`
	ImpactPoint int                      `json:"impact_point"`
	CreatedAt   string                   `json:"createdAt"`
	UpdatedAt   string                   `json:"updatedAt"`
}

type ImpactCategoryResponse struct {
	ImpactCategory ImpactCategory `json:"impact_category"`
}

type ImpactCategory struct {
	Name     string `json:"name"`
	ImageURL string `json:"image_url"`
}

func ToResponse(data []orders.OrdersProduct) []OrdersProductResponse {
	var responses []OrdersProductResponse
	for _, order := range data {
		var helps []ImpactCategoryResponse
		for _, category := range order.ImpactCategories {
			helps = append(helps, ImpactCategoryResponse{
				ImpactCategory: ImpactCategory{
					Name:     category.Name,
					ImageURL: category.ImageURL,
				},
			})
		}

		totalCoin := order.Coin * order.Qty
		Total := int(order.TotalHarga)
		response := OrdersProductResponse{
			Username:    order.Username,
			Email:       order.Email,
			Qty:         order.Qty,
			Coin:        totalCoin,
			TotalHarga:  float64(Total),
			ProductName: order.ProductName,
			Helps:       helps,
			ImpactPoint: order.ImpactPointTotal,
			CreatedAt:   order.CreatedAt.Format("02/01/06"),
			UpdatedAt:   order.UpdatedAt.Format("02/01/06"),
		}
		responses = append(responses, response)
	}
	return responses
}

// Challenge
type ChallengeConfirmationResponse struct {
	Username       string                   `json:"username"`
	Email          string                   `json:"email"`
	ChallengeName  string                   `json:"challenge_name"`
	Exp            int                      `json:"exp"`
	Difficulty     string                   `json:"difficulty"`
	ImpactPoint    int                      `json:"impact_point"`
	ImpactCategory []ImpactCategoryResponse `json:"helps"`
	CreatedAt      string                   `json:"createdAt"`
	UpdatedAt      string                   `json:"updatedAt"`
}

func ToChallengeConfirmationResponse(data []orders.ChallengeConfirmation) []ChallengeConfirmationResponse {
	var responses []ChallengeConfirmationResponse
	for _, confirmation := range data {
		var helps []ImpactCategoryResponse
		for _, category := range confirmation.ImpactCategories {
			helps = append(helps, ImpactCategoryResponse{
				ImpactCategory: ImpactCategory{
					Name:     category.Name,
					ImageURL: category.ImageURL,
				},
			})
		}

		response := ChallengeConfirmationResponse{
			Username:       confirmation.Username,
			Email:          confirmation.Email,
			ChallengeName:  confirmation.Challenge.Title,
			Exp:            confirmation.Challenge.Exp,
			Difficulty:     confirmation.Challenge.Difficulty,
			ImpactPoint:    confirmation.ImpactPointTotal,
			ImpactCategory: helps,
			CreatedAt:      confirmation.CreatedAt.Format("02/01/06"),
			UpdatedAt:      confirmation.UpdatedAt.Format("02/01/06"),
		}
		responses = append(responses, response)
	}
	return responses
}
