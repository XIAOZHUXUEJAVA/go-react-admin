package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHashPassword(t *testing.T) {
	tests := []struct {
		name     string
		password string
		wantErr  bool
	}{
		{
			name:     "valid password",
			password: "password123",
			wantErr:  false,
		},
		{
			name:     "empty password",
			password: "",
			wantErr:  false, // bcrypt can hash empty strings
		},
		{
			name:     "long password",
			password: "this_is_a_very_long_password_that_should_still_work_fine_123456789",
			wantErr:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hash, err := HashPassword(tt.password)
			
			if tt.wantErr {
				assert.Error(t, err)
				assert.Empty(t, hash)
			} else {
				assert.NoError(t, err)
				assert.NotEmpty(t, hash)
				assert.NotEqual(t, tt.password, hash) // Hash should be different from original
				assert.True(t, len(hash) > 50) // bcrypt hashes are typically 60 chars
			}
		})
	}
}

func TestCheckPassword(t *testing.T) {
	password := "testPassword123"
	hash, err := HashPassword(password)
	assert.NoError(t, err)

	tests := []struct {
		name     string
		password string
		hash     string
		want     bool
	}{
		{
			name:     "correct password",
			password: password,
			hash:     hash,
			want:     true,
		},
		{
			name:     "incorrect password",
			password: "wrongPassword",
			hash:     hash,
			want:     false,
		},
		{
			name:     "empty password",
			password: "",
			hash:     hash,
			want:     false,
		},
		{
			name:     "empty hash",
			password: password,
			hash:     "",
			want:     false,
		},
		{
			name:     "invalid hash format",
			password: password,
			hash:     "invalid_hash",
			want:     false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := CheckPassword(tt.password, tt.hash)
			assert.Equal(t, tt.want, result)
		})
	}
}

func TestPasswordHashingConsistency(t *testing.T) {
	password := "consistencyTest123"
	
	// Hash the same password multiple times
	hash1, err1 := HashPassword(password)
	hash2, err2 := HashPassword(password)
	
	assert.NoError(t, err1)
	assert.NoError(t, err2)
	
	// Hashes should be different (bcrypt uses salt)
	assert.NotEqual(t, hash1, hash2)
	
	// But both should validate the original password
	assert.True(t, CheckPassword(password, hash1))
	assert.True(t, CheckPassword(password, hash2))
}