package user

import (
	"context"

	"github.com/jeferagudeloc/grpc-http-gateway/src/server/application/usecase"
)

type Server struct {
	findUsersUseCase usecase.FindUsersUseCase
}

func NewServer(findUsersUseCase usecase.FindUsersUseCase) *Server {
	return &Server{
		findUsersUseCase: findUsersUseCase,
	}
}

func (s *Server) GetUsers(context.Context, *GetUsersRequest) (*GetUsersResponse, error) {

	orders, err := s.findUsersUseCase.Execute(context.Background())

	if err != nil {
		return nil, err
	}

	orderResponse := make([]*User, 0)

	for _, o := range orders {
		orderResponse = append(orderResponse, &User{
			Name:     o.Name,
			Lastname: o.LastName,
			Email:    o.Email,
			Role: &Role{
				Name:        o.Role.Name,
				Permissions: nil,
			},
			Status: o.Status,
		})
	}

	return &GetUsersResponse{
		Users: orderResponse,
	}, nil
}
