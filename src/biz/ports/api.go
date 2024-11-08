package ports

import "github.com/huynhminhtruong/go-store-user/src/biz/application/core/domain"

type DBPort interface {
	Save(user *domain.User) *domain.CreateUserResponse
	GetUser(id int64) *domain.GetUserResponse
	GetListUsers() *domain.ListUserResponse
}
