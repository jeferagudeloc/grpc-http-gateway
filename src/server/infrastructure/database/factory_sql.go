package database

import (
	"errors"

	"github.com/jeferagudeloc/grpc-http-gateway/src/server/application/adapter/repository"
)

var (
	errInvalidSQLDatabaseInstance = errors.New("invalid sql db instance")
)

const (
	InstanceMysql int = iota
)

func NewDatabaseSQLFactory(instance int) (repository.SQL, error) {
	switch instance {
	case InstanceMysql:
		return NewMysqlHandler(newConfigMysql())
	default:
		return nil, errInvalidSQLDatabaseInstance
	}
}
