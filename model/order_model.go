package model

import (
	helper "go-micro/utils"
	"time"

	"gorm.io/gorm"
)

type Order struct {
	ID          string    `gorm:"primaryKey" json:"id"`
	ProductID   string    `json:"product_id"`
	ProductName string    `json:"product_name"`
	OrderBy     string    `json:"order_by"`
	Date        time.Time `json:"order_date"`
	CreatedAt   time.Time `json:"created_at" gorm:"string"`
}

type OrderRepository interface {
	Insert(order *Order) error
	Read(order *Order) (*Order, error)
	// Update(order *Order) error
	Delete(order *Order) error
}

type orderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) OrderRepository {
	return &orderRepository{db}
}

func (o *orderRepository) Insert(order *Order) error {
	if order.ID == "" {
		order.ID = helper.NewUUID()
	}
	return o.db.Save(order).Error
}
func (o *orderRepository) Read(order *Order) (*Order, error) {
	if err := o.db.Model(Order{ID: order.ID}).First(&order); err != nil {
		return nil, err.Error
	}
	return order, nil
}

// func (o *orderRepository) Update(order *Order) error {
// 	return o.db.Save(&order).Error
// }

func (o *orderRepository) Delete(order *Order) error {
	return o.db.Delete(&order).Error
}
