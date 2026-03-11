package service

import (
	"Hexa/domain/entity"
)

type BookingServiceInterface interface {
	AddBooking(booking *entity.Booking) error
	GetBooking(id string) (*entity.Booking, error)
	UpdateBooking(booking *entity.Booking) error
	DeleteBooking(id string) error
	GetBookingsByUserID(userID string) ([]*entity.Booking, error)
}

type UserServiceInterface interface {
	AddUser(user *entity.User) error
	GetUser(id string) (*entity.User, error)
	UpdateUser(user *entity.User) error
	DeleteUser(id string) error
	GetAllUsers() ([]*entity.User, error)
}

type RoomServiceInterface interface {
	AddRoom(room *entity.Room) error
	GetRoom(id string) (*entity.Room, error)
	UpdateRoom(room *entity.Room) error
	DeleteRoom(id string) error
	GetAllRooms() ([]*entity.Room, error)
}
