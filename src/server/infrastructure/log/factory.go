package log

import (
	"errors"

	"github.com/jeferagudeloc/grpc-http-gateway/src/server/application/adapter/logger"
)

const (
	InstanceLogrusLogger int = iota
)

var (
	errInvalidLoggerInstance = errors.New("invalid log instance")
)

func NewLoggerFactory(instance int) (logger.Logger, error) {
	switch instance {
	case InstanceLogrusLogger:
		return NewLogrusLogger(), nil
	default:
		return nil, errInvalidLoggerInstance
	}
}
