package middleware

import (
	"os"
	"testing"
	"time"

	"github.com/BrooitsFeiskJR/digital-wallet-wallet-service/infra/config/auth"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestVerifyTokenExpired(t *testing.T) {
	secretKey := []byte(os.Getenv("SECRECT_KEY"))

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
	err = auth.VerifyToken(tokenString, claims)
	assert.Error(t, err)
	assert.Equal(t, "Token is expired", err.Error())
}
