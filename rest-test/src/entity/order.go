package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Order struct {
	OrderID         string         `json:"order_id" gorm:"primaryKey;unique;not null;type:varchar(64)"`
	CustomerID      string         `json:"customer_id" gorm:"type:varchar(64);not null"`
	OrderNumber     string         `json:"order_number" gorm:"type:varchar(40);not null"`
	OrderDate       time.Time      `json:"order_date" gorm:"type:datetime;not null;autoCreateTime"`
	PaymentMethodID string         `json:"payment_method_id" gorm:"type:varchar(64);not null" validate:"required"`
	OrderDetails    []*OrderDetail `json:"order_details" gorm:"foreignKey:OrderID" validate:"required,dive"`
}

func (o *Order) BeforeCreate(tx *gorm.DB) (err error) {
	o.OrderID = uuid.New().String()

	for _, od := range o.OrderDetails {
		od.OrderID = o.OrderID
	}
	return
}
