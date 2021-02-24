package product

import (
	"landing-back/entities"
)

//Reader interface
type Reader interface {
	List() ([]entities.Product, error)
}

//Writer book writer
type Writer interface {
	Create(i *entities.Product) (*entities.Product, error)
	Delete(id int64) error
}

//Repository interface
type Repository interface {
	Reader
	Writer
}

//UseCase interface
type UseCase interface {
	// GetProduct(id int64) (*entities.Product, error)
	// SearchProducts(query string) ([]*entities.Product, error)
	ListProducts() ([]entities.Product, error)
	CreateProduct(title string, description string, image string, price float64, availableSizes []string) (*entities.Product, error)
	// UpdateProduct(e *entities.Product) error
	DeleteProduct(id int64) error
}
