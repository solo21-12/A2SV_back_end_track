package bootstrap

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewMongoDatabase(env *Env) *mongo.Client {
	// creating the context
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	dbUser := env.MONGO_USER
	dbPassword := env.MONGO_PASSWORD

	dbURL := fmt.Sprintf("mongodb+srv://%s:%s@mongodb-university.fs4tab8.mongodb.net/?retryWrites=true&w=majority&appName=mongodb-university", dbUser, dbPassword)
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
