package services

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sync"

	"english-tutor/backend/models"

	"google.golang.org/genai"
)

// GeminiService manages the Gemini client and active chat sessions.
type GeminiService struct {
	client   *genai.Client
	sessions map[string]*genai.Chat
	mu       sync.RWMutex
}

var Service *GeminiService

// InitGemini creates the Gemini client using the GEMINI_API_KEY env var.
func InitGemini() {
	apiKey := os.Getenv("GEMINI_API_KEY")
	if apiKey == "" {
		log.Fatal("GEMINI_API_KEY environment variable is required")
	}

	ctx := context.Background()
	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey:  apiKey,
		Backend: genai.BackendGeminiAPI,
	})
	if err != nil {
		log.Fatalf("Failed to create Gemini client: %v", err)
	}

	Service = &GeminiService{
		client:   client,
		sessions: make(map[string]*genai.Chat),
	}

	log.Println("Gemini client initialized successfully")
}

// buildSystemInstruction returns the system prompt for a given scenario.
func buildSystemInstruction(scenario models.Scenario) string {
	return fmt.Sprintf(`You are an English conversation practice partner. You are roleplaying in the following scenario:

**Scenario: %s**
%s

RULES:
1. Stay in character for the scenario. Respond naturally as the other party in the conversation.
2. Keep your responses conversational, friendly, and at an intermediate English level.
3. After each user message, analyze their English for grammar, vocabulary, and naturalness.
4. You MUST respond with ONLY a valid JSON object (no markdown, no code fences, no extra text) in this exact format:
{
  "reply": "Your in-character conversational response here",
  "corrections": [
    {
      "original": "what the user wrote incorrectly",
      "corrected": "the corrected version",
      "explanation": "brief explanation of the error"
    }
  ],
  "suggestions": ["alternative phrase or vocabulary the user could try"]
}

5. If the user's English is correct, return an empty "corrections" array.
6. Always include 1-2 vocabulary suggestions relevant to the scenario.
7. Keep replies concise (1-3 sentences for the conversation part).
8. Be encouraging and supportive — this is a learning environment.`, scenario.Title, scenario.Description)
}

// StartSession creates a new Gemini chat session for a given scenario.
// It sends an initial greeting prompt and returns the AI's opening message.
func (s *GeminiService) StartSession(ctx context.Context, sessionID string, scenario models.Scenario) (models.GeminiResponsePayload, error) {
	systemInstruction := buildSystemInstruction(scenario)

	config := &genai.GenerateContentConfig{
		SystemInstruction: &genai.Content{
			Parts: []*genai.Part{{Text: systemInstruction}},
		},
		Temperature: genai.Ptr(float32(0.8)),
		TopP:        genai.Ptr(float32(0.9)),
	}

	chat, err := s.client.Chats.Create(ctx, "gemini-2.5-flash", config, nil)
	if err != nil {
		return models.GeminiResponsePayload{}, fmt.Errorf("failed to create chat: %w", err)
	}

	s.mu.Lock()
	s.sessions[sessionID] = chat
	s.mu.Unlock()

	// Send an initial prompt to get the AI to greet the user in-character.
	greetingPrompt := genai.Part{Text: `The user has just started this conversation scenario. Greet them in character and set the scene. Remember to respond with ONLY a valid JSON object. Since this is the opening message, "corrections" should be empty and "suggestions" should include 1-2 useful phrases for this scenario.`}

	resp, err := chat.SendMessage(ctx, greetingPrompt)
	if err != nil {
		return models.GeminiResponsePayload{}, fmt.Errorf("failed to get greeting: %w", err)
	}

	return parseGeminiResponse(resp)
}

// SendMessage sends a user message to an existing chat session and returns the parsed response.
func (s *GeminiService) SendMessage(ctx context.Context, sessionID string, message string) (models.GeminiResponsePayload, error) {
	s.mu.RLock()
	chat, exists := s.sessions[sessionID]
	s.mu.RUnlock()

	if !exists {
		return models.GeminiResponsePayload{}, fmt.Errorf("session %s not found", sessionID)
	}

	part := genai.Part{Text: message}
	resp, err := chat.SendMessage(ctx, part)
	if err != nil {
		return models.GeminiResponsePayload{}, fmt.Errorf("failed to send message: %w", err)
	}

	return parseGeminiResponse(resp)
}

// RemoveSession cleans up a chat session from memory.
func (s *GeminiService) RemoveSession(sessionID string) {
	s.mu.Lock()
	delete(s.sessions, sessionID)
	s.mu.Unlock()
}

// parseGeminiResponse extracts the JSON payload from the Gemini response.
func parseGeminiResponse(resp *genai.GenerateContentResponse) (models.GeminiResponsePayload, error) {
	if resp == nil || len(resp.Candidates) == 0 || resp.Candidates[0].Content == nil {
		return models.GeminiResponsePayload{}, fmt.Errorf("empty response from Gemini")
	}

	var rawText string
	for _, part := range resp.Candidates[0].Content.Parts {
		if part.Text != "" {
			rawText += part.Text
		}
	}

	if rawText == "" {
		return models.GeminiResponsePayload{}, fmt.Errorf("no text in Gemini response")
	}

	// Try to parse the response as JSON. Gemini sometimes wraps in markdown code fences.
	cleaned := cleanJSON(rawText)

	var payload models.GeminiResponsePayload
	if err := json.Unmarshal([]byte(cleaned), &payload); err != nil {
		// Fallback: treat the whole response as a plain reply
		log.Printf("Warning: could not parse Gemini response as JSON: %v\nRaw: %s", err, rawText)
		return models.GeminiResponsePayload{
			Reply:       rawText,
			Corrections: []models.Correction{},
			Suggestions: []string{},
		}, nil
	}

	// Ensure non-nil slices for JSON serialization
	if payload.Corrections == nil {
		payload.Corrections = []models.Correction{}
	}
	if payload.Suggestions == nil {
		payload.Suggestions = []string{}
	}

	return payload, nil
}

// cleanJSON strips markdown code fences from a JSON string.
func cleanJSON(s string) string {
	// Remove ```json ... ``` wrapping
	start := 0
	end := len(s)

	for i := 0; i < len(s); i++ {
		if s[i] == '{' {
			start = i
			break
		}
	}
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '}' {
			end = i + 1
			break
		}
	}

	if start < end {
		return s[start:end]
	}
	return s
}
