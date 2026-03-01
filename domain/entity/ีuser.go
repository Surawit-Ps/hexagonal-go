package entity

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `json:"name" gorm:"type:varchar(255);not null"`
}