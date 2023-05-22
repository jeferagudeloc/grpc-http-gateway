package usecase

import (
	"context"
	"time"

	"github.com/jeferagudeloc/grpc-http-gateway/src/gateway/application/adapter/logger"
	"github.com/jeferagudeloc/grpc-http-gateway/src/gateway/domain"
	"github.com/jeferagudeloc/grpc-http-gateway/src/server/application/adapter/grpc/order"
	"google.golang.org/grpc"
)

type (
	CreateOrderRequest struct {
		OrderNumber int64 `json:"orderNumber"`
	}

	CreateOrderUseCase interface {
		Execute(context.Context, domain.OrderRequest) (CreateOrderOutput, error)
	}

	CreateOrderPresenter interface {
		Output(domain.Order) CreateOrderOutput
	}

	CreateOrderOutput struct {
		OrderNumber int64 `json:"orderNumber"`
	}

	CreateOrderInteractor struct {
		presenter  CreateOrderPresenter
		ctxTimeout time.Duration
		logger     logger.Logger
		grpcClient *grpc.ClientConn
	}
)

func NewCreateOrderInteractor(
	presenter CreateOrderPresenter,
	t time.Duration,
	l logger.Logger,
	grpcClient *grpc.ClientConn,
) CreateOrderUseCase {
	return CreateOrderInteractor{
		presenter:  presenter,
		ctxTimeout: t,
		logger:     l,
		grpcClient: grpcClient,
	}
}

func (a CreateOrderInteractor) Execute(ctx context.Context, orderRequest domain.OrderRequest) (CreateOrderOutput, error) {

	orderClient := order.NewOrderHandlerClient(a.grpcClient)

	orderProto := order.CreateOrderRequest{
		OrderNumber: orderRequest.OrderNumber,
	}

	a.logger.Infof("send order to server: Order Number[%v]", orderProto.OrderNumber)
	response, err := orderClient.CreateOrder(context.Background(), &orderProto)

	if err != nil {
		a.logger.Fatalln("Error when trying to say hello: %v", err)
	}

	a.logger.Infof("response from server: %v", response)

	return a.presenter.Output(domain.Order{
		OrderNumber: response.ID,
	}), nil
}
