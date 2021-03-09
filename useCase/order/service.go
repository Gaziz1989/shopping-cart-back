package order

import (
	"fmt"
	"landing-back/entities"
)

//Service product usecase
type Service struct {
	repo Repository
}

//NewService create new service
func NewService(r Repository) *Service {
	return &Service{
		repo: r,
	}
}

//DeleteOrder delete a product
func (s *Service) DeleteOrder(id int64) error {
	err := s.repo.Delete(id)
	if err != nil {
		return err
	}
	return s.repo.Delete(id)
}

//CreateBook create a product
func (s *Service) CreateOrder(order entities.Order) error {
	err := order.Validate()
	if err != nil {
		return err
	}

	err = s.repo.Create(&order)
	if err != nil {
		return err
	}

	return nil
}

//ListOrders list products
func (s *Service) ListOrders() ([]entities.Order, error) {
	products, err := s.repo.List()
	if err != nil {
		return nil, err
	}
	if len(products) == 0 {
		return nil, fmt.Errorf("Orders not found")
	}
	return products, nil
}