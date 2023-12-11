package router

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"sushkof_test/internal/controller"
)

func NewRouter(
	userController *controller.UserController,
	taskController *controller.TaskController,
	htmlController *controller.HTMLController,
	taskUserController *controller.TaskUserController) *httprouter.Router {

	router := httprouter.New()

	router.ServeFiles("/static/*filepath", http.Dir("../../static"))

	router.GET("/api/user/register", htmlController.RegisterHandler)
	router.GET("/api/user/list", htmlController.ListHandler)
	router.GET("/api/user/task", htmlController.TaskHandler)

	router.GET("/api/user/findAll", userController.FindAll)
	router.POST("/api/user/register_user", userController.Save)

	router.POST("/api/user/task/add", taskUserController.AddTaskToUser)
	router.POST("/api/tasks/getTask", taskUserController.GetTaskByName)

	router.POST("/api/user/task/delete", taskController.Delete)
	router.POST("/api/user/task/update", taskController.Update)

	return router
}
