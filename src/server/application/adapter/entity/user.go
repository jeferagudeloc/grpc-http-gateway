package entity

import (
	"github.com/jeferagudeloc/grpc-http-gateway/src/server/domain"
	"gorm.io/gorm"
)

type (
	User struct {
		gorm.Model
		ID       string `gorm:"primary_key; unique;"`
		Name     string
		LastName string `gorm:"column:lastname" json:"lastname"`
		Email    string
		Status   string
		RoleID   string
		Role     Role
	}

	Role struct {
		gorm.Model
		ID          string `gorm:"primary_key; unique;"`
		Name        string
		Permissions string `json:"permissions"`
	}
)

func (User) TableName() string {
	return "user"
}

func (Role) TableName() string {
	return "role"
}

func (u *User) ToDomain() *domain.User {
	return &domain.User{
		Name:     u.Name,
		LastName: u.LastName,
		Email:    u.Email,
		Status:   u.Status,
		Role: domain.Role{
			Name:        u.Role.Name,
			Permissions: []string{u.Role.Permissions},
		},
	}
}

func ToUsersDomainList(users []User) []domain.User {
	domainUsers := make([]domain.User, len(users))
	for i, u := range users {
		domainUsers[i] = *u.ToDomain()
	}
	return domainUsers
}
