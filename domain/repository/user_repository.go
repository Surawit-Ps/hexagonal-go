package repository

import (
	"Hexa/domain/entity"
)

type UserRepository interface {
	CreateUser(user *entity.User) error
	GetUserByID(id string) (*entity.User, error)
	UpdateUser(user *entity.User) error
	DeleteUser(id string) error
	GetAllUsers() ([]*entity.User, error)
}