package db

import (
	"context"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func TestConnectMongoDB(t *testing.T) {
	t.Run("should return error when URI is empty", func(t *testing.T) {
		client, err := ConnectMongoDB("")
		assert.Error(t, err)
		assert.Nil(t, client)
		assert.Equal(t, "env MONGO_URI is required", err.Error())
	})

	t.Run("should return error when URI is invalid", func(t *testing.T) {
		client, err := ConnectMongoDB("invaliduri")
		assert.Error(t, err)
		assert.Nil(t, client)
	})

	t.Run("should connect successfully with valid URI", func(t *testing.T) {
		err := godotenv.Load("../../.env")
		if err != nil {
			t.Fatal("Error loading .env file")
		}
		mongoURI := os.Getenv("MONGO_TEST_URI")

		client, err := ConnectMongoDB(mongoURI)
		assert.NoError(t, err)
		assert.NotNil(t, client)

		// Clean up
		if client != nil {
			err := client.Disconnect(context.TODO())
			assert.NoError(t, err)
		}
	})
}
