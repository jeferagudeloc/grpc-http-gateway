package repository

import (
	"context"

	"github.com/jeferagudeloc/grpc-http-gateway/src/server/application/adapter/model"
)

type SQL interface {
	SaveOrder(context.Context, model.Order) (*model.Order, error)
}
