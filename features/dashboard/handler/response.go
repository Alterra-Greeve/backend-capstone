package handler

import (
	"backendgreeve/features/dashboard"
	productHandler "backendgreeve/features/product/handler"
)

type DashboardResponse struct {
	TotalProduct                       int                              `json:"total_product"`
	TotalProductPercentage             string                           `json:"total_product_percentage"`
	TotalNewProductThisMonth           int                              `json:"total_new_product_this_month"`
	TotalNewProductThisMonthPercentage string                           `json:"total_new_product_this_month_percentage"`
	TotalUser                          int                              `json:"total_user"`
	TotalUserPercentage                string                           `json:"total_user_percentage"`
	NewProducts                        []productHandler.ProductResponse `json:"new_products"`
}

func (d *DashboardResponse) ToResponse(data dashboard.Dashboard) *DashboardResponse {
	newProducts := make([]productHandler.ProductResponse, len(data.NewProduct))
	for i, product := range data.NewProduct {
		newProducts[i] = new(productHandler.ProductResponse).ToResponse(product)
	}

	return &DashboardResponse{
		TotalProduct:                       data.TotalProduct,
		TotalProductPercentage:             data.TotalProductPercentage,
		TotalNewProductThisMonth:           data.TotalNewProductThisMonth,
		TotalNewProductThisMonthPercentage: data.TotalNewProductPercentage,
		TotalUser:                          data.TotalUser,
		TotalUserPercentage:                data.NewUserPercentage,
		NewProducts:                        newProducts,
	}
}
