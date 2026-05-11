package handlers

import (
	"context"
	"encoding/json"
	"time"

	"english-tutor/backend/database"
	"english-tutor/backend/models"
	"english-tutor/backend/services"

	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
)

// Scenarios defines the 6 built-in conversation scenarios.
var Scenarios = []models.Scenario{
	{
		ID:          "ordering_coffee",
		Title:       "Ordering Coffee",
		Description: "You're at a cozy coffee shop and want to order your favorite drink. Practice ordering, customizing your drink, and making small talk with the barista.",
		Icon:        "☕",
		SystemPrompt: "You are a friendly barista at a cozy coffee shop. The customer has just walked in.",
	},
	{
		ID:          "job_interview",
		Title:       "Job Interview",
		Description: "You're at a job interview for your dream position. Practice answering common interview questions, talking about your experience, and asking questions.",
		Icon:        "💼",
		SystemPrompt: "You are a friendly but professional interviewer conducting a job interview.",
	},
	{
		ID:          "hotel_checkin",
		Title:       "Hotel Check-In",
		Description: "You've arrived at a hotel after a long trip. Practice checking in, asking about amenities, and handling common hotel situations.",
		Icon:        "🏨",
		SystemPrompt: "You are a polite and helpful hotel receptionist. A guest has just arrived to check in.",
	},
	{
		ID:          "doctor_visit",
		Title:       "Doctor's Visit",
		Description: "You need to see a doctor about a minor health concern. Practice describing symptoms, understanding medical advice, and asking follow-up questions.",
		Icon:        "🏥",
		SystemPrompt: "You are a caring and patient doctor. A patient has come in for a consultation.",
	},
	{
		ID:          "giving_directions",
		Title:       "Giving Directions",
		Description: "A tourist is asking you for directions to a nearby landmark. Practice giving clear directions, using prepositions of place, and being helpful.",
		Icon:        "🗺️",
		SystemPrompt: "You are a tourist in a new city, asking a local for directions to a famous landmark.",
	},
	{
		ID:          "casual_smalltalk",
		Title:       "Casual Small Talk",
		Description: "You've met someone new at a social gathering. Practice making small talk, sharing interests, and keeping a conversation going naturally.",
		Icon:        "💬",
		SystemPrompt: "You are a friendly person at a social gathering who has just met someone new.",
	},
}

// GetScenarios returns all available conversation scenarios.
func GetScenarios(c fiber.Ctx) error {
	// Return scenarios without the system prompt
	type publicScenario struct {
		ID          string `json:"id"`
		Title       string `json:"title"`
		Description string `json:"description"`
		Icon        string `json:"icon"`
	}

	result := make([]publicScenario, len(Scenarios))
	for i, s := range Scenarios {
		result[i] = publicScenario{
			ID:          s.ID,
			Title:       s.Title,
			Description: s.Description,
			Icon:        s.Icon,
		}
	}

	return c.JSON(result)
}

// NewSession creates a new conversation session.
func NewSession(c fiber.Ctx) error {
	var req models.NewSessionRequest
	if err := c.Bind().JSON(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	// Find the scenario
	var scenario models.Scenario
	found := false
	for _, s := range Scenarios {
		if s.ID == req.ScenarioID {
			scenario = s
			found = true
			break
		}
	}
	if !found {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Unknown scenario ID",
		})
	}

	userIDStr := c.Locals("userID").(string)
	userUUID, _ := uuid.Parse(userIDStr)

	// Create DB session
	sessionID := uuid.New()
	dbSession := models.Session{
		ID:         sessionID,
		UserID:     &userUUID,
		ScenarioID: req.ScenarioID,
	}
	if err := database.DB.Create(&dbSession).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create session",
		})
	}

	// Start Gemini chat and get greeting
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	payload, err := services.Service.StartSession(ctx, sessionID.String(), scenario)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to start AI conversation: " + err.Error(),
		})
	}

	// Save the greeting message to DB
	correctionsJSON, _ := json.Marshal(payload.Corrections)
	suggestionsJSON, _ := json.Marshal(payload.Suggestions)

	greetingMsg := models.Message{
		SessionID:   sessionID,
		Role:        "assistant",
		Content:     payload.Reply,
		Corrections: correctionsJSON,
		Suggestions: suggestionsJSON,
	}
	database.DB.Create(&greetingMsg)

	return c.JSON(models.NewSessionResponse{
		SessionID: sessionID.String(),
		Greeting:  payload.Reply,
	})
}

// SendMessage handles a user message in an existing session.
func SendMessage(c fiber.Ctx) error {
	var req models.ChatRequest
	if err := c.Bind().JSON(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	sessionUUID, err := uuid.Parse(req.SessionID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid session ID",
		})
	}

	userIDStr := c.Locals("userID").(string)

	// Verify session exists in DB and belongs to user
	var session models.Session
	if err := database.DB.First(&session, "id = ?", sessionUUID).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Session not found",
		})
	}

	if session.UserID == nil || session.UserID.String() != userIDStr {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error": "Access denied",
		})
	}

	// Save user message to DB
	userMsg := models.Message{
		SessionID: sessionUUID,
		Role:      "user",
		Content:   req.Message,
	}
	database.DB.Create(&userMsg)

	// Send to Gemini
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	payload, err := services.Service.SendMessage(ctx, req.SessionID, req.Message)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to get AI response: " + err.Error(),
		})
	}

	// Save AI response to DB
	correctionsJSON, _ := json.Marshal(payload.Corrections)
	suggestionsJSON, _ := json.Marshal(payload.Suggestions)

	aiMsg := models.Message{
		SessionID:   sessionUUID,
		Role:        "assistant",
		Content:     payload.Reply,
		Corrections: correctionsJSON,
		Suggestions: suggestionsJSON,
	}
	database.DB.Create(&aiMsg)

	return c.JSON(models.ChatResponse{
		Reply:       payload.Reply,
		Corrections: payload.Corrections,
		Suggestions: payload.Suggestions,
	})
}

// GetMessages returns all messages for a session.
func GetMessages(c fiber.Ctx) error {
	sessionID := c.Params("id")
	sessionUUID, err := uuid.Parse(sessionID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid session ID",
		})
	}

	userIDStr := c.Locals("userID").(string)

	// Verify session belongs to user
	var session models.Session
	if err := database.DB.First(&session, "id = ?", sessionUUID).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Session not found",
		})
	}

	if session.UserID == nil || session.UserID.String() != userIDStr {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error": "Access denied",
		})
	}

	var messages []models.Message
	if err := database.DB.Where("session_id = ?", sessionUUID).Order("created_at ASC").Find(&messages).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch messages",
		})
	}

	return c.JSON(messages)
}
