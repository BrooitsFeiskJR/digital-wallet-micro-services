package db

import (
	"errors"
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func ConnectToDB(connString string) (*sqlx.DB, error) {
	if connString == "" {
		return nil, errors.New("connection string is empty")
	}
	db, err := sqlx.Connect("postgres", connString)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	fmt.Println("Connected to database!")
	return db, nil
}
