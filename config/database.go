package config

import (
	"context"
	"log"
	"time"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

var Client *mongo.Client

func ConnectDb() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(
		options.Client().ApplyURI("mongodb://localhost:27017"),
	)

	if err != nil {
		log.Fatal("Failed to connect to MongoDB:", err)
	}

	err = client.Ping(ctx, nil)

	if err != nil {
		log.Fatal("MongoDB ping failed:", err)
	}

	Client = client

	log.Println("MongoDB Connected Successfully")
}