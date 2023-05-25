package usecase

import (
	"context"
	"time"

	"github.com/jeferagudeloc/grpc-http-gateway/src/gateway/application/adapter/logger"
	"github.com/jeferagudeloc/grpc-http-gateway/src/gateway/domain"
	"google.golang.org/grpc"
)

type (
	GetUsersGrpcUseCase interface {
		Execute(context.Context) ([]domain.User, error)
	}

	GetUserInteractor struct {
		ctxTimeout time.Duration
		logger     logger.Logger
		grpcClient *grpc.ClientConn
	}
)

func NewGetUserGrpcInteractor(
	t time.Duration,
	l logger.Logger,
	grpcClient *grpc.ClientConn,
) GetUsersGrpcUseCase {
	return GetUserInteractor{
		ctxTimeout: t,
		logger:     l,
		grpcClient: grpcClient,
	}
}

func (a GetUserInteractor) Execute(ctx context.Context) ([]domain.User, error) {
	return nil, nil
}
