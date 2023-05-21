package server

import (
	"errors"

	"github.com/jeferagudeloc/grpc-http-gateway/src/server/application/adapter/logger"
)

type Server interface {
	Listen()
}

type Port int64

var (
	errInvalidServerInstance = errors.New("invalid server instance")
)

const (
	InstanceGRPC int = iota
)

func NewGrpcServerFactory(
	instance int,
	log logger.Logger,
) (Server, error) {
	switch instance {
	case InstanceGRPC:
		return NewGRPCServer(log), nil
	default:
		return nil, errInvalidServerInstance
	}
}
