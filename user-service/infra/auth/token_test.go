package auth

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestCreateToken(t *testing.T) {

	id := uuid.New()
	name := "John Doe"
	email := "john.doe@example.com"

	tokenString, err := CreateAccessToken(id, name, email)
	assert.NoError(t, err)
	assert.NotEmpty(t, tokenString)
}

func TestCreateRefreshToken(t *testing.T) {
	id := uuid.New()
	name := "John Doe"
	email := "john.doe@example.com"

	tokenString, err := CreateRefreshAcessToken(id, name, email)
	assert.NoError(t, err)
	assert.NotEmpty(t, tokenString)
}
