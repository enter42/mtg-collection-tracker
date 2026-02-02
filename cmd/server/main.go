package main

import (
	"html/template"
	"log"
	"os"

	"github.com/enter42/mtg-collection-tracker/internal/handler"
	"github.com/enter42/mtg-collection-tracker/internal/handler/middleware"
	"github.com/enter42/mtg-collection-tracker/internal/infrastructure/database"
	"github.com/enter42/mtg-collection-tracker/internal/infrastructure/repository"
	"github.com/enter42/mtg-collection-tracker/internal/usecase"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables")
	}

	// Initialize database
	db, err := database.NewDatabase()
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	// Initialize repositories
	userRepo := repository.NewUserRepository(db)
	cardRepo := repository.NewCardRepository(db)

	// Initialize use cases
	authUseCase := usecase.NewAuthUseCase(userRepo)
	cardUseCase := usecase.NewCardUseCase(cardRepo)

	// Initialize handlers
	authHandler := handler.NewAuthHandler(authUseCase)
	cardHandler := handler.NewCardHandler(cardUseCase)

	// Initialize Gin
	router := gin.Default()

	// Session middleware
	sessionSecret := os.Getenv("SESSION_SECRET")
	if sessionSecret == "" {
		sessionSecret = "default-secret-change-this"
	}
	store := cookie.NewStore([]byte(sessionSecret))
	router.Use(sessions.Sessions("mtg_session", store))

	// Custom template functions
	funcMap := template.FuncMap{
		"sub": func(a, b int) int {
			return a - b
		},
		"add": func(a, b int) int {
			return a + b
		},
		"until": func(count int) []int {
			result := make([]int, count)
			for i := 0; i < count; i++ {
				result[i] = i
			}
			return result
		},
	}

	// Load templates with custom functions
	router.SetFuncMap(funcMap)
	router.LoadHTMLGlob("templates/**/*.html")

	// Public routes
	router.GET("/", func(c *gin.Context) {
		c.Redirect(302, "/login")
	})
	router.GET("/login", authHandler.ShowLoginPage)
	router.POST("/login", authHandler.Login)
	router.GET("/register", authHandler.ShowRegisterPage)
	router.POST("/register", authHandler.Register)

	// Protected routes
	protected := router.Group("/")
	protected.Use(middleware.AuthRequired())
	{
		protected.GET("/logout", authHandler.Logout)
		protected.GET("/cards", cardHandler.ListCards)
		protected.GET("/cards/add", cardHandler.ShowAddCardPage)
		protected.POST("/cards/add", cardHandler.AddCard)
		protected.GET("/cards/edit/:id", cardHandler.ShowEditCardPage)
		protected.POST("/cards/edit/:id", cardHandler.EditCard)
		protected.POST("/cards/delete/:id", cardHandler.DeleteCard)
	}

	// Start server
	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server starting on port %s...", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
