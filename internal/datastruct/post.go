package datastruct

import (
	"github.com/google/uuid"
	"time"
)

const PostsTableName = "posts"

type Category string

const (
	NEWS Category = "news"
)

// Post describes post struct
type Post struct {
	PostID           uuid.UUID `json:"post_id" db:"post_id" bson:"post_id" validate:"omitempty,uuid"`
	AuthorID         uuid.UUID `json:"author_id,omitempty" db:"author_id" bson:"author_id" validate:"required,uuid"`
	Title            string    `json:"title" db:"title" bson:"title" validate:"omitempty,lte=512"`
	Content          string    `json:"content" db:"content" bson:"content"`
	ImageURL         string    `json:"image_url,omitempty" db:"image_url" bson:"image_url" validate:"omitempty,lte=512"`
	IntroDescription string    `json:"intro_description,omitempty" db:"intro_description" bson:"intro_description" validate:"omitempty,lte=1024"`
	Category         string    `json:"category,omitempty" db:"category" bson:"category" validate:"omitempty,lte=10"`
	Published        bool      `json:"published" db:"published" bson:"published"`
	FirstPage        bool      `json:"first_page" db:"first_page" bson:"first_page"`
	MetaDescription  string    `json:"meta_description" db:"meta_description" bson:"meta_description"`
	CreatedAt        time.Time `json:"created_at,omitempty" db:"created_at" bson:"created_at"`
	UpdatedAt        time.Time `json:"updated_at,omitempty" db:"updated_at" bson:"updated_at"`
	PublishedAt      time.Time `json:"published_at,omitempty" db:"published_at" bson:"published_at"`
	Author           User      `json:"author" bson:"author"`
}

// PostsList a struct containing all Post, used for pagination
type PostsList struct {
	TotalCount int    `json:"total_count"`
	TotalPages int    `json:"total_pages"`
	Page       int    `json:"page"`
	Size       int    `json:"size"`
	HasMore    bool   `json:"has_more"`
	Posts      []Post `json:"posts"`
}
