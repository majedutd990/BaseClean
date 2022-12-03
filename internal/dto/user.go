package dto

import "github.com/google/uuid"

// User dto
type User struct {
	UserID      uuid.UUID
	FirstName   string
	LastName    string
	Email       string
	Password    string
	PhoneNumber string
}
