package services

import (
	"errors"
	"testing"

	"github.com/BrooitsFeiskJR/digital-wallet-user-service/domain/dto"
	"github.com/BrooitsFeiskJR/digital-wallet-user-service/infra/mocks"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Common error variables for consistency
var (
	ErrUserNotFound = errors.New("user not found")
	ErrDatabase     = errors.New("database error")
)

func TestNewUserService(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mockRepo := &mocks.MockUserRepository{}
		service, err := NewUserService(mockRepo)
		assert.NoError(t, err)
		assert.NotNil(t, service)
	})

	t.Run("user repository is required", func(t *testing.T) {
		service, err := NewUserService(nil)
		assert.Error(t, err)
		assert.Equal(t, ErrNilRepository, err)
		assert.Nil(t, service)
	})
}

func TestUserService_GetUserById(t *testing.T) {
	tests := []struct {
		name      string
		userId    string
		mockFunc  func(id string) (*dto.UserDTO, error)
		wantUser  *dto.UserDTO
		wantError error
	}{
		{
			name:   "successful get user",
			userId: uuid.New().String(),
			mockFunc: func(id string) (*dto.UserDTO, error) {
				return &dto.UserDTO{
					ID:          uuid.MustParse(id),
					Name:        "Test User",
					Email:       "test@example.com",
					PhoneNumber: "1234567890",
					CPF:         "12345678901",
				}, nil
			},
			wantUser: &dto.UserDTO{
				Name:        "Test User",
				Email:       "test@example.com",
				PhoneNumber: "1234567890",
				CPF:         "12345678901",
			},
			wantError: nil,
		},
		{
			name:   "user not found",
			userId: uuid.New().String(),
			mockFunc: func(id string) (*dto.UserDTO, error) {
				return nil, ErrUserNotFound
			},
			wantUser:  nil,
			wantError: ErrUserNotFound,
		},
		{
			name:   "empty user id",
			userId: "",
			mockFunc: func(id string) (*dto.UserDTO, error) {
				return nil, errors.New("user id is required")
			},
			wantUser:  nil,
			wantError: errors.New("user id is required"),
		},
		{
			name:   "invalid user id format",
			userId: "invalid-uuid",
			mockFunc: func(id string) (*dto.UserDTO, error) {
				return nil, errors.New("invalid user id format")
			},
			wantUser:  nil,
			wantError: errors.New("invalid user id format"),
		},
		{
			name:   "database error",
			userId: uuid.New().String(),
			mockFunc: func(id string) (*dto.UserDTO, error) {
				return nil, ErrDatabase
			},
			wantUser:  nil,
			wantError: ErrDatabase,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup mock repository
			mockRepo := &mocks.MockUserRepository{
				GetUserByIdFunc: tt.mockFunc,
			}

			// Create service with mock repository
			userService, err := NewUserService(mockRepo)
			require.NoError(t, err)
			require.NotNil(t, userService)

			// Call method under test
			user, err := userService.GetUserById(tt.userId)

			// Check results
			if tt.wantError != nil {
				assert.Error(t, err)
				assert.Equal(t, tt.wantError.Error(), err.Error())
				assert.Nil(t, user)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, user)
				// Check ID only if the test expects to set it
				if tt.wantUser.ID != uuid.Nil {
					assert.Equal(t, tt.wantUser.ID, user.ID)
				}
				assert.Equal(t, tt.wantUser.Name, user.Name)
				assert.Equal(t, tt.wantUser.Email, user.Email)
				assert.Equal(t, tt.wantUser.PhoneNumber, user.PhoneNumber)
				assert.Equal(t, tt.wantUser.CPF, user.CPF)
			}
		})
	}
}

