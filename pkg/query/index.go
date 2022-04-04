package query

import "github.com/graphql-go/graphql"

func Query() graphql.ObjectConfig {
	fields := graphql.Fields{
		"hello":     Hello(),
		"databases": Databases(),
	}

	rootQuery := graphql.ObjectConfig{Name: "RootQuery", Fields: fields}
	return rootQuery
}
