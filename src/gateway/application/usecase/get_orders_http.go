package usecase

import (
	"context"
	"time"

	"github.com/jeferagudeloc/grpc-http-gateway/src/gateway/application/adapter/logger"
	"github.com/jeferagudeloc/grpc-http-gateway/src/gateway/domain"
	"github.com/jeferagudeloc/grpc-http-gateway/src/gateway/infrastructure/client"
)

type (
	GetOrdersHttpUseCase interface {
		Execute(context.Context) ([]domain.Order, error)
	}

	GetOrdersHttpInteractor struct {
		ctxTimeout time.Duration
		logger     logger.Logger
		httpClient client.HttpClient
	}
)

func NewGetOrdersHttpInteractor(
	t time.Duration,
	l logger.Logger,
	httpClient client.HttpClient,
) GetOrdersHttpUseCase {
	return GetOrdersHttpInteractor{
		ctxTimeout: t,
		logger:     l,
		httpClient: httpClient,
	}
}

func (a GetOrdersHttpInteractor) Execute(ctx context.Context) ([]domain.Order, error) {

	output, err := a.httpClient.GetOrdersFromBackend()
	if err != nil {
		return nil, err
	}

	return output, nil
}
