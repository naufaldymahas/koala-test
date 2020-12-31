package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Customer struct {
	CustomerID   string    `gorm:"primaryKey;unique;not null;type:varchar(64)"`
	CustomerName string    `gorm:"type:varchar(80);not null"`
	Email        string    `gorm:"type:varchar(50);not null"`
	PhoneNumber  string    `gorm:"type:varchar(20);not null"`
	DOB          time.Time `gorm:"column:dob;type:date;not null"`
	Sex          bool      `gorm:"not null"`
	Salt         string    `gorm:"type:varchar(80);not null"`
	Password     string    `gorm:"type:varchar(50);not null"`
	CreatedDate  time.Time `gorm:"type:datetime;not null;autoCreateTime"`
}

func (c *Customer) BeforeCreate(tx *gorm.DB) (err error) {
	c.CustomerID = uuid.New().String()
	return
}
