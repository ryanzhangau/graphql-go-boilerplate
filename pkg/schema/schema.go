package schema

import (
	"log"

	"github.com/graphql-go/graphql"
	"github.com/ryanzhangau/graphql-go/pkg/query"
)

func Schema() graphql.Schema {

	schemaConf := graphql.SchemaConfig{Query: graphql.NewObject(query.Query())}

	schema, err := graphql.NewSchema(schemaConf)

	if err != nil {
		log.Fatalf("failed to creaqte new schema, error: %v", err)
	}

	return schema
}
