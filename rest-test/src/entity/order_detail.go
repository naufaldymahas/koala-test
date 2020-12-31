package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type OrderDetail struct {
	OrderDetailID string    `json:"order_detail_id" gorm:"primaryKey;unique;not null;type:varchar(64)"`
	OrderID       string    `json:"order_id" gorm:"type:varchar(64);not null"`
	ProductID     string    `json:"product_id" gorm:"type:varchar(64);not null" validate:"required,uuid4"`
	QTY           int32     `json:"qty" gorm:"type:not null" validate:"required,gte=1"`
	CreatedDate   time.Time `json:"created_date" gorm:"type:datetime;not null;autoCreateTime"`
}

func (od *OrderDetail) BeforeCreate(tx *gorm.DB) (err error) {
	od.OrderDetailID = uuid.New().String()
	return
}
