package validation

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateEmailField(t *testing.T) {
	tests := []struct {
		name     string
		email    string
		expected bool
		err      error
	}{
		{
			name:     "Valid email",
			email:    "example@example.com",
			expected: true,
			err:      nil,
		},
		{
			name:     "Invalid email",
			email:    "example",
			expected: false,
			err:      nil,
		},
		{
			name:     "Empty email",
			email:    "",
			expected: false,
			err:      errors.New("email is required"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ValidateEmailField(tt.email)
			assert.Equal(t, tt.expected, result)
		})
	}
}
