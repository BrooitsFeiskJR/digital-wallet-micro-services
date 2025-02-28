package repositories

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/BrooitsFeiskJR/digital-wallet-user-service/domain/dto"
	"github.com/BrooitsFeiskJR/digital-wallet-user-service/domain/entities"
	"github.com/jmoiron/sqlx"
)

type AuthRepository struct {
	db *sqlx.DB
}

func NewAuthRepository(db *sqlx.DB) (*AuthRepository, error) {
	if db == nil {
		return nil, errors.New("db is required")
	}
	return &AuthRepository{
		db: db,
	}, nil
}

func (ar *AuthRepository) Login(email string) (*entities.User, error) {
	query := `
	SELECT * FROM users 
	WHERE email = $1
`
	var user entities.User
	err := ar.db.Get(&user, query, email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("user not found")
		}
		return nil, fmt.Errorf("failed to get user: %w", err)
	}
	return &user, nil
}

func (ar *AuthRepository) Register(req *dto.CreateUserDTO) (*dto.UserDTO, error) {
	user, err := entities.NewUser(req)
	if err != nil {
		return nil, err
	}
	tx, err := ar.db.Beginx()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()
	query := `
		INSERT INTO users (id, name, email, password, phone_number, cpf, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		RETURNING id, name, email, phone_number, cpf, created_at
	`
	var userDTO dto.UserDTO
	if err = tx.QueryRowx(query, user.ID, user.Name, user.Email, user.Password, user.PhoneNumber, user.CPF, user.CreatedAt, user.UpdateAt).StructScan(&userDTO); err != nil {
		return nil, err
	}
	if err = tx.Commit(); err != nil {
		return nil, err
	}

	return &userDTO, nil
}
