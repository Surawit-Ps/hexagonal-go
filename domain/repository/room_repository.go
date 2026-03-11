package repository

import (
	"Hexa/domain/entity"
)

type RoomRepository interface {
	CreateRoom(room *entity.Room) error
	GetRoomByID(id string) (*entity.Room, error)
	UpdateRoom(room *entity.Room) error
	DeleteRoom(id string) error
	GetAllRooms() ([]*entity.Room, error)
}