package service

import (
	"Hexa/domain/entity"
	"Hexa/domain/repository"
)

type UserService struct {
	Repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *UserService {
	return &UserService{Repo: repo}
}

func (s *UserService) AddUser(user *entity.User) error {
	return s.Repo.CreateUser(user)
}

func (s *UserService) GetUser(id uint) (*entity.User, error) {
	return s.Repo.GetUserByID(id)
}

func (s *UserService) UpdateUser(user *entity.User) error {
	return s.Repo.UpdateUser(user)
}

func (s *UserService) DeleteUser(id uint) error {
	return s.Repo.DeleteUser(id)
}

func (s *UserService) GetAllUsers() ([]*entity.User, error) {
	return s.Repo.GetAllUsers()
}

// RoomService
type RoomService struct {
	Repo repository.RoomRepository
}

func NewRoomService(repo repository.RoomRepository) *RoomService {
	return &RoomService{Repo: repo}
}

func (s *RoomService) AddRoom(room *entity.Room) error {
	return s.Repo.CreateRoom(room)
}

func (s *RoomService) GetRoom(id uint) (*entity.Room, error) {
	return s.Repo.GetRoomByID(id)
}

func (s *RoomService) UpdateRoom(room *entity.Room) error {
	return s.Repo.UpdateRoom(room)
}

func (s *RoomService) DeleteRoom(id uint) error {
	return s.Repo.DeleteRoom(id)
}

func (s *RoomService) GetAllRooms() ([]*entity.Room, error) {
	return s.Repo.GetAllRooms()
}
