package api

import (
	"github.com/huynhminhtruong/go-store-user/src/biz/application/core/domain"
	"github.com/huynhminhtruong/go-store-user/src/biz/ports"
)

type Application struct {
	db ports.DBPort
}

// NewApplication will decide type of database server will be used in application
func NewApplication(db ports.DBPort) *Application {
	return &Application{
		db: db,
	}
}

func (a Application) CreateUser(user domain.User) *domain.CreateUserResponse {
	// call this method from ports.DBPort interface
	result := a.db.Save(&user)
	if result.ErrorMessage != nil {
		return &domain.CreateUserResponse{
			ErrorMessage: result.ErrorMessage,
		}
	}
	return result
}

func (a Application) GetUser(id int64) *domain.GetUserResponse {
	// call this method from ports.DBPort interface
	result := a.db.GetUser(id)
	if result.ErrorMessage != nil {
		return &domain.GetUserResponse{
			ErrorMessage: result.ErrorMessage,
		}
	}
	return result
}

func (a Application) GetUsers() *domain.ListUserResponse {
	// call this method from ports.DBPort interface
	result := a.db.GetListUsers()
	if result.ErrorMessage != nil {
		return &domain.ListUserResponse{
			ErrorMessage: result.ErrorMessage,
		}
	}
	return result
}
