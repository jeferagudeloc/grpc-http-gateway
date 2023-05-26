package http

import (
	"errors"
	"time"

	"github.com/jeferagudeloc/grpc-http-gateway/src/server/application/adapter/logger"
	"github.com/jeferagudeloc/grpc-http-gateway/src/server/application/adapter/repository"
)

type Server interface {
	Listen()
}

type Port int64

var (
	errInvalidWebServerInstance = errors.New("invalid router server instance")
)

const (
	InstanceGorillaMux int = iota
)

func NewWebServerFactory(
	instance int,
	log logger.Logger,
	port Port,
	ctxTimeout time.Duration,
	dbSQL repository.SQL,
) (Server, error) {
	switch instance {
	case InstanceGorillaMux:
		return newGorillaMux(log, port, ctxTimeout, dbSQL), nil
	default:
		return nil, errInvalidWebServerInstance
	}
}
