package action

import (
	"net/http"

	"github.com/jeferagudeloc/grpc-http-gateway/src/server/application/adapter/api/response"
	"github.com/jeferagudeloc/grpc-http-gateway/src/server/application/adapter/logger"
	"github.com/jeferagudeloc/grpc-http-gateway/src/server/application/adapter/logging"
	"github.com/jeferagudeloc/grpc-http-gateway/src/server/application/usecase"
)

type FindOrdersAction struct {
	uc  usecase.FindOrdersUseCase
	log logger.Logger
}

func NewFindOrdersAction(uc usecase.FindOrdersUseCase, log logger.Logger) FindOrdersAction {
	return FindOrdersAction{
		uc:  uc,
		log: log,
	}
}

func (fova FindOrdersAction) Execute(w http.ResponseWriter, r *http.Request) {
	const logKey = "find_all_orders"

	output, err := fova.uc.Execute(r.Context())
	if err != nil {
		switch err {
		default:
			logging.NewError(
				fova.log,
				err,
				logKey,
				http.StatusInternalServerError,
			).Log("error find all orders")

			response.NewError(err, http.StatusInternalServerError).Send(w)
			return
		}
	}
	logging.NewInfo(fova.log, logKey, http.StatusOK).Log("find all orders success")

	response.NewSuccess(output, http.StatusOK).Send(w)
}
