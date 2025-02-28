package auth

import (
	"os"
	"testing"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestCreateToken(t *testing.T) {
	secretKey = []byte(os.Getenv("SECRECT_KEY"))

	id := uuid.New()
	name := "John Doe"
	email := "john.doe@example.com"

	tokenString, err := CreateToken(id, name, email)
	assert.NoError(t, err)
	assert.NotEmpty(t, tokenString)
}

func TestVerifyToken(t *testing.T) {
	secretKey = []byte(os.Getenv("SECRECT_KEY"))

	id := uuid.New()
	name := "John Doe"
	email := "john.doe@example.com"

	tokenString, err := CreateToken(id, name, email)
	assert.NoError(t, err)
	assert.NotEmpty(t, tokenString)

	claims := jwt.MapClaims{}
	err = VerifyToken(tokenString, claims)
	assert.NoError(t, err)
	assert.Equal(t, id.String(), claims["id"])
	assert.Equal(t, name, claims["name"])
	assert.Equal(t, email, claims["email"])
}

func TestVerifyTokenExpired(t *testing.T) {
	secretKey = []byte(os.Getenv("SECRECT_KEY"))

	id := uuid.New()
	name := "John Doe"
	email := "john.doe@example.com"

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":    id,
		"name":  name,
		"email": email,
		"exp":   time.Now().Add(-time.Hour).Unix(),
	})

	tokenString, err := token.SignedString(secretKey)
	assert.NoError(t, err)
	assert.NotEmpty(t, tokenString)

	claims := jwt.MapClaims{}
	err = VerifyToken(tokenString, claims)
	assert.Error(t, err)
	assert.Equal(t, "Token is expired", err.Error())
}
