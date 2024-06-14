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

		response := OrdersProductResponse{
			Username:    order.Username,
			Email:       order.Email,
			Coin:        order.TotalCoin,
			Qty:         order.Qty,
			ProductName: order.ProductName,
			Helps:       helps,
			ImpactPoint: order.ImpactPointTotal,
			CreatedAt:   order.CreatedAt.String(),
			UpdatedAt:   order.UpdatedAt.String(),
		}
		responses = append(responses, response)
	}
	return responses
}

// type ChallengeConfirmationResponse struct {
// 	Username       string                   `json:"username"`
// 	ChallengeName  string                   `json:"nama_challenge"`
// 	ImpactPoint    int                      `json:"impact_point"`
// 	ImpactCategory []ImpactCategoryResponse `json:"impact_category_helps"`
// 	CreatedAt      string                   `json:"createdAt"`
// 	UpdatedAt      string                   `json:"updatedAt"`
// }

// func ToChallengeConfirmationResponse(data []orders.OrderChallengeConfirmation) []ChallengeConfirmationResponse {
// 	var responses []ChallengeConfirmationResponse
// 	for _, confirmation := range data {
// 		var helps []ImpactCategoryResponse
// 		for _, category := range confirmation.ImpactCategories {
// 			helps = append(helps, ImpactCategoryResponse{
// 				ID:   category.ID,
// 				Name: category.Name,
// 			})
// 		}

// 		response := ChallengeConfirmationResponse{
// 			Username:       confirmation.Username, // Make sure to fetch User's Username
// 			ChallengeName:  confirmation.Challenge.Title,
// 			ImpactPoint:    confirmation.ImpactPointTotal,
// 			ImpactCategory: helps,
// 			CreatedAt:      confirmation.CreatedAt.Format(time.RFC3339),
// 			UpdatedAt:      confirmation.UpdatedAt.Format(time.RFC3339),
// 		}
// 		responses = append(responses, response)
// 	}
// 	return responses
// }
