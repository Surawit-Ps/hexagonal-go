package mockdata

import (
	"Hexa/domain/entity"
	"strconv"
	"time"
)

type MockData struct {
	Users     []*entity.User
	Rooms     []*entity.Room
	Bookings  []*entity.Booking
	userID    uint
	roomID    uint
	bookingID uint
}

func NewMockData() *MockData {
	m := &MockData{
		Users:     make([]*entity.User, 0),
		Rooms:     make([]*entity.Room, 0),
		Bookings:  make([]*entity.Booking, 0),
		userID:    0,
		roomID:    0,
		bookingID: 0,
	}

	// Initialize with sample data
	m.initializeSampleData()
	return m
}

func (m *MockData) initializeSampleData() {
	// Add sample users
	m.Users = []*entity.User{
		{ID: strconv.FormatUint(1, 10), Name: "John Doe", Email: "john@example.com"},
		{ID: strconv.FormatUint(2, 10), Name: "Jane Smith", Email: "jane@example.com"},
		{ID: strconv.FormatUint(3, 10), Name: "Bob Johnson", Email: "bob@example.com"},
	}

	// Add sample rooms
	m.Rooms = []*entity.Room{
		{ID: strconv.FormatUint(1, 10), Name: "Room A", Capacity: 10, Status: "available"},
		{ID: strconv.FormatUint(2, 10), Name: "Room B", Capacity: 20, Status: "booked"},
		{ID: strconv.FormatUint(3, 10), Name: "Room C", Capacity: 15, Status: "available"},
	}

	// Add sample bookings
	m.Bookings = []*entity.Booking{
		{
			ID:        strconv.FormatUint(1, 10),
			UserID:    strconv.FormatUint(1, 10),
			RoomID:    strconv.FormatUint(2, 10),
			StartTime: time.Date(2024, time.June, 1, 9, 0, 0, 0, time.UTC),
			EndTime:   time.Date(2024, time.June, 1, 10, 0, 0, 0, time.UTC),
		},
		{
			ID:        strconv.FormatUint(2, 10),
			UserID:    strconv.FormatUint(2, 10),
			RoomID:    strconv.FormatUint(1, 10),
			StartTime: time.Date(2024, time.June, 2, 14, 0, 0, 0, time.UTC),
			EndTime:   time.Date(2024, time.June, 2, 15, 0, 0, 0, time.UTC),
		},
		{
			ID:        strconv.FormatUint(3, 10),
			UserID:    strconv.FormatUint(3, 10),
			RoomID:    strconv.FormatUint(3, 10),
			StartTime: time.Date(2024, time.June, 3, 11, 0, 0, 0, time.UTC),
			EndTime:   time.Date(2024, time.June, 3, 12, 0, 0, 0, time.UTC),
		},
	}

	m.userID = 3
	m.roomID = 3
	m.bookingID = 3
}

// ===== Booking Repository Implementation =====
func (m *MockData) CreateBooking(booking *entity.Booking) error {
	m.bookingID++
	booking.ID = strconv.FormatUint(uint64(m.bookingID), 10)
	m.Bookings = append(m.Bookings, booking)
	return nil
}

func (m *MockData) GetBookingByID(id string) (*entity.Booking, error) {
	for _, booking := range m.Bookings {
		if booking.ID == id {
			return booking, nil
		}
	}
	return nil, nil
}

func (m *MockData) UpdateBooking(booking *entity.Booking) error {
	for i, b := range m.Bookings {
		if b.ID == booking.ID {
			m.Bookings[i] = booking
			return nil
		}
	}
	return nil
}

func (m *MockData) DeleteBooking(id string) error {
	for i, booking := range m.Bookings {
		if booking.ID == id {
			m.Bookings = append(m.Bookings[:i], m.Bookings[i+1:]...)
			return nil
		}
	}
	return nil
}

func (m *MockData) GetBookingsByUserID(userID string) ([]*entity.Booking, error) {
	var bookings []*entity.Booking
	for _, booking := range m.Bookings {
		if booking.UserID == userID {
			bookings = append(bookings, booking)
		}
	}
	return bookings, nil
}

// ===== User Repository Implementation =====
func (m *MockData) CreateUser(user *entity.User) error {
	m.userID++
	user.ID = strconv.FormatUint(uint64(m.userID), 10)
	m.Users = append(m.Users, user)
	return nil
}

func (m *MockData) GetUserByID(id string) (*entity.User, error) {
	for _, user := range m.Users {
		if user.ID == id {
			return user, nil
		}
	}
	return nil, nil
}

func (m *MockData) UpdateUser(user *entity.User) error {
	for i, u := range m.Users {
		if u.ID == user.ID {
			m.Users[i] = user
			return nil
		}
	}
	return nil
}

func (m *MockData) DeleteUser(id string) error {
	for i, user := range m.Users {
		if user.ID == id {
			m.Users = append(m.Users[:i], m.Users[i+1:]...)
			return nil
		}
	}
	return nil
}

func (m *MockData) GetAllUsers() ([]*entity.User, error) {
	return m.Users, nil
}

// ===== Room Repository Implementation =====
func (m *MockData) CreateRoom(room *entity.Room) error {
	m.roomID++
	room.ID = strconv.FormatUint(uint64(m.roomID), 10)
	m.Rooms = append(m.Rooms, room)
	return nil
}

func (m *MockData) GetRoomByID(id string) (*entity.Room, error) {
	for _, room := range m.Rooms {
		if room.ID == id {
			return room, nil
		}
	}
	return nil, nil
}

func (m *MockData) UpdateRoom(room *entity.Room) error {
	for i, r := range m.Rooms {
		if r.ID == room.ID {
			m.Rooms[i] = room
			return nil
		}
	}
	return nil
}

func (m *MockData) DeleteRoom(id string) error {
	for i, room := range m.Rooms {
		if room.ID == id {
			m.Rooms = append(m.Rooms[:i], m.Rooms[i+1:]...)
			return nil
		}
	}
	return nil
}

func (m *MockData) GetAllRooms() ([]*entity.Room, error) {
	return m.Rooms, nil
}
