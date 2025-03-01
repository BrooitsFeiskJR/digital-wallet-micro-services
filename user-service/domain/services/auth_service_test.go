package services

import (
	"errors"
	"testing"

	"github.com/BrooitsFeiskJR/digital-wallet-user-service/domain/dto"
	"github.com/BrooitsFeiskJR/digital-wallet-user-service/domain/entities"
	valueobject "github.com/BrooitsFeiskJR/digital-wallet-user-service/domain/value_object"
	"github.com/BrooitsFeiskJR/digital-wallet-user-service/infra/mocks"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

func TestAuthService_Login(t *testing.T) {
	tests := []struct {
		name      string
		email     string
		password  string
		mockFunc  func(email string) (*entities.User, error)
		wantUser  *entities.User
		wantError error
	}{
		{
			name:     "successful login",
			email:    "test@example.com",
			password: "password",
			mockFunc: func(email string) (*entities.User, error) {
				hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)
				return &entities.User{Email: email, Password: string(hashedPassword)}, nil
			},
			wantUser:  &entities.User{Email: "test@example.com"},
			wantError: nil,
		},
		{
			name:     "user not found",
			email:    "notfound@example.com",
			password: "password",
			mockFunc: func(email string) (*entities.User, error) {
				return nil, errors.New("user not found")
			},
			wantUser:  nil,
			wantError: errors.New("user not found"),
		},
		{
			name:     "invalid password",
			email:    "test@example.com",
			password: "wrongpassword",
			mockFunc: func(email string) (*entities.User, error) {
				hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)
				return &entities.User{Email: email, Password: string(hashedPassword)}, nil
			},
			wantUser:  nil,
			wantError: errors.New("invalid password"),
		},
		{
			name:     "database error",
			email:    "dberror@example.com",
			password: "password",
			mockFunc: func(email string) (*entities.User, error) {
				return nil, errors.New("database error")
			},
			wantUser:  nil,
			wantError: errors.New("database error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := &mocks.MockAuthRepository{
				LoginFunc: tt.mockFunc,
			}
			authService, err := NewAuthService(mockRepo)
			if err != nil {
				t.Fatal(err)
			}
			login := valueobject.LoginRequest{
				Email:    tt.email,
				Password: tt.password,
			}

			token, err := authService.Login(login)
			if err != nil {
				assert.Equal(t, tt.wantError.Error(), err.Error())
			} else {
				assert.Nil(t, tt.wantError)
			}
			assert.NotNil(t, token)
		})
	}
}

func TestAuthService_Register(t *testing.T) {
	tests := []struct {
		name      string
		req       *dto.CreateUserDTO
		mockFunc  func(req *dto.CreateUserDTO) (*dto.UserDTO, error)
		wantUser  *dto.UserDTO
		wantError error
	}{
		{
			name: "successful registration",
			req:  &dto.CreateUserDTO{Email: "test@example.com"},
			mockFunc: func(req *dto.CreateUserDTO) (*dto.UserDTO, error) {
				return &dto.UserDTO{Email: req.Email}, nil
			},
			wantUser:  &dto.UserDTO{Email: "test@example.com"},
			wantError: nil,
		},
		{
			name: "registration error",
			req:  &dto.CreateUserDTO{Email: "error@example.com"},
			mockFunc: func(req *dto.CreateUserDTO) (*dto.UserDTO, error) {
				return nil, errors.New("registration error")
			},
			wantUser:  nil,
			wantError: errors.New("registration error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := &mocks.MockAuthRepository{
				RegisterFunc: tt.mockFunc,
			}
			authService, err := NewAuthService(mockRepo)
			if err != nil {
				t.Fatal(err)
			}

			user, err := authService.Register(tt.req)
			assert.Equal(t, tt.wantUser, user)
			assert.Equal(t, tt.wantError, err)
		})
	}
}
