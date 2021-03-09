package order

import (
	"landing-back/entities"
)

//Reader interface
type Reader interface {
	List() ([]entities.Order, error)
}

//Writer book writer
type Writer interface {
	Create(i *entities.Order) error
	Delete(id int64) error
}

//Repository interface
type Repository interface {
	Reader
	Writer
}

//UseCase interface
type UseCase interface {
	// GetOrder(id int64) (*entities.Order, error)
	// SearchOrders(query string) ([]*entities.Order, error)
	ListOrders() ([]entities.Order, error)
	CreateOrder(order entities.Order) error
	// UpdateOrder(e *entities.Order) error
	DeleteOrder(id int64) error
}
