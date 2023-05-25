package usecase

import (
	"context"
	"time"

	"github.com/jeferagudeloc/grpc-http-gateway/src/gateway/application/adapter/logger"
	"github.com/jeferagudeloc/grpc-http-gateway/src/gateway/domain"
	"google.golang.org/grpc"
)

type (
	GetOrdersHttpUseCase interface {
		Execute(context.Context) ([]domain.Order, error)
	}

	GetOrdersHttpInteractor struct {
		ctxTimeout time.Duration
		logger     logger.Logger
		grpcClient *grpc.ClientConn
	}
)

func NewGetOrdersHttpInteractor(
	t time.Duration,
	l logger.Logger,
	grpcClient *grpc.ClientConn,
) GetOrdersHttpUseCase {
	return GetOrdersHttpInteractor{
		ctxTimeout: t,
		logger:     l,
		grpcClient: grpcClient,
	}
}

func (a GetOrdersHttpInteractor) Execute(ctx context.Context) ([]domain.Order, error) {
	return nil, nil
}
