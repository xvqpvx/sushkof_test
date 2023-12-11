package service

import (
	"context"
	"sushkof_test/internal/data/request"
	"sushkof_test/internal/data/response"
)

type TaskUserService interface {
	AddTaskToUser(ctx context.Context, req request.AddTaskRequest)
	GetTaskByName(ctx context.Context, name string) []response.TasksResponse
}
