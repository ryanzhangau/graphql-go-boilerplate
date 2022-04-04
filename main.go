package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/graphql-go/handler"
	"github.com/ryanzhangau/graphql-go/pkg/schema"
	"github.com/ryanzhangau/graphql-go/pkg/utilFunc"
)

func main() {

	// load env file according application environment
	err := utilFunc.GetEnv()

	if err != nil {
		log.Fatalf("get env file failed, error: %v", err)
	}

	// GraphQL schema
	schema := schema.Schema()

	handler := handler.New(&handler.Config{
		Schema:     &schema,
		Pretty:     true,
		GraphiQL:   true,
		Playground: os.Getenv("APP_ENV") == "development",
	})

	http.Handle("/graphql", handler)

	port := utilFunc.GetPort()

	fmt.Println("Server starts at port", port)
	http.ListenAndServe(port, nil)
}
