package action

import (
	"net/http"

	"github.com/jeferagudeloc/grpc-http-gateway/src/server/application/adapter/api/response"
	"github.com/jeferagudeloc/grpc-http-gateway/src/server/application/adapter/logger"
	"github.com/jeferagudeloc/grpc-http-gateway/src/server/application/adapter/logging"
	"github.com/jeferagudeloc/grpc-http-gateway/src/server/application/usecase"
)

type FindUsersAction struct {
	uc  usecase.FindUsersUseCase
	log logger.Logger
}

func NewFindUsersAction(uc usecase.FindUsersUseCase, log logger.Logger) FindUsersAction {
	return FindUsersAction{
		uc:  uc,
		log: log,
	}
}

func (fova FindUsersAction) Execute(w http.ResponseWriter, r *http.Request) {
	const logKey = "find_all_users"

	output, err := fova.uc.Execute(r.Context())
	if err != nil {
		switch err {
		default:
			logging.NewError(
				fova.log,
				err,
				logKey,
				http.StatusInternalServerError,
			).Log("error find all users")

			response.NewError(err, http.StatusInternalServerError).Send(w)
			return
		}
	}
	logging.NewInfo(fova.log, logKey, http.StatusOK).Log("find all users success")

	response.NewSuccess(output, http.StatusOK).Send(w)
}
