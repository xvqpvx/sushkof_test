package main

import (
	_ "github.com/go-sql-driver/mysql"
	"net/http"
	"sushkof_test/internal/controller"
	"sushkof_test/internal/db_config"
	"sushkof_test/internal/helper"
	"sushkof_test/internal/repository"
	"sushkof_test/internal/router"
	"sushkof_test/internal/service"
)

func main() {

	db := db_config.DatabaseConnection()

	userRepo := repository.NewUserRepository(db)
	taskRepo := repository.NewTaskRepository(db)

	userService := service.NewUserServiceImpl(userRepo)
	taskService := service.NewTaskServiceImpl(taskRepo)
	taskUserService := service.NewTaskUserServiceImpl(taskService, userService)

	userController := controller.NewUserController(userService)
	taskController := controller.NewTaskController(taskService)
	taskUserController := controller.NewTaskUserController(taskUserService)

	htmlController := controller.NewHTMLController()

	routes := router.NewRouter(
		userController,
		taskController,
		htmlController,
		taskUserController,
	)

	serv := http.Server{Addr: "localhost:8080", Handler: routes}
	err := serv.ListenAndServe()
	helper.PanicIfError(err)
}
