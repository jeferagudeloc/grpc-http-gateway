package usecase

import (
	"context"
	"fmt"
	"time"

	"github.com/jeferagudeloc/grpc-http-gateway/src/gateway/application/adapter/logger"
	"github.com/jeferagudeloc/grpc-http-gateway/src/gateway/domain"
	"github.com/jeferagudeloc/grpc-http-gateway/src/server/application/adapter/grpc/order"
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

	output := make([]domain.Order, 0)
	orderClient := order.NewOrderHandlerClient(a.grpcClient)

	a.logger.Infof("get orders calling grpc client")
	response, err := orderClient.GetOrders(ctx, &order.GetOrdersRequest{})

	for _, o := range response.Orders {
		output = append(output, domain.Order{
			ID:           o.Id,
			OrderType:    o.OrderType,
			Store:        o.Store,
			Address:      o.Address,
			CreationDate: o.CreationDate,
			Status:       o.Status,
		})
	}

	fmt.Printf("output server: %v\n", len(output))

	if err != nil {
		a.logger.Fatalln("Error when trying to say hello: %v", err)
	}

	a.logger.Infof("response from server: %v", response)

	return output, nil
}
