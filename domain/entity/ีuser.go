package entity


type User struct {
	ID  string
	Name     string 
	Email    string 
}

type UserResponse struct {
	ID    string   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}