package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PaymentMethod struct {
	PaymentMethodID string    `gorm:"primaryKey;unique;not null;type:varchar(64)"`
	MethodName      string    `gorm:"type:varchar(70);not null"`
	Code            string    `gorm:"type:varchar(10);not null"`
	CreatedDate     time.Time `gorm:"type:datetime;not null;autoCreateTime"`
}

func (pm *PaymentMethod) BeforeCreate(tx *gorm.DB) (err error) {
	pm.PaymentMethodID = uuid.New().String()
	return
}
