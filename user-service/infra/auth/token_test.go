package auth

import (
	"os"
	"testing"

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
