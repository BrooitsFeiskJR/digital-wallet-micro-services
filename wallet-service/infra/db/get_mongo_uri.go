package db

import (
	"errors"
	"os"
)

func GetMongoURI() (string, error) {
	uri := os.Getenv("MONGO_URI")
	if uri == "" {
		return "", errors.New("env MONGO_URI is required")
	}

	return uri, nil
}
