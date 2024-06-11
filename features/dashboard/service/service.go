package service

import "backendgreeve/features/dashboard"

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
