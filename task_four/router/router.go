package router

import (
	"example.com/task_manager_api/controllers"

	"github.com/gin-gonic/gin"
)

func Run() {
	router := gin.Default()
	router.GET("/tasks", controller.GetTasksController)
	router.GET("/tasks/:id", controller.GetTaskByIDController)
	router.POST("/tasks", controller.PostTaskController)
	router.PUT("/tasks/:id", controller.UpdateTaskController)
	router.DELETE("/tasks/:id", controller.DeleteTaskController)

	router.Run(":8081")
}
