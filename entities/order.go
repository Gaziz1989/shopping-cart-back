package entities

import (
	"time"
	"fmt"
	"encoding/json"
)

type Order struct {
	ID int64 `json:"id"`
	Email string `json:"email,omitempty"`
	Name string `json:"name,omitempty"`
	Address string `json:"address,omitempty"`
	Total float64 `json:"total,omitempty"`
	StartDate time.Time `json:"start_date,omitempty" gorm:"default:CURRENT_TIMESTAMP"`
	EndDate *time.Time `json:"end_date,omitempty"`
	HistoryID int64 `json:"history_id"`
	
	OrderProducts []Product `json:"cartItems,omitempty" gorm:"-"`
}

type OrderProducts struct {
	ID int64 `json:"id"`
	OrderID int64 `json:"order_id,omitempty"`
	ProductID int64 `json:"product_id,omitempty"`
	StartDate time.Time `json:"start_date,omitempty" gorm:"default:CURRENT_TIMESTAMP"`
	EndDate *time.Time `json:"end_date,omitempty"`
	HistoryID int64 `json:"history_id"`
}

func NewOrder(email string, name string, address string, total float64) (*Order, error) {
	p := &Order{
		Email: email,
		Name: name,
		Address: address,
		Total: total,
	}
	err := p.Validate()
	if err != nil {
		return nil, err
	}
	return p, nil
}

//Validate validate order
func (p *Order) Validate() error {
	if p.Email == "" {
		return fmt.Errorf("Email is empty")
	}
	if p.Name == "" {
		return fmt.Errorf("Name is empty")
	}
	if p.Address == "" {
		return fmt.Errorf("Address is empty")
	}
	if p.Total == 0 {
		return fmt.Errorf("Total is empty")
	}
	if len(p.OrderProducts) == 0 {
		return fmt.Errorf("No products selected")
	}
	return nil
}

func (u *Order) String() string {
	s, err := u.PrettyPrint()
	if err != nil {
		return fmt.Sprintf("Error: %s", err.Error())
	}
	return s
}

func (u *Order) PrettyPrint() (string, error) {
	b, err := json.MarshalIndent(u, "", "  ")
	if err != nil {
		return "", err
	}
	return string(b), nil
}