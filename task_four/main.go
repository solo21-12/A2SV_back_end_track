package main

import (
	"example.com/task_manager_api/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/tasks", controller.GetTasksController)
	router.GET("/tasks/:id", controller.GetTaskByIDController)
	router.POST("/tasks", controller.PostTaskController)
	router.DELETE("/tasks/:id", controller.DeleteTaskController)

	router.Run(":8081")
}
