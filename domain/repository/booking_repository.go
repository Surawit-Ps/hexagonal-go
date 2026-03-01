package repository

import "Hexa/domain/entity"

type BookingRepository interface {
	CreateBooking(booking *entity.Booking) error
	GetBookingByID(id uint) (*entity.Booking, error)
	UpdateBooking(booking *entity.Booking) error
	DeleteBooking(id uint) error
	GetBookingsByUserID(userID uint) ([]*entity.Booking, error)
}