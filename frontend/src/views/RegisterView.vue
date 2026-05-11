<script setup lang="ts">
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import { useAuth } from '../composables/useAuth'

const router = useRouter()
const { register, isLoading, error } = useAuth()

const username = ref('')
const email = ref('')
const password = ref('')

const handleRegister = async () => {
  const success = await register(username.value, email.value, password.value)
  if (success) {
    router.push('/')
  }
}
</script>

<template>
  <div class="flex h-full w-full items-center justify-center p-4">
    <div class="glass animate-fade-in-up w-full max-w-md rounded-2xl p-8 shadow-2xl">
      <div class="mb-8 text-center">
        <h2 class="text-3xl font-bold text-[var(--color-text-primary)]">Create Account</h2>
        <p class="mt-2 text-[var(--color-text-secondary)]">Start your English journey today</p>
      </div>

      <form @submit.prevent="handleRegister" class="space-y-6">
        <div v-if="error" class="rounded-lg bg-red-500/10 p-3 text-sm text-red-400">
          {{ error }}
        </div>

        <div>
          <label class="mb-2 block text-sm font-medium text-[var(--color-text-primary)]">Username</label>
          <input
            v-model="username"
            type="text"
            required
            class="w-full rounded-xl border border-[var(--color-border)] bg-[var(--color-bg-input)] px-4 py-3 text-[var(--color-text-primary)] transition-colors focus:border-[var(--color-accent)] focus:outline-none focus:ring-1 focus:ring-[var(--color-accent)]"
            placeholder="johndoe"
          />
        </div>

        <div>
          <label class="mb-2 block text-sm font-medium text-[var(--color-text-primary)]">Email</label>
          <input
            v-model="email"
            type="email"
            required
            class="w-full rounded-xl border border-[var(--color-border)] bg-[var(--color-bg-input)] px-4 py-3 text-[var(--color-text-primary)] transition-colors focus:border-[var(--color-accent)] focus:outline-none focus:ring-1 focus:ring-[var(--color-accent)]"
            placeholder="you@example.com"
          />
        </div>

        <div>
          <label class="mb-2 block text-sm font-medium text-[var(--color-text-primary)]">Password</label>
          <input
            v-model="password"
            type="password"
            required
            class="w-full rounded-xl border border-[var(--color-border)] bg-[var(--color-bg-input)] px-4 py-3 text-[var(--color-text-primary)] transition-colors focus:border-[var(--color-accent)] focus:outline-none focus:ring-1 focus:ring-[var(--color-accent)]"
            placeholder="••••••••"
          />
        </div>

        <button
          type="submit"
          :disabled="isLoading"
          class="flex w-full items-center justify-center rounded-xl bg-[var(--color-accent)] px-4 py-3 font-semibold text-white shadow-lg transition-all hover:bg-[var(--color-accent-hover)] disabled:opacity-50"
        >
          <i v-if="isLoading" class="pi pi-spinner animate-spin mr-2"></i>
          {{ isLoading ? 'Creating account...' : 'Sign Up' }}
        </button>
      </form>

      <div class="mt-6 text-center text-sm text-[var(--color-text-secondary)]">
        Already have an account?
        <router-link to="/login" class="font-medium text-[var(--color-accent)] hover:underline">
          Log in
        </router-link>
      </div>
    </div>
  </div>
</template>
