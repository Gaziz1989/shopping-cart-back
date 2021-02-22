package product

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

//CreateBook create a product
func (s *Service) CreateProduct(title string, description string, image string, price float64, availableSizes []string) (*entities.Product, error) {
	b, err := entities.NewProduct(title, description, image, price, availableSizes)
	if err != nil {
		return nil, err
	}
	return s.repo.Create(b)
}

//ListProducts list products
func (s *Service) ListProducts() ([]entities.Product, error) {
	products, err := s.repo.List()
	if err != nil {
		return nil, err
	}
	if len(products) == 0 {
		return nil, fmt.Errorf("Products not found")
	}
	return products, nil
}