package mockdata

import (
	"Hexa/domain/entity"
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
		{Name: "John Doe", Email: "john@example.com", Password: "pass123"},
		{Name: "Jane Smith", Email: "jane@example.com", Password: "pass123"},
		{Name: "Bob Johnson", Email: "bob@example.com", Password: "pass123"},
	}
	m.Users[0].ID = 1
	m.Users[1].ID = 2
	m.Users[2].ID = 3

	// Add sample rooms
	m.Rooms = []*entity.Room{
		{Name: "Room A", Capacity: 10, Status: "available"},
		{Name: "Room B", Capacity: 20, Status: "booked"},
		{Name: "Room C", Capacity: 15, Status: "available"},
	}
	m.Rooms[0].ID = 1
	m.Rooms[1].ID = 2
	m.Rooms[2].ID = 3

	// Add sample bookings
	m.Bookings = []*entity.Booking{
		{
			UserID:    1,
			RoomID:    2,
			StartTime: time.Date(2024, time.June, 1, 9, 0, 0, 0, time.UTC),
			EndTime:   time.Date(2024, time.June, 1, 10, 0, 0, 0, time.UTC),
		},
		{
			UserID:    2,
			RoomID:    1,
			StartTime: time.Date(2024, time.June, 2, 14, 0, 0, 0, time.UTC),
			EndTime:   time.Date(2024, time.June, 2, 15, 0, 0, 0, time.UTC),
		},
		{
			UserID:    3,
			RoomID:    3,
			StartTime: time.Date(2024, time.June, 3, 11, 0, 0, 0, time.UTC),
			EndTime:   time.Date(2024, time.June, 3, 12, 0, 0, 0, time.UTC),
		},
	}
	m.Bookings[0].ID = 1
	m.Bookings[1].ID = 2
	m.Bookings[2].ID = 3

	m.userID = 3
	m.roomID = 3
	m.bookingID = 3
}

// ===== Booking Repository Implementation =====
func (m *MockData) CreateBooking(booking *entity.Booking) error {
	m.bookingID++
	booking.ID = m.bookingID
	m.Bookings = append(m.Bookings, booking)
	return nil
}

func (m *MockData) GetBookingByID(id uint) (*entity.Booking, error) {
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

func (m *MockData) DeleteBooking(id uint) error {
	for i, booking := range m.Bookings {
		if booking.ID == id {
			m.Bookings = append(m.Bookings[:i], m.Bookings[i+1:]...)
			return nil
		}
	}
	return nil
}

func (m *MockData) GetBookingsByUserID(userID uint) ([]*entity.Booking, error) {
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
	user.ID = m.userID
	m.Users = append(m.Users, user)
	return nil
}

func (m *MockData) GetUserByID(id uint) (*entity.User, error) {
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

func (m *MockData) DeleteUser(id uint) error {
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
	room.ID = m.roomID
	m.Rooms = append(m.Rooms, room)
	return nil
}

func (m *MockData) GetRoomByID(id uint) (*entity.Room, error) {
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

func (m *MockData) DeleteRoom(id uint) error {
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
