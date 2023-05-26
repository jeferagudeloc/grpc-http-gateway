package order

import (
	"context"

	"github.com/jeferagudeloc/grpc-http-gateway/src/server/application/usecase"
)

type Server struct {
	findOrdersUseCase usecase.FindOrdersUseCase
}

func NewServer(findOrdersUseCase usecase.FindOrdersUseCase) *Server {
	return &Server{
		findOrdersUseCase: findOrdersUseCase,
	}
}

func (s *Server) GetOrders(context context.Context, getOrdersRequest *GetOrdersRequest) (*GetOrdersResponse, error) {

	orders, err := s.findOrdersUseCase.Execute(context)
	if err != nil {
		return nil, err
	}

	orderResponse := make([]*Order, 0)

	for _, o := range orders {
		orderResponse = append(orderResponse, &Order{
			Id:           o.ID,
			OrderType:    o.OrderType,
			Store:        o.Store,
			Address:      o.Address,
			CreationDate: o.CreationDate,
			Status:       o.Status,
		})
	}

	return &GetOrdersResponse{
		Orders: orderResponse,
	}, nil
}
