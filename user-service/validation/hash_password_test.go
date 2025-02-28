package validation

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

func TestHashPassword(t *testing.T) {
	password := "password"
	hash, err := HashPassword(password)
	if err != nil {
		t.Errorf("Error hashing password: %v", err)
	}

	err = bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		t.Errorf("Error comparing hash and password: %v", err)
	}
	assert.Equal(t, nil, err)
}
