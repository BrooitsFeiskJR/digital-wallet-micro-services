package repositories

import (
	"errors"

	"github.com/jmoiron/sqlx"
)

type TransactionRepository struct {
	db *sqlx.DB
}

func NewTransactionRepository(db *sqlx.DB) (*TransactionRepository, error) {
	if db == nil {
		return nil, errors.New("db is required")
	}
	return &TransactionRepository{
		db: db,
	}, nil
}
