package data

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectMongo() *mongo.Database {
	mongoURL := os.Getenv("MONGO_URL")

	if mongoURL == "" {
		log.Fatal("MONGO_URL is not set in the environment")
	}

	// options
	clientOptions := options.Client().ApplyURI(mongoURL)

	// client
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	// ping for connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	return client.Database("task_manager")
}
