package service

import (
	"context"
	"sushkof_test/internal/data/request"
	"sushkof_test/internal/data/response"
)

type TaskUserServiceImpl struct {
	TaskService TaskService
	UserService UserService
}

func NewTaskUserServiceImpl(taskService TaskService, userService UserService) *TaskUserServiceImpl {
	return &TaskUserServiceImpl{
		TaskService: taskService,
		UserService: userService,
	}
}

func (tus *TaskUserServiceImpl) GetTaskByName(ctx context.Context, name string) []response.TasksResponse {

	id := tus.UserService.GetIdByName(ctx, name)

	tasks := tus.TaskService.GetTasksById(ctx, id)

	var taskResponse []response.TasksResponse

	for _, value := range tasks {
		task := response.TasksResponse{
			Title:       value.Title,
			Description: value.Description,
			Status:      value.Status,
		}

		taskResponse = append(taskResponse, task)
	}

	return taskResponse
}

func (tus *TaskUserServiceImpl) AddTaskToUser(ctx context.Context, req request.AddTaskRequest) {

	task := request.SaveTaskRequest{
		Title:       req.Title,
		Description: req.Description,
		Status:      req.Status,
	}

	tus.TaskService.Save(ctx, task)

	userId := tus.UserService.GetIdByName(ctx, req.Username)
	
	tus.TaskService.AssignTaskToUser(ctx, task.Title, userId)
}
