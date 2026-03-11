package entity

import (
	"time")

type Booking struct {
	ID string
	UserID string  
	RoomID string  
	StartTime time.Time  
	EndTime time.Time  
}

type BookingResponse struct {
	ID string `json:"id"`
	UserID string `json:"user_id"`
	RoomID string `json:"room_id"`
	StartTime time.Time `json:"start_time"`
	EndTime time.Time `json:"end_time"`
}