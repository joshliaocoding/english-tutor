# SpeakEasy: AI English Tutor

SpeakEasy is a full-stack web application designed to help users practice conversational English with an AI tutor. The application provides various real-world scenarios, engaging the user in a dialogue and offering real-time grammar corrections and vocabulary suggestions powered by Google's Gemini AI.

## Features

- **Real-time Roleplay:** Practice English in curated scenarios like ordering coffee, checking into a hotel, or passing a job interview.
- **Smart Feedback:** Get instant grammar corrections and suggested phrasing for every message sent.
- **Premium UI:** Features a sleek, responsive, WhatsApp-style messaging interface built with PrimeVue and a dark glassmorphism aesthetic.
- **Persistent Chat:** Backend uses PostgreSQL to save chat sessions.

## Tech Stack

### Frontend
- **Framework:** Vue 3 (Composition API)
- **Build Tool:** Vite
- **Styling:** Tailwind CSS v4, custom glassmorphism utilities
- **UI Components:** PrimeVue v4 (Aura Theme with custom Sky Blue preset)
- **Icons:** PrimeIcons

### Backend
- **Language:** Go (1.24+)
- **Framework:** Go Fiber
- **Database:** PostgreSQL (with GORM)
- **AI Integration:** Google GenAI SDK (Gemini 2.5 Flash)

## Prerequisites

Before you begin, ensure you have the following installed:
- Node.js (v18+)
- Go (v1.24+)
- PostgreSQL (v18+)
- A Google Gemini API Key

## Setup & Installation

### 1. Database Setup
1. Ensure your PostgreSQL server is running.
2. Create a new database named `english_tutor`.
3. The backend uses GORM to automatically migrate the schema on startup.

### 2. Backend Setup
1. Navigate to the backend directory:
   ```bash
   cd backend
   ```
2. Create a `.env` file in the `backend` directory with the following configuration:
   ```env
   DB_HOST=localhost
   DB_USER=postgres
   DB_PASSWORD=your_postgres_password
   DB_NAME=english_tutor
   DB_PORT=5432
   GEMINI_API_KEY=your_gemini_api_key
   ```
3. Download Go dependencies:
   ```bash
   go mod tidy
   ```
4. Start the backend server (it runs on port `8081`):
   ```bash
   go run main.go
   ```

### 3. Frontend Setup
1. Open a new terminal and navigate to the frontend directory:
   ```bash
   cd frontend
   ```
2. Install Node dependencies:
   ```bash
   npm install
   ```
3. Start the Vite development server:
   ```bash
   npm run dev
   ```
4. Open your browser and navigate to the URL provided by Vite (usually `http://localhost:5173/`).

## Project Structure

```
english-tutor/
├── backend/
│   ├── handlers/      # API route handlers
│   ├── models/        # GORM database models
│   ├── services/      # Gemini AI and DB business logic
│   ├── main.go        # Fiber app entrypoint
│   └── .env           # Environment variables
└── frontend/
    ├── src/
    │   ├── components/# Vue UI components (ChatWindow, Header, etc.)
    │   ├── composables/# Shared Vue logic (useChat)
    │   ├── App.vue    # Root Vue component
    │   └── main.ts    # Frontend entrypoint & PrimeVue config
    ├── tailwind.css   # Tailwind v4 configuration
    └── vite.config.ts # Vite configuration & API proxy
```

## License
MIT License
