package router

import (
	"example.com/task_manager_api/controllers"
	"example.com/task_manager_api/data"
	"example.com/task_manager_api/services"

	"github.com/gin-gonic/gin"
)

func Run() {
	db := data.ConnectMongo()

	taskService := services.NewTaskService(db)
	taskController := controller.NewTaskController(taskService)

	userService := services.NewUserService(db)
	userController := controller.NewUserController(userService)

	router := gin.Default()
	router.GET("/tasks", taskController.GetTasksController)
	router.GET("/tasks/:id", taskController.GetTaskByIDController)
	router.POST("/tasks", taskController.PostTaskController)
	router.PUT("/tasks/:id", taskController.UpdateTaskController)
	router.DELETE("/tasks/:id", taskController.DeleteTaskController)

	router.POST("/register", userController.RegisterUser)
	router.POST("/login", userController.LoginUser)

	router.Run(":8081")
}
