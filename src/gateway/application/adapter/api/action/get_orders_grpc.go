package action

import (
	"net/http"

	"github.com/jeferagudeloc/grpc-http-gateway/src/gateway/application/adapter/api/response"
	"github.com/jeferagudeloc/grpc-http-gateway/src/gateway/application/adapter/logger"
	"github.com/jeferagudeloc/grpc-http-gateway/src/gateway/application/adapter/logging"
	"github.com/jeferagudeloc/grpc-http-gateway/src/gateway/application/usecase"
)

type GetOrdersGrpcAction struct {
	uc  usecase.GetOrdersGrpcUseCase
	log logger.Logger
}

func NewGetOrdersGrpcAction(uc usecase.GetOrdersGrpcUseCase, log logger.Logger) GetOrdersGrpcAction {
	return GetOrdersGrpcAction{
		uc:  uc,
		log: log,
	}
}

func (goga GetOrdersGrpcAction) Execute(w http.ResponseWriter, r *http.Request) {
	const logKey = "get_orders_grpc_action"

	output, err := goga.uc.Execute(r.Context())
	if err != nil {
		switch err {
		default:
			logging.NewError(
				goga.log,
				err,
				logKey,
				http.StatusInternalServerError,
			).Log("error getting orders")

			response.NewError(err, http.StatusInternalServerError).Send(w)
			return
		}
	}
	logging.NewInfo(goga.log, logKey, http.StatusOK).Log("getting orders successfully")

	response.NewSuccess(output, http.StatusOK).Send(w)
}
