package repositories

import (
	"context"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/jmoiron/sqlx"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

var testDB *sqlx.DB

func TestMain(m *testing.M) {
	ctx := context.Background()

	container, dsn, err := startPostgresContainer(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer container.Terminate(ctx)

	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err = createTables(db); err != nil {
		log.Fatal(err)
	}

	testDB = db

	os.Exit(m.Run())
}

func startPostgresContainer(ctx context.Context) (testcontainers.Container, string, error) {
	req := testcontainers.ContainerRequest{
		// Changed image from "postgres:15.10-alpine" to "postgres:15.10" so that uuid-ossp extension is available
		Image:        "postgres:15.10",
		ExposedPorts: []string{"5432/tcp"},
		Env: map[string]string{
			"POSTGRES_DB":       "testDB",
			"POSTGRES_USER":     "user",
			"POSTGRES_PASSWORD": "password",
		},
		WaitingFor: wait.ForListeningPort("5432/tcp"),
	}

	postgresContainer, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		return nil, "", fmt.Errorf("failed to start container: %w", err)
	}

	port, err := postgresContainer.MappedPort(ctx, "5432")
	if err != nil {
		return nil, "", fmt.Errorf("failed to get mapped port: %w", err)
	}

	host, err := postgresContainer.Host(ctx)
	if err != nil {
		return nil, "", fmt.Errorf("failed to get host: %w", err)
	}

	dsn := fmt.Sprintf("postgres://user:password@%s:%s/testDB?sslmode=disable", host, port.Port())

	return postgresContainer, dsn, nil
}

func createTables(db *sqlx.DB) error {
	_, err := db.Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp" WITH SCHEMA public;`)
	if err != nil {
		return fmt.Errorf("failed to create uuid-ossp extension: %w", err)
	}

	// Create the users table
	usersQuery := `
        CREATE TABLE IF NOT EXISTS users (
            id UUID PRIMARY KEY,
            name TEXT NOT NULL,
            email TEXT UNIQUE NOT NULL,
            password TEXT NOT NULL,
			phone_number TEXT NOT NULL,
			cpf TEXT NOT NULL,
            created_at TIMESTAMP NOT NULL,
            updated_at TIMESTAMP NOT NULL
        );
    `
	_, err = db.Exec(usersQuery)
	if err != nil {
		return fmt.Errorf("failed to create users table: %w", err)
	}
	return nil
}
