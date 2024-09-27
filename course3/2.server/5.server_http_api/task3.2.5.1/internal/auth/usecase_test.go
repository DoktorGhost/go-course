package auth

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

// MockStorageRepository is a mock implementation of StorageRepository
type MockStorageRepository struct{}

func (m *MockStorageRepository) Create(login, password string) error {
	if login == "" || password == "" {
		return errors.New("error create")
	}
	return nil
}

func (m *MockStorageRepository) Read(login string) (string, error) {
	if login == "admin" {
		pass, _ := HashPassword("password")
		return pass, nil
	}
	return "", errors.New("user already exists")
}

func newMockRepo() *MockStorageRepository {
	return &MockStorageRepository{}
}

// Test Register method
func TestAuthUseCase_Register(t *testing.T) {
	mockRepo := new(MockStorageRepository)
	authUseCase := NewAuthUseCase(mockRepo, "secretKey")

	tests := []struct {
		name      string
		loginData Login
		wantErr   bool
	}{
		{
			name: "successful registration",
			loginData: Login{
				Username: "testuser",
				Password: "password123",
			},
			wantErr: false,
		},
		{
			name: "user already exists",
			loginData: Login{
				Username: "",
				Password: "",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := authUseCase.Register(tt.loginData)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

// Test Login method
func TestAuthUseCase_Login(t *testing.T) {
	mockRepo := newMockRepo()
	authUseCase := NewAuthUseCase(mockRepo, "secretKey")

	// Pre-register a user for login tests
	hashedPassword, _ := HashPassword("password123")
	mockRepo.Create("testuser", hashedPassword)

	tests := []struct {
		name      string
		loginData Login
		wantErr   bool
	}{
		{
			name: "successful login",
			loginData: Login{
				Username: "admin",
				Password: "password",
			},
			wantErr: false,
		},
		{
			name: "user not found",
			loginData: Login{
				Username: "nonexistentuser",
				Password: "password123",
			},
			wantErr: true,
		},
		{
			name: "invalid password",
			loginData: Login{
				Username: "testuser",
				Password: "wrongpassword",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := authUseCase.Login(tt.loginData)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
