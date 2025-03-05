package db

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func ConnectMongoDB(uri string) (*mongo.Client, error) {
	if uri == "" {
		return nil, errors.New("env MONGO_URI is required")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOptions := options.Client().
		ApplyURI(uri).
		SetConnectTimeout(5 * time.Second).
		SetServerSelectionTimeout(5 * time.Second)

	if poolSize := os.Getenv("MONGO_MAX_POOL_SIZE"); poolSize != "" {
		var maxPoolSize uint64
		fmt.Sscanf(poolSize, "%d", &maxPoolSize)
		if maxPoolSize > 0 {
			clientOptions.SetMaxPoolSize(maxPoolSize)
		}
	}

	log.Printf("Attempting to connect to MongoDB at: %s", uri)

	var client *mongo.Client
	var err error

	for i := 0; i < 5; i++ {
		client, err = mongo.Connect(ctx, clientOptions)
		if err == nil {
			pingCtx, pingCancel := context.WithTimeout(context.Background(), 2*time.Second)
			err = client.Ping(pingCtx, readpref.Primary())
			pingCancel()

			if err == nil {
				log.Println("Successfully connected to MongoDB!")
				return client, nil
			}
		}

		log.Printf("Failed to connect to MongoDB (attempt %d/5): %v", i+1, err)
		time.Sleep(2 * time.Second)
	}

	return nil, fmt.Errorf("failed to connect to MongoDB after multiple attempts: %w", err)
}
