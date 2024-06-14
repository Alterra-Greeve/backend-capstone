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
	Helps       []ImpactCategoryResponse `json:"helps"`
	ImpactPoint int                      `json:"impact_point"`
	CreatedAt   string                   `json:"createdAt"`
	UpdatedAt   string                   `json:"updatedAt"`
}

type ImpactCategoryResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func ToResponse(data []orders.OrdersProduct) []OrdersProductResponse {
	var responses []OrdersProductResponse
	for _, order := range data {
		var helps []ImpactCategoryResponse
		for _, category := range order.ImpactCategories {
			helps = append(helps, ImpactCategoryResponse{
				ID:   category.ID,
				Name: category.Name,
			})
		}

		totalCoin := order.Coin * order.Qty // Calculate the total coin

		response := OrdersProductResponse{
			Username:    order.Username,
			Email:       order.Email,
			Qty:         order.Qty,
			Coin:        totalCoin, // Set the total coin here
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
	ChallengeName  string                   `json:"challenge_name"`
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
				ID:   category.ID,
				Name: category.Name,
			})
		}

		response := ChallengeConfirmationResponse{
			Username:       confirmation.Username,
			ChallengeName:  confirmation.Challenge.Title,
			ImpactPoint:    confirmation.ImpactPointTotal,
			ImpactCategory: helps,
			CreatedAt:      confirmation.CreatedAt.Format("02/01/06"),
			UpdatedAt:      confirmation.UpdatedAt.Format("02/01/06"),
		}
		responses = append(responses, response)
	}
	return responses
}
