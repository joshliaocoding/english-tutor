<script setup lang="ts">
import { ref } from 'vue'
import Textarea from 'primevue/textarea';
import Button from 'primevue/button';

defineProps<{
  isLoading: boolean
}>()

const emit = defineEmits<{
  send: [message: string]
}>()

const inputText = ref('')

function handleSend() {
  if (!inputText.value.trim()) return
  emit('send', inputText.value)
  inputText.value = ''
}

function handleKeydown(e: KeyboardEvent) {
  if (e.key === 'Enter' && !e.shiftKey) {
    e.preventDefault()
    handleSend()
  }
}
</script>

<template>
  <div class="bg-[var(--color-bg-primary)] p-2 sm:p-3 border-t border-[var(--color-border)]">
    <div class="mx-auto flex w-full items-end gap-2">
      <div class="relative flex-1">
        <Textarea
          id="chat-input"
          v-model="inputText"
          rows="1"
          autoResize
          placeholder="Message"
          class="w-full !rounded-3xl !bg-[var(--color-bg-input)] !text-[var(--color-text-primary)] !border-[var(--color-border)] !px-4 !py-3 !min-h-[44px] !max-h-[120px] focus:!border-[var(--color-border-accent)] transition-all duration-200"
          :disabled="isLoading"
          @keydown="handleKeydown"
        />
      </div>

      <Button
        id="send-btn"
        icon="pi pi-send"
        class="!h-[44px] !w-[44px] shrink-0 !rounded-full !bg-gradient-to-br !from-[var(--color-accent)] !to-sky-600 !border-none !text-white shadow-md transition-all duration-200 hover:scale-105 active:scale-95"
        :loading="isLoading"
        :disabled="isLoading || !inputText.trim()"
        @click="handleSend"
      />
    </div>
  </div>
</template>
