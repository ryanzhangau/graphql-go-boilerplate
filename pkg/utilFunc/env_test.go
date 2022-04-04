package utilFunc

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetEnvFunc(t *testing.T) {
	os.Setenv("APP_ENV", "production")
	err := GetEnv()
	require.NoError(t, err)

	os.Setenv("APP_ENV", "development")
	err = GetEnv()
	require.NoError(t, err)

	os.Setenv("APP_ENV", "ci")
	err = GetEnv()
	require.NoError(t, err)

	os.Setenv("APP_ENV", "noEnv")
	err = GetEnv()
	require.EqualError(t, err, "cannot get app environnment, please specify one")
}

func TestGetPortFunc(t *testing.T) {
	os.Setenv("PORT", "")
	port := GetPort()
	require.Equal(t, port, ":8080")

	os.Setenv("PORT", "1234")
	port = GetPort()
	require.Equal(t, port, ":1234")
}
