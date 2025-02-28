package validation

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckHashPassword(t *testing.T) {
	hashPassword, _ := HashPassword("password")
	providedPassword := "password"
	err := CheckHashPassword(hashPassword, providedPassword)
	if err != nil {
		t.Errorf("Error checking hash password: %v", err)
	}
	assert.Equal(t, nil, err)
}
