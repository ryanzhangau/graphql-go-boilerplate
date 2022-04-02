package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/graphql-go/handler"
	"github.com/joho/godotenv"
	"github.com/ryanzhangau/graphql-go/pkg/schema"
)

// GetEnv load env file based on app environment
func GetEnv() error {

	env := os.Getenv("APP_ENV")
	var err error
	switch env {
	case "production":
		err = godotenv.Load(".env")
	case "development":
		err = godotenv.Load(".env.dev")
	case "ci":
		err = godotenv.Load(".env.ci")
	default:
		err = errors.New("cannot get app environnment, please specify one")
	}

	return err
}

// GetPort get port for graphql endpoint, default is :8080
func GetPort() string {
	port := os.Getenv("PORT")

	if port != "" {
		return ":" + port
	} else {
		return ":8080"
	}
}

func main() {

	// load env file according application environment
	err := GetEnv()

	if err != nil {
		log.Fatalf("get env file failed, error: %v", err)
	}

	schema := schema.Schema()

	handler := handler.New(&handler.Config{
		Schema:     &schema,
		Pretty:     true,
		GraphiQL:   true,
		Playground: os.Getenv("APP_ENV") == "development",
	})

	http.Handle("/graphql", handler)

	port := GetPort()

	fmt.Println("Server starts at port", port)
	http.ListenAndServe(port, nil)
}
