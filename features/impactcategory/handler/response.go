package handler

type ImpactCategoryResponse struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	ImpactPoint int    `json:"impact_point"`
	IconURL     string `json:"icon_url"`
}
