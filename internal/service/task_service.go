package service

import (
	"context"
	"sushkof_test/internal/data/request"
	"sushkof_test/internal/data/response"
)

type TaskService interface {
	FindByName(ctx context.Context, title string) response.TasksResponse
	GetTasksById(ctx context.Context, id int) []response.TasksResponse
	AssignTaskToUser(ctx context.Context, title string, userId int)
	FindAll(ctx context.Context) []response.TasksResponse
	Save(ctx context.Context, task request.SaveTaskRequest)
	Update(ctx context.Context, task request.UpdateTaskRequest)
	Delete(ctx context.Context, title string)
}
