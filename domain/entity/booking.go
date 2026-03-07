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