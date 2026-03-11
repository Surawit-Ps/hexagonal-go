package mongoDB

import (
	"Hexa/domain/entity"
	"time"
)

type MongoDBBooking struct {
	ID     string   `bson:"_id,omitempty"`
	UserID string   `bson:"user_id"`
	RoomID string   `bson:"room_id"`
	StartTime string `bson:"start_time"`
	EndTime   string `bson:"end_time"`
}

type MongoDBUser struct {
	ID   string   `bson:"_id,omitempty"`
	Name string `bson:"name"`
	Email string `bson:"email"`
}

type MongoDBRoom struct {
	ID   string   `bson:"_id,omitempty"`
 	Name string `bson:"name"`
	Capacity int    `bson:"capacity"`
	Status   string `bson:"status"`
}  

func toMongoDBBooking(booking *entity.Booking) *MongoDBBooking {
	return &MongoDBBooking{
		ID: booking.ID,	
		UserID: booking.UserID,
		RoomID: booking.RoomID,
		StartTime: booking.StartTime.Format("2006-01-02T15:04:05Z07:00"),
		EndTime: booking.EndTime.Format("2006-01-02T15:04:05Z07:00"),
	}
}

func toEntityBooking(mongoBooking *MongoDBBooking) *entity.Booking {
	startTime, _ := time.Parse("2006-01-02T15:04:05Z07:00", mongoBooking.StartTime)
	endTime, _ := time.Parse("2006-01-02T15:04:05Z07:00", mongoBooking.EndTime)
	return &entity.Booking{
		ID: mongoBooking.ID,
		UserID: mongoBooking.UserID,
		RoomID: mongoBooking.RoomID,
		StartTime: startTime,
		EndTime: endTime,
	}
}

func toMongoDBUser(user *entity.User) *MongoDBUser {
	return &MongoDBUser{
		ID: user.ID,	
		Name: user.Name,
		Email: user.Email,
	}
}

func toEntityUser(mongoUser *MongoDBUser) *entity.User {
	return &entity.User{
		ID: mongoUser.ID,
		Name: mongoUser.Name,
		Email: mongoUser.Email,
	}
}

func toMongoDBRoom(room *entity.Room) *MongoDBRoom {
	return &MongoDBRoom{
		ID: room.ID,
		Name: room.Name,
		Capacity: room.Capacity,
		Status: room.Status,
	}
}

func toEntityRoom(mongoRoom *MongoDBRoom) *entity.Room {
	return &entity.Room{
		ID: mongoRoom.ID,	
		Name: mongoRoom.Name,
		Capacity: mongoRoom.Capacity,
		Status: mongoRoom.Status,
	}
}