package repositories

import (
	"testing"
	"time"

	"github.com/BrooitsFeiskJR/digital-wallet-user-service/domain/dto"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

func TestAuthRepository_LoginWithInvalidDB(t *testing.T) {
	_, err := NewAuthRepository(nil)
	assert.Error(t, err)
}

func TestAuthRepository_Login(t *testing.T) {
	repo, err := NewAuthRepository(testDB)
	assert.NoError(t, err)

	// Insert a test user
	_, err = testDB.Exec(`INSERT INTO users (id, name, email, password, phone_number, cpf, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`,
		uuid.New(), "Test User", "test@example.com", "password", "43999999999", "12345678901", time.Now(), time.Now())
	assert.NoError(t, err)

	user, err := repo.Login("test@example.com")
	assert.NoError(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, "Test User", user.Name)
}

func TestAuthRepository_Register(t *testing.T) {
	repo, err := NewAuthRepository(testDB)
	assert.NoError(t, err)

	req := &dto.CreateUserDTO{
		Name:        "New User",
		Email:       "newuser@example.com",
		Password:    "password",
		PhoneNumber: "43999999999",
		CPF:         "12345678909",
	}

	userDTO, err := repo.Register(req)
	assert.NoError(t, err)
	assert.NotNil(t, userDTO)
	assert.Equal(t, "New User", userDTO.Name)
	assert.Equal(t, "newuser@example.com", userDTO.Email)
}
