<script setup lang="ts">
import { ref } from 'vue'

defineProps<{
  isLoading: boolean
}>()

const emit = defineEmits<{
  send: [message: string]
}>()

const inputText = ref('')

function handleSend() {
  if (!inputText.value.trim()) return
  emit('send', inputText.value)
  inputText.value = ''
}

function handleKeydown(e: KeyboardEvent) {
  if (e.key === 'Enter' && !e.shiftKey) {
    e.preventDefault()
    handleSend()
  }
}
</script>

<template>
  <div class="border-t border-[var(--color-border)] bg-[var(--color-bg-secondary)] p-3 sm:p-4">
    <div class="mx-auto flex max-w-3xl items-end gap-3">
      <div class="relative flex-1">
        <textarea
          id="chat-input"
          v-model="inputText"
          rows="1"
          placeholder="Type your message in English..."
          class="w-full resize-none rounded-xl border border-[var(--color-border)] bg-[var(--color-bg-input)] px-4 py-3 text-sm text-[var(--color-text-primary)] placeholder-[var(--color-text-muted)] outline-none transition-all duration-200 focus:border-[var(--color-border-accent)] focus:shadow-[0_0_20px_var(--color-accent-glow)]"
          :disabled="isLoading"
          @keydown="handleKeydown"
        />
      </div>

      <button
        id="send-btn"
        class="flex h-11 w-11 shrink-0 items-center justify-center rounded-xl bg-gradient-to-r from-[var(--color-accent)] to-emerald-600 text-white shadow-lg transition-all duration-200 hover:shadow-[0_0_20px_var(--color-accent-glow-strong)] hover:scale-105 active:scale-95 disabled:opacity-40 disabled:cursor-not-allowed disabled:hover:scale-100 disabled:hover:shadow-none"
        :disabled="isLoading || !inputText.trim()"
        @click="handleSend"
      >
        <!-- Send icon -->
        <svg
          v-if="!isLoading"
          width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"
        >
          <path d="M22 2L11 13" /><path d="M22 2L15 22L11 13L2 9L22 2Z" />
        </svg>
        <!-- Loading spinner -->
        <svg
          v-else
          class="animate-spin" width="18" height="18" viewBox="0 0 24 24" fill="none"
        >
          <circle cx="12" cy="12" r="10" stroke="currentColor" stroke-width="3" stroke-dasharray="31.4 31.4" stroke-linecap="round" />
        </svg>
      </button>
    </div>
  </div>
</template>
