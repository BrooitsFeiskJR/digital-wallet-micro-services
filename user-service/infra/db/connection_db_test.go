package db

import (
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func TestConnectToDB(t *testing.T) {
	// Load environment variables from .env file
	err := godotenv.Load("../../.env")
	if err != nil {
		t.Fatalf("Error loading .env file")
	}

	// Test connection string is empty
	db, err := ConnectToDB("")
	assert.NotNil(t, err)
	assert.Nil(t, db)

	// Test invalid connection string
	invalidConnString := "invalid-connection-string"
	db, err = ConnectToDB(invalidConnString)
	assert.NotNil(t, err)
	assert.Nil(t, db)

	// Test valid connection string from environment variable
	validConnString := os.Getenv("DATABASE_URL")
	db, err = ConnectToDB(validConnString)
	assert.Nil(t, err)
	assert.NotNil(t, db)
}
