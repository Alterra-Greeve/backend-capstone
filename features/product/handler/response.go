package handler

type ProductResponse struct {
	ID          string                  `json:"product_id"`
	Name        string                  `json:"name"`
	Description string                  `json:"description"`
	Price       float64                 `json:"price"`
	Coin        int                     `json:"coin"`
	CurrentPage int                     `json:"current_page"`
	TotalPage   int                     `json:"total_page"`
	Category    []ProductImpactCategory `json:"category"`
	Images      []ProductImage          `json:"images"`
}

type ProductImpactCategory struct {
	ImpactCategory ImpactCategory `json:"impact_category"`
}

type ProductImage struct {
	ImageURL string `json:"image_url"`
	Position int    `json:"position"`
}

type ImpactCategory struct {
	Name        string `json:"name"`
	ImpactPoint int    `json:"impact_point"`
}
