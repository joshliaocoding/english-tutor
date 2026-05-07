<script setup lang="ts">
import type { Scenario } from '../types'

defineProps<{
  scenarios: Scenario[]
  isLoading: boolean
}>()

defineEmits<{
  select: [scenario: Scenario]
}>()
</script>

<template>
  <div class="flex flex-1 flex-col items-center justify-center px-4 py-8 sm:py-12">
    <!-- Hero section -->
    <div class="mb-10 text-center animate-fade-in-up sm:mb-14">
      <div class="mx-auto mb-5 flex h-16 w-16 items-center justify-center rounded-2xl bg-gradient-to-br from-[var(--color-accent)] to-emerald-600 text-3xl shadow-lg glow-accent sm:h-20 sm:w-20 sm:text-4xl">
        💬
      </div>
      <h2 class="mb-3 text-2xl font-bold text-[var(--color-text-primary)] sm:text-3xl">
        Choose a Scenario
      </h2>
      <p class="mx-auto max-w-md text-sm text-[var(--color-text-secondary)] sm:text-base">
        Pick a real-world conversation to practice. Our AI tutor will roleplay with you and help improve your English.
      </p>
    </div>

    <!-- Scenario cards grid -->
    <div class="mx-auto grid w-full max-w-3xl grid-cols-1 gap-4 sm:grid-cols-2 lg:grid-cols-3">
      <button
        v-for="(scenario, index) in scenarios"
        :key="scenario.id"
        :id="`scenario-${scenario.id}`"
        class="group glass rounded-2xl p-5 text-left transition-all duration-300 hover:bg-[var(--color-bg-card-hover)] hover:border-[var(--color-border-accent)] hover:glow-accent-strong hover:scale-[1.02] active:scale-[0.98] disabled:opacity-50 disabled:cursor-not-allowed"
        :style="{ animationDelay: `${index * 80}ms` }"
        :class="['animate-fade-in-up']"
        :disabled="isLoading"
        @click="$emit('select', scenario)"
      >
        <div class="mb-3 text-3xl sm:text-4xl">{{ scenario.icon }}</div>
        <h3 class="mb-1.5 text-base font-semibold text-[var(--color-text-primary)] group-hover:text-[var(--color-accent)] transition-colors duration-200">
          {{ scenario.title }}
        </h3>
        <p class="text-xs leading-relaxed text-[var(--color-text-secondary)] sm:text-sm">
          {{ scenario.description }}
        </p>

        <!-- Arrow indicator -->
        <div class="mt-3 flex items-center gap-1 text-xs font-medium text-[var(--color-accent)] opacity-0 transition-opacity duration-200 group-hover:opacity-100">
          Start conversation
          <svg width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
            <path d="M5 12h14" /><path d="M12 5l7 7-7 7" />
          </svg>
        </div>
      </button>
    </div>
  </div>
</template>
