package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/solo21-12/A2SV_back_end_track/tree/main/task_seven/Delivery/controllers"
	repositories "github.com/solo21-12/A2SV_back_end_track/tree/main/task_seven/Repositories"
	usecases "github.com/solo21-12/A2SV_back_end_track/tree/main/task_seven/UseCases"
	"github.com/solo21-12/A2SV_back_end_track/tree/main/task_seven/bootstrap"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewTaskRouter(env *bootstrap.Env, db *mongo.Database, group *gin.RouterGroup) {
	taskRepo := repositories.NewTaskRepository(db, env.TASK_COLLECTION)
	taskUseCase := usecases.NewTaskUseCase(taskRepo)

	taskController := controllers.TaskController{
		TaskUseCase: taskUseCase,
	}

	group.POST("/tasks", taskController.Create)
	group.GET("/tasks", taskController.GetAll)
	group.GET("/tasks/:id", taskController.Get)
	group.DELETE("/tasks/:id", taskController.Delete)
	group.PUT("/tasks/:id", taskController.Update)
}
