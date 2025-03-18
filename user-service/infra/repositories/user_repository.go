package repositories

import (
	"errors"

	"github.com/BrooitsFeiskJR/digital-wallet-user-service/domain/dto"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type UserRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) (*UserRepository, error) {
	if db == nil {
		return nil, errors.New("database is required")
	}
	return &UserRepository{
		db: db,
	}, nil
}

func (ur *UserRepository) GetUserById(id string) (*dto.UserDTO, error) {
	tx, err := ur.db.Beginx()
	if err != nil {
		return nil, errors.New("failed to begin transaction")
	}
	defer tx.Rollback()

	query := `SELECT id, name, email, phone_number, cpf, created_at from users WHERE id = $1`
	var userDTO dto.UserDTO
	if err := tx.Get(&userDTO, query, id); err != nil {
		return nil, errors.New("failed to get user")
	}
	return &userDTO, nil
}

// TODO: Add validation for check if the update data is not the same as the current storage data
func (ur *UserRepository) UpdateUserById(id uuid.UUID, updateData *dto.UpdateUserDTO) (*dto.UpdatedUserDTO, error) {
	if id == uuid.Nil {
		return nil, errors.New("id is required")
	}
	if updateData == nil {
		return nil, errors.New("update data is required")
	}

	tx, err := ur.db.Beginx()
	if err != nil {
		return nil, errors.New("failed to begin transaction")
	}
	defer tx.Rollback()

	query := "UPDATE users SET "
	params := []interface{}{}
	paramCount := 1

	setClauses := []string{}

	if updateData.Name != "" {
		setClauses = append(setClauses, "name = $"+string(rune('0'+paramCount)))
		params = append(params, updateData.Name)
		paramCount++
	}

	if updateData.PhoneNumber != "" {
		setClauses = append(setClauses, "phone_number = $"+string(rune('0'+paramCount)))
		params = append(params, updateData.PhoneNumber)
		paramCount++
	}

	setClauses = append(setClauses, "updated_at = NOW()")

	if len(setClauses) <= 1 {
		return nil, errors.New("no fields to update")
	}

	query += getStrings(setClauses)
	query += " WHERE id = $" + string(rune('0'+paramCount))
	params = append(params, id)

	_, err = tx.Exec(query, params...)
	if err != nil {
		return nil, errors.New("failed to update user: " + err.Error())
	}

	getUserQuery := `SELECT id, name, email, phone_number, cpf, created_at, updated_at FROM users WHERE id = $1`
	var updatedUser dto.UpdatedUserDTO
	err = tx.Get(&updatedUser, getUserQuery, id)
	if err != nil {
		return nil, errors.New("failed to retrieve updated user: " + err.Error())
	}

	if err = tx.Commit(); err != nil {
		return nil, errors.New("failed to commit transaction: " + err.Error())
	}

	return &updatedUser, nil
}

func getStrings(clauses []string) string {
	result := ""
	for i, clause := range clauses {
		if i > 0 {
			result += ", "
		}
		result += clause
	}
	return result
}
