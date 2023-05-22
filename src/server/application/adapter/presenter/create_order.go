package presenter

import "github.com/jeferagudeloc/grpc-http-gateway/src/server/application/usecase"

type createOrderPresenter struct{}

func NewCreateOrderPresenter() usecase.CreateOrderPresenter {
	return createOrderPresenter{}
}

func (a createOrderPresenter) Output(saved bool) usecase.CreateOrderOutput {
	return usecase.CreateOrderOutput{
		Saved: saved,
	}
}
