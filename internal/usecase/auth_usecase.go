package usecase

import (
	"errors"

	"github.com/enter42/mtg-collection-tracker/internal/domain/entity"
	"github.com/enter42/mtg-collection-tracker/internal/domain/repository"
	"golang.org/x/crypto/bcrypt"
)

type AuthUseCase struct {
	userRepo repository.UserRepository
}

func NewAuthUseCase(userRepo repository.UserRepository) *AuthUseCase {
	return &AuthUseCase{userRepo: userRepo}
}

func (uc *AuthUseCase) Register(username, password string) error {
	// Check if user already exists
	_, err := uc.userRepo.FindByUsername(username)
	if err == nil {
		return errors.New("username already exists")
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user := &entity.User{
		Username: username,
		Password: string(hashedPassword),
	}

	return uc.userRepo.Create(user)
}

func (uc *AuthUseCase) Login(username, password string) (*entity.User, error) {
	user, err := uc.userRepo.FindByUsername(username)
	if err != nil {
		return nil, errors.New("invalid username or password")
	}

	// Check password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, errors.New("invalid username or password")
	}

	return user, nil
}

func (uc *AuthUseCase) GetUserByID(id uint) (*entity.User, error) {
	return uc.userRepo.FindByID(id)
}
