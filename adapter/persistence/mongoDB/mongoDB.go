package mongoDB

import (
	"context"
	"Hexa/domain/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDB struct {
	db *mongo.Client
}

func NewMongoDB(uri string) (*MongoDB, error) {
	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return nil, err
	}
	return &MongoDB{db: client}, nil
}

func (m *MongoDB) CreateBooking(booking *entity.Booking) error {
	collection := m.db.Database("bookingdb").Collection("bookings")
	mongoBooking := toMongoDBBooking(booking)
	_, err := collection.InsertOne(context.Background(), mongoBooking)
	return err
}

func (m *MongoDB) GetBookingByID(id uint) (*entity.Booking, error) {
	collection := m.db.Database("bookingdb").Collection("bookings")
	var booking entity.Booking
	mongoBooking := toMongoDBBooking(&booking)
	err := collection.FindOne(context.Background(), bson.M{"id": id}).Decode(&mongoBooking)
	if err != nil {
		return nil, err
	}
	return &booking, nil
}

func (m *MongoDB) UpdateBooking(booking *entity.Booking) error {
	collection := m.db.Database("bookingdb").Collection("bookings")
	mongoBooking := toMongoDBBooking(booking)
	_, err := collection.UpdateOne(context.Background(), bson.M{"id": booking.ID}, bson.M{"$set": mongoBooking})
	return err
}

func (m *MongoDB) DeleteBooking(id string) error {
	collection := m.db.Database("bookingdb").Collection("bookings")
	mongoBooking := toMongoDBBooking(&entity.Booking{ID: id})
	_, err := collection.DeleteOne(context.Background(), bson.M{"id": mongoBooking.ID})
	return err
}

func (m *MongoDB) GetBookingsByUserID(userID string) ([]*entity.Booking, error) {
	collection := m.db.Database("bookingdb").Collection("bookings")
	var mongoBookings []MongoDBBooking	
	cursor, err := collection.Find(context.Background(), bson.M{"user_id": userID})
	if err != nil {
		return nil, err
	}
	if err := cursor.All(context.Background(), &mongoBookings); err != nil {
		return nil, err
	}
	var bookings []*entity.Booking
	for _, mongoBooking := range mongoBookings {
		bookings = append(bookings, toEntityBooking(&mongoBooking))
	}
	return bookings, nil
}

/////// User Repository Implementation
func (m *MongoDB) CreateUser(user *entity.User) error {
	collection := m.db.Database("bookingdb").Collection("users")
	mongoUser := toMongoDBUser(user)
	_, err := collection.InsertOne(context.Background(), mongoUser)
	return err
}

func (m *MongoDB) GetUserByID(id string) (*entity.User, error) {
	collection := m.db.Database("bookingdb").Collection("users")
	var user entity.User
	mongoUser := toMongoDBUser(&user)
	err := collection.FindOne(context.Background(), bson.M{"id": id}).Decode(&mongoUser)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (m *MongoDB) UpdateUser(user *entity.User) error {
	collection := m.db.Database("bookingdb").Collection("users")
	mongoUser := toMongoDBUser(user)
	_, err := collection.UpdateOne(context.Background(), bson.M{"id": user.ID}, bson.M{"$set": mongoUser})
	return err
}

func (m *MongoDB) DeleteUser(id string) error {
	collection := m.db.Database("bookingdb").Collection("users")
	mongoUser := toMongoDBUser(&entity.User{ID: id})
	_, err := collection.DeleteOne(context.Background(), bson.M{"id": mongoUser.ID})
	return err
}

//// Room Repository Implementation
func (m *MongoDB) CreateRoom(room *entity.Room) error {
	collection := m.db.Database("bookingdb").Collection("rooms")
	mongoRoom := toMongoDBRoom(room)
	_, err := collection.InsertOne(context.Background(), mongoRoom)
	return err
}

func (m *MongoDB) GetRoomByID(id uint) (*entity.Room, error) {
	collection := m.db.Database("bookingdb").Collection("rooms")
	var room entity.Room
	mongoRoom := toMongoDBRoom(&room)
	err := collection.FindOne(context.Background(), bson.M{"id": id}).Decode(&mongoRoom)
	if err != nil {
		return nil, err
	}
	return &room, nil
}

func (m *MongoDB) UpdateRoom(room *entity.Room) error {
	collection := m.db.Database("bookingdb").Collection("rooms")
	mongoRoom := toMongoDBRoom(room)
	_, err := collection.UpdateOne(context.Background(), bson.M{"id": room.ID}, bson.M{"$set": mongoRoom})
	return err
}

func (m *MongoDB) DeleteRoom(id string) error {
	collection := m.db.Database("bookingdb").Collection("rooms")
	mongoRoom := toMongoDBRoom(&entity.Room{ID: id})
	_, err := collection.DeleteOne(context.Background(), bson.M{"id": mongoRoom.ID})
	return err
}
