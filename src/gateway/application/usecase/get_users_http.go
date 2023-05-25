package usecase

import (
	"context"
	"time"

	"github.com/jeferagudeloc/grpc-http-gateway/src/gateway/application/adapter/logger"
	"github.com/jeferagudeloc/grpc-http-gateway/src/gateway/domain"
	"google.golang.org/grpc"
)

type (
	GetUsersHttpUseCase interface {
		Execute(context.Context) ([]domain.User, error)
	}

	GetUserHttpInteractor struct {
		ctxTimeout time.Duration
		logger     logger.Logger
		grpcClient *grpc.ClientConn
	}
)

func NewGetUserHttpInteractor(
	t time.Duration,
	l logger.Logger,
	grpcClient *grpc.ClientConn,
) GetUsersHttpUseCase {
	return GetUserInteractor{
		ctxTimeout: t,
		logger:     l,
		grpcClient: grpcClient,
	}
}

func (a GetUserHttpInteractor) Execute(ctx context.Context) ([]domain.User, error) {
	return nil, nil
}
