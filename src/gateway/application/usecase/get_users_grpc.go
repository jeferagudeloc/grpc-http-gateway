package usecase

import (
	"context"
	"time"

	"github.com/jeferagudeloc/grpc-http-gateway/src/gateway/application/adapter/logger"
	"github.com/jeferagudeloc/grpc-http-gateway/src/gateway/domain"
	"github.com/jeferagudeloc/grpc-http-gateway/src/server/application/adapter/grpc/user"
	"google.golang.org/grpc"
)

type (
	GetUsersGrpcUseCase interface {
		Execute(context.Context) ([]domain.User, error)
	}

	GetUserGrpcInteractor struct {
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
	return GetUserGrpcInteractor{
		ctxTimeout: t,
		logger:     l,
		grpcClient: grpcClient,
	}
}

func (a GetUserGrpcInteractor) Execute(ctx context.Context) ([]domain.User, error) {

	output := make([]domain.User, 0)
	userClient := user.NewUserHandlerClient(a.grpcClient)

	a.logger.Infof("get orders calling grpc client")
	response, err := userClient.GetUsers(ctx, &user.GetUsersRequest{})

	for _, o := range response.Users {
		output = append(output, domain.User{
			ID:       o.Id,
			Name:     o.Name,
			LastName: o.Lastname,
			Email:    o.Email,
			Status:   o.Status,
			Role:     mapRolesOfUser(o.Role),
		})
	}

	if err != nil {
		a.logger.Fatalln("Error when trying to say hello: %v", err)
	}

	return output, nil
}

func mapRolesOfUser(grpcRole *user.Role) domain.Role {
	return domain.Role{
		Name:        grpcRole.Name,
		Permissions: grpcRole.Permissions,
	}
}
