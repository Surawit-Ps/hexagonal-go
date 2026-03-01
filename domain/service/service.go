package service

import (
	"errors"
	"time"
	"Hexa/domain/entity"
	"Hexa/domain/repository"
)

type BookingService struct {
	bookingRepo repository.BookingRepository
	roomRepo    repository.RoomRepository
	userRepo    repository.UserRepository
}

func NewBookingService(bookingRepo repository.BookingRepository, roomRepo repository.RoomRepository, userRepo repository.UserRepository) *BookingService {
	return &BookingService{
		bookingRepo: bookingRepo,
		roomRepo:    roomRepo,
		userRepo:    userRepo,
	}
}

func (s *BookingService) CreateBooking(userID uint, roomID uint, startTime time.Time, endTime time.Time) error {
	user, err := s.userRepo.GetUserByID(userID)
	if err != nil {
		return errors.New("user not found")
	}
	
	room, err := s.roomRepo.GetRoomByID(roomID)
	if err != nil {
		return errors.New("room not found")
	}

	booking := &entity.Booking{
		UserID:      user.ID,
		RoomID:      room.ID,
		StartTime: startTime,
		EndTime:   endTime,
	}

	return s.bookingRepo.CreateBooking(booking)
}

func (s *BookingService) GetBookingByID(id uint) (*entity.Booking, error) {
	return s.bookingRepo.GetBookingByID(id)
}

func (s *BookingService) UpdateBooking(booking *entity.Booking) error {
	return s.bookingRepo.UpdateBooking(booking)
}

func (s *BookingService) DeleteBooking(id uint) error {
	return s.bookingRepo.DeleteBooking(id)
}

func (s *BookingService) GetBookingsByUserID(userID uint) ([]*entity.Booking, error) {
	return s.bookingRepo.GetBookingsByUserID(userID)
}

func (s *BookingService) CreateUser(name string) (*entity.User, error) {
	user := &entity.User{
		Name:  name, 
	}
	err := s.userRepo.CreateUser(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}