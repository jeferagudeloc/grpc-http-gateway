package user

import (
	"context"

	"github.com/google/uuid"
)

type Server struct {
}

func (s *Server) GetUsers(context.Context, *GetUsersRequest) (*GetUsersResponse, error) {
	users := []*User{
		{
			Id:       uuid.New().String(),
			Name:     "John",
			Lastname: "Doe",
			Email:    "john.doe@example.com",
			Status:   "Active",
			Role: &Role{
				Name:        "Admin",
				Permissions: []string{"read", "write"},
			},
		},
		{
			Id:       uuid.New().String(),
			Name:     "Jane",
			Lastname: "Smith",
			Email:    "jane.smith@example.com",
			Status:   "Inactive",
			Role: &Role{
				Name:        "User",
				Permissions: []string{"read"},
			},
		},
		{
			Id:       uuid.New().String(),
			Name:     "Johan",
			Lastname: "Bell",
			Email:    "johan.bell@example.com",
			Status:   "Active",
			Role: &Role{
				Name:        "User",
				Permissions: []string{"read"},
			},
		},
	}
	return &GetUsersResponse{
		Users: users,
	}, nil
}
