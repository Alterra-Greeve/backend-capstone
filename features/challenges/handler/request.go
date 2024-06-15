package handler

type ChallengeParticipateRequest struct {
	ChallengeId   string `json:"challenge_id"`
	ChallengeType string `json:"type"`
}

type ChallengeCreateRequest struct {
	Title            string   `json:"title"`
	Description      string   `json:"description"`
	Exp              int      `json:"exp"`
	Coin             int      `json:"coin"`
	DateStart        string   `json:"date_start"`
	DateEnd          string   `json:"date_end"`
	Difficulty       string   `json:"difficulty"`
	ImageURL         string   `json:"image_url"`
	ImpactCategories []string `json:"category"`
}

type ChallengeImpactCategory struct {
	ImpactCategoryID string `json:"impact_category_id"`
}

type EditChallengeConfirmationForUser struct {
	ID    string   `json:"id"`
	Image []string `json:"image"`
}
