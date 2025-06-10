package models

// User defines user struct
type User struct {
	ID    int    `json:"id" example:"1" format:"int"`
	Name  string `json:"name" example:"Dmytro" minLength:"1" maxLength:"30"`
	Email string `json:"email" example:"sample@gmail.com"`
}
