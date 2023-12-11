package controller

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"sushkof_test/internal/data/request"
	"sushkof_test/internal/service"
)

type TaskController struct {
	TaskService service.TaskService
}

func NewTaskController(taskService service.TaskService) *TaskController {
	return &TaskController{TaskService: taskService}
}

var updatePayload struct {
	OldTitle    string `json:"oldTitle"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

var deletePayload struct {
	Title string `json:"title"`
}

func (tc *TaskController) Update(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	err := json.NewDecoder(r.Body).Decode(&updatePayload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	taskUpd := tc.TaskService.FindByName(r.Context(), updatePayload.OldTitle)

	task := request.UpdateTaskRequest{
		IdTask:      taskUpd.IdTask,
		Title:       updatePayload.Title,
		Description: updatePayload.Description,
		Status:      updatePayload.Status,
	}

	tc.TaskService.Update(r.Context(), task)
}

func (tc *TaskController) Delete(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	err := json.NewDecoder(r.Body).Decode(&deletePayload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	title := deletePayload.Title

	tc.TaskService.Delete(r.Context(), title)
}
