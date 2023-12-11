package repository

import (
	"context"
	"sushkof_test/internal/model"
)

type TaskRepo interface {
	FindByName(ctx context.Context, title string) (model.Task, error)
	FindById(ctx context.Context, id int) (model.Task, error)
	GetTasksById(ctx context.Context, id int) []model.Task
	AssignTaskToUser(ctx context.Context, title string, userId int)
	FindAll(ctx context.Context) []model.Task
	Save(ctx context.Context, task model.Task)
	Update(ctx context.Context, task model.Task)
	Delete(ctx context.Context, idTask int)
}
