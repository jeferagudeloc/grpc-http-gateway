package action

import (
	"net/http"

	"github.com/jeferagudeloc/grpc-http-gateway/src/gateway/application/adapter/api/response"
	"github.com/jeferagudeloc/grpc-http-gateway/src/gateway/application/adapter/logger"
	"github.com/jeferagudeloc/grpc-http-gateway/src/gateway/application/adapter/logging"
	"github.com/jeferagudeloc/grpc-http-gateway/src/gateway/application/usecase"
)

type GetUsersHttpAction struct {
	uc  usecase.GetUsersHttpUseCase
	log logger.Logger
}

func NewGetUsersHttpAction(uc usecase.GetUsersHttpUseCase, log logger.Logger) GetUsersHttpAction {
	return GetUsersHttpAction{
		uc:  uc,
		log: log,
	}
}

func (goga GetUsersHttpAction) Execute(w http.ResponseWriter, r *http.Request) {
	const logKey = "get_users_http_action"

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
