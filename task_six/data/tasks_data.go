package data

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectMongo() *mongo.Database {

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	mongoURL := os.Getenv("MONGO_URL")

	if mongoURL == "" {
		log.Fatal("MONGO_URL is not set in .env file")
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
