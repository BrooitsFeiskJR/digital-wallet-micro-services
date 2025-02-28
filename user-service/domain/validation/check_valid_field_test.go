package validation

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsValidField(t *testing.T) {
	tests := []struct {
		name     string
		field    string
		expected bool
	}{
		{
			name:     "Valid field",
			field:    "example",
			expected: true,
		},
		{
			name:     "Empty field",
			field:    "",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := isValidField(tt.field)
			assert.Equal(t, tt.expected, result)
		})
	}
}
