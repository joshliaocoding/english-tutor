import { ref, nextTick } from 'vue'
import type { Scenario, ChatMessage, NewSessionResponse, ChatResponse } from '../types'

const scenarios = ref<Scenario[]>([])
const messages = ref<ChatMessage[]>([])
const sessionId = ref<string | null>(null)
const activeScenario = ref<Scenario | null>(null)
const isLoading = ref(false)
const error = ref<string | null>(null)

export function useChat() {
  async function fetchScenarios() {
    try {
      const res = await fetch('/api/scenarios')
      if (!res.ok) throw new Error('Failed to fetch scenarios')
      scenarios.value = await res.json()
    } catch (err: any) {
      error.value = err.message
    }
  }

  async function startSession(scenario: Scenario) {
    isLoading.value = true
    error.value = null
    activeScenario.value = scenario
    messages.value = []

    try {
      const token = localStorage.getItem('token')
      const res = await fetch('/api/chat/new', {
        method: 'POST',
        headers: { 
          'Content-Type': 'application/json',
          'Authorization': `Bearer ${token}`
        },
        body: JSON.stringify({ scenarioId: scenario.id }),
      })

      if (!res.ok) throw new Error('Failed to start session')

      const data: NewSessionResponse = await res.json()
      sessionId.value = data.sessionId

      // Add AI greeting as first message
      messages.value.push({
        role: 'assistant',
        content: data.greeting,
        corrections: [],
        suggestions: [],
      })
    } catch (err: any) {
      error.value = err.message
      activeScenario.value = null
    } finally {
      isLoading.value = false
    }
  }

  async function sendMessage(text: string) {
    if (!sessionId.value || !text.trim()) return

    // Add user message immediately
    messages.value.push({
      role: 'user',
      content: text.trim(),
      corrections: [],
      suggestions: [],
    })

    isLoading.value = true
    error.value = null

    // Scroll to bottom after adding user message
    await nextTick()

    try {
      const token = localStorage.getItem('token')
      const res = await fetch('/api/chat', {
        method: 'POST',
        headers: { 
          'Content-Type': 'application/json',
          'Authorization': `Bearer ${token}`
        },
        body: JSON.stringify({
          sessionId: sessionId.value,
          message: text.trim(),
        }),
      })

      if (!res.ok) throw new Error('Failed to send message')

      const data: ChatResponse = await res.json()

      // Add AI response
      messages.value.push({
        role: 'assistant',
        content: data.reply,
        corrections: data.corrections || [],
        suggestions: data.suggestions || [],
      })
    } catch (err: any) {
      error.value = err.message
    } finally {
      isLoading.value = false
    }
  }

  function endSession() {
    sessionId.value = null
    activeScenario.value = null
    messages.value = []
    error.value = null
  }

  return {
    scenarios,
    messages,
    sessionId,
    activeScenario,
    isLoading,
    error,
    fetchScenarios,
    startSession,
    sendMessage,
    endSession,
  }
}
