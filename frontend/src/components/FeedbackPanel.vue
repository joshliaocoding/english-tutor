<script setup lang="ts">
import { ref } from 'vue'
import type { Correction } from '../types'

const props = defineProps<{
  corrections: Correction[]
  suggestions: string[]
}>()

const isOpen = ref(true)
</script>

<template>
  <div
    v-if="corrections.length > 0 || suggestions.length > 0"
    class="mx-auto w-full max-w-3xl px-3 sm:px-4 animate-fade-in-up"
  >
    <div class="glass rounded-xl overflow-hidden">
      <!-- Toggle header -->
      <button
        id="feedback-toggle"
        class="flex w-full items-center justify-between px-4 py-2.5 text-xs font-medium text-[var(--color-text-secondary)] hover:text-[var(--color-text-primary)] transition-colors"
        @click="isOpen = !isOpen"
      >
        <span class="flex items-center gap-2">
          <span class="flex h-5 w-5 items-center justify-center rounded-md bg-[var(--color-accent-glow)] text-[var(--color-accent)]">
            ✨
          </span>
          Feedback & Suggestions
        </span>
        <svg
          width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor"
          stroke-width="2" stroke-linecap="round" stroke-linejoin="round"
          class="transition-transform duration-200"
          :class="{ 'rotate-180': isOpen }"
        >
          <path d="M6 9l6 6 6-6" />
        </svg>
      </button>

      <!-- Content -->
      <div v-if="isOpen" class="animate-slide-down border-t border-[var(--color-border)] px-4 py-3 space-y-3">
        <!-- Corrections -->
        <div v-if="corrections.length > 0">
          <h4 class="mb-2 text-xs font-semibold uppercase tracking-wider text-[var(--color-warning)]">
            📝 Corrections
          </h4>
          <div
            v-for="(c, i) in corrections"
            :key="i"
            class="mb-2 rounded-lg bg-[rgba(245,158,11,0.06)] border border-[rgba(245,158,11,0.15)] p-3"
          >
            <div class="flex flex-wrap items-center gap-2 text-sm">
              <span class="line-through text-[var(--color-error)] opacity-80">{{ c.original }}</span>
              <svg width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="var(--color-text-muted)" stroke-width="2">
                <path d="M5 12h14" /><path d="M12 5l7 7-7 7" />
              </svg>
              <span class="font-medium text-[var(--color-success)]">{{ c.corrected }}</span>
            </div>
            <p class="mt-1.5 text-xs text-[var(--color-text-secondary)]">
              {{ c.explanation }}
            </p>
          </div>
        </div>

        <!-- Suggestions -->
        <div v-if="suggestions.length > 0">
          <h4 class="mb-2 text-xs font-semibold uppercase tracking-wider text-[var(--color-accent)]">
            💡 Try These Phrases
          </h4>
          <div class="flex flex-wrap gap-2">
            <span
              v-for="(s, i) in suggestions"
              :key="i"
              class="inline-block rounded-lg bg-[var(--color-accent-glow)] border border-[rgba(0,212,170,0.2)] px-3 py-1.5 text-xs font-medium text-[var(--color-accent)]"
            >
              {{ s }}
            </span>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
