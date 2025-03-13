package db

import (
	"errors"
	"sync"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var dbConn *sqlx.DB
var once sync.Once

func GetDBConnection(connString string) (*sqlx.DB, error) {
	if connString == "" {
		return nil, errors.New("connection string is empty")
	}
	var err error
	once.Do(func() {
		var db *sqlx.DB
		db, err = sqlx.Connect("postgres", connString)
		if err == nil {
			dbConn = db
		}
	})
	if err != nil {
		return nil, err
	}

	err = dbConn.Ping()
	if err != nil {
		return nil, err
	}
	return dbConn, nil
}
