package handler

import (
	challenge "backendgreeve/features/challenges"
)

type ChallengeResponse struct {
	ID          string                      `json:"id"`
	Title       string                      `json:"title"`
	Description string                      `json:"description"`
	Coin        int                         `json:"coin"`
	Exp         int                         `json:"exp"`
	Participant int                         `json:"participant"`
	Difficulty  string                      `json:"difficulty"`
	ImageURL    string                      `json:"image_url"`
	DateStart   string                      `json:"date_start"`
	DateEnd     string                      `json:"date_end"`
	Categories  []ChallengeImpactCategories `json:"categories"`
}

type ChallengeImpactCategories struct {
	ImpactCategory ImpactCategory `json:"impact_category"`
}

type ImpactCategory struct {
	Name        string `json:"name"`
	ImpactPoint int    `json:"impact_point"`
	IconURL     string `json:"icon_url"`
}

type MetadataResponse struct {
	CurrentPage int `json:"current_page"`
	TotalPage   int `json:"total_page"`
}

type UserChallengeConfirmationResponse struct {
	ID        string                       `json:"challenge_confirmation_id"`
	UserID    string                       `json:"user_id"`
	Status    string                       `json:"status"`
	Challenge ChallengeResponse            `json:"challenge"`
	Images    []ChallengeConfirmationImage `json:"images"`
}

type ChallengeConfirmationImage struct {
	ID       string `json:"id"`
	ImageURL string `json:"image_url"`
}

func (cd UserChallengeConfirmationResponse) ToResponse(challenge challenge.ChallengeConfirmation) UserChallengeConfirmationResponse {
	challengeResponse := ChallengeResponse{
		ID:          challenge.Challenge.ID,
		Title:       challenge.Challenge.Title,
		Difficulty:  challenge.Challenge.Difficulty,
		Description: challenge.Challenge.Description,
		Exp:         challenge.Challenge.Exp,
		Coin:        challenge.Challenge.Coin,
		ImageURL:    challenge.Challenge.ImageURL,
		DateStart:   challenge.Challenge.DateStart.Format("02/01/06"),
		DateEnd:     challenge.Challenge.DateEnd.Format("02/01/06"),
	}

	for _, category := range challenge.Challenge.ImpactCategories {
		challengeResponse.Categories = append(challengeResponse.Categories, ChallengeImpactCategories{
			ImpactCategory: ImpactCategory{
				Name:        category.ImpactCategory.Name,
				ImpactPoint: category.ImpactCategory.ImpactPoint,
				IconURL:     category.ImpactCategory.IconURL,
			},
		})
	}

	var images []ChallengeConfirmationImage
	for _, image := range challenge.ChallengeConfirmationImages {
		images = append(images, ChallengeConfirmationImage{
			ID:       image.ID,
			ImageURL: image.ImageURL,
		})
	}

	response := UserChallengeConfirmationResponse{
		ID:        challenge.ID,
		UserID:    challenge.UserID,
		Status:    challenge.Status,
		Challenge: challengeResponse,
		Images:    images,
	}

	return response
}
