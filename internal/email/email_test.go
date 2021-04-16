package email

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidate(t *testing.T) {
	examples := []struct {
		email string
		valid bool
	}{
		{"test@test.com", true},
		{"test@test@.com", false},
		{"test@.com", false},
		{"test@test.c", false},
	}

	for _, tt := range examples {
		t.Run(tt.email, func(t *testing.T) {
			v := Validate(tt.email)
			assert.Equal(t, tt.valid, v)
		})
	}
}

func TestSplit_ValidEmail(t *testing.T) {
	examples := []struct {
		email       string
		local, host string
	}{
		{"test@test.com", "test", "test.com"},
		{"user@gmail.com", "user", "gmail.com"},
	}

	for _, tt := range examples {
		t.Run(tt.email, func(t *testing.T) {
			local, host, err := Split(tt.email)
			assert.Nil(t, err)
			assert.Equal(t, tt.local, local)
			assert.Equal(t, tt.host, host)
		})
	}
}

func TestSplit_InvalidEmail(t *testing.T) {
	input := "invalid@email"
	_, _, err := Split(input)
	assert.NotNil(t, err)

}
