package service

import (
	"context"
	"sushkof_test/internal/data/request"
	"sushkof_test/internal/data/response"
	"sushkof_test/internal/helper"
	"sushkof_test/internal/model"
	"sushkof_test/internal/repository"
)

type TaskServiceImpl struct {
	TaskRepo repository.TaskRepo
}

func NewTaskServiceImpl(taskRepo repository.TaskRepo) TaskService {
	return &TaskServiceImpl{TaskRepo: taskRepo}
}

func (t *TaskServiceImpl) FindByName(ctx context.Context, title string) response.TasksResponse {

	task, err := t.TaskRepo.FindByName(ctx, title)
	helper.PanicIfError(err)
	return response.TasksResponse{
		IdTask:      task.IdTask,
		Title:       task.Title,
		Description: task.Description,
		Status:      task.Status,
	}

}

func (t *TaskServiceImpl) GetTasksById(ctx context.Context, id int) []response.TasksResponse {
	tasks := t.TaskRepo.GetTasksById(ctx, id)

	var tasksResponse []response.TasksResponse

	for _, value := range tasks {
		task := response.TasksResponse{
			Title:       value.Title,
			Description: value.Description,
			Status:      value.Status,
		}

		tasksResponse = append(tasksResponse, task)
	}

	return tasksResponse
}

func (t *TaskServiceImpl) FindAll(ctx context.Context) []response.TasksResponse {
	tasks := t.TaskRepo.FindAll(ctx)

	var tasksResponse []response.TasksResponse

	for _, value := range tasks {
		task := response.TasksResponse{
			Title:       value.Title,
			Description: value.Description,
			Status:      value.Status,
		}

		tasksResponse = append(tasksResponse, task)
	}

	return tasksResponse

}

func (t *TaskServiceImpl) Save(ctx context.Context, task request.SaveTaskRequest) {
	taskToSave := model.Task{
		Title:       task.Title,
		Description: task.Description,
		Status:      task.Status,
	}

	t.TaskRepo.Save(ctx, taskToSave)
}

func (t *TaskServiceImpl) Update(ctx context.Context, task request.UpdateTaskRequest) {
	taskToUpdate := model.Task{
		IdTask:      task.IdTask,
		Title:       task.Title,
		Description: task.Description,
		Status:      task.Status,
	}

	t.TaskRepo.Update(ctx, taskToUpdate)
}

func (t *TaskServiceImpl) Delete(ctx context.Context, title string) {
	taskToDelete, err := t.TaskRepo.FindByName(ctx, title)
	helper.PanicIfError(err)

	t.TaskRepo.Delete(ctx, taskToDelete.IdTask)
}

func (t *TaskServiceImpl) AssignTaskToUser(ctx context.Context, title string, userId int) {
	t.TaskRepo.AssignTaskToUser(ctx, title, userId)
}
