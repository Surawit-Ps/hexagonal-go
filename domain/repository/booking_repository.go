package repository

import "Hexa/domain/entity"

type BookingRepository interface {
	CreateBooking(booking *entity.Booking) error
	GetBookingByID(id string) (*entity.Booking, error)
	UpdateBooking(booking *entity.Booking) error
	DeleteBooking(id string) error
	GetBookingsByUserID(userID string) ([]*entity.Booking, error)
}