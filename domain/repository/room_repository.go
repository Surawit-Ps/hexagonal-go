package repository

import (
	"Hexa/domain/entity"
)

type RoomRepository interface {
	CreateRoom(room *entity.Room) error
	GetRoomByID(id uint) (*entity.Room, error)
	UpdateRoom(room *entity.Room) error
	DeleteRoom(id uint) error
	GetAllRooms() ([]*entity.Room, error)
}