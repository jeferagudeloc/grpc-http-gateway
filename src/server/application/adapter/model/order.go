package model

import "time"

type Order struct {
	ID           string `gorm:"primaryKey"`
	OrderNumber  string `gorm:"not null"`
	CreationDate time.Time
	UpdationDate time.Time
	Status       string
}

func (Order) TableName() string {
	return "order"
}
