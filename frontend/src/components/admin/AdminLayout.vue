<template>
  <div class="min-h-screen bg-gray-100">
    <!-- Sidebar -->
    <aside
      :class="[
        'fixed inset-y-0 left-0 z-50 w-64 bg-gray-900 text-white transform transition-transform duration-300 ease-in-out',
        sidebarOpen ? 'translate-x-0' : '-translate-x-full lg:translate-x-0'
      ]"
    >
      <div class="flex flex-col h-full">
        <!-- Logo -->
        <div class="flex items-center justify-between h-16 px-6 border-b border-gray-800">
          <div class="flex items-center space-x-3">
            <div class="w-8 h-8 bg-primary-600 rounded-lg flex items-center justify-center">
              <svg class="w-5 h-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 6V4m0 2a2 2 0 100 4m0-4a2 2 0 110 4m-6 8a2 2 0 100-4m0 4a2 2 0 110-4m0 4v2m0-6V4m6 6v10m6-2a2 2 0 100-4m0 4a2 2 0 110-4m0 4v2m0-6V4"></path>
              </svg>
            </div>
            <span class="text-lg font-bold">Admin Panel</span>
          </div>
          <button
            @click="sidebarOpen = false"
            class="lg:hidden text-gray-400 hover:text-white"
          >
            <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"></path>
            </svg>
          </button>
        </div>

        <!-- Navigation -->
        <nav class="flex-1 px-4 py-6 space-y-2 overflow-y-auto">
          <router-link
            to="/admin/dashboard"
            class="flex items-center space-x-3 px-4 py-3 rounded-lg transition-colors"
            :class="$route.path === '/admin/dashboard' ? 'bg-primary-600 text-white' : 'text-gray-300 hover:bg-gray-800 hover:text-white'"
          >
            <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 12l2-2m0 0l7-7 7 7M5 10v10a1 1 0 001 1h3m10-11l2 2m-2-2v10a1 1 0 01-1 1h-3m-6 0a1 1 0 001-1v-4a1 1 0 011-1h2a1 1 0 011 1v4a1 1 0 001 1m-6 0h6"></path>
            </svg>
            <span class="font-medium">Dashboard</span>
          </router-link>

          <router-link
            to="/admin/gallery"
            class="flex items-center space-x-3 px-4 py-3 rounded-lg transition-colors"
            :class="$route.path === '/admin/gallery' ? 'bg-primary-600 text-white' : 'text-gray-300 hover:bg-gray-800 hover:text-white'"
          >
            <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 16l4.586-4.586a2 2 0 012.828 0L16 16m-2-2l1.586-1.586a2 2 0 012.828 0L20 14m-6-6h.01M6 20h12a2 2 0 002-2V6a2 2 0 00-2-2H6a2 2 0 00-2 2v12a2 2 0 002 2z"></path>
            </svg>
            <span class="font-medium">Gallery</span>
          </router-link>

          <router-link
            to="/admin/events"
            class="flex items-center space-x-3 px-4 py-3 rounded-lg transition-colors"
            :class="$route.path === '/admin/events' ? 'bg-primary-600 text-white' : 'text-gray-300 hover:bg-gray-800 hover:text-white'"
          >
            <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 7V3m8 4V3m-9 8h10M5 21h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v12a2 2 0 002 2z"></path>
            </svg>
            <span class="font-medium">Events</span>
          </router-link>

          <router-link
            to="/admin/content"
            class="flex items-center space-x-3 px-4 py-3 rounded-lg transition-colors"
            :class="$route.path === '/admin/content' ? 'bg-primary-600 text-white' : 'text-gray-300 hover:bg-gray-800 hover:text-white'"
          >
            <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11 5H6a2 2 0 00-2 2v11a2 2 0 002 2h11a2 2 0 002-2v-5m-1.414-9.414a2 2 0 112.828 2.828L11.828 15H9v-2.828l8.586-8.586z"></path>
            </svg>
            <span class="font-medium">Content</span>
          </router-link>

          <router-link
            to="/admin/contacts"
            class="flex items-center space-x-3 px-4 py-3 rounded-lg transition-colors"
            :class="$route.path === '/admin/contacts' ? 'bg-primary-600 text-white' : 'text-gray-300 hover:bg-gray-800 hover:text-white'"
          >
            <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 8l7.89 5.26a2 2 0 002.22 0L21 8M5 19h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v10a2 2 0 002 2z"></path>
            </svg>
            <span class="font-medium">Messages</span>
          </router-link>

          <router-link
            to="/admin/settings"
            class="flex items-center space-x-3 px-4 py-3 rounded-lg transition-colors"
            :class="$route.path === '/admin/settings' ? 'bg-primary-600 text-white' : 'text-gray-300 hover:bg-gray-800 hover:text-white'"
          >
            <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10.325 4.317c.426-1.756 2.924-1.756 3.35 0a1.724 1.724 0 002.573 1.066c1.543-.94 3.31.826 2.37 2.37a1.724 1.724 0 001.065 2.572c1.756.426 1.756 2.924 0 3.35a1.724 1.724 0 00-1.066 2.573c.94 1.543-.826 3.31-2.37 2.37a1.724 1.724 0 00-2.572 1.065c-.426 1.756-2.924 1.756-3.35 0a1.724 1.724 0 00-2.573-1.066c-1.543.94-3.31-.826-2.37-2.37a1.724 1.724 0 00-1.065-2.572c-1.756-.426-1.756-2.924 0-3.35a1.724 1.724 0 001.066-2.573c-.94-1.543.826-3.31 2.37-2.37.996.608 2.296.07 2.572-1.065z"></path>
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z"></path>
            </svg>
            <span class="font-medium">Settings</span>
          </router-link>
        </nav>

        <!-- Footer -->
        <div class="p-4 border-t border-gray-800">
          <div class="flex items-center space-x-3 mb-4">
            <div class="w-10 h-10 bg-primary-600 rounded-full flex items-center justify-center">
              <span class="text-white font-semibold">
                {{ displayInitial }}
              </span>
            </div>
            <div class="flex-1">
              <p class="text-sm font-medium text-white">
                {{ displayName }}
              </p>
              <p class="text-xs text-gray-400">
                {{ displayEmail }}
              </p>
            </div>
          </div>
          <button
            @click="handleLogout"
            class="w-full flex items-center justify-center space-x-2 px-4 py-2 bg-gray-800 text-gray-300 rounded-lg hover:bg-gray-700 hover:text-white transition-colors"
          >
            <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 16l4-4m0 0l-4-4m4 4H7m6 4v1a3 3 0 01-3 3H6a3 3 0 01-3-3V7a3 3 0 013-3h4a3 3 0 013 3v1"></path>
            </svg>
            <span class="text-sm font-medium">Logout</span>
          </button>
        </div>
      </div>
    </aside>

    <!-- Overlay for mobile -->
    <div
      v-if="sidebarOpen"
      class="fixed inset-0 bg-black bg-opacity-50 z-40 lg:hidden"
      @click="sidebarOpen = false"
    ></div>

    <!-- Main Content -->
    <div class="lg:pl-64">
      <!-- Top Bar -->
      <header class="sticky top-0 z-30 bg-white shadow-sm border-b border-gray-200">
        <div class="flex items-center justify-between h-16 px-4 sm:px-6 lg:px-8">
          <button
            @click="sidebarOpen = true"
            class="lg:hidden text-gray-600 hover:text-gray-900"
          >
            <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 6h16M4 12h16M4 18h16"></path>
            </svg>
          </button>
          
          <div class="flex-1 lg:ml-0 ml-4">
            <h1 class="text-xl font-semibold text-gray-900">{{ pageTitle }}</h1>
          </div>

          <div class="flex items-center space-x-4">
            <button class="p-2 text-gray-600 hover:text-gray-900 hover:bg-gray-100 rounded-lg transition-colors">
              <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 17h5l-1.405-1.405A2.032 2.032 0 0118 14.158V11a6.002 6.002 0 00-4-5.659V5a2 2 0 10-4 0v.341C7.67 6.165 6 8.388 6 11v3.159c0 .538-.214 1.055-.595 1.436L4 17h5m6 0v1a3 3 0 11-6 0v-1m6 0H9"></path>
              </svg>
            </button>
            <router-link
              to="/"
              target="_blank"
              class="px-4 py-2 text-sm text-gray-600 hover:text-gray-900 hover:bg-gray-100 rounded-lg transition-colors"
            >
              View Site
            </router-link>
          </div>
        </div>
      </header>

      <!-- Page Content -->
      <main class="p-4 sm:p-6 lg:p-8">
        <slot></slot>
      </main>
    </div>
  </div>
