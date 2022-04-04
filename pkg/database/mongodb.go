package database

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/ryanzhangau/graphql-go/pkg/utilFunc"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Database struct {
	Client *mongo.Client
}

func (d Database) CreateContext() (context.Context, context.CancelFunc) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	return ctx, cancel
}

// GetMongoDBConnect connect mongo db and return database struct
func GetMongoDBConnect() (*Database, error) {
	err := utilFunc.GetEnv()

	if err != nil {
		log.Fatalf("get env file failed, error: %v", err)
	}

	// mongoDB connection
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://" + os.Getenv("DB_USER") + ":" + os.Getenv("DB_PASS") + "@localhost:27017"))

	if err != nil {
		return nil, err
	}

	db := Database{client}

	ctx, cancel := db.CreateContext()
	defer cancel()

	err = client.Connect(ctx)

	if err != nil {
		return nil, err
	}

	return &Database{Client: client}, nil
}
