package service

import (
	"backendgreeve/features/dashboard"
)

type DashboardService struct {
	d dashboard.DashboardDataInterface
}

func New(d dashboard.DashboardDataInterface) dashboard.DashboardServiceInterface {
	return &DashboardService{
		d: d,
	}
}

func (s *DashboardService) GetDashboard() (dashboard.Dashboard, error) {
	return s.d.GetDashboard()
}

func (s *DashboardService) GetMonthlyImpact() ([]dashboard.MonthlyImpact, error) {
	months := []string{"2024-01", "2024-02", "2024-03", "2024-04", "2024-05", "2024-06"}

	var monthlyImpacts []dashboard.MonthlyImpact

	for _, month := range months {
		impactPoints, err := s.getMonthlyImpactPoints(month)
		if err != nil {
			return nil, err
		}

		monthlyImpact := dashboard.MonthlyImpact{
			Date:        month,
			ImpactPoint: impactPoints,
		}

		monthlyImpacts = append(monthlyImpacts, monthlyImpact)
	}
	return monthlyImpacts, nil
}

func (s *DashboardService) getMonthlyImpactPoints(month string) ([]dashboard.ImpactPoint, error) {
	productPoints := []dashboard.ImpactPoint{}
	challengePoints := []dashboard.ImpactPoint{}

	impactPoints, err := s.d.GetMonthlyImpactChallenge(month)
	if err != nil {
		return nil, err
	}
	challengePoints = append(challengePoints, impactPoints...)

	impactPoints, err = s.d.GetMonthlyImpactProduct(month)
	if err != nil {
		return nil, err
	}
	productPoints = append(productPoints, impactPoints...)

	// Calculate total point
	totalPoints := make(map[string]int)
	for _, point := range challengePoints {
		totalPoints[point.Name] += point.TotalPoint
	}
	for _, point := range productPoints {
		totalPoints[point.Name] += point.TotalPoint
	}

	var totalImpactPoints []dashboard.ImpactPoint
	for name, totalPoint := range totalPoints {
		totalImpactPoints = append(totalImpactPoints, dashboard.ImpactPoint{Name: name, TotalPoint: totalPoint})
	}

	return totalImpactPoints, nil
}
