<script setup lang="ts">
import { onMounted, watch } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { useChat } from '../composables/useChat'
import ChatWindow from '../components/ChatWindow.vue'

const route = useRoute()
const router = useRouter()
const { messages, isLoading, sendMessage, sessionId, activeScenario } = useChat()

onMounted(() => {
  // In a real app we might fetch the specific session if not loaded
  // For now, if there is no active session, redirect to home
  if (!sessionId.value || route.params.id !== sessionId.value) {
    router.push('/')
  }
})
</script>

<template>
  <ChatWindow
    v-if="sessionId"
    :messages="messages"
    :is-loading="isLoading"
    @send="sendMessage"
    class="flex-1 overflow-hidden"
  />
</template>
