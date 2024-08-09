package router

import (
	"example.com/task_manager_api/controllers"
	"example.com/task_manager_api/data"

	"github.com/gin-gonic/gin"
)

func Run() {
	db := data.ConnectMongo()

	taskService := data.NewTaskService(db)
	taskController := controller.NewTaskController(taskService)
	router := gin.Default()
	router.GET("/tasks", taskController.GetTasksController)
	router.GET("/tasks/:id", taskController.GetTaskByIDController)
	router.POST("/tasks", taskController.PostTaskController)
	router.PUT("/tasks/:id", taskController.UpdateTaskController)
	router.DELETE("/tasks/:id", taskController.DeleteTaskController)

	router.Run(":8081")
}
