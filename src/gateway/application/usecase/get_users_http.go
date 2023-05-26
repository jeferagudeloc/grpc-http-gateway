package usecase

import (
	"context"
	"time"

	"github.com/jeferagudeloc/grpc-http-gateway/src/gateway/application/adapter/logger"
	"github.com/jeferagudeloc/grpc-http-gateway/src/gateway/domain"
	"github.com/jeferagudeloc/grpc-http-gateway/src/gateway/infrastructure/client"
)

type (
	GetUsersHttpUseCase interface {
		Execute(context.Context) ([]domain.User, error)
	}

	GetUserHttpInteractor struct {
		ctxTimeout time.Duration
		logger     logger.Logger
		httpClient client.HttpClient
	}
)

func NewGetUserHttpInteractor(
	t time.Duration,
	l logger.Logger,
	httpClient client.HttpClient,
) GetUsersHttpUseCase {
	return GetUserHttpInteractor{
		ctxTimeout: t,
		logger:     l,
		httpClient: httpClient,
	}
}

func (a GetUserHttpInteractor) Execute(ctx context.Context) ([]domain.User, error) {

	output, err := a.httpClient.GetUsersFromBackend()
	if err != nil {
		return nil, err
	}

	return output, nil
}
