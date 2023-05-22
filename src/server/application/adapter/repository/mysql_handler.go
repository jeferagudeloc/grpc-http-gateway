package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/jeferagudeloc/grpc-http-gateway/src/server/application/adapter/model"
	"github.com/jeferagudeloc/grpc-http-gateway/src/server/domain"
	"github.com/pkg/errors"
)

type MysqlSQL struct {
	db SQL
}

func NewMysqlSQL(db SQL) MysqlSQL {
	return MysqlSQL{
		db: db,
	}
}

func (a MysqlSQL) SaveOrder(ctx context.Context, order domain.OrderData) (domain.OrderData, error) {

	orderEntity := model.Order{
		ID:           uuid.New().String(),
		OrderNumber:  order.OrderNumber,
		CreationDate: order.CreationDate,
		UpdationDate: order.UpdationDate,
		Status:       order.Status,
	}

	_, err := a.db.SaveOrder(
		ctx,
		orderEntity,
	)

	if err != nil {
		return order, errors.Wrap(err, "error creating account")
	}

	return order, nil
}
