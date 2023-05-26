package grpc

import (
	"errors"

	"github.com/jeferagudeloc/grpc-http-gateway/src/server/application/adapter/logger"
	"github.com/jeferagudeloc/grpc-http-gateway/src/server/application/adapter/repository"
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
	sql repository.SQL,
) (Server, error) {
	switch instance {
	case InstanceGRPC:
		return NewGRPCServer(log, sql), nil
	default:
		return nil, errInvalidServerInstance
	}
}
