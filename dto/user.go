package dto

type User struct {
	FirstName string `json:"firstName" binding:"required,max=20"`
	LastName  string `json:"lastName" binding:"required,max=20"`
}
