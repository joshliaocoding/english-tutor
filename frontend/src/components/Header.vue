<script setup lang="ts">
import type { Scenario } from '../types'

defineProps<{
  activeScenario: Scenario | null
}>()

defineEmits<{
  back: []
}>()
</script>

<template>
  <header
    class="sticky top-0 z-50 glass border-b border-[var(--color-border)] px-4 py-3 sm:px-6"
  >
    <div class="mx-auto flex max-w-4xl items-center gap-3">
      <!-- Back button (visible when in a chat) -->
      <button
        v-if="activeScenario"
        id="header-back-btn"
        class="flex items-center justify-center rounded-lg p-2 text-[var(--color-text-secondary)] transition-all duration-200 hover:bg-[var(--color-bg-card-hover)] hover:text-[var(--color-text-primary)]"
        @click="$emit('back')"
      >
        <svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
          <path d="M19 12H5" /><path d="M12 19l-7-7 7-7" />
        </svg>
      </button>

      <!-- Logo + Title -->
      <div class="flex items-center gap-2.5">
        <div
          class="flex h-9 w-9 items-center justify-center rounded-xl bg-gradient-to-br from-[var(--color-accent)] to-emerald-600 text-lg font-bold text-white shadow-lg"
        >
          S
        </div>
        <div>
          <h1 class="text-lg font-semibold leading-tight text-[var(--color-text-primary)]">
            SpeakEasy
          </h1>
          <p
            v-if="activeScenario"
            class="text-xs text-[var(--color-text-secondary)] leading-tight"
          >
            {{ activeScenario.title }}
          </p>
          <p
            v-else
            class="text-xs text-[var(--color-text-secondary)] leading-tight"
          >
            Practice natural English
          </p>
        </div>
      </div>
    </div>
  </header>
</template>
