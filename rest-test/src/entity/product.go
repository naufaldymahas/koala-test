package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Product struct {
	ProductID   string    `gorm:"primaryKey;unique;not null;type:varchar(64)"`
	ProductName string    `gorm:"type:varchar(80);not null"`
	BasePrice   int64     `gorm:"not null"`
	CreatedDate time.Time `gorm:"type:datetime;not null;autoCreateTime"`
}

func (p *Product) BeforeCreate(tx *gorm.DB) (err error) {
	p.ProductID = uuid.New().String()
	return
}
