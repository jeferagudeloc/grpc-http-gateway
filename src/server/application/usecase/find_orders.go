package usecase

import (
	"context"

	"github.com/jeferagudeloc/grpc-http-gateway/src/server/domain"
	"github.com/sirupsen/logrus"
)

type (
	FindOrdersUseCase interface {
		Execute(ctx context.Context) ([]domain.Order, error)
	}

	FindOrdersInteractor struct {
		repo domain.OrderRepository
	}
)

func NewFindOrdersInteractor(
	repo domain.OrderRepository,
) FindOrdersUseCase {
	return FindOrdersInteractor{
		repo: repo,
	}
}

func (a FindOrdersInteractor) Execute(ctx context.Context) ([]domain.Order, error) {
	logrus.Info("find all orders")
	output, err := a.repo.GetOrders()

	if err != nil {
		return nil, err
	}
	return output, nil
}
