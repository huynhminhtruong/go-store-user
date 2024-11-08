package db

import (
	"fmt"
	"github.com/huynhminhtruong/go-store-user/src/biz/application/core/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// User is used for mapping with db table
type User struct {
	// Adds entity metadata such as ID to struct
	gorm.Model
	UserID   int64
	Username string
	Password string
	Email    string
	Phone    string
	Location string
}

// Adapter database
type Adapter struct {
	db *gorm.DB
}

// NewAdapter initialize
func NewAdapter(dataSourceURL string) (*Adapter, error) {
	db, openErr := gorm.Open(postgres.Open(dataSourceURL), &gorm.Config{})
	if openErr != nil {
		return nil, fmt.Errorf("db connection error: %v", openErr)
	}

	// create database table
	err := db.AutoMigrate(&User{})
	if err != nil {
		return nil, fmt.Errorf("db migration error: %v", err)
	}

	return &Adapter{db: db}, nil
}

func (a Adapter) GetUser(id int64) *domain.GetUserResponse {
	var user domain.User

	err := a.db.First(&user, "id = ?", id).Error
	if err != nil {
		return &domain.GetUserResponse{
			ErrorMessage: err,
		}
	}
	return &domain.GetUserResponse{
		UserInfo: user,
	}
}

func (a Adapter) GetListUsers() *domain.ListUserResponse {
	var users []domain.User

	if err := a.db.Find(&users).Error; err != nil {
		return &domain.ListUserResponse{}
	}

	return &domain.ListUserResponse{
		Users: users,
	}
}

func (a Adapter) Save(user *domain.User) *domain.CreateUserResponse {
	usr := a.db.Create(user)
	if usr.Error != nil {
		return &domain.CreateUserResponse{}
	}

	return &domain.CreateUserResponse{
		UserID:       user.ID,
		RowsAffected: usr.RowsAffected,
		ErrorMessage: usr.Error,
	}
}
