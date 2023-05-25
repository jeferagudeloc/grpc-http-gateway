package action

import (
	"net/http"

	"github.com/jeferagudeloc/grpc-http-gateway/src/gateway/application/adapter/api/response"
	"github.com/jeferagudeloc/grpc-http-gateway/src/gateway/application/adapter/logger"
	"github.com/jeferagudeloc/grpc-http-gateway/src/gateway/application/adapter/logging"
	"github.com/jeferagudeloc/grpc-http-gateway/src/gateway/application/usecase"
)

type GetOrdersHttpAction struct {
	uc  usecase.GetOrdersHttpUseCase
	log logger.Logger
}

func NewGetOrdersHttpAction(uc usecase.GetOrdersHttpUseCase, log logger.Logger) GetOrdersHttpAction {
	return GetOrdersHttpAction{
		uc:  uc,
		log: log,
	}
}

func (goga GetOrdersHttpAction) Execute(w http.ResponseWriter, r *http.Request) {
	const logKey = "get_orders_http_action"

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
