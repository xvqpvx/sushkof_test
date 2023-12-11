package service

import (
	"context"
	"sushkof_test/internal/data/request"
	"sushkof_test/internal/data/response"
)

type UserService interface {
	Save(ctx context.Context, user request.SaveUserRequest)
	FindAll(ctx context.Context) []response.UserResponse
	GetIdByName(ctx context.Context, name string) int
}
