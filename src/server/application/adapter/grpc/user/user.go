package user

import (
	"context"
	"log"
)

type Server struct {
}

func (s *Server) CreateUser(ctx context.Context, user *CreateUserRequest) (*User, error) {
	log.Printf("Creating User: %v", user.Email)
	return &User{ID: 1}, nil
}
