package usecase

import (
	"context"
	"time"

	"github.com/jeferagudeloc/grpc-http-gateway/src/gateway/application/adapter/logger"
	"github.com/jeferagudeloc/grpc-http-gateway/src/gateway/domain"
	"google.golang.org/grpc"
)

type (
	GetOrdersGrpcUseCase interface {
		Execute(context.Context) ([]domain.Order, error)
	}

	GetOrdersGrpcInteractor struct {
		ctxTimeout time.Duration
		logger     logger.Logger
		grpcClient *grpc.ClientConn
	}
)

func NewGetOrdersGrpcInteractor(
	t time.Duration,
	l logger.Logger,
	grpcClient *grpc.ClientConn,
) GetOrdersGrpcUseCase {
	return GetOrdersGrpcInteractor{
		ctxTimeout: t,
		logger:     l,
		grpcClient: grpcClient,
	}
}

func (a GetOrdersGrpcInteractor) Execute(ctx context.Context) ([]domain.Order, error) {
	return nil, nil
}
