package domain

import (
	"time"
)

/*
	Business logic here
*/

// User this is used for response data to client
type User struct {
	UserInformation
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt bool      `json:"deleted_at"`
}

type UserInformation struct {
	ID       int64  `json:"id"`
	UserID   int64  `json:"user_id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Location string `json:"location"`
}

// CreateUserResponse this is used for response data to client
type CreateUserResponse struct {
	UserID       int64 `json:"user_id"`
	RowsAffected int64 `json:"rows_affected"`
	ErrorMessage error `json:"error_message"`
}

// GetUserResponse this is used for response data to client
type GetUserResponse struct {
	UserInfo     User  `json:"user_info"`
	RowsAffected int64 `json:"rows_affected"`
	ErrorMessage error `json:"error_message"`
}

// ListUserResponse this is used for response data to client
type ListUserResponse struct {
	Users        []User `json:"users"`
	RowsAffected int64  `json:"rows_affected"`
	ErrorMessage error  `json:"error_message"`
}

func NewUser(user *UserInformation) *User {
	return &User{
		UserInformation: *user,
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
		DeletedAt:       false,
	}
}
