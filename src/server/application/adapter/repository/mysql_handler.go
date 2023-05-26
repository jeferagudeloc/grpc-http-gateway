package repository

import "github.com/jeferagudeloc/grpc-http-gateway/src/server/domain"

type MysqlSQL struct {
	db SQL
}

func NewMysqlSQL(db SQL) MysqlSQL {
	return MysqlSQL{
		db: db,
	}
}

func (m MysqlSQL) GetOrders() ([]domain.Order, error) {
	return m.db.GetOrders()
}

func (m MysqlSQL) GetUsers() ([]domain.User, error) {
	return m.db.GetUsers()
}
