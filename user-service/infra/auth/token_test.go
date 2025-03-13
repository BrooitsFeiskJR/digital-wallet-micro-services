package auth

import (
	"testing"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestMissingPrivateKey(t *testing.T) {
	originalPrivateKey := privateKey
	privateKey = nil
	defer func() { privateKey = originalPrivateKey }()

	id := uuid.New()
	name := "John Doe"
	email := "john.doe@example.com"

	accessToken, err := CreateAccessToken(id, name, email)
	assert.Error(t, err)
	assert.Equal(t, "private key not initialized", err.Error())
	assert.Empty(t, accessToken)

	refreshToken, err := CreateRefreshAcessToken(id, name, email)
	assert.Error(t, err)
	assert.Equal(t, "private key not initialized", err.Error())
	assert.Empty(t, refreshToken)
}

func TestTokenWithValidData(t *testing.T) {
	if privateKey == nil {
		t.Skip("Private key not initialized, skipping test")
	}

	id := uuid.New()
	name := "John Doe"
	email := "john.doe@example.com"

	accessToken, err := CreateAccessToken(id, name, email)
	assert.NoError(t, err)
	assert.NotEmpty(t, accessToken)

	token, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
		return &privateKey.PublicKey, nil
	})

	assert.NoError(t, err)
	assert.True(t, token.Valid)
	assert.NotNil(t, accessToken)

	claims, ok := token.Claims.(jwt.MapClaims)
	assert.True(t, ok)
	assert.Equal(t, id.String(), claims["id"])
	assert.Equal(t, name, claims["name"])
	assert.Equal(t, email, claims["email"])
	assert.NotEmpty(t, claims["exp"])
}
