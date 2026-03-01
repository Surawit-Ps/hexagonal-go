package repository

import (
	"Hexa/domain/entity"
)

type UserRepository interface {
	CreateUser(user *entity.User) error
	GetUserByID(id uint) (*entity.User, error)
	UpdateUser(user *entity.User) error
	DeleteUser(id uint) error
	GetAllUsers() ([]*entity.User, error)
}