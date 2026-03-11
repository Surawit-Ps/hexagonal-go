package sql

import (
	"Hexa/domain/entity"

	"gorm.io/gorm"
)

// ===== Booking Repository =====
type Sql struct {
	DB *gorm.DB
}

func NewSql(db *gorm.DB) *Sql {
	return &Sql{DB: db}
}

func (s *Sql) CreateBooking(booking *entity.Booking) error {
	return s.DB.Create(booking).Error
}

func (s *Sql) GetBookingByID(id string) (*entity.Booking, error) {
	var booking entity.Booking
	if err := s.DB.First(&booking, id).Error; err != nil {
		return nil, err
	}
	return &booking, nil
}

func (s *Sql) UpdateBooking(booking *entity.Booking) error {
	SQLBooking := toSQLBooking(booking)
	return s.DB.Save(&SQLBooking).Error
}

func (s *Sql) DeleteBooking(id string) error {
	return s.DB.Delete(&entity.Booking{}, id).Error
}

func (s *Sql) GetBookingsByUserID(userID string) ([]*entity.Booking, error) {
	var bookings []*entity.Booking
	if err := s.DB.Where("user_id = ?", userID).Find(&bookings).Error; err != nil {
		return nil, err
	}
	return bookings, nil
}

// ===== User Repository =====
type SqlUser struct {
	DB *gorm.DB
}

func NewSqlUser(db *gorm.DB) *SqlUser {
	return &SqlUser{DB: db}
}

func (s *SqlUser) CreateUser(user *entity.User) error {
	return s.DB.Create(user).Error
}

func (s *SqlUser) GetUserByID(id string) (*entity.User, error) {
	var user entity.User
	if err := s.DB.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (s *SqlUser) UpdateUser(user *entity.User) error {
	return s.DB.Save(user).Error
}

func (s *SqlUser) DeleteUser(id string) error {
	return s.DB.Delete(&entity.User{}, id).Error
}

func (s *SqlUser) GetAllUsers() ([]*entity.User, error) {
	var users []*entity.User
	if err := s.DB.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

// ===== Room Repository =====
type SqlRoom struct {
	DB *gorm.DB
}

func NewSqlRoom(db *gorm.DB) *SqlRoom {
	return &SqlRoom{DB: db}
}

func (s *SqlRoom) CreateRoom(room *entity.Room) error {
	return s.DB.Create(room).Error
}

func (s *SqlRoom) GetRoomByID(id string) (*entity.Room, error) {
	var room entity.Room
	if err := s.DB.First(&room, id).Error; err != nil {
		return nil, err
	}
	return &room, nil
}

func (s *SqlRoom) UpdateRoom(room *entity.Room) error {
	return s.DB.Save(room).Error
}

func (s *SqlRoom) DeleteRoom(id string) error {
	return s.DB.Delete(&entity.Room{}, id).Error
}

func (s *SqlRoom) GetAllRooms() ([]*entity.Room, error) {
	var rooms []*entity.Room
	if err := s.DB.Find(&rooms).Error; err != nil {
		return nil, err
	}
	return rooms, nil
}
