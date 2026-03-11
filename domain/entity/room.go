package entity


type Room struct {
	ID string
	Name     string 
	Capacity int    
	Status   string `validate:"required,oneof='available' 'occupied' 'maintenance'"`
}

type RoomResponse struct {
	ID       string   `json:"id"`
	Name     string `json:"name"`
	Capacity int    `json:"capacity"`
	Status   string `json:"status"`
}
