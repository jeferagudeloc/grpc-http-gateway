package entity

import (
	"github.com/jeferagudeloc/grpc-http-gateway/src/server/domain"
	"gorm.io/gorm"
)

type (
	Order struct {
		gorm.Model
		ID           string `gorm:"primary_key; unique;"`
		OrderType    string `gorm:"column:orderType"`
		Store        string
		Address      string
		CreationDate string `gorm:"column:creationDate" `
		Status       string
	}
)

func (Order) TableName() string {
	return "order"
}

func (o *Order) ToDomain() *domain.Order {
	return &domain.Order{
		ID:           o.ID,
		OrderType:    o.OrderType,
		Store:        o.Store,
		Address:      o.Address,
		CreationDate: o.CreationDate,
		Status:       o.Status,
	}
}

func ToOrdersDomainList(orders []Order) []domain.Order {
	domainOrders := make([]domain.Order, len(orders))
	for i, o := range orders {
		domainOrders[i] = *o.ToDomain()
	}
	return domainOrders
}
