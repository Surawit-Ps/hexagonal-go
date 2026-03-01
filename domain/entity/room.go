package entity

import "gorm.io/gorm"

type Room struct {
	gorm.Model
	Capacity int `json:"capacity" gorm:"type:int;not null"`
	Status   string `json:"status" gorm:"type:varchar(255);not null"`
}