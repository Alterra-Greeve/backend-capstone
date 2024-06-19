package service

import (
	"backendgreeve/features/orders"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockOrdersData struct {
	mock.Mock
}

func (m *MockOrdersData) GetOrdersProduct() ([]orders.OrdersProduct, error) {
	args := m.Called()
	return args.Get(0).([]orders.OrdersProduct), args.Error(1)
}

func (m *MockOrdersData) GetOrdersChallenge() ([]orders.ChallengeConfirmation, error) {
	args := m.Called()
	return args.Get(0).([]orders.ChallengeConfirmation), args.Error(1)
}

func (m *MockOrdersData) GetTotalCoin() (int, error) {
	args := m.Called()
	return args.Int(0), args.Error(1)
}
func TestOrderService_GetOrdersProduct(t *testing.T) {
	mockData := new(MockOrdersData)
	service := New(mockData)

	expectedOrders := []orders.OrdersProduct{
		{UserID: "1", Username: "User 1", Email: "user1@example.com", ProductName: "Product 1", Qty: 1, Coin: 10, TotalHarga: 100.0, CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{UserID: "2", Username: "User 2", Email: "user2@example.com", ProductName: "Product 2", Qty: 2, Coin: 20, TotalHarga: 200.0, CreatedAt: time.Now(), UpdatedAt: time.Now()},
	}
	mockData.On("GetOrdersProduct").Return(expectedOrders, nil)

	orders, err := service.GetOrdersProduct()

	assert.NoError(t, err)
	assert.Equal(t, expectedOrders, orders)
	mockData.AssertExpectations(t)
}

func TestOrderService_GetOrdersProduct_Error(t *testing.T) {
	mockData := new(MockOrdersData)
	service := New(mockData)

	expectedError := errors.New("error fetching orders product")

	mockData.On("GetOrdersProduct").Return([]orders.OrdersProduct{}, expectedError)

	orders, err := service.GetOrdersProduct()

	assert.Error(t, err)
	assert.Empty(t, orders)
	assert.Equal(t, expectedError, err)
	mockData.AssertExpectations(t)
}

func TestOrderService_GetOrdersChallenge(t *testing.T) {
	mockData := new(MockOrdersData)
	service := New(mockData)

	expectedChallenges := []orders.ChallengeConfirmation{
		{ID: "1", Status: "Confirmed"},
		{ID: "2", Status: "Pending"},
	}

	mockData.On("GetOrdersChallenge").Return(expectedChallenges, nil)

	challenges, err := service.GetOrdersChallenge()

	assert.NoError(t, err)
	assert.Equal(t, expectedChallenges, challenges)
	mockData.AssertExpectations(t)
}

func TestOrderService_GetOrdersChallenge_Error(t *testing.T) {
	mockData := new(MockOrdersData)
	service := New(mockData)

	expectedError := errors.New("error fetching orders challenge")

	mockData.On("GetOrdersChallenge").Return([]orders.ChallengeConfirmation{}, expectedError)

	challenges, err := service.GetOrdersChallenge()

	assert.Error(t, err)
	assert.Empty(t, challenges)
	assert.Equal(t, expectedError, err)
	mockData.AssertExpectations(t)
}
