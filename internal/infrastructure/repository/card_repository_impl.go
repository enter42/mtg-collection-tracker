package repository

import (
	"github.com/enter42/mtg-collection-tracker/internal/domain/entity"
	"github.com/enter42/mtg-collection-tracker/internal/domain/repository"
	"gorm.io/gorm"
)

type cardRepository struct {
	db *gorm.DB
}

func NewCardRepository(db *gorm.DB) repository.CardRepository {
	return &cardRepository{db: db}
}

func (r *cardRepository) Create(card *entity.Card) error {
	return r.db.Create(card).Error
}

func (r *cardRepository) Update(card *entity.Card) error {
	return r.db.Save(card).Error
}

func (r *cardRepository) Delete(id uint, userID uint) error {
	return r.db.Where("id = ? AND user_id = ?", id, userID).Delete(&entity.Card{}).Error
}

func (r *cardRepository) FindByID(id uint, userID uint) (*entity.Card, error) {
	var card entity.Card
	err := r.db.Where("user_id = ?", userID).First(&card, id).Error
	if err != nil {
		return nil, err
	}
	return &card, nil
}

func (r *cardRepository) FindByUserID(userID uint, page, pageSize int, search string) ([]entity.Card, int64, error) {
	var cards []entity.Card
	var total int64

	query := r.db.Model(&entity.Card{}).Where("user_id = ?", userID)

	// Apply search filter if provided
	if search != "" {
		searchPattern := "%" + search + "%"
		query = query.Where("card_name LIKE ? OR set_code LIKE ? OR collector_number LIKE ?",
			searchPattern, searchPattern, searchPattern)
	}

	// Get total count
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Get paginated results
	offset := (page - 1) * pageSize
	if err := query.Order("created_at DESC").Offset(offset).Limit(pageSize).Find(&cards).Error; err != nil {
		return nil, 0, err
	}

	return cards, total, nil
}
