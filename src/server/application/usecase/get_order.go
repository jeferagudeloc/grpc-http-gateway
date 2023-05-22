package usecase

import (
	"context"
	"time"

	"github.com/jeferagudeloc/grpc-http-gateway/src/server/domain"
)

type (
	GetOrderUseCase interface {
		Execute(context.Context, GetOrderInput) (GetOrderOutput, error)
	}

	GetOrderInput struct {
		OrderNumber string `json:"orderNumber" validate:"required"`
	}

	GetOrderPresenter interface {
		Output(bool) GetOrderOutput
	}

	GetOrderOutput struct {
		Saved bool `json:"saved"`
	}

	GetOrderInteractor struct {
		repo       domain.OrderRepository
		presenter  GetOrderPresenter
		ctxTimeout time.Duration
	}
)

func NewGetOrderInteractor(
	repo domain.OrderRepository,
	presenter GetOrderPresenter,
	t time.Duration,
) GetOrderUseCase {
	return GetOrderInteractor{
		repo:       repo,
		presenter:  presenter,
		ctxTimeout: t,
	}
}

func (a GetOrderInteractor) Execute(ctx context.Context, order GetOrderInput) (GetOrderOutput, error) {
	return a.presenter.Output(true), nil
}
