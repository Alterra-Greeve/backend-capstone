package handler

type ProductResponse struct {
	ID          string                  `json:"product_id"`
	Name        string                  `json:"name"`
	Description string                  `json:"description"`
	Price       float64                 `json:"price"`
	Coin        int                     `json:"coin"`
	CreatedAt   string                  `json:"created_at"`
	UpdatedAt   string                  `json:"updated_at"`
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

type MetadataResponse struct {
	TotalPage int `json:"total_page"`
	Page      int `json:"current_page"`
}
