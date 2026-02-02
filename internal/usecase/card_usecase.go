package usecase

import (
	"time"

	"github.com/enter42/mtg-collection-tracker/internal/domain/entity"
	"github.com/enter42/mtg-collection-tracker/internal/domain/repository"
)

type CardUseCase struct {
	cardRepo repository.CardRepository
}

func NewCardUseCase(cardRepo repository.CardRepository) *CardUseCase {
	return &CardUseCase{cardRepo: cardRepo}
}

type CreateCardInput struct {
	UserID          uint
	CardName        string
	CardImageURL    string
	SetCode         string
	CollectorNumber string
	Language        string
	Quantity        int
	BuyingPrice     float64
	BoughtDate      *time.Time
	SellDate        *time.Time
}

type UpdateCardInput struct {
	ID              uint
	UserID          uint
	CardName        string
	CardImageURL    string
	SetCode         string
	CollectorNumber string
	Language        string
	Quantity        int
	BuyingPrice     float64
	BoughtDate      *time.Time
	SellDate        *time.Time
}

func (uc *CardUseCase) CreateCard(input CreateCardInput) error {
	card := &entity.Card{
		UserID:          input.UserID,
		CardName:        input.CardName,
		CardImageURL:    input.CardImageURL,
		SetCode:         input.SetCode,
		CollectorNumber: input.CollectorNumber,
		Language:        input.Language,
		Quantity:        input.Quantity,
		BuyingPrice:     input.BuyingPrice,
		BoughtDate:      input.BoughtDate,
		SellDate:        input.SellDate,
	}

	return uc.cardRepo.Create(card)
}

func (uc *CardUseCase) UpdateCard(input UpdateCardInput) error {
	// First check if card belongs to user
	card, err := uc.cardRepo.FindByID(input.ID, input.UserID)
	if err != nil {
		return err
	}

	card.CardName = input.CardName
	card.CardImageURL = input.CardImageURL
	card.SetCode = input.SetCode
	card.CollectorNumber = input.CollectorNumber
	card.Language = input.Language
	card.Quantity = input.Quantity
	card.BuyingPrice = input.BuyingPrice
	card.BoughtDate = input.BoughtDate
	card.SellDate = input.SellDate

	return uc.cardRepo.Update(card)
}

func (uc *CardUseCase) DeleteCard(id uint, userID uint) error {
	return uc.cardRepo.Delete(id, userID)
}

func (uc *CardUseCase) GetCard(id uint, userID uint) (*entity.Card, error) {
	return uc.cardRepo.FindByID(id, userID)
}

func (uc *CardUseCase) ListCards(userID uint, page, pageSize int, search string) ([]entity.Card, int64, error) {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 20
	}

	return uc.cardRepo.FindByUserID(userID, page, pageSize, search)
}
