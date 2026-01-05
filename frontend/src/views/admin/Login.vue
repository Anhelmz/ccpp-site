<template>
  <div class="min-h-screen bg-main flex items-center justify-center px-4">
    <div class="max-w-md w-full bg-white rounded-lg shadow-xl p-8">
      <div class="text-center mb-8">
        <h1 class="text-3xl font-bold text-gray-900 mb-2">Admin Login</h1>
        <p class="text-gray-600">Calvary Chapel Phnom Penh</p>
      </div>

      <div
        v-if="error"
        class="mb-4 rounded-lg border border-red-200 bg-red-50 px-4 py-3 text-sm text-red-700"
      >
        {{ error }}
      </div>
      
      <form @submit.prevent="handleLogin" class="space-y-6">
        <div>
          <label for="email" class="block text-sm font-medium text-gray-700 mb-2">
            Email
          </label>
          <input
            id="email"
            v-model="email"
            type="email"
            required
            class="w-full px-4 py-3 border border-gray-300 rounded-lg focus:ring-2 focus:ring-main focus:border-transparent"
            placeholder="admin@example.com"
          />
        </div>
        
        <div>
          <label for="password" class="block text-sm font-medium text-gray-700 mb-2">
            Password
          </label>
          <input
            id="password"
            v-model="password"
            type="password"
            required
            class="w-full px-4 py-3 border border-gray-300 rounded-lg focus:ring-2 focus:ring-main focus:border-transparent"
            placeholder="Enter your password"
          />
        </div>
        
        <button
          type="submit"
          :disabled="loading"
          class="w-full bg-main text-white py-3 rounded-lg font-semibold hover:opacity-90 transition-colors shadow-lg"
          :class="{ 'opacity-70 cursor-not-allowed': loading }"
        >
          {{ loading ? 'Signing in...' : 'Sign In' }}
        </button>
      </form>
      
      <div class="mt-6 text-center">
        <router-link to="/" class="text-sm text-white hover:opacity-80">
          ← Back to Website
        </router-link>
      </div>
    </div>
  </div>
</template>

<script>
import { ref } from 'vue'
import { useRouter, useRoute } from 'vue-router'
import { login } from '@/services/auth'

export default {
  name: 'AdminLogin',
  setup() {
    const router = useRouter()
    const route = useRoute()
    const email = ref('')
    const password = ref('')
    const loading = ref(false)
    const error = ref('')
    
    const handleLogin = async () => {
      error.value = ''
      loading.value = true
      try {
        const { token, user } = await login(email.value, password.value)
        localStorage.setItem('authToken', token)
        localStorage.setItem('auth_token', token)
        localStorage.setItem('authUser', JSON.stringify(user))

        const redirectTo = route.query.redirect || '/admin/dashboard'
        router.push(redirectTo)
      } catch (err) {
        error.value = err.message || 'Unable to sign in'
      } finally {
        loading.value = false
      }
    }
    
    return {
      email,
      password,
      loading,
      error,
      handleLogin
    }
  }
}
</script>

