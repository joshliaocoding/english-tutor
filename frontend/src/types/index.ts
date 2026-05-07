export interface Scenario {
  id: string
  title: string
  description: string
  icon: string
}

export interface Correction {
  original: string
  corrected: string
  explanation: string
}

export interface ChatMessage {
  id?: number
  sessionId?: string
  role: 'user' | 'assistant'
  content: string
  corrections: Correction[]
  suggestions: string[]
  createdAt?: string
}

export interface NewSessionResponse {
  sessionId: string
  greeting: string
}

export interface ChatResponse {
  reply: string
  corrections: Correction[]
  suggestions: string[]
}
