package bootstrap

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type Application struct {
	Mongo *mongo.Client
	Env   *Env
}

func App() Application {

	app := Application{}
	app.Env = NewEnv()
	app.Mongo = NewMongoDatabase(app.Env)

	return app
}

func (app *Application) CloseDBConnection() {
	CloseMongoDBConnection(app.Mongo, context.TODO())
}
