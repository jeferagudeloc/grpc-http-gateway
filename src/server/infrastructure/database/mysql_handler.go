package database

import (
	"fmt"

	"github.com/jeferagudeloc/grpc-http-gateway/src/server/application/adapter/entity"
	"github.com/jeferagudeloc/grpc-http-gateway/src/server/domain"
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
		return nil, err
	}

	err = Migration(db)
	if err != nil {
		fmt.Errorf("there was a migration error", err)
	}

	sqlDB.SetMaxIdleConns(5)
	sqlDB.SetMaxOpenConns(10)

	return &MysqlHandler{db}, nil
}

func (mysqlHandler MysqlHandler) GetOrders() ([]domain.Order, error) {
	var orders []entity.Order
	if err := mysqlHandler.db.Find(&orders).Error; err != nil {
		return nil, err
	}
	return entity.ToOrdersDomainList(orders), nil
}

func (mysqlHandler MysqlHandler) GetUsers() ([]domain.User, error) {
	var users []entity.User
	err := mysqlHandler.db.Model(&entity.User{}).Preload("Role").Find(&users).Error
	if err != nil {
		return nil, err
	}
	return entity.ToUsersDomainList(users), nil

}
