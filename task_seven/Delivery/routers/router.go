package routers

import (
	"github.com/gin-gonic/gin"
	infrastructure "github.com/solo21-12/A2SV_back_end_track/tree/main/task_seven/Infrastructure"
	"github.com/solo21-12/A2SV_back_end_track/tree/main/task_seven/bootstrap"
	"go.mongodb.org/mongo-driver/mongo"
)

func Setup(env *bootstrap.Env, db *mongo.Database, gin *gin.Engine) {

	publicRouter := gin.Group("")
	NewSignupRouter(env, db, publicRouter)
	NewLoginRouter(env, db, publicRouter)

	protectedRouter := gin.Group("")
	protectedRouter.Use(infrastructure.AuthMiddleware())

	adminGroup := protectedRouter.Group("")
	adminGroup.Use(infrastructure.RoleBasedMiddleWare())
}
