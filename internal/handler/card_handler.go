package handler

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/enter42/mtg-collection-tracker/internal/usecase"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type CardHandler struct {
	cardUseCase *usecase.CardUseCase
}

func NewCardHandler(cardUseCase *usecase.CardUseCase) *CardHandler {
	return &CardHandler{cardUseCase: cardUseCase}
}

func (h *CardHandler) ListCards(c *gin.Context) {
	session := sessions.Default(c)
	userID := session.Get("user_id").(uint)
	username := session.Get("username").(string)

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize := 20
	search := c.Query("search")

	cards, total, err := h.cardUseCase.ListCards(userID, page, pageSize, search)
	if err != nil {
		log.Printf("Error listing cards: %v", err)
		c.HTML(http.StatusInternalServerError, "error.html", gin.H{
			"title": "Error",
			"error": "Failed to load cards",
		})
		return
	}

	totalPages := int((total + int64(pageSize) - 1) / int64(pageSize))

	c.HTML(http.StatusOK, "cards.html", gin.H{
		"title":      "My Card Collection",
		"username":   username,
		"cards":      cards,
		"page":       page,
		"totalPages": totalPages,
		"search":     search,
		"total":      total,
	})
}

func (h *CardHandler) ShowAddCardPage(c *gin.Context) {
	session := sessions.Default(c)
	username := session.Get("username").(string)

	c.HTML(http.StatusOK, "add_card.html", gin.H{
		"title":    "Add Card",
		"username": username,
	})
}

func (h *CardHandler) AddCard(c *gin.Context) {
	session := sessions.Default(c)
	userID := session.Get("user_id").(uint)
	username := session.Get("username").(string)

	quantity, _ := strconv.Atoi(c.PostForm("quantity"))
	if quantity < 1 {
		quantity = 1
	}

	buyingPrice, _ := strconv.ParseFloat(c.PostForm("buying_price"), 64)

	var boughtDate *time.Time
	if boughtDateStr := c.PostForm("bought_date"); boughtDateStr != "" {
		if t, err := time.Parse("2006-01-02", boughtDateStr); err == nil {
			boughtDate = &t
		}
	}

	var sellDate *time.Time
	if sellDateStr := c.PostForm("sell_date"); sellDateStr != "" {
		if t, err := time.Parse("2006-01-02", sellDateStr); err == nil {
			sellDate = &t
		}
	}

	input := usecase.CreateCardInput{
		UserID:          userID,
		CardName:        c.PostForm("card_name"),
		CardImageURL:    c.PostForm("card_image_url"),
		SetCode:         c.PostForm("set_code"),
		CollectorNumber: c.PostForm("collector_number"),
		Language:        c.PostForm("language"),
		Quantity:        quantity,
		BuyingPrice:     buyingPrice,
		BoughtDate:      boughtDate,
		SellDate:        sellDate,
	}

	if err := h.cardUseCase.CreateCard(input); err != nil {
		log.Printf("Error creating card: %v", err)
		c.HTML(http.StatusOK, "add_card.html", gin.H{
			"title":    "Add Card",
			"username": username,
			"error":    "Failed to add card",
		})
		return
	}

	c.Redirect(http.StatusFound, "/cards")
}

func (h *CardHandler) ShowEditCardPage(c *gin.Context) {
	session := sessions.Default(c)
	userID := session.Get("user_id").(uint)
	username := session.Get("username").(string)

	cardID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.Redirect(http.StatusFound, "/cards")
		return
	}

	card, err := h.cardUseCase.GetCard(uint(cardID), userID)
	if err != nil {
		log.Printf("Error getting card: %v", err)
		c.Redirect(http.StatusFound, "/cards")
		return
	}

	var boughtDateStr string
	if card.BoughtDate != nil {
		boughtDateStr = card.BoughtDate.Format("2006-01-02")
	}

	var sellDateStr string
	if card.SellDate != nil {
		sellDateStr = card.SellDate.Format("2006-01-02")
	}

	c.HTML(http.StatusOK, "edit_card.html", gin.H{
		"title":         "Edit Card",
		"username":      username,
		"card":          card,
		"boughtDateStr": boughtDateStr,
		"sellDateStr":   sellDateStr,
	})
}

func (h *CardHandler) EditCard(c *gin.Context) {
	session := sessions.Default(c)
	userID := session.Get("user_id").(uint)

	cardID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.Redirect(http.StatusFound, "/cards")
		return
	}

	quantity, _ := strconv.Atoi(c.PostForm("quantity"))
	if quantity < 1 {
		quantity = 1
	}

	buyingPrice, _ := strconv.ParseFloat(c.PostForm("buying_price"), 64)

	var boughtDate *time.Time
	if boughtDateStr := c.PostForm("bought_date"); boughtDateStr != "" {
		if t, err := time.Parse("2006-01-02", boughtDateStr); err == nil {
			boughtDate = &t
		}
	}

	var sellDate *time.Time
	if sellDateStr := c.PostForm("sell_date"); sellDateStr != "" {
		if t, err := time.Parse("2006-01-02", sellDateStr); err == nil {
			sellDate = &t
		}
	}

	input := usecase.UpdateCardInput{
		ID:              uint(cardID),
		UserID:          userID,
		CardName:        c.PostForm("card_name"),
		CardImageURL:    c.PostForm("card_image_url"),
		SetCode:         c.PostForm("set_code"),
		CollectorNumber: c.PostForm("collector_number"),
		Language:        c.PostForm("language"),
		Quantity:        quantity,
		BuyingPrice:     buyingPrice,
		BoughtDate:      boughtDate,
		SellDate:        sellDate,
	}

	if err := h.cardUseCase.UpdateCard(input); err != nil {
		log.Printf("Error updating card: %v", err)
		c.Redirect(http.StatusFound, "/cards")
		return
	}

	c.Redirect(http.StatusFound, "/cards")
}

func (h *CardHandler) DeleteCard(c *gin.Context) {
	session := sessions.Default(c)
	userID := session.Get("user_id").(uint)

	cardID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.Redirect(http.StatusFound, "/cards")
		return
	}

	if err := h.cardUseCase.DeleteCard(uint(cardID), userID); err != nil {
		log.Printf("Error deleting card: %v", err)
	}

	c.Redirect(http.StatusFound, "/cards")
}
