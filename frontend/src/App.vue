<script setup lang="ts">
import { onMounted } from 'vue'
import { useChat } from './composables/useChat'
import Header from './components/Header.vue'
import ScenarioPicker from './components/ScenarioPicker.vue'
import ChatWindow from './components/ChatWindow.vue'

const {
  scenarios,
  messages,
  activeScenario,
  isLoading,
  fetchScenarios,
  startSession,
  sendMessage,
  endSession,
} = useChat()

onMounted(() => {
  fetchScenarios()
})
</script>

<template>
  <div class="flex h-dvh w-full items-center justify-center overflow-hidden">
    <!-- App Container -->
    <div class="relative flex h-full w-full max-w-2xl flex-col bg-[var(--color-bg-primary)] sm:border-x sm:border-[var(--color-border)] sm:shadow-2xl">
      <Header
        :active-scenario="activeScenario"
        @back="endSession"
      />

      <!-- Scenario picker (no active session) -->
      <ScenarioPicker
        v-if="!activeScenario"
        :scenarios="scenarios"
        :is-loading="isLoading"
        @select="startSession"
        class="overflow-y-auto"
      />

      <!-- Chat window (active session) -->
      <ChatWindow
        v-else
        :messages="messages"
        :is-loading="isLoading"
        @send="sendMessage"
        class="flex-1 overflow-hidden"
      />
    </div>
  </div>
</template>
