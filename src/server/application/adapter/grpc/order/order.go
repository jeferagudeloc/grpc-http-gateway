package order

import (
	"context"

	"github.com/jeferagudeloc/grpc-http-gateway/src/server/application/usecase"
)

type Server struct {
	createOrderUseCase usecase.CreateOrderUseCase
	getOrderUseCase    usecase.GetOrderUseCase
}

func NewServer(createOrderUseCase usecase.CreateOrderUseCase, getOrderUseCase usecase.GetOrderUseCase) *Server {
	return &Server{
		createOrderUseCase: createOrderUseCase,
		getOrderUseCase:    getOrderUseCase,
	}
}

func (s *Server) GetOrders(context context.Context, getOrdersRequest *GetOrdersRequest) (*GetOrdersResponse, error) {

	orders := []*Order{
		{
			Id:           "1",
			OrderType:    "Type A",
			Store:        "Store A",
			Address:      "Address A",
			CreationDate: "2022-01-01",
			Status:       "Pending",
		},
		{
			Id:           "2",
			OrderType:    "Type B",
			Store:        "Store B",
			Address:      "Address B",
			CreationDate: "2022-02-02",
			Status:       "Completed",
		},
	}

	return &GetOrdersResponse{
		Orders: orders,
	}, nil
}
