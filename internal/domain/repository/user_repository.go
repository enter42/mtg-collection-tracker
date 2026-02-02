package repository

import "github.com/enter42/mtg-collection-tracker/internal/domain/entity"

type UserRepository interface {
	Create(user *entity.User) error
	FindByUsername(username string) (*entity.User, error)
	FindByID(id uint) (*entity.User, error)
}
