package controller

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"sushkof_test/internal/data/request"
	"sushkof_test/internal/service"
)

type TaskUserController struct {
	TaskUserService service.TaskUserService
}

var requestPayload struct {
	Username string `json:"username"`
}

func NewTaskUserController(tuService service.TaskUserService) *TaskUserController {
	return &TaskUserController{TaskUserService: tuService}
}

func (tuc *TaskUserController) GetTaskByName(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	err := json.NewDecoder(r.Body).Decode(&requestPayload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	username := requestPayload.Username

	tasks := tuc.TaskUserService.GetTaskByName(r.Context(), username)

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(tasks)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (tuc *TaskUserController) AddTaskToUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	var taskPayload request.AddTaskRequest

	err := json.NewDecoder(r.Body).Decode(&taskPayload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	tuc.TaskUserService.AddTaskToUser(r.Context(), taskPayload)
}
