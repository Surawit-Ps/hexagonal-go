package entity

import ("gorm.io/gorm"
	"time")

type Booking struct {
	gorm.Model
	UserID uint  `gorm:"not null"`
	RoomID uint  `gorm:"not null"`
	StartTime time.Time  `gorm:"not null"`
	EndTime time.Time  `gorm:"not null"`
}

type BookingResponse struct {
	ID uint `json:"id"`
	UserID uint `json:"user_id"`
	RoomID uint `json:"room_id"`
	StartTime time.Time `json:"start_time"`
	EndTime time.Time `json:"end_time"`
}