package action

import (
	"net/http"

	"github.com/jeferagudeloc/grpc-http-gateway/src/gateway/application/adapter/api/response"
	"github.com/jeferagudeloc/grpc-http-gateway/src/gateway/application/adapter/logger"
	"github.com/jeferagudeloc/grpc-http-gateway/src/gateway/application/adapter/logging"
	"github.com/jeferagudeloc/grpc-http-gateway/src/gateway/application/usecase"
)

type GetUsersGrpcAction struct {
	uc  usecase.GetUsersGrpcUseCase
	log logger.Logger
}

func NewGetUsersGrpcAction(uc usecase.GetUsersGrpcUseCase, log logger.Logger) GetUsersGrpcAction {
	return GetUsersGrpcAction{
		uc:  uc,
		log: log,
	}
}

func (goga GetUsersGrpcAction) Execute(w http.ResponseWriter, r *http.Request) {
	const logKey = "get_users_grpc_action"

	output, err := goga.uc.Execute(r.Context())
	if err != nil {
		switch err {
		default:
			logging.NewError(
				goga.log,
				err,
				logKey,
				http.StatusInternalServerError,
			).Log("error getting users")

			response.NewError(err, http.StatusInternalServerError).Send(w)
			return
		}
	}
	logging.NewInfo(goga.log, logKey, http.StatusOK).Log("getting users successfully")

	response.NewSuccess(output, http.StatusOK).Send(w)
}
