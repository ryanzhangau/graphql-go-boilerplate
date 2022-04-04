package database

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func CreateContext() (context.Context, context.CancelFunc) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	return ctx, cancel
}

func GetMongoDBConnect() (*mongo.Client, error) {
	// mongoDB connection
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://root:123456@localhost:27017"))

	if err != nil {
		return nil, err
	}

	ctx, cancel := CreateContext()
	defer cancel()

	err = client.Connect(ctx)

	if err != nil {
		return nil, err
	}

	return client, nil
}
