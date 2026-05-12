import { createRouter, createWebHistory } from 'vue-router'
import { useAuth } from '../composables/useAuth'

import HomeView from '../views/HomeView.vue'
import LoginView from '../views/LoginView.vue'
import RegisterView from '../views/RegisterView.vue'
import ChatSessionView from '../views/ChatSessionView.vue'

const router = createRouter({
  history: createWebHistory(),
  routes: [
    {
      path: '/login',
      name: 'login',
      component: LoginView,
      meta: { requiresGuest: true }
    },
    {
      path: '/register',
      name: 'register',
      component: RegisterView,
      meta: { requiresGuest: true }
    },
    {
      path: '/',
      name: 'home',
      component: HomeView,
      meta: { requiresAuth: true }
    },
    {
      path: '/chat/:id',
      name: 'chat',
      component: ChatSessionView,
      meta: { requiresAuth: true }
    }
  ]
})

router.beforeEach(async (to, from, next) => {
  const { isAuthenticated, fetchUser, token } = useAuth()
  
  if (token.value && !isAuthenticated.value) {
    await fetchUser()
  }

  if (to.meta.requiresAuth && !isAuthenticated.value) {
    next({ name: 'login' })
  } else if (to.meta.requiresGuest && isAuthenticated.value) {
    next({ name: 'home' })
  } else {
    next()
  }
})

export default router
