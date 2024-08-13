package main

import (
	"github.com/gin-gonic/gin"
	"github.com/solo21-12/A2SV_back_end_track/tree/main/task_seven/Delivery/routers"
	"github.com/solo21-12/A2SV_back_end_track/tree/main/task_seven/bootstrap"
)

func main() {
	app := bootstrap.App()
	env := app.Env
	db := app.Mongo.Database(env.MONGO_DATABASE)
	defer app.CloseDBConnection()

	gin := gin.Default()

	routers.Setup(env, db, gin)

	gin.Run(env.SERVER_ADDRESS)

}
