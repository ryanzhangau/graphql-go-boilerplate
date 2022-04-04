package query

import (
	"log"

	"github.com/graphql-go/graphql"
	"github.com/ryanzhangau/graphql-go/pkg/database"
	"go.mongodb.org/mongo-driver/bson"
)

func Hello() *graphql.Field {
	return &graphql.Field{
		Type: graphql.String,
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			return "world", nil
		},
	}
}

func Databases() *graphql.Field {
	return &graphql.Field{
		Type: graphql.NewList(graphql.String),
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			db, err := database.GetMongoDBConnect()

			if err != nil {
				log.Fatalf("database connect error: %v", err)
			}

			ctx, cancel := db.CreateContext()
			defer cancel()
			defer db.Client.Disconnect(ctx)

			databases, err := db.Client.ListDatabaseNames(ctx, bson.M{})
			if err != nil {
				log.Fatalf("get database names error: %v", err)
			}

			return databases, err
		},
	}
}
