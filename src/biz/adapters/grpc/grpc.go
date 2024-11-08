package grpc

import (
	"context"
	"github.com/huynhminhtruong/go-store-user/src/biz/application/core/domain"
	"github.com/huynhminhtruong/go-store-user/src/services/user"
)

func (a Adapter) Create(ctx context.Context, request *user.RegisterUserRequest) (*user.RegisterUserResponse, error) {
	newUser := domain.NewUser(&domain.UserInformation{
		Username: request.GetUsername(),
		Password: request.GetPassword(),
		Email:    request.GetEmail(),
		Phone:    request.GetPhoneNumber(),
		Location: request.GetLocation(),
	})
	// call this method from ports.APIPort interface and Application is already implement it
	result := a.api.CreateUser(newUser)
	if result.ErrorMessage != nil {
		return nil, result.ErrorMessage
	}
	return &user.RegisterUserResponse{UserId: result.UserID}, nil
}

func (a Adapter) ListUsers(ctx context.Context, request *user.ListUsersRequest) (*user.ListUsersResponse, error) {
	result := a.api.GetListUsers()
	if result.ErrorMessage != nil {
		return nil, result.ErrorMessage
	}

	var users []*user.GetUserResponse
	for _, usr := range result.Users {
		userResponse := &user.GetUserResponse{
			Username:    usr.Username,
			Email:       usr.Email,
			PhoneNumber: usr.Phone,
			Location:    usr.Location,
		}
		users = append(users, userResponse)
	}
	return &user.ListUsersResponse{Users: users}, nil
}

func (a Adapter) GetUser(ctx context.Context, request *user.GetUserRequest) (*user.GetUserResponse, error) {
	userID := request.GetUserId()
	result := a.api.GetUser(userID)
	if result.ErrorMessage != nil {
		return nil, result.ErrorMessage
	}
	usr := result.UserInfo
	userResponse := &user.GetUserResponse{
		Username:    usr.Username,
		Email:       usr.Email,
		PhoneNumber: usr.Phone,
		Location:    usr.Location,
	}
	return userResponse, nil
}
