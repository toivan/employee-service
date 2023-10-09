package model

type Employee struct {
	ID       int    `json:"id"`
	FullName string `json:"full_name"`
	Phone    string `json:"phone"`
	Gender   string `json:"gender"`
	Age      int    `json:"age"`
	Email    string `json:"email"`
	Address  string `json:"address"`
}
