package handler

type ProductRequest struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Price       float64  `json:"price"`
	Coin        int      `json:"coin"`
	Stock       int      `json:"stock"`
	Category    []string `json:"category"`
	ImageURL    []string `json:"image_url"`
}
