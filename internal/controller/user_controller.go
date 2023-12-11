package controller

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
	"sushkof_test/internal/data/request"
	"sushkof_test/internal/helper"
	"sushkof_test/internal/service"
)

type UserController struct {
	UserService service.UserService
}

func NewUserController(userService service.UserService) *UserController {
	return &UserController{UserService: userService}
}

func (uc *UserController) Save(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	name := r.FormValue("name")
	age := r.FormValue("age")
	email := r.FormValue("email")
	password := r.FormValue("password")

	ageInt, err := strconv.Atoi(age)
	if err != nil {
		http.Error(w, "Invalid age value", http.StatusBadRequest)
		return
	}

	userSaveRequest := request.SaveUserRequest{
		Name:     name,
		Age:      uint8(ageInt),
		Email:    email,
		Password: password,
	}

	uc.UserService.Save(r.Context(), userSaveRequest)

	http.Redirect(w, r, "/api/user/list", http.StatusSeeOther)
}

func (uc *UserController) FindAll(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	users := uc.UserService.FindAll(r.Context())

	jsonResponse, err := json.Marshal(users)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	_, err = w.Write(jsonResponse)
	helper.PanicIfError(err)
}
