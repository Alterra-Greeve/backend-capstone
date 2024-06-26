package service

import (
	"backendgreeve/features/dashboard"
	"time"
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

	// Static categories
	staticCategories := []string{"Mengurangi Pemanasan Global", "Hemat Uang", "Mengurangi Limbah", "Perluas Wawasan"}

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
	for _, name := range staticCategories {
		totalPoint, exists := totalPoints[name]
		if !exists {
			totalPoint = 0
		}
		totalImpactPoints = append(totalImpactPoints, dashboard.ImpactPoint{Name: name, TotalPoint: totalPoint})
	}

	return totalImpactPoints, nil
}

func (s *DashboardService) GetUserCoin(userID string) ([]dashboard.UserCoin, error) {
	productEarning, err := s.d.GetTransactionCoinEarned(userID)
	if err != nil {
		return nil, err
	}
	challengeEarning, err := s.d.GetChallengeCoinEarned(userID)
	if err != nil {
		return nil, err
	}

	result := []dashboard.UserCoin{}
	result = append(result, productEarning...)
	result = append(result, challengeEarning...)
	for i, res := range result {
		date, _ := time.Parse(time.RFC3339, res.Date)
		result[i].Date = date.Format("2006-01-02")
	}
	return result, nil
}

func (s *DashboardService) GetUserSpending(userID string) ([]dashboard.UserSpending, error) {
	userSpending, err := s.d.GetUserSpending(userID)
	if err != nil {
		return nil, err
	}
	for i, res := range userSpending {
		date, _ := time.Parse(time.RFC3339, res.Date)
		userSpending[i].Date = date.Format("2006-01-02")
	}
	return userSpending, nil
}