</template>

<script>
import { ref, computed, onMounted } from 'vue'
import { useRoute, useRouter } from 'vue-router'

export default {
  name: 'AdminLayout',
  setup() {
    const route = useRoute()
    const router = useRouter()
    const sidebarOpen = ref(false)
    const user = ref(null)

    const pageTitle = computed(() => {
      const titles = {
        '/admin/dashboard': 'Dashboard',
        '/admin/gallery': 'Gallery Management',
        '/admin/events': 'Events & Calendar',
        '/admin/content': 'Content Management',
        '/admin/contacts': 'Contact Messages',
        '/admin/settings': 'Settings'
      }
      return titles[route.path] || 'Admin Panel'
    })

    const displayName = computed(() => user.value?.name || 'Admin User')
    const displayEmail = computed(() => user.value?.email || 'admin@ccpp.org')
    const displayInitial = computed(() => displayName.value.charAt(0).toUpperCase())

    const handleLogout = () => {
      localStorage.removeItem('authToken')
      localStorage.removeItem('auth_token')
      localStorage.removeItem('authUser')
      router.push('/admin/login')
    }

    onMounted(() => {
      const savedUser = localStorage.getItem('authUser')
      if (savedUser) {
        try {
          user.value = JSON.parse(savedUser)
        } catch {
          user.value = null
        }
      }
    })

    return {
      sidebarOpen,
      pageTitle,
      user,
      displayName,
      displayEmail,
      displayInitial,
      handleLogout
    }
  }
}
</script>

