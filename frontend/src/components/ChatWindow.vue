<script setup lang="ts">
import { ref, watch, nextTick, computed } from 'vue'
import type { ChatMessage as ChatMessageType } from '../types'
import ChatMessage from './ChatMessage.vue'
import ChatInput from './ChatInput.vue'
import FeedbackPanel from './FeedbackPanel.vue'

const props = defineProps<{
  messages: ChatMessageType[]
  isLoading: boolean
}>()

const emit = defineEmits<{
  send: [message: string]
}>()

const messagesContainer = ref<HTMLElement | null>(null)

// Get the latest AI message's feedback
const latestFeedback = computed(() => {
  const aiMessages = props.messages.filter(m => m.role === 'assistant')
  if (aiMessages.length === 0) return null
  const last = aiMessages[aiMessages.length - 1]
  if ((last.corrections && last.corrections.length > 0) || (last.suggestions && last.suggestions.length > 0)) {
    return last
  }
  return null
})

// Auto-scroll to bottom when new messages arrive
watch(
  () => props.messages.length,
  async () => {
    await nextTick()
    scrollToBottom()
  }
)

watch(
  () => props.isLoading,
  async () => {
    await nextTick()
    scrollToBottom()
  }
)

function scrollToBottom() {
  if (messagesContainer.value) {
    messagesContainer.value.scrollTop = messagesContainer.value.scrollHeight
  }
}
</script>

<template>
  <div class="flex flex-1 flex-col overflow-hidden">
    <!-- Messages area -->
    <div
      ref="messagesContainer"
      class="flex-1 overflow-y-auto px-3 py-4 sm:px-4 sm:py-6"
    >
      <div class="mx-auto max-w-3xl space-y-4">
        <ChatMessage
          v-for="(msg, index) in messages"
          :key="index"
          :message="msg"
        />

        <!-- Loading indicator -->
        <div
          v-if="isLoading"
          class="flex items-center gap-3 animate-fade-in"
        >
          <div class="flex h-8 w-8 shrink-0 items-center justify-center rounded-full bg-gradient-to-br from-[var(--color-accent)] to-emerald-600 text-sm font-bold text-white shadow-md">
            AI
          </div>
          <div class="glass-strong rounded-2xl rounded-bl-md px-4 py-3">
            <div class="flex items-center gap-1.5">
              <span
                v-for="i in 3"
                :key="i"
                class="h-2 w-2 rounded-full bg-[var(--color-accent)]"
                :style="{ animation: `pulse-dot 1.4s ease-in-out ${(i - 1) * 0.2}s infinite` }"
              />
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- Feedback panel (shows below messages, above input) -->
    <FeedbackPanel
      v-if="latestFeedback"
      :corrections="latestFeedback.corrections"
      :suggestions="latestFeedback.suggestions"
      class="pb-2"
    />

    <!-- Input area -->
    <ChatInput
      :is-loading="isLoading"
      @send="$emit('send', $event)"
    />
  </div>
</template>
