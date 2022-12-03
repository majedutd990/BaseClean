package datastruct

import (
	"github.com/google/uuid"
	"time"
)

const UserTableName = "users"

type Role string

const (
	ADMIN  Role = "admin"
	USER   Role = "user"
	AUTHOR Role = "author"
)

type Status string

const (
	ACTIVE  Status = "active"
	BLOCKED Status = "blocked"
	PENDING Status = "pending"
)

type Gender string

const (
	MALE   Gender = "male"
	FEMALE Gender = "female"
	OTHER  Gender = "other"
)

// User describes User struct
type User struct {
	UserID        uuid.UUID `json:"user_id" db:"user_id" bson:"user_id" validate:"omitempty,uuid"`
	FirstName     string    `json:"first_name" db:"first_name" bson:"first_name" validate:"lte=30"`
	LastName      string    `json:"last_name" db:"last_name" bson:"last_name" validate:"lte=30"`
	Email         string    `json:"email,omitempty" db:"email" bson:"email" validate:"email,required,lte=30"`
	Password      string    `json:"password" db:"password" bson:"password" validate:"omitempty,required,gte=6"`
	PhoneNumber   string    `json:"phone_number,omitempty" db:"phone_number" bson:"phone_number" validate:"omitempty,lte=20"`
	Role          Role      `json:"role" db:"role" bson:"role" validate:"omitempty,lte=10"`
	About         string    `json:"about,omitempty" db:"about" bson:"about" validate:"omitempty,lte=1024"`
	Avatar        string    `json:"avatar,omitempty" db:"avatar" bson:"avatar" validate:"omitempty,lte=512"`
	Verified      bool      `json:"verified" db:"verified" bson:"verified" validate:"omitempty"`
	EmailCode     string    `json:"email_code" db:"email_code" bson:"email_code" validate:"omitempty"`
	Status        Status    `json:"status" db:"status" bson:"status" validate:"omitempty"`
	Address       string    `json:"address,omitempty" db:"address" bson:"address" validate:"omitempty,lte=250"`
	City          string    `json:"city,omitempty" db:"city" bson:"city" validate:"omitempty,lte=24"`
	Country       string    `json:"country,omitempty" db:"country" bson:"country" validate:"omitempty,lte=24"`
	Gender        Gender    `json:"gender,omitempty" db:"gender" bson:"gender" validate:"omitempty,lte=10"`
	Postcode      string    `json:"postcode,omitempty" db:"postcode" bson:"postcode" validate:"omitempty"`
	Birthday      time.Time `json:"birthday,omitempty" db:"birthday" bson:"birthday" validate:"omitempty"`
	CreatedAt     time.Time `json:"created_at,omitempty" db:"created_at" bson:"created_at"`
	UpdatedAt     time.Time `json:"updated_at,omitempty" db:"updated_at" bson:"updated_at"`
	LastLoginDate time.Time `json:"last_login_date" db:"last_login_date" bson:"last_login_date"`
}

// UsersList a struct containing all users, used for pagination
type UsersList struct {
	TotalCount int    `json:"total_count"`
	TotalPages int    `json:"total_pages"`
	Page       int    `json:"page"`
	Size       int    `json:"size"`
	HasMore    bool   `json:"has_more"`
	Users      []User `json:"users"`
}
