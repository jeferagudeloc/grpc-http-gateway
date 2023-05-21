package usecase

import (
	"context"
	"fmt"
	"time"

	"github.com/jeferagudeloc/grpc-http-gateway/src/server/domain"
	"github.com/sirupsen/logrus"
)

type (
	CreateOrderUseCase interface {
		Execute(context.Context, CreateOrderInput) (CreateOrderOutput, error)
	}

	CreateOrderInput struct {
		OrderNumber int64            `json:"orderNumber" validate:"required"`
		Products    []domain.Product `json:"products" validate:"required"`
	}

	CreateOrderPresenter interface {
		Output(bool) CreateOrderOutput
	}

	CreateOrderOutput struct {
		Saved bool `json:"saved"`
	}

	CreateOrderInteractor struct {
		repo       domain.OrderRepository
		presenter  CreateOrderPresenter
		ctxTimeout time.Duration
	}
)

func NewCreateOrderInteractor(
	repo domain.OrderRepository,
	presenter CreateOrderPresenter,
	t time.Duration,
) CreateOrderUseCase {
	return CreateOrderInteractor{
		repo:       repo,
		presenter:  presenter,
		ctxTimeout: t,
	}
}

func (a CreateOrderInteractor) Execute(ctx context.Context, order CreateOrderInput) (CreateOrderOutput, error) {

	orderDataToSave := domain.OrderData{
		OrderNumber: fmt.Sprint(order.OrderNumber),
		Products:    order.Products,
	}

	orderDataSaved, err := a.repo.SaveOrder(orderDataToSave)
	if err != nil {
		return a.presenter.Output(false), err
	}

	logrus.Info("orderDataSaved", orderDataSaved)

	return a.presenter.Output(true), nil
}
