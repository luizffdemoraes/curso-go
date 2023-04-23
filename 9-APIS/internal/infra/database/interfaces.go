package database

import "github.com/devfullcycle/goexpert/9-APIS/internal/entity"

type UserInterface interface {
	Create(user *entity.User) error
	FindByEmail(email string) (*entity.User, error)
}
