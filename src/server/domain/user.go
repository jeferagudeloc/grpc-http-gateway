package domain

import "github.com/jeferagudeloc/grpc-http-gateway/src/server/application/adapter/model"

type (
	UserRepository interface {
		SaveUser(UserData) (model.Order, error)
	}

	UserData struct {
		Email    string
		Name     string
		LastName string
	}
)
