<script setup lang="ts">
import type { Correction } from '../types'
import Panel from 'primevue/panel';

const props = defineProps<{
  corrections: Correction[]
  suggestions: string[]
}>()

</script>

<template>
  <div
    v-if="corrections.length > 0 || suggestions.length > 0"
    class="mx-auto w-full px-2 sm:px-4 pb-2 animate-fade-in-up"
  >
    <Panel 
      toggleable 
      :collapsed="true"
      class="!bg-[var(--color-bg-card)] !backdrop-blur-xl !border !border-[var(--color-border)] !rounded-2xl !shadow-lg overflow-hidden"
      :pt="{
        header: '!bg-transparent !border-b !border-[var(--color-border)] !px-4 !py-2',
        title: '!text-xs !font-medium !text-[var(--color-text-secondary)]',
        content: '!bg-transparent !px-4 !py-3 !space-y-3'
      }"
    >
      <template #header>
        <div class="flex items-center gap-2 text-xs font-medium text-[var(--color-text-secondary)]">
          <span class="flex h-5 w-5 items-center justify-center rounded-md bg-[var(--color-accent-glow)] text-[var(--color-accent)]">
            ✨
          </span>
          Feedback & Suggestions
        </div>
      </template>

      <!-- Content -->
      <div>
        <!-- Corrections -->
        <div v-if="corrections.length > 0">
          <h4 class="mb-2 text-[11px] font-semibold uppercase tracking-wider text-[var(--color-warning)]">
            📝 Corrections
          </h4>
          <div
            v-for="(c, i) in corrections"
            :key="i"
            class="mb-2 rounded-lg bg-[rgba(245,158,11,0.06)] border border-[rgba(245,158,11,0.15)] p-2.5"
          >
            <div class="flex flex-wrap items-center gap-2 text-[13px]">
              <span class="line-through text-[var(--color-error)] opacity-80">{{ c.original }}</span>
              <svg width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="var(--color-text-muted)" stroke-width="2">
                <path d="M5 12h14" /><path d="M12 5l7 7-7 7" />
              </svg>
              <span class="font-medium text-[var(--color-success)]">{{ c.corrected }}</span>
            </div>
            <p class="mt-1 text-[11px] text-[var(--color-text-secondary)]">
              {{ c.explanation }}
            </p>
          </div>
        </div>

        <!-- Suggestions -->
        <div v-if="suggestions.length > 0" :class="{ 'mt-3': corrections.length > 0 }">
          <h4 class="mb-2 text-[11px] font-semibold uppercase tracking-wider text-[var(--color-accent)]">
            💡 Try These Phrases
          </h4>
          <div class="flex flex-wrap gap-1.5">
            <span
              v-for="(s, i) in suggestions"
              :key="i"
              class="inline-block rounded-lg bg-[var(--color-accent-glow)] border border-[rgba(14,165,233,0.2)] px-2.5 py-1 text-[11px] font-medium text-[var(--color-accent)]"
            >
              {{ s }}
            </span>
          </div>
        </div>
      </div>
    </Panel>
  </div>
</template>
