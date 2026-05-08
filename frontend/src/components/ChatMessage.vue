<script setup lang="ts">
import type { ChatMessage } from '../types'
import Message from 'primevue/message';

defineProps<{
  message: ChatMessage
}>()
</script>

<template>
  <div
    class="flex w-full gap-2 px-2 sm:px-4"
    :class="[
      message.role === 'user' ? 'justify-end' : 'justify-start',
      message.role === 'user' ? 'animate-slide-in-right' : 'animate-slide-in-left',
    ]"
  >
    <!-- Message bubble using PrimeVue Message -->
    <div class="max-w-[85%] sm:max-w-[75%]">
      <Message
        :closable="false"
        :unstyled="true"
        :pt="{
          root: [
            'p-3 sm:p-4 rounded-2xl shadow-sm m-0',
            message.role === 'user'
              ? 'bg-gradient-to-br from-[var(--color-accent)] to-sky-600 text-white rounded-tr-sm'
              : 'bg-[var(--color-bg-card)] backdrop-blur-xl border border-[var(--color-border)] text-[var(--color-text-primary)] rounded-tl-sm'
          ],
          content: 'p-0 m-0',
          text: 'text-[15px] leading-relaxed whitespace-pre-wrap p-0 m-0'
        }"
      >
        {{ message.content }}
      </Message>
    </div>
  </div>
</template>
