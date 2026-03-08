package service

import (
	"Hexa/domain/entity"
	"Hexa/domain/repository"
	"log"
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
	booking ,err := s.Repo.GetBookingByID(id)
	if err != nil {
		return nil, err
	}
	if booking == nil {
		log.Printf("Booking with ID %d not found", id)
		return nil, err
	}
	return booking, err
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



