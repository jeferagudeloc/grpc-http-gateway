package usecase

import (
	"context"

	"github.com/jeferagudeloc/grpc-http-gateway/src/server/domain"
	"github.com/sirupsen/logrus"
)

type (
	FindUsersUseCase interface {
		Execute(context.Context) ([]domain.User, error)
	}

	FindUsersInteractor struct {
		repo domain.UserRepository
	}
)

func NewFindUsersInteractor(
	repo domain.UserRepository,
) FindUsersUseCase {
	return FindUsersInteractor{
		repo: repo,
	}
}

func (a FindUsersInteractor) Execute(ctx context.Context) ([]domain.User, error) {
	logrus.Info("find users")
	output, err := a.repo.GetUsers()
	if err != nil {
		return nil, err
	}

	return output, nil
}
