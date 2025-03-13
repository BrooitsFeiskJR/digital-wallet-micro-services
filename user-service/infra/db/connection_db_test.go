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

func TestGetDBConnectionWithEmptyConnString(t *testing.T) {
	validConnString := ""
	db, err := GetDBConnection(validConnString)
	assert.Nil(t, db)
	assert.NotNil(t, err)
	assert.Equal(t, "connection string is empty", err.Error())
}

func TestGetDBConnectionWithMultipleCalls(t *testing.T) {
	err := godotenv.Load("../../.env")
	if err != nil {
		t.Fatalf("Error loading .env file")
	}
	validConnString := os.Getenv("DATABASE_URL")
	db1, err1 := GetDBConnection(validConnString)
	assert.Nil(t, err1)
	assert.NotNil(t, db1)

	db2, err2 := GetDBConnection(validConnString)
	assert.Nil(t, err2)
	assert.NotNil(t, db2)

	assert.Equal(t, db1, db2)
}
