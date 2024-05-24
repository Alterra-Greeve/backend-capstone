package handler

type CreateImpactCategoryRequest struct {
	Name        string `json:"name"`
	ImpactPoint int    `json:"impact_point"`
	IconURL     string `json:"icon_url"`
}
