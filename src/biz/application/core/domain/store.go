package domain

import (
	"time"
)

/*
	Business logic here
*/

type User struct {
	ID        int64     `json:"id"`
	UserID    int64     `json:"user_id"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
	Location  string    `json:"location"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt bool      `json:"deleted_at"`
}

type CreateUserResponse struct {
	UserID       int64 `json:"user_id"`
	RowsAffected int64 `json:"rows_affected"`
	ErrorMessage error `json:"error_message"`
}

type GetUserResponse struct {
	UserInfo     User  `json:"user_info"`
	RowsAffected int64 `json:"rows_affected"`
	ErrorMessage error `json:"error_message"`
}

type ListUserResponse struct {
	Users        []User `json:"users"`
	RowsAffected int64  `json:"rows_affected"`
	ErrorMessage error  `json:"error_message"`
}

func NewUser(username, password, email, phone, location string) *User {
	return &User{
		Username:  username,
		Password:  password,
		Email:     email,
		Phone:     phone,
		Location:  location,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		DeletedAt: false,
	}
}
