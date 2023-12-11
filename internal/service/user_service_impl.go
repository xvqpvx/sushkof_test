package service

import (
	"context"
	"sushkof_test/internal/data/request"
	"sushkof_test/internal/data/response"
	"sushkof_test/internal/model"
	"sushkof_test/internal/repository"
)

type UserServiceImpl struct {
	UserRepo repository.UserRepo
}

func NewUserServiceImpl(userRepo repository.UserRepo) UserService {
	return &UserServiceImpl{UserRepo: userRepo}
}

func (u *UserServiceImpl) Save(ctx context.Context, request request.SaveUserRequest) {
	user := model.User{
		Name:     request.Name,
		Age:      request.Age,
		Email:    request.Email,
		Password: request.Password,
	}

	u.UserRepo.Save(ctx, user)

}

func (u *UserServiceImpl) GetIdByName(ctx context.Context, name string) int {
	id := u.UserRepo.GetIdByName(ctx, name)
	return id
}

func (u *UserServiceImpl) FindAll(ctx context.Context) []response.UserResponse {
	users := u.UserRepo.FindAll(ctx)

	var usersResponse []response.UserResponse

	for _, value := range users {
		user := response.UserResponse{
			Name:  value.Name,
			Age:   value.Age,
			Email: value.Email,
		}

		usersResponse = append(usersResponse, user)
	}

	return usersResponse
}
