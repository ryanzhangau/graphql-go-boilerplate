package utilFunc

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
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
