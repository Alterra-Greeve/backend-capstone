package handler

type ProductResponse struct {
	ID          string                  `json:"id"`
	Name        string                  `json:"name"`
	Description string                  `json:"description"`
	Price       float64                 `json:"price"`
	Coin        int                     `json:"coin"`
	Category    []ProductImpactCategory `json:"category"`
	Images      []ProductImage          `json:"images"`
}

type ProductImpactCategory struct {
	ID               string `json:"id"`
	ProductID        string `json:"product_id"`
	ImpactCategoryID string `json:"impact_category_id"`
}

type ProductImage struct {
	ID        string `json:"id"`
	ProductID string `json:"product_id"`
	ImageURL  string `json:"image_url"`
	Position  int    `json:"position"`
}
