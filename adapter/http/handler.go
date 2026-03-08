package http

import (
	"Hexa/domain/entity"
	"Hexa/domain/service"
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	Service *service.BookingService
}

func NewHandler(service *service.BookingService) *Handler {
	return &Handler{
		Service: service,
	}
}

func (h *Handler) CreateBooking(c *fiber.Ctx) error {
	var booking entity.Booking
	if err := c.BodyParser(&booking); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}
	if err := h.Service.AddBooking(&booking); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create booking"})
	}
	return c.Status(fiber.StatusCreated).JSON(booking)
}

func (h *Handler) GetBooking(c *fiber.Ctx) error {
	id := c.Params("id")
	idUint, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}
	booking, err := h.Service.GetBooking(uint(idUint))
	if err != nil {	
		log.Printf("Error retrieving booking with ID %d: %v", idUint, err)
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Booking not found"})
	}
	return c.JSON(booking)
}

func (h *Handler) UpdateBooking(c *fiber.Ctx) error {
	id := c.Params("id")
	idUint, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	var booking entity.Booking
	if err := c.BodyParser(&booking); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}
	booking.ID = uint(idUint)
	if err := h.Service.UpdateBooking(&booking); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update booking"})
	}
	return c.JSON(booking)
}

func (h *Handler) DeleteBooking(c *fiber.Ctx) error {
	id := c.Params("id")
	idUint, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}
	if err := h.Service.DeleteBooking(uint(idUint)); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete booking"})
	}
	c.SendStatus(fiber.StatusNoContent)
	return nil
}

func (h *Handler) GetBookingsByUserID(c *fiber.Ctx) error {
	userID := c.Params("userID")
	userIDUint, err := strconv.ParseUint(userID, 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid userID"})
	}
	bookings, err := h.Service.GetBookingsByUserID(uint(userIDUint))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to retrieve bookings"})
	}
	return c.JSON(bookings)
}