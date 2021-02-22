package entities

import (
	"time"
	"fmt"
)

type Product struct {
	ID int64 `json:"id"`
	Title string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
	Image string `json:"image,omitempty"`
	Price float64 `json:"price,omitempty"`
	AvailableSizes []string `json:"availableSizes,omitempty"`
	StartDate time.Time `json:"start_date,omitempty"`
	EndDate time.Time `json:"end_date,omitempty"`
	HistoryID int64 `json:"history_id"`
}

func NewProduct(title string, description string, image string, price float64, availableSizes []string) (*Product, error) {
	p := &Product{
		Title: title,
		Description: description,
		Image: image,
		Price: price,
		AvailableSizes: availableSizes,
	}
	err := p.Validate()
	if err != nil {
		return nil, err
	}
	return p, nil
}

//Validate validate book
func (p *Product) Validate() error {
	if p.Title == "" {
		return fmt.Errorf("Title is empty")
	}
	if p.Image == "" {
		return fmt.Errorf("Image is empty")
	}
	if p.Price == 0.0 {
		return fmt.Errorf("Price is empty")
	}
	if len(p.AvailableSizes) == 0 {
		return fmt.Errorf("AvailableSizes is empty")
	}
	return nil
}