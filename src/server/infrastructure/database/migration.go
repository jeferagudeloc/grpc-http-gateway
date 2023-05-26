package database

import (
	"github.com/google/uuid"
	"github.com/jeferagudeloc/grpc-http-gateway/src/server/application/adapter/entity"
	"gorm.io/gorm"
)

func Migration(db *gorm.DB) error {
	db.AutoMigrate(&entity.User{}, &entity.Role{}, &entity.Order{})

	roleId := uuid.NewString()
	err := db.Create(&entity.Role{ID: roleId, Name: "default", Permissions: "view-x"}).Error
	if err != nil {
		return err
	}

	for i := 1; i <= 200; i++ {

		err = db.Create(&entity.User{ID: uuid.NewString(), Name: "John", LastName: "Doe", Email: "john@example.com", Status: "active", RoleID: roleId}).Error
		if err != nil {
			return err
		}

		err = db.Create(&entity.Order{
			ID:           uuid.NewString(),
			OrderType:    "food",
			Store:        "chin",
			Address:      "barrancas",
			CreationDate: "2023-05-25",
			Status:       "loaded",
		}).Error
		if err != nil {
			return err
		}
	}
	return nil
}
