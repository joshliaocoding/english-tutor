<script setup lang="ts">
import type { ChatMessage } from '../types'

defineProps<{
  message: ChatMessage
}>()
</script>

<template>
  <div
    class="flex w-full gap-3"
    :class="[
      message.role === 'user' ? 'justify-end' : 'justify-start',
      message.role === 'user' ? 'animate-slide-in-right' : 'animate-slide-in-left',
    ]"
  >
    <!-- AI avatar -->
    <div
      v-if="message.role === 'assistant'"
      class="mt-1 flex h-8 w-8 shrink-0 items-center justify-center rounded-full bg-gradient-to-br from-[var(--color-accent)] to-emerald-600 text-sm font-bold text-white shadow-md"
    >
      AI
    </div>

    <!-- Message bubble -->
    <div
      class="max-w-[80%] rounded-2xl px-4 py-3 sm:max-w-[70%]"
      :class="[
        message.role === 'user'
          ? 'bg-gradient-to-br from-[var(--color-accent)] to-emerald-600 text-white rounded-br-md'
          : 'glass-strong text-[var(--color-text-primary)] rounded-bl-md',
      ]"
    >
      <p class="text-sm leading-relaxed whitespace-pre-wrap">{{ message.content }}</p>
    </div>

    <!-- User avatar -->
    <div
      v-if="message.role === 'user'"
      class="mt-1 flex h-8 w-8 shrink-0 items-center justify-center rounded-full bg-gradient-to-br from-violet-500 to-[var(--color-secondary)] text-sm font-bold text-white shadow-md"
    >
      U
    </div>
  </div>
</template>
