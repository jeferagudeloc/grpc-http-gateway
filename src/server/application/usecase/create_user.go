package usecase

import (
	"context"
	"time"

	"github.com/jeferagudeloc/grpc-http-gateway/src/server/domain"
	"github.com/sirupsen/logrus"
)

type (
	CreateUserUseCase interface {
		Execute(context.Context, CreateUserInput) (CreateUserOutput, error)
	}

	CreateUserInput struct {
		Email    string `json:"email" validate:"required"`
		Name     string `json:"name" validate:"required"`
		LastName string `json:"lastName" validate:"required"`
	}

	CreateUserPresenter interface {
		Output(bool) CreateUserOutput
	}

	CreateUserOutput struct {
		Saved bool `json:"saved"`
	}

	CreateUserInteractor struct {
		repo       domain.UserRepository
		presenter  CreateUserPresenter
		ctxTimeout time.Duration
	}
)

func NewCreateUserInteractor(
	repo domain.UserRepository,
	presenter CreateUserPresenter,
	t time.Duration,
) CreateUserUseCase {
	return CreateUserInteractor{
		repo:       repo,
		presenter:  presenter,
		ctxTimeout: t,
	}
}

func (a CreateUserInteractor) Execute(ctx context.Context, user CreateUserInput) (CreateUserOutput, error) {

	userDataToSave := domain.UserData{
		Email:    user.Email,
		Name:     user.Name,
		LastName: user.LastName,
	}

	userDataSaved, err := a.repo.SaveUser(userDataToSave)
	if err != nil {
		return a.presenter.Output(false), err
	}

	logrus.Info("userDataSaved", userDataSaved)

	return a.presenter.Output(true), nil
}
