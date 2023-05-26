package repository

import "github.com/jeferagudeloc/grpc-http-gateway/src/server/domain"

type SQL interface {
	GetOrders() ([]domain.Order, error)
	GetUsers() ([]domain.User, error)
}
