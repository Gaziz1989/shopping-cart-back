package repositories

import (
	"time"

	"github.com/jinzhu/gorm"
	"landing-back/entities"
)

type ProductRepository interface {
	Delete(id int64) error
	Create(i *entities.Product) (*entities.Product, error)
	List() ([]entities.Product, error)
}

type ProductPSQL struct {
	DB *gorm.DB
}

// New FireStore repository
func NewProductPSQLRepository(db *gorm.DB) ProductRepository {
	return &ProductPSQL{
		DB: db,
	}
}

func (r *ProductPSQL) Delete(id int64) error {
	tx := r.DB.Begin()
	if tx.Error != nil {
		return tx.Error
	}
	err := tx.Table("public.products").Where("id = ?", id).Update("end_date", time.Now()).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	err = tx.Commit().Error
	if err != nil {
		tx.Rollback()
		return err
	}
	return nil 
}

func(r *ProductPSQL) Create(i *entities.Product) (*entities.Product, error) {
	tx := r.DB.Begin()
	if tx.Error != nil {
		return i, tx.Error
	}
	err := tx.Table("public.products").Create(i).Error
	if err != nil {
		tx.Rollback()
		return i, err
	}
	err = tx.Table("public.products").Model(i).Update("history_id", i.ID).Error
	if err != nil {
		tx.Rollback()
		return i, err
	}
	err = tx.Commit().Error
	if err != nil {
		tx.Rollback()
		return i, err
	}
	return i, err
}

func(r *ProductPSQL) List() ([]entities.Product, error) {
	products := []entities.Product{}
	err := r.DB.
		Table("public.products").
		Where(`end_date ISNULL`).
		Select(`id, 
		title, 
		description, 
		image, 
		price, 
		available_sizes,
		start_date,
		end_date,
		history_id`).
		Find(&products).
		Error
	return products, err
}
