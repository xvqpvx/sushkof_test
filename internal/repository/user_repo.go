package repository

import (
	"context"
	"sushkof_test/internal/model"
)

type UserRepo interface {
	Save(ctx context.Context, user model.User)
	GetIdByName(ctx context.Context, name string) int
	FindAll(ctx context.Context) []model.User
}