func TestUserService_UpdateUserById(t *testing.T) {
	// Generate test user ID
	userId := uuid.New().String()

	tests := []struct {
		name       string
		userId     string
		updateData *dto.UpdateUserDTO
		mockFunc   func(updateData *dto.UpdateUserDTO) (*dto.UpdatedUserDTO, error)
		wantUser   *dto.UpdatedUserDTO
		wantError  error
	}{
		{
			name:   "successful update",
			userId: userId,
			updateData: &dto.UpdateUserDTO{
				Name:        "Updated Name",
				PhoneNumber: "9876543210",
			},
			mockFunc: func(updateData *dto.UpdateUserDTO) (*dto.UpdatedUserDTO, error) {
				return &dto.UpdatedUserDTO{
					ID:          uuid.MustParse(userId),
					Name:        "Updated Name",
					Email:       "existing@example.com", // Unchanged field
					PhoneNumber: "9876543210",
					CPF:         "12345678901", // Unchanged field
				}, nil
			},
			wantUser: &dto.UpdatedUserDTO{
				ID:          uuid.MustParse(userId),
				Name:        "Updated Name",
				Email:       "existing@example.com",
				PhoneNumber: "9876543210",
				CPF:         "12345678901",
			},
			wantError: nil,
		},
		{
			name:   "user not found",
			userId: uuid.New().String(),
			updateData: &dto.UpdateUserDTO{
				Name: "Updated Name",
			},
			mockFunc: func(updateData *dto.UpdateUserDTO) (*dto.UpdatedUserDTO, error) {
				return nil, ErrUserNotFound
			},
			wantUser:  nil,
			wantError: ErrUserNotFound,
		},
		{
			name:       "nil update data",
			userId:     userId,
			updateData: nil,
			mockFunc: func(updateData *dto.UpdateUserDTO) (*dto.UpdatedUserDTO, error) {
				return nil, errors.New("update data is required")
			},
			wantUser:  nil,
			wantError: errors.New("update data is required"),
		},
		{
			name:   "empty user id",
			userId: "",
			updateData: &dto.UpdateUserDTO{
				Name: "Updated Name",
			},
			mockFunc: func(updateData *dto.UpdateUserDTO) (*dto.UpdatedUserDTO, error) {
				return nil, errors.New("user id is required")
			},
			wantUser:  nil,
			wantError: errors.New("user id is required"),
		},
		{
			name:       "no fields to update",
			userId:     userId,
			updateData: &dto.UpdateUserDTO{
				// No fields to update
			},
			mockFunc: func(updateData *dto.UpdateUserDTO) (*dto.UpdatedUserDTO, error) {
				return nil, errors.New("no fields to update")
			},
			wantUser:  nil,
			wantError: errors.New("no fields to update"),
		},
		{
			name:   "database error",
			userId: userId,
			updateData: &dto.UpdateUserDTO{
				Name: "Updated Name",
			},
			mockFunc: func(updateData *dto.UpdateUserDTO) (*dto.UpdatedUserDTO, error) {
				return nil, ErrDatabase
			},
			wantUser:  nil,
			wantError: ErrDatabase,
		},
		{
			name:   "mismatched ID",
			userId: userId,
			updateData: &dto.UpdateUserDTO{
				Name: "Updated Name",
			},
			mockFunc: func(updateData *dto.UpdateUserDTO) (*dto.UpdatedUserDTO, error) {
				// This should not be called due to validation
				return nil, errors.New("unexpected call")
			},
			wantUser:  nil,
			wantError: errors.New("mismatched user ids"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup mock repository
			mockRepo := &mocks.MockUserRepository{
				UpdateUserByIdFunc: tt.mockFunc,
			}

			// Create service with mock repository
			userService, err := NewUserService(mockRepo)
			require.NoError(t, err)
			require.NotNil(t, userService)

			// Call method under test
			updatedUser, err := userService.UpdateUserById(tt.userId, tt.updateData)

			// Check results
			if tt.wantError != nil {
				assert.Error(t, err)
				if tt.name == "mismatched ID" {
					assert.Equal(t, tt.wantError.Error(), err.Error())
				} else {
					assert.Contains(t, err.Error(), tt.wantError.Error())
				}
				assert.Nil(t, updatedUser)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, updatedUser)
				assert.Equal(t, tt.wantUser.ID, updatedUser.ID)
				assert.Equal(t, tt.wantUser.Name, updatedUser.Name)
				assert.Equal(t, tt.wantUser.Email, updatedUser.Email)
				assert.Equal(t, tt.wantUser.PhoneNumber, updatedUser.PhoneNumber)
				assert.Equal(t, tt.wantUser.CPF, updatedUser.CPF)
			}
		})
	}
}
