package handlers

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/BrooitsFeiskJR/digital-wallet-user-service/domain/dto"
	valueobject "github.com/BrooitsFeiskJR/digital-wallet-user-service/domain/value_object"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockAuthService struct {
	mock.Mock
}

func (m *MockAuthService) Register(createUserDTO *dto.CreateUserDTO) (*dto.UserDTO, error) {
	args := m.Called(createUserDTO)
	if args.Get(0) != nil {
		return args.Get(0).(*dto.UserDTO), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockAuthService) Login(loginRequest valueobject.LoginRequest) (string, string, error) {
	args := m.Called(loginRequest)
	return args.String(0), args.String(0), args.Error(1)
}

func TestRegisterHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("should return 400 if request body is invalid", func(t *testing.T) {
		mockService := new(MockAuthService)
		handler, _ := NewAuthHandler(mockService)
		router := gin.Default()
		router.POST("/register", handler.RegisterHandler)

		req, _ := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer([]byte(`invalid`)))
		resp := httptest.NewRecorder()

		router.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusBadRequest, resp.Code)
	})

	t.Run("should return 500 if service returns an error", func(t *testing.T) {
		mockService := new(MockAuthService)
		handler, _ := NewAuthHandler(mockService)
		router := gin.Default()
		router.POST("/register", handler.RegisterHandler)

		createUserDTO := dto.CreateUserDTO{Email: "test@example.com", Password: "password"}
		mockService.On("Register", &createUserDTO).Return(nil, errors.New("service error"))

		body, _ := json.Marshal(createUserDTO)
		req, _ := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(body))
		resp := httptest.NewRecorder()

		router.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusInternalServerError, resp.Code)
	})

	t.Run("should return 200 if registration is successful", func(t *testing.T) {
		mockService := new(MockAuthService)
		handler, _ := NewAuthHandler(mockService)
		router := gin.Default()
		router.POST("/register", handler.RegisterHandler)

		createUserDTO := dto.CreateUserDTO{Name: "Luiz Antonio", Email: "test@example.com", Password: "password", ConfirmPassword: "password", PhoneNumber: "43999999999", CPF: "12345678909"}
		userDTO := &dto.UserDTO{ID: uuid.New(), Email: "test@example.com"}
		mockService.On("Register", &createUserDTO).Return(userDTO, nil)

		body, _ := json.Marshal(createUserDTO)
		req, _ := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(body))
		resp := httptest.NewRecorder()

		router.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusOK, resp.Code)
	})
}

func TestLoginUser(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("should return 400 if request body is invalid", func(t *testing.T) {
		mockService := new(MockAuthService)
		handler, _ := NewAuthHandler(mockService)
		router := gin.Default()
		router.POST("/login", handler.LoginUser)

		req, _ := http.NewRequest(http.MethodPost, "/login", bytes.NewBuffer([]byte(`invalid`)))
		resp := httptest.NewRecorder()

		router.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusBadRequest, resp.Code)
	})

	t.Run("should return 400 if login fails", func(t *testing.T) {
		mockService := new(MockAuthService)
		handler, _ := NewAuthHandler(mockService)
		router := gin.Default()
		router.POST("/login", handler.LoginUser)

		loginRequest := valueobject.LoginRequest{Email: "test@example.com", Password: "password"}
		mockService.On("Login", loginRequest).Return("", errors.New("login error"))

		body, _ := json.Marshal(loginRequest)
		req, _ := http.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(body))
		resp := httptest.NewRecorder()

		router.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusBadRequest, resp.Code)
	})

	t.Run("should return 200 if login is successful", func(t *testing.T) {
		mockService := new(MockAuthService)
		handler, _ := NewAuthHandler(mockService)
		router := gin.Default()
		router.POST("/login", handler.LoginUser)

		loginRequest := valueobject.LoginRequest{Email: "test@example.com", Password: "password"}
		mockService.On("Login", loginRequest).Return("token", nil)

		body, _ := json.Marshal(loginRequest)
		req, _ := http.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(body))
		resp := httptest.NewRecorder()

		router.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusOK, resp.Code)
	})
}

func TestNewAuthHandler(t *testing.T) {
	t.Run("should return an error if service is nil", func(t *testing.T) {
		handler, err := NewAuthHandler(nil)
		assert.Nil(t, handler)
		assert.NotNil(t, err)
	})

	t.Run("should return an instance of AuthHandler", func(t *testing.T) {
		mockService := new(MockAuthService)
		handler, err := NewAuthHandler(mockService)
		assert.NotNil(t, handler)
		assert.Nil(t, err)
	})
}
