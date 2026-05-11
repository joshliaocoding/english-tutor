<script setup lang="ts">
import { useRouter, useRoute } from 'vue-router'
import { useChat } from './composables/useChat'
import { useAuth } from './composables/useAuth'
import Header from './components/Header.vue'

const router = useRouter()
const route = useRoute()
const { activeScenario, endSession } = useChat()
const { logout } = useAuth()

const handleBack = () => {
  endSession()
  router.push('/')
}
</script>

<template>
  <div class="flex h-dvh w-full items-center justify-center overflow-hidden">
    <!-- App Container -->
    <div class="relative flex h-full w-full flex-col bg-[var(--color-bg-primary)] sm:border-x sm:border-[var(--color-border)] sm:shadow-2xl">
      <Header
        v-if="!['login', 'register'].includes(route.name as string)"
        :active-scenario="activeScenario"
        @back="handleBack"
        @logout="logout"
      />

      <router-view />
    </div>
  </div>
</template>
