package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/datatypes"
)

// ============================================================
// GORM Models (database tables)
// ============================================================

// Session represents a conversation session tied to a scenario.
type Session struct {
	ID         uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	ScenarioID string    `gorm:"not null" json:"scenarioId"`
	Messages   []Message `gorm:"foreignKey:SessionID" json:"messages,omitempty"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
}

// Message represents a single message in a conversation.
type Message struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	SessionID   uuid.UUID      `gorm:"type:uuid;not null;index" json:"sessionId"`
	Role        string         `gorm:"not null" json:"role"` // "user" or "assistant"
	Content     string         `gorm:"type:text;not null" json:"content"`
	Corrections datatypes.JSON `gorm:"type:jsonb" json:"corrections"`
	Suggestions datatypes.JSON `gorm:"type:jsonb" json:"suggestions"`
	CreatedAt   time.Time      `json:"createdAt"`
}

// ============================================================
// DTOs (request/response payloads)
// ============================================================

// Scenario describes a conversation scenario the user can pick.
type Scenario struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
	SystemPrompt string `json:"-"` // not exposed to client
}

// NewSessionRequest is the payload for POST /api/chat/new.
type NewSessionRequest struct {
	ScenarioID string `json:"scenarioId"`
}

// NewSessionResponse is returned when a new session is created.
type NewSessionResponse struct {
	SessionID string `json:"sessionId"`
	Greeting  string `json:"greeting"`
}

// ChatRequest is the payload for POST /api/chat.
type ChatRequest struct {
	SessionID string `json:"sessionId"`
	Message   string `json:"message"`
}

// ChatResponse is returned after each user message.
type ChatResponse struct {
	Reply       string       `json:"reply"`
	Corrections []Correction `json:"corrections"`
	Suggestions []string     `json:"suggestions"`
}

// Correction describes a grammar/usage fix.
type Correction struct {
	Original    string `json:"original"`
	Corrected   string `json:"corrected"`
	Explanation string `json:"explanation"`
}

// GeminiResponsePayload is the structured JSON we ask Gemini to return.
type GeminiResponsePayload struct {
	Reply       string       `json:"reply"`
	Corrections []Correction `json:"corrections"`
	Suggestions []string     `json:"suggestions"`
}
