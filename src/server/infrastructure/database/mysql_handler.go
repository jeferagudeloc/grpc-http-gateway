package database

import (
	"context"
	"errors"
	"fmt"

	"github.com/jeferagudeloc/grpc-http-gateway/src/server/application/adapter/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type IMysqlHandler interface {
	GetConnection() *gorm.DB
}

type MysqlHandler struct {
	db *gorm.DB
}

func NewMysqlHandler(c *config) (*MysqlHandler, error) {
	var dsn = fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		c.user,
		c.password,
		c.host,
		c.port,
		c.database,
	)

	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,
		DefaultStringSize:         256,
		DisableDatetimePrecision:  true,
		DontSupportRenameIndex:    true,
		DontSupportRenameColumn:   true,
		SkipInitializeWithVersion: false,
	}), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	sqlDB, err := db.DB()

	if err != nil {
		fmt.Errorf("there was an error creating database", err)
	}

	sqlDB.SetMaxIdleConns(5)
	sqlDB.SetMaxOpenConns(10)

	return &MysqlHandler{db}, nil
}

func (mysqlHandler MysqlHandler) SaveOrder(ctx context.Context, orderToSave model.Order) (*model.Order, error) {

	result := mysqlHandler.db.Create(&orderToSave)

	if result.Error != nil {
		return nil, errors.New("there was an error saving the order")
	}

	return &orderToSave, nil
}
