package entity

import "gorm.io/gorm"

type Room struct {
	gorm.Model
	Name     string `gorm:"type:varchar(255);not null"`
	Capacity int    `gorm:"not null"`
	Status   string `gorm:"type:varchar(255);not null"`
}

type RoomResponse struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Capacity int    `json:"capacity"`
	Status   string `json:"status"`
}
