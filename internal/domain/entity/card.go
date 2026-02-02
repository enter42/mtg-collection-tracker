package entity

import (
	"time"

	"gorm.io/gorm"
)

type Card struct {
	ID              uint           `gorm:"primarykey" json:"id"`
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at"`
	DeletedAt       gorm.DeletedAt `gorm:"index" json:"-"`
	UserID          uint           `gorm:"not null;index" json:"user_id"`
	CardName        string         `gorm:"size:255;not null" json:"card_name"`
	CardImageURL    string         `gorm:"size:500" json:"card_image_url"`
	SetCode         string         `gorm:"size:20" json:"set_code"`
	CollectorNumber string         `gorm:"size:20" json:"collector_number"`
	Language        string         `gorm:"size:50" json:"language"`
	Quantity        int            `gorm:"default:1" json:"quantity"`
	BuyingPrice     float64        `gorm:"type:decimal(10,2)" json:"buying_price"`
	BoughtDate      *time.Time     `json:"bought_date"`
	SellDate        *time.Time     `json:"sell_date"`
	User            User           `gorm:"foreignKey:UserID" json:"-"`
}
