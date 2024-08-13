package bootstrap

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewMongoDatabase(env *Env) *mongo.Client {
	// creating the context
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	dbURL := env.MONGO_URL
	clientOptions := options.Client().ApplyURI(dbURL)

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err.Error())
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatalf("Error pinging the database: %v", err.Error())
	}

	return client
}

func CloseMongoDBConnection(client *mongo.Client, ctx context.Context) {
	if client == nil {
		return
	}

	err := client.Disconnect(ctx)
	if err != nil {
		log.Fatalf("Error disconnecting from the database: %v", err.Error())
	}
}
