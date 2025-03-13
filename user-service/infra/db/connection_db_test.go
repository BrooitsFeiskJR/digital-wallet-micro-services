package db

import (
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func TestGetDBConnection(t *testing.T) {
	// Load environment variables from .env file
	err := godotenv.Load("../../.env")
	if err != nil {
		t.Fatalf("Error loading .env file")
	}

	validConnString := os.Getenv("DATABASE_URL")
	db, err := GetDBConnection(validConnString)
	assert.Nil(t, err)
	assert.NotNil(t, db)
}
