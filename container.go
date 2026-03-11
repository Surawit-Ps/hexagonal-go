package main

import (
	"log"

	"Hexa/adapter/http"
	"Hexa/adapter/persistence/mockdata"
	"Hexa/adapter/persistence/sql"
	"Hexa/domain/entity"
	"Hexa/domain/service"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)


type Container struct {
	DB             *gorm.DB
	BookingService *service.BookingService
	UserService    *service.UserService
	RoomService    *service.RoomService
	BookingHandler *http.Handler
	UserHandler    *http.UserHandler
	RoomHandler    *http.RoomHandler
}


func NewContainer(useMock bool) *Container {

	var db *gorm.DB
	if !useMock {
		var err error
		db, err = initDatabase()
		if err != nil {
			log.Fatal("Failed to initialize database:", err)
		}
	}

	

	var bookingService *service.BookingService
	var userService *service.UserService
	var roomService *service.RoomService

	if useMock {
		// Use Mock Data
		mockRepo := mockdata.NewMockData()
		bookingService = service.NewBookingService(mockRepo)
		userService = service.NewUserService(mockRepo)
		roomService = service.NewRoomService(mockRepo)
	} else {
		// Use SQL Repositories
		bookingRepo := sql.NewSql(db)
		userRepo := sql.NewSqlUser(db)
		roomRepo := sql.NewSqlRoom(db)

		bookingService = service.NewBookingService(bookingRepo)
		userService = service.NewUserService(userRepo)
		roomService = service.NewRoomService(roomRepo)
	}


	bookingHandler := http.NewHandler(bookingService)
	userHandler := http.NewUserHandler(userService)
	roomHandler := http.NewRoomHandler(roomService)

	return &Container{
		DB:             db,
		BookingService: bookingService,
		UserService:    userService,
		RoomService:    roomService,
		BookingHandler: bookingHandler,
		UserHandler:    userHandler,
		RoomHandler:    roomHandler,
	}
}

func initDatabase() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	User := []entity.User{
		{Name: "Alice", Email: "alice@example.com"},
		{Name: "Bob", Email: "bob@example.com"},
		{Name: "Charlie", Email: "charlie@example.com"},
		{Name: "David", Email: "david@example.com"},
	}
	Room := []entity.Room{
		{Name: "Conference Room A", Capacity: 10, Status: "available"},
		{Name: "Conference Room B", Capacity: 20, Status: "available"},
		{Name: "Conference Room C", Capacity: 15, Status: "available"},
		{Name: "Conference Room D", Capacity: 25, Status: "available"},
	}


	err = db.AutoMigrate(&entity.User{}, &entity.Room{}, &entity.Booking{})
	if err != nil {
		return nil, err
	}
	for _, user := range User {
		db.FirstOrCreate(&user, entity.User{Email: user.Email})
	}
	for _, room := range Room {
		db.FirstOrCreate(&room, entity.Room{Name: room.Name})
	}

	return db, nil
}
