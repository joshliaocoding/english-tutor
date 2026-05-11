<script setup lang="ts">
import { onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { useChat } from '../composables/useChat'
import ScenarioPicker from '../components/ScenarioPicker.vue'

const router = useRouter()
const { scenarios, isLoading, fetchScenarios, startSession } = useChat()

onMounted(() => {
  fetchScenarios()
})

const handleSelect = async (scenario: any) => {
  await startSession(scenario)
  // Assuming startSession sets the sessionId globally
  const { sessionId } = useChat()
  if (sessionId.value) {
    router.push(`/chat/${sessionId.value}`)
  }
}
</script>

<template>
  <ScenarioPicker
    :scenarios="scenarios"
    :is-loading="isLoading"
    @select="handleSelect"
    class="overflow-y-auto"
  />
</template>
