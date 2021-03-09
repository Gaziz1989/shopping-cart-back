package entities

import (
	"time"
	"fmt"
	"github.com/lib/pq"
	"encoding/json"
)

type Product struct {
	ID int64 `json:"id"`
	Title string `json:"title,omitempty"`
	Description *string `json:"description,omitempty"`
	Image string `json:"image,omitempty"`
	Price float64 `json:"price,omitempty"`
	AvailableSizes pq.StringArray  `json:"availableSizes,omitempty" gorm:"type:text[]"`
	StartDate time.Time `json:"start_date,omitempty" gorm:"default:CURRENT_TIMESTAMP"`
	EndDate *time.Time `json:"end_date,omitempty"`
	HistoryID int64 `json:"history_id"`
}

func NewProduct(title string, description string, image string, price float64, availableSizes pq.StringArray) (*Product, error) {
	p := &Product{
		Title: title,
		Description: &description,
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

func (u *Product) String() string {
	s, err := u.PrettyPrint()
	if err != nil {
		return fmt.Sprintf("Error: %s", err.Error())
	}
	return s
}

func (u *Product) PrettyPrint() (string, error) {
	b, err := json.MarshalIndent(u, "", "  ")
	if err != nil {
		return "", err
	}
	return string(b), nil
}