<script setup lang="ts">
import type { Scenario } from '../types'
import Button from 'primevue/button';
import Avatar from 'primevue/avatar';

defineProps<{
  activeScenario: Scenario | null
}>()

defineEmits<{
  back: []
}>()
</script>

<template>
  <header
    class="sticky top-0 z-50 bg-[var(--color-bg-secondary)] border-b border-[var(--color-border)] px-2 py-2 sm:px-4"
  >
    <div class="mx-auto flex w-full items-center gap-2">
      <!-- Back button (visible when in a chat) -->
      <Button
        v-if="activeScenario"
        id="header-back-btn"
        icon="pi pi-arrow-left"
        text
        severity="secondary"
        rounded
        class="!w-10 !h-10 !p-0 text-[var(--color-text-secondary)] hover:bg-[var(--color-bg-card-hover)] hover:text-[var(--color-text-primary)] transition-all duration-200"
        @click="$emit('back')"
      />

      <!-- Logo + Title -->
      <div class="flex items-center gap-3 ml-1 cursor-pointer">
        <Avatar
          v-if="activeScenario"
          label="AI"
          shape="circle"
          class="bg-gradient-to-br from-[var(--color-accent)] to-sky-600 text-white font-bold shadow-md shrink-0 !w-10 !h-10 text-base"
        />
        <div
          v-else
          class="flex h-10 w-10 shrink-0 items-center justify-center rounded-xl bg-gradient-to-br from-[var(--color-accent)] to-sky-600 text-lg font-bold text-white shadow-lg"
        >
          S
        </div>

        <div class="flex flex-col justify-center">
          <h1 class="text-base font-semibold leading-tight text-[var(--color-text-primary)]">
            {{ activeScenario ? 'AI Tutor' : 'SpeakEasy' }}
          </h1>
          <p class="text-[13px] text-[var(--color-text-secondary)] mt-0.5 leading-tight flex items-center gap-1.5">
            <span v-if="activeScenario" class="w-1.5 h-1.5 rounded-full bg-[var(--color-success)] shadow-[0_0_5px_var(--color-success)]"></span>
            <span class="truncate max-w-[200px] sm:max-w-[300px]">
              {{ activeScenario ? 'online • ' + activeScenario.title : 'Practice natural English' }}
            </span>
          </p>
        </div>
      </div>
    </div>
  </header>
</template>
