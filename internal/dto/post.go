package dto

import (
	"github.com/google/uuid"
)

// Post dto
type Post struct {
	PostID   uuid.UUID
	AuthorID uuid.UUID
	Title    string
	Author   User
}
