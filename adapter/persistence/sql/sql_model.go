package sql

import (
	"Hexa/domain/entity"
	"time"
)

type SQLBooking struct {
	ID string `gorm:"primaryKey"`
	UserID string
	RoomID string
	StartTime time.Time
	EndTime time.Time
}

type SQLUser struct {
	ID string `gorm:"primaryKey"`
	Name string
	Email string
}

type SQLRoom struct {
	ID string `gorm:"primaryKey"`
	Name string
	Capacity int
	Status string
}

func toSQLBooking(booking *entity.Booking) *SQLBooking {
	return &SQLBooking{
		ID: booking.ID,
		UserID: booking.UserID,
		RoomID: booking.RoomID,
		StartTime: booking.StartTime,
		EndTime: booking.EndTime,
	}
}

func toEntityBooking(sqlBooking *SQLBooking) *entity.Booking {
	return &entity.Booking{
		ID: sqlBooking.ID,
		UserID: sqlBooking.UserID,
		RoomID: sqlBooking.RoomID,
		StartTime: sqlBooking.StartTime,
		EndTime: sqlBooking.EndTime,
	}
}

func toSQLUser(user *entity.User) *SQLUser {
	return &SQLUser{
		ID: user.ID,	
		Name: user.Name,
		Email: user.Email,
	}
}

func toEntityUser(sqlUser *SQLUser) *entity.User {
	return &entity.User{
		ID: sqlUser.ID,
		Name: sqlUser.Name,
		Email: sqlUser.Email,
	}
}

func toSQLRoom(room *entity.Room) *SQLRoom {
	return &SQLRoom{
		ID: room.ID,
		Name: room.Name,
		Capacity: room.Capacity,
		Status: room.Status,
	}
}

func toEntityRoom(sqlRoom *SQLRoom) *entity.Room {
	return &entity.Room{
		ID: sqlRoom.ID,	
		Name: sqlRoom.Name,
		Capacity: sqlRoom.Capacity,
		Status: sqlRoom.Status,
	}
}



