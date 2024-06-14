package service

import "backendgreeve/features/orders"

type OrderService struct {
	d orders.OrdersDataInterface
}

func New(d orders.OrdersDataInterface) orders.OrdersServiceInterface {
	return &OrderService{
		d: d,
	}
}

func (s *OrderService) GetOrdersProduct() ([]orders.OrdersProduct, error) {
	return s.d.GetOrdersProduct()
}

// func (s *OrderService) GetOrdersChallenge() ([]orders.OrderChallengeConfirmation, error) {
// 	return s.d.GetOrdersChallenge()
// }
