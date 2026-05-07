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
  <div class="flex min-h-dvh flex-col">
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
    />

    <!-- Chat window (active session) -->
    <ChatWindow
      v-else
      :messages="messages"
      :is-loading="isLoading"
      @send="sendMessage"
    />
  </div>
</template>
