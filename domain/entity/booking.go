package entity

import ("gorm.io/gorm"
	"time")

type Booking struct {
	gorm.Model
	UserID uint `json:"user_id" gorm:"type:int;not null"`
	RoomID uint `json:"room_id" gorm:"type:int;not null"`
	StartTime time.Time `json:"start_time" gorm:"type:datetime;not null"`
	EndTime time.Time `json:"end_time" gorm:"type:datetime;not null"`
}