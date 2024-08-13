package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/solo21-12/A2SV_back_end_track/tree/main/task_seven/Delivery/controllers"
	repositories "github.com/solo21-12/A2SV_back_end_track/tree/main/task_seven/Repositories"
	usecases "github.com/solo21-12/A2SV_back_end_track/tree/main/task_seven/UseCases"
	"github.com/solo21-12/A2SV_back_end_track/tree/main/task_seven/bootstrap"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewPromoteRouter(env *bootstrap.Env, db *mongo.Database, group *gin.RouterGroup) {
	userRepo := repositories.NewUserRepository(db, env.USER_COLLECTION)
	promoteUseCase := usecases.NewPromoteUseCase(userRepo)

	promoteController := controllers.PromoteController{
		PromoteUseCase: promoteUseCase,
		Env:            env,
	}

	group.PATCH("/promote/:id", promoteController.PromoteUser)
}
