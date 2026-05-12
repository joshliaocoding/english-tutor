import { ref, computed } from 'vue'
import { useRouter } from 'vue-router'

const token = ref<string | null>(localStorage.getItem('token'))
const user = ref<any>(null)
const isLoading = ref(false)
const error = ref<string | null>(null)

export function useAuth() {
  const router = useRouter()

  const isAuthenticated = computed(() => !!token.value)

  const setAuth = (newToken: string, userData: any) => {
    token.value = newToken
    user.value = userData
    localStorage.setItem('token', newToken)
  }

  const logout = () => {
    token.value = null
    user.value = null
    localStorage.removeItem('token')
    router.push('/login')
  }

  const fetchUser = async () => {
    if (!token.value) return

    try {
      const res = await fetch('/api/auth/me', {
        headers: {
          'Authorization': `Bearer ${token.value}`
        }
      })
      if (res.ok) {
        user.value = await res.json()
      } else {
        logout()
      }
    } catch (err) {
      console.error('Failed to fetch user', err)
      logout()
    }
  }

  const login = async (email: string, password: string) => {
    isLoading.value = true
    error.value = null
    try {
      const res = await fetch('/api/auth/login', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ email, password })
      })
      const data = await res.json()
      if (!res.ok) {
        error.value = data.error || 'Login failed'
        return false
      }
      setAuth(data.token, data.user)
      return true
    } catch (err: any) {
      error.value = 'Network error'
      return false
    } finally {
      isLoading.value = false
    }
  }

  const register = async (username: string, email: string, password: string) => {
    isLoading.value = true
    error.value = null
    try {
      const res = await fetch('/api/auth/register', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ username, email, password })
      })
      const data = await res.json()
      if (!res.ok) {
        error.value = data.error || 'Registration failed'
        return false
      }
      setAuth(data.token, data.user)
      return true
    } catch (err: any) {
      error.value = 'Network error'
      return false
    } finally {
      isLoading.value = false
    }
  }

  return {
    token,
    user,
    isAuthenticated,
    isLoading,
    error,
    login,
    register,
    logout,
    fetchUser
  }
}
