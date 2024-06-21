package service

import (
	"backendgreeve/features/dashboard"
	"backendgreeve/features/product"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockDashboardData struct {
	mock.Mock
}

func (m *MockDashboardData) GetDashboard() (dashboard.Dashboard, error) {
	args := m.Called()
	return args.Get(0).(dashboard.Dashboard), args.Error(1)
}

func (m *MockDashboardData) GetMonthlyImpactChallenge(month string) ([]dashboard.ImpactPoint, error) {
	args := m.Called(month)
	return args.Get(0).([]dashboard.ImpactPoint), args.Error(1)
}

func (m *MockDashboardData) GetMonthlyImpactProduct(month string) ([]dashboard.ImpactPoint, error) {
	args := m.Called(month)
	return args.Get(0).([]dashboard.ImpactPoint), args.Error(1)
}

func (m *MockDashboardData) GetTransactionCoinEarned(userID string) ([]dashboard.UserCoin, error) {
	args := m.Called(userID)
	return args.Get(0).([]dashboard.UserCoin), args.Error(1)
}

func (m *MockDashboardData) GetChallengeCoinEarned(userID string) ([]dashboard.UserCoin, error) {
	args := m.Called(userID)
	return args.Get(0).([]dashboard.UserCoin), args.Error(1)
}

func (m *MockDashboardData) GetUserSpending(userID string) ([]dashboard.UserSpending, error) {
	args := m.Called(userID)
	return args.Get(0).([]dashboard.UserSpending), args.Error(1)
}

func (m *MockDashboardData) GetNewProduct() ([]product.Product, error) {
	args := m.Called()
	return args.Get(0).([]product.Product), args.Error(1)
}
func (m *MockDashboardData) GetNewUserPercentage() (string, error) {
	args := m.Called()
	return args.String(0), args.Error(1)
}
func (m *MockDashboardData) GetTotalMembership() (int, error) {
	args := m.Called()
	return args.Int(0), args.Error(1)
}
func (m *MockDashboardData) GetTotalNewProductPercentage() (string, error) {
	args := m.Called()
	return args.String(0), args.Error(1)
}

func (m *MockDashboardData) GetTotalNewProductThisMonth() (int, error) {
	args := m.Called()
	return args.Int(0), args.Error(1)
}
func (m *MockDashboardData) GetTotalProductPercentage() (string, error) {
	args := m.Called()
	return args.String(0), args.Error(1)
}
func (m *MockDashboardData) GetTotalProduct() (int, error) {
	args := m.Called()
	return args.Int(0), args.Error(1)
}

func (m *MockDashboardData) GetTotalUser() (int, error) {
	args := m.Called()
	return args.Int(0), args.Error(1)
}

func TestDashboardService_GetUserCoin(t *testing.T) {
	mockData := new(MockDashboardData)
	service := New(mockData)

	userID := "user123"

	mockData.On("GetTransactionCoinEarned", userID).Return([]dashboard.UserCoin{
		{Date: "2024-06-01T00:00:00Z", Coin: 10},
		{Date: "2024-06-02T00:00:00Z", Coin: 20},
	}, nil)

	mockData.On("GetChallengeCoinEarned", userID).Return([]dashboard.UserCoin{
		{Date: "2024-06-01T00:00:00Z", Coin: 5},
		{Date: "2024-06-02T00:00:00Z", Coin: 15},
	}, nil)

	expectedUserCoins := []dashboard.UserCoin{
		{Date: "2024-06-01", Coin: 10},
		{Date: "2024-06-02", Coin: 20},
		{Date: "2024-06-01", Coin: 5},
		{Date: "2024-06-02", Coin: 15},
	}

	result, err := service.GetUserCoin(userID)

	assert.NoError(t, err)
	assert.Equal(t, expectedUserCoins, result)
	mockData.AssertExpectations(t)
}

func TestDashboardService_GetUserSpending(t *testing.T) {
	mockData := new(MockDashboardData)
	service := New(mockData)

	userID := "user123"

	mockData.On("GetUserSpending", userID).Return([]dashboard.UserSpending{
		{Date: "2024-06-01T00:00:00Z"},
		{Date: "2024-06-02T00:00:00Z"},
	}, nil)

	expectedUserSpending := []dashboard.UserSpending{
		{Date: "2024-06-01"},
		{Date: "2024-06-02"},
	}

	result, err := service.GetUserSpending(userID)

	assert.NoError(t, err)
	assert.Equal(t, expectedUserSpending, result)
	mockData.AssertExpectations(t)
}

func TestDashboardService_GetDashboard(t *testing.T) {
	mockData := new(MockDashboardData)
	service := New(mockData)

	expectedDashboard := dashboard.Dashboard{
		TotalProduct:              100,
		TotalProductPercentage:    "50%",
		TotalNewProductThisMonth:  20,
		TotalNewProductPercentage: "10%",
		TotalUser:                 500,
		NewUserPercentage:         "5%",
		TotalMembership:           200,
	}

	mockData.On("GetDashboard").Return(expectedDashboard, nil)

	result, err := service.GetDashboard()

	assert.NoError(t, err)
	assert.Equal(t, expectedDashboard, result)
	mockData.AssertExpectations(t)
}
