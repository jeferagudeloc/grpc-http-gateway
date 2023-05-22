package presenter

import (
	"github.com/jeferagudeloc/grpc-http-gateway/src/gateway/application/usecase"
	"github.com/jeferagudeloc/grpc-http-gateway/src/gateway/domain"
)

type createOrderPresenter struct{}

func NewCreateOrderPresenter() usecase.CreateOrderPresenter {
	return createOrderPresenter{}
}

func (a createOrderPresenter) Output(domain domain.Order) usecase.CreateOrderOutput {
	return usecase.CreateOrderOutput{
		OrderNumber: domain.OrderNumber,
	}
}
