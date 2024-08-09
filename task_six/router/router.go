package router

import (
	"example.com/task_manager_api/controllers"
	"example.com/task_manager_api/data"
	"example.com/task_manager_api/middleware"
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

	protected := router.Group("/")
	protected.Use(middleware.AuthMiddleware())
	{
		protected.GET("/tasks", taskController.GetTasksController)
		protected.GET("/tasks/:id", taskController.GetTaskByIDController)

		adminGroup := protected.Group("/")
		adminGroup.Use(middleware.RoleBasedMiddleWare("admin"))

		adminGroup.POST("/tasks", taskController.PostTaskController)
		adminGroup.PUT("/tasks/:id", taskController.UpdateTaskController)
		adminGroup.DELETE("/tasks/:id", taskController.DeleteTaskController)
		adminGroup.PATCH("/promote", userController.PromoteUser)
	}

	router.POST("/register", userController.RegisterUser)
	router.POST("/login", userController.LoginUser)

	router.Run(":8081")
}
