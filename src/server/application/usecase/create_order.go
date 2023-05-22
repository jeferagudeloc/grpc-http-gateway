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
		OrderNumber int64 `json:"orderNumber" validate:"required"`
	}

	CreateOrderPresenter interface {
		Output(bool) CreateOrderOutput
	}

	CreateOrderOutput struct {
		Saved bool `json:"saved"`
	}

	CreateOrderInteractor struct {
		repo      domain.OrderRepository
		presenter CreateOrderPresenter
	}
)

func NewCreateOrderInteractor(
	repo domain.OrderRepository,
	presenter CreateOrderPresenter,
) CreateOrderUseCase {
	return CreateOrderInteractor{
		repo:      repo,
		presenter: presenter,
	}
}

func (a CreateOrderInteractor) Execute(ctx context.Context, order CreateOrderInput) (CreateOrderOutput, error) {

	orderDataToSave := domain.OrderData{
		OrderNumber:  fmt.Sprint(order.OrderNumber),
		CreationDate: time.Now(),
		UpdationDate: time.Now(),
		Status:       "idle",
	}

	orderDataSaved, err := a.repo.SaveOrder(ctx, orderDataToSave)
	if err != nil {
		return a.presenter.Output(false), err
	}

	logrus.Info("orderDataSaved", orderDataSaved)

	return a.presenter.Output(true), nil
}
