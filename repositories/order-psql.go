package repositories

import (
	"time"

	"github.com/lib/pq"

	"github.com/jinzhu/gorm"
	"landing-back/entities"
)

type OrderRepository interface {
	Delete(id int64) error
	Create(i *entities.Order) error
	List() ([]entities.Order, error)
}

type OrderPSQL struct {
	DB *gorm.DB
}

// New FireStore repository
func NewOrderPSQLRepository(db *gorm.DB) OrderRepository {
	return &OrderPSQL{
		DB: db,
	}
}

func (r *OrderPSQL) Delete(id int64) error {
	tx := r.DB.Begin()
	if tx.Error != nil {
		return tx.Error
	}
	err := tx.Table("public.orders").Where("id = ?", id).Update("end_date", time.Now()).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Table("public.order_products").Where("order_id = ?", id).Update("end_date", time.Now()).Error
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

func(r *OrderPSQL) Create(i *entities.Order) error {
	tx := r.DB.Begin()
	if tx.Error != nil {
		return tx.Error
	}
	err := tx.Table("public.orders").Create(i).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	err = tx.Table("public.orders").Model(i).Update("history_id", i.ID).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	for p := range i.OrderProducts {
		op := entities.OrderProducts{
			OrderID: i.HistoryID,
			ProductID: i.OrderProducts[p].HistoryID,
		}

		err := tx.Table("public.order_products").Create(&op).Error
		if err != nil {
			tx.Rollback()
			return err
		}
		err = tx.Table("public.order_products").Model(&op).Update("history_id", op.ID).Error
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	err = tx.Commit().Error
	if err != nil {
		tx.Rollback()
		return err
	}
	return err
}

func(r *OrderPSQL) List() ([]entities.Order, error) {
	orders := []entities.Order{}
	err := r.DB.
		Table("public.orders").
		Where(`end_date ISNULL`).
		Select(`id,
			email,
			name,
			address,
			total,
			start_date,
			end_date,
			history_id`).
		Find(&orders).
		Error
	if err != nil {
		return nil, err
	}
	for o := range orders {
		rows, err := r.DB.Raw(`SELECT 
			p.id, 
			p.title, 
			p.description, 
			p.image, 
			p.price, 
			p.available_sizes,
			p.start_date,
			p.end_date,
			p.history_id
			from public.order_products as op
			left join public.products as p
			on op.product_id = p.history_id
			where op.order_id = $1`,
			orders[o].HistoryID,
		).Rows()
		if err != nil {
			return nil, err
		}
		defer rows.Close()
		for rows.Next() {
			var ID int64
			var Title string
			var Description *string
			var Image string
			var Price float64
			var AvailableSizes pq.StringArray
			var StartDate time.Time
			var EndDate *time.Time
			var HistoryID int64
			rows.Scan(&ID,
				&Title,
				&Description,
				&Image,
				&Price,
				&AvailableSizes,
				&StartDate,
				&EndDate,
				&HistoryID)
			product := entities.Product{
				ID: ID,
				Title: Title,
				Description: Description,
				Image: Image,
				Price: Price,
				AvailableSizes: AvailableSizes,
				StartDate: StartDate,
				EndDate: EndDate,
				HistoryID: HistoryID,
			}
			orders[o].OrderProducts = append(orders[o].OrderProducts, product)
		}
	}
	return orders, err
}