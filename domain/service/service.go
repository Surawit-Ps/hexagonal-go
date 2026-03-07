package service

import (
	"Hexa/domain/repository"
	"Hexa/domain/entity"
)

type BookingService struct {
	Repo repository.BookingRepository
}

func NewBookingService(repo repository.BookingRepository) *BookingService {
	return &BookingService{Repo: repo}
}

func (s *BookingService) AddBooking(booking *entity.Booking) error {
	return s.Repo.CreateBooking(booking)
}

func (s *BookingService) GetBooking(id uint) (*entity.Booking, error) {
	return s.Repo.GetBookingByID(id)
}

func (s *BookingService) UpdateBooking(booking *entity.Booking) error {
	return s.Repo.UpdateBooking(booking)
}

func (s *BookingService) DeleteBooking(id uint) error {
	return s.Repo.DeleteBooking(id)
}

func (s *BookingService) GetBookingsByUserID(userID uint) ([]*entity.Booking, error) {
	return s.Repo.GetBookingsByUserID(userID)
}



