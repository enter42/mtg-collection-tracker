package repository

import "github.com/enter42/mtg-collection-tracker/internal/domain/entity"

type CardRepository interface {
	Create(card *entity.Card) error
	Update(card *entity.Card) error
	Delete(id uint, userID uint) error
	FindByID(id uint, userID uint) (*entity.Card, error)
	FindByUserID(userID uint, page, pageSize int, search string) ([]entity.Card, int64, error)
}
