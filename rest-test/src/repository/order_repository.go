package repository

import (
	"rest-test/src/entity"

	"gorm.io/gorm"
)

type OrderRepository struct {
	DB *gorm.DB
}

func ProvideOrderRepository(db *gorm.DB) OrderRepository {
	return OrderRepository{DB: db}
}

func (or *OrderRepository) CreateOrder(order *entity.Order) error {
	if err := or.DB.Create(order).Error; err != nil {
		return err
	}
	return nil
}
