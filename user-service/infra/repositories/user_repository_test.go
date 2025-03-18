package repositories

import (
	"testing"
	"time"

	"github.com/BrooitsFeiskJR/digital-wallet-user-service/domain/dto"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUserRepository_NewUserRepositoryWithInvalidDB(t *testing.T) {
	// Test repository creation with nil database
	repo, err := NewUserRepository(nil)
	assert.Error(t, err)
	assert.Nil(t, repo)
	assert.Contains(t, err.Error(), "database is required")
}

func TestUserRepository_GetUserById(t *testing.T) {
	// Create a valid repository instance
	repo, err := NewUserRepository(testDB)
	require.NoError(t, err)
	require.NotNil(t, repo)

	// Setup test data
	userId := uuid.New()
	userName := "Get User Test"
	userEmail := "getuser@example.com"

	// Insert a test user
	_, err = testDB.Exec(`
        INSERT INTO users (id, name, email, password, phone_number, cpf, created_at, updated_at) 
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`,
		userId, userName, userEmail, "password123", "55999999999", "12345678901",
		time.Now(), time.Now())
	require.NoError(t, err)

	// Test cases
	t.Run("Get existing user", func(t *testing.T) {
		user, err := repo.GetUserById(userId.String())
		assert.NoError(t, err)
		assert.NotNil(t, user)
		assert.Equal(t, userId, user.ID)
		assert.Equal(t, userName, user.Name)
		assert.Equal(t, userEmail, user.Email)
		assert.Equal(t, "55999999999", user.PhoneNumber)
		assert.Equal(t, "12345678901", user.CPF)
	})

	t.Run("Get non-existing user", func(t *testing.T) {
		nonExistingId := uuid.New().String()
		user, err := repo.GetUserById(nonExistingId)
		assert.Error(t, err)
		assert.Nil(t, user)
		assert.Contains(t, err.Error(), "failed to get user")
	})

	t.Run("Get user with empty ID", func(t *testing.T) {
		user, err := repo.GetUserById("")
		assert.Error(t, err)
		assert.Nil(t, user)
		// The specific error message may vary based on your DB driver
	})

	t.Run("Get user with invalid ID format", func(t *testing.T) {
		user, err := repo.GetUserById("not-a-uuid")
		assert.Error(t, err)
		assert.Nil(t, user)
	})
}

func TestUserRepository_UpdateUserById(t *testing.T) {
	// Create a valid repository instance
	repo, err := NewUserRepository(testDB)
	require.NoError(t, err)
	require.NotNil(t, repo)

	// Setup test data
	userId := uuid.New()
	originalName := "Update User Test"
	originalEmail := "updateuser@example.com"
	originalPhone := "55888888888"
	originalCPF := "98765432101"

	// Insert a test user
	_, err = testDB.Exec(`
        INSERT INTO users (id, name, email, password, phone_number, cpf, created_at, updated_at) 
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`,
		userId, originalName, originalEmail, "password123", originalPhone, originalCPF,
		time.Now(), time.Now())
	require.NoError(t, err)

	// Test cases
	t.Run("Update user name", func(t *testing.T) {
		updateData := &dto.UpdateUserDTO{
			Name: "Updated Name",
		}

		updatedUser, err := repo.UpdateUserById(userId, updateData)
		assert.NoError(t, err)
		assert.NotNil(t, updatedUser)
		assert.Equal(t, userId, updatedUser.ID)
		assert.Equal(t, "Updated Name", updatedUser.Name)       // Verify name was updated
		assert.Equal(t, originalEmail, updatedUser.Email)       // Verify email wasn't changed
		assert.Equal(t, originalPhone, updatedUser.PhoneNumber) // Verify phone wasn't changed
	})

	t.Run("Update user phone number", func(t *testing.T) {
		updateData := &dto.UpdateUserDTO{
			PhoneNumber: "55777777777",
		}

		updatedUser, err := repo.UpdateUserById(userId, updateData)
		assert.NoError(t, err)
		assert.NotNil(t, updatedUser)
		assert.Equal(t, userId, updatedUser.ID)
		assert.Equal(t, "Updated Name", updatedUser.Name)       // Should still have the name from previous test
		assert.Equal(t, "55777777777", updatedUser.PhoneNumber) // Verify phone was updated
	})

	t.Run("Update multiple fields", func(t *testing.T) {
		updateData := &dto.UpdateUserDTO{
			Name:        "Multiple Update Name",
			PhoneNumber: "55666666666",
		}

		updatedUser, err := repo.UpdateUserById(userId, updateData)
		assert.NoError(t, err)
		assert.NotNil(t, updatedUser)
		assert.Equal(t, userId, updatedUser.ID)
		assert.Equal(t, "Multiple Update Name", updatedUser.Name)
		assert.Equal(t, "55666666666", updatedUser.PhoneNumber)
		assert.Equal(t, originalEmail, updatedUser.Email) // Verify email wasn't changed
	})

	t.Run("Update with nil update data", func(t *testing.T) {
		updatedUser, err := repo.UpdateUserById(uuid.Nil, nil)
		assert.Error(t, err)
		assert.Nil(t, updatedUser)
		assert.Contains(t, err.Error(), "id is required")
	})

	t.Run("Update non-existent user", func(t *testing.T) {
		nonExistingId := uuid.New()
		updateData := &dto.UpdateUserDTO{
			Name: "Non-existent User",
		}

		updatedUser, err := repo.UpdateUserById(nonExistingId, updateData)
		assert.Error(t, err)
		assert.Nil(t, updatedUser)
		assert.Contains(t, err.Error(), "failed to retrieve updated user")
	})

	t.Run("Update with no fields to update", func(t *testing.T) {
		updateData := &dto.UpdateUserDTO{
			// No fields to update
		}

		updatedUser, err := repo.UpdateUserById(userId, updateData)
		assert.Error(t, err)
		assert.Nil(t, updatedUser)
		assert.Contains(t, err.Error(), "no fields to update")
	})
}

func TestUserRepository_Integration(t *testing.T) {
	// Create a valid repository instance
	repo, err := NewUserRepository(testDB)
	require.NoError(t, err)
	require.NotNil(t, repo)

	// Setup: Insert a new user via register repository
	authRepo, err := NewAuthRepository(testDB)
	require.NoError(t, err)

	// Create user via auth repository
	createUserDTO := &dto.CreateUserDTO{
		Name:            "Integration Test User",
		Email:           "integration@example.com",
		Password:        "securepassword",
		ConfirmPassword: "securepassword",
		PhoneNumber:     "55444444444",
		CPF:             "05161891076",
	}

	userDTO, err := authRepo.Register(createUserDTO)
	require.NoError(t, err)
	require.NotNil(t, userDTO)
	userId := userDTO.ID

	// Test: Get user by ID
	retrievedUser, err := repo.GetUserById(userId.String())
	require.NoError(t, err)
	require.NotNil(t, retrievedUser)
	assert.Equal(t, "Integration Test User", retrievedUser.Name)
	assert.Equal(t, "integration@example.com", retrievedUser.Email)

	// Test: Update user
	updateData := &dto.UpdateUserDTO{
		Name:        "Updated Integration User",
		PhoneNumber: "55333333333",
	}

	updatedUser, err := repo.UpdateUserById(userId, updateData)
	require.NoError(t, err)
	require.NotNil(t, updatedUser)
	assert.Equal(t, "Updated Integration User", updatedUser.Name)
	assert.Equal(t, "55333333333", updatedUser.PhoneNumber)
	assert.Equal(t, "integration@example.com", updatedUser.Email) // Email should be unchanged
}

// Helper function to clean up the database between tests
func clearUserTable() error {
	_, err := testDB.Exec("DELETE FROM users")
	return err
}

// Optional: If you want to run each test with a clean state
func TestUserRepository_WithCleanState(t *testing.T) {
	// Clean up before test
	err := clearUserTable()
	require.NoError(t, err)

	repo, err := NewUserRepository(testDB)
	require.NoError(t, err)

	// Setup test data
	userId := uuid.New()
	userName := "Clean State Test"
	userEmail := "cleanstate@example.com"

	// Insert a test user
	_, err = testDB.Exec(`
        INSERT INTO users (id, name, email, password, phone_number, cpf, created_at, updated_at) 
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`,
		userId, userName, userEmail, "password123", "5911111111", "11111111111",
		time.Now(), time.Now())
	require.NoError(t, err)

	// Verify user was inserted
	user, err := repo.GetUserById(userId.String())
	assert.NoError(t, err)
	assert.Equal(t, userName, user.Name)

	// Clean up after test
	err = clearUserTable()
	require.NoError(t, err)
}
