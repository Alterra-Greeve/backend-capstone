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
	TotalMembership                    int                              `json:"total_membership"`
	NewProducts                        []productHandler.ProductResponse `json:"new_products"`
	MonthlyImpact                      []MonthlyImpactResponse          `json:"monthly_impact"`
}

type MonthlyImpactResponse struct {
	Date  string        `json:"month"`
	Point []ImpactPoint `json:"point"`
}

type ImpactPoint struct {
	Name  string `json:"name"`
	Point int64  `json:"point"`
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
		TotalMembership:                    data.TotalMembership,
		NewProducts:                        newProducts,
	}
}

func (m *MonthlyImpactResponse) ToResponse(data dashboard.MonthlyImpact) MonthlyImpactResponse {
	impactPoints := make([]ImpactPoint, len(data.ImpactPoint))
	for i, point := range data.ImpactPoint {
		impactPoints[i] = ImpactPoint{
			Name:  point.Name,
			Point: int64(point.TotalPoint),
		}
	}
	return MonthlyImpactResponse{
		Date:  data.Date,
		Point: impactPoints,
	}
}
