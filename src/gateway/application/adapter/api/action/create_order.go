package action

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/jeferagudeloc/grpc-http-gateway/src/gateway/application/adapter/api/response"
	"github.com/jeferagudeloc/grpc-http-gateway/src/gateway/application/adapter/logger"
	"github.com/jeferagudeloc/grpc-http-gateway/src/gateway/application/adapter/logging"
	"github.com/jeferagudeloc/grpc-http-gateway/src/gateway/application/usecase"
	"github.com/jeferagudeloc/grpc-http-gateway/src/gateway/domain"
)

type CreateOrderAction struct {
	uc  usecase.CreateOrderUseCase
	log logger.Logger
}

func NewCreateOrderAction(uc usecase.CreateOrderUseCase, log logger.Logger) CreateOrderAction {
	return CreateOrderAction{
		uc:  uc,
		log: log,
	}
}

func (fova CreateOrderAction) Execute(w http.ResponseWriter, r *http.Request) {
	const logKey = "save_device_token"

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Unable to read request body", http.StatusBadRequest)
		return
	}

	var dt domain.OrderRequest
	err = json.Unmarshal(body, &dt)
	if err != nil {
		http.Error(w, "Unable to parse JSON data", http.StatusBadRequest)
		return
	}

	output, err := fova.uc.Execute(r.Context(), dt)
	if err != nil {
		switch err {
		default:
			logging.NewError(
				fova.log,
				err,
				logKey,
				http.StatusInternalServerError,
			).Log("error saving device token")

			response.NewError(err, http.StatusInternalServerError).Send(w)
			return
		}
	}
	logging.NewInfo(fova.log, logKey, http.StatusOK).Log("device token has been saved")

	response.NewSuccess(output, http.StatusOK).Send(w)
}
