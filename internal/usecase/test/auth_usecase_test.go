package usecase_test

import (
"errors"
"testing"

"github.com/enter42/mtg-collection-tracker/internal/domain/entity"
"github.com/enter42/mtg-collection-tracker/internal/usecase"
)

// Mock repository for testing
type mockUserRepository struct {
users map[string]*entity.User
}

func newMockUserRepository() *mockUserRepository {
return &mockUserRepository{
users: make(map[string]*entity.User),
}
}

func (m *mockUserRepository) Create(user *entity.User) error {
m.users[user.Username] = user
return nil
}

func (m *mockUserRepository) FindByUsername(username string) (*entity.User, error) {
user, ok := m.users[username]
if !ok {
return nil, errors.New("record not found")
}
return user, nil
}

func (m *mockUserRepository) FindByID(id uint) (*entity.User, error) {
for _, user := range m.users {
if user.ID == id {
return user, nil
}
}
return nil, errors.New("record not found")
}

func TestAuthUseCase_Register(t *testing.T) {
repo := newMockUserRepository()
authUseCase := usecase.NewAuthUseCase(repo)

// Test successful registration
err := authUseCase.Register("testuser", "password123")
if err != nil {
t.Errorf("Expected no error, got %v", err)
}

// Verify user was created
user, err := repo.FindByUsername("testuser")
if err != nil {
t.Fatalf("Expected user to be found, got error: %v", err)
}
if user.Username != "testuser" {
t.Errorf("Expected username 'testuser', got %s", user.Username)
}

// Test duplicate username
err = authUseCase.Register("testuser", "password456")
if err == nil {
t.Error("Expected error for duplicate username, got nil")
}
}

func TestAuthUseCase_Login(t *testing.T) {
repo := newMockUserRepository()
authUseCase := usecase.NewAuthUseCase(repo)

// Register a user first
authUseCase.Register("testuser", "password123")

// Test successful login
user, err := authUseCase.Login("testuser", "password123")
if err != nil {
t.Errorf("Expected no error, got %v", err)
}
if user == nil {
t.Error("Expected user to be returned")
}
if user.Username != "testuser" {
t.Errorf("Expected username 'testuser', got %s", user.Username)
}

// Test login with wrong password
_, err = authUseCase.Login("testuser", "wrongpassword")
if err == nil {
t.Error("Expected error for wrong password, got nil")
}

// Test login with non-existent user
_, err = authUseCase.Login("nonexistent", "password123")
if err == nil {
t.Error("Expected error for non-existent user, got nil")
}
}
