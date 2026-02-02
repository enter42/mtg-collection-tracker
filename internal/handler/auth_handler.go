package handler

import (
	"log"
	"net/http"

	"github.com/enter42/mtg-collection-tracker/internal/usecase"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	authUseCase *usecase.AuthUseCase
}

func NewAuthHandler(authUseCase *usecase.AuthUseCase) *AuthHandler {
	return &AuthHandler{authUseCase: authUseCase}
}

func (h *AuthHandler) ShowLoginPage(c *gin.Context) {
	session := sessions.Default(c)
	userID := session.Get("user_id")
	if userID != nil {
		c.Redirect(http.StatusFound, "/cards")
		return
	}

	c.HTML(http.StatusOK, "login.html", gin.H{
		"title": "Login",
	})
}

func (h *AuthHandler) Login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	user, err := h.authUseCase.Login(username, password)
	if err != nil {
		c.HTML(http.StatusOK, "login.html", gin.H{
			"title": "Login",
			"error": err.Error(),
		})
		return
	}

	session := sessions.Default(c)
	session.Set("user_id", user.ID)
	session.Set("username", user.Username)
	if err := session.Save(); err != nil {
		log.Printf("Failed to save session: %v", err)
		c.HTML(http.StatusInternalServerError, "login.html", gin.H{
			"title": "Login",
			"error": "Failed to save session",
		})
		return
	}

	c.Redirect(http.StatusFound, "/cards")
}

func (h *AuthHandler) Logout(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()
	c.Redirect(http.StatusFound, "/login")
}

func (h *AuthHandler) ShowRegisterPage(c *gin.Context) {
	c.HTML(http.StatusOK, "register.html", gin.H{
		"title": "Register",
	})
}

func (h *AuthHandler) Register(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	confirmPassword := c.PostForm("confirm_password")

	if password != confirmPassword {
		c.HTML(http.StatusOK, "register.html", gin.H{
			"title": "Register",
			"error": "Passwords do not match",
		})
		return
	}

	if err := h.authUseCase.Register(username, password); err != nil {
		c.HTML(http.StatusOK, "register.html", gin.H{
			"title": "Register",
			"error": err.Error(),
		})
		return
	}

	c.Redirect(http.StatusFound, "/login")
}
