package handler

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
	ID        string            `json:"challenge_confirmation_id"`
	UserID    string            `json:"user_id"`
	Status    string            `json:"status"`
	Challenge ChallengeResponse `json:"challenge"`
}
