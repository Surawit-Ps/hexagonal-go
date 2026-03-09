package main

import (
	"log"
	"github.com/gofiber/fiber/v2"
)

func main() {
	// Initialize dependency container with mock data
	container := NewContainer(false) // Set to false to use SQL database

	// Initialize Fiber app
	app := fiber.New()

	// Setup routes
	setupRoutes(app, container)

	// Start server
	log.Println("Starting server on :3000")
	if err := app.Listen(":3000"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}

// setupRoutes configures all HTTP routes
func setupRoutes(app *fiber.App, container *Container) {
	
	booking := app.Group("/api/bookings")
	booking.Post("/", container.BookingHandler.CreateBooking)
	booking.Get("/:id", container.BookingHandler.GetBooking)
	booking.Put("/:id", container.BookingHandler.UpdateBooking)
	booking.Delete("/:id", container.BookingHandler.DeleteBooking)
	booking.Get("/user/:userID", container.BookingHandler.GetBookingsByUserID)

	
	user := app.Group("/api/users")
	user.Post("/", container.UserHandler.CreateUser)
	user.Get("/:id", container.UserHandler.GetUser)
	user.Put("/:id", container.UserHandler.UpdateUser)
	user.Delete("/:id", container.UserHandler.DeleteUser)
	user.Get("/", container.UserHandler.GetAllUsers)


	room := app.Group("/api/rooms")
	room.Post("/", container.RoomHandler.CreateRoom)
	room.Get("/:id", container.RoomHandler.GetRoom)
	room.Put("/:id", container.RoomHandler.UpdateRoom)
	room.Delete("/:id", container.RoomHandler.DeleteRoom)
	room.Get("/", container.RoomHandler.GetAllRooms)

	
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"status": "ok"})
	})
}
