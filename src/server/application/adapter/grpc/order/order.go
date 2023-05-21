package order

import (
	"context"
	"fmt"
	"log"

	"github.com/jeferagudeloc/grpc-http-gateway/src/server/application/usecase"
	"github.com/jeferagudeloc/grpc-http-gateway/src/server/domain"
)

type Server struct {
	createOrderUseCase usecase.CreateOrderUseCase
	getOrderUseCase    usecase.GetOrderUseCase
}

func (s *Server) CreateOrder(ctx context.Context, req *CreateOrderRequest) (*Order, error) {
	output, err := s.createOrderUseCase.Execute(ctx, usecase.CreateOrderInput{
		OrderNumber: req.OrderNumber,
	})

	if err != nil {
		log.Fatalf("there was an error creating order: %v", err)
	}

	return &Order{Content: fmt.Sprint(output.Saved)}, nil
}

func (s *Server) GetOrder(ctx context.Context, req *GetOrderRequest) (*Order, error) {

	output, err := s.getOrderUseCase.Execute(ctx, usecase.GetOrderInput{
		OrderNumber: fmt.Sprint(req.OrderNumber),
		Products:    make([]domain.Product, 0),
	})

	if err != nil {
		log.Fatalf("there was an error creating order: %v", err)
	}

	return &Order{Content: fmt.Sprint(output.Saved)}, nil
}
