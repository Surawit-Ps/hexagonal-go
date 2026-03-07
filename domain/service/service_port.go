package service

import (
	"Hexa/domain/entity"
)

type BookingServiceInterface interface {
	AddBooking(booking *entity.Booking) error
	GetBooking(id uint) (*entity.Booking, error)
	UpdateBooking(booking *entity.Booking) error
	DeleteBooking(id uint) error
	GetBookingsByUserID(userID uint) ([]*entity.Booking, error)
}

type UserServiceInterface interface {
	AddUser(user *entity.User) error
	GetUser(id uint) (*entity.User, error)
	UpdateUser(user *entity.User) error
	DeleteUser(id uint) error
	GetAllUsers() ([]*entity.User, error)
}

type RoomServiceInterface interface {
	AddRoom(room *entity.Room) error
	GetRoom(id uint) (*entity.Room, error)
	UpdateRoom(room *entity.Room) error
	DeleteRoom(id uint) error
	GetAllRooms() ([]*entity.Room, error)
}
