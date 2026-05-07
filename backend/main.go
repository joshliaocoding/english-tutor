package main

import (
	"log"
	"os"

	"english-tutor/backend/database"
	"english-tutor/backend/handlers"
	"english-tutor/backend/services"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/gofiber/fiber/v3/middleware/logger"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env file (ignore error if not present, e.g. in production)
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables")
	}

	// Initialize database
	database.Connect()

	// Initialize Gemini
	services.InitGemini()

	// Create Fiber app
	app := fiber.New(fiber.Config{
		AppName: "SpeakEasy - English Tutor",
	})

	// Middleware
	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:5173", "http://localhost:3000"},
		AllowHeaders: []string{"Origin", "Content-Type", "Accept"},
		AllowMethods: []string{"GET", "POST", "OPTIONS"},
	}))

	// API routes
	api := app.Group("/api")
	api.Get("/scenarios", handlers.GetScenarios)
	api.Post("/chat/new", handlers.NewSession)
	api.Post("/chat", handlers.SendMessage)
	api.Get("/sessions/:id/messages", handlers.GetMessages)

	// Get port from env or default to 8080
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server starting on :%s", port)
	log.Fatal(app.Listen(":" + port))
}
