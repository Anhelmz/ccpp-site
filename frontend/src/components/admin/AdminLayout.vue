<template>
  <div :class="['admin-root flex flex-col h-screen', isDarkMode ? 'bg-[#050505] text-white admin-dark-mode' : 'bg-gray-50']">
    <!-- Top Bar -->
    <nav
      class="border-b shadow-sm"
      :class="isDarkMode ? 'bg-[#000000] border-[#0C94AB] text-white' : 'bg-gray-50 border-gray-200'"
    >
      <div class="px-4 sm:px-6 flex items-center justify-between h-12">
        <div class="flex items-center">
          <button
            @click="toggleSidebar"
            class="lg:hidden mr-3 p-2 rounded-md transition-colors"
            aria-label="Toggle menu"
            :class="isDarkMode ? 'text-white hover:bg-[#0C94AB1A]' : 'text-zinc-600 hover:text-zinc-900 hover:bg-gray-100'"
          >
            <svg class="h-5 w-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 6h16M4 12h16M4 18h16" />
            </svg>
          </button>
          <div class="flex items-center">
            <div class="w-8 sm:w-10 flex flex-col items-center justify-center mr-1 mb-0.5">
              <div class="flex gap-1 sm:gap-1.5 mb-0.5">
                <div class="w-1 sm:w-1.5 h-1 sm:h-1.5 bg-cyan-400 rounded-full"></div>
                <div class="w-1 sm:w-1.5 h-1 sm:h-1.5 bg-cyan-400 rounded-full"></div>
              </div>
              <div class="w-1 sm:w-1.5 h-1 sm:h-1.5 bg-cyan-400 rounded-full"></div>
            </div>
            <h1 class="text-xl sm:text-2xl font-light tracking-wider leading-tight" :class="isDarkMode ? 'text-white' : 'text-zinc-700'">
              ANHELM
            </h1>
          </div>
        </div>

        <div class="flex items-center">
          <button
            type="button"
            @click="toggleTheme"
            :aria-pressed="isDarkMode"
            class="flex items-center gap-2 px-3 py-2 text-sm font-medium rounded-md border transition-colors"
            :class="isDarkMode ? 'bg-[#0a0a0a] border-[#0C94AB] text-white hover:bg-[#0C94AB26]' : 'border-gray-200 text-zinc-600 hover:text-zinc-800 hover:bg-white'"
          >
            <svg v-if="isDarkMode" class="h-4 w-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 3v1m0 16v1m9-9h-1M4 12H3m15.364 6.364l-.707-.707M6.343 6.343l-.707-.707m12.728 0l-.707.707M6.343 17.657l-.707.707M16 12a4 4 0 11-8 0 4 4 0 018 0z" />
            </svg>
            <svg v-else class="h-4 w-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 12.79A9 9 0 1111.21 3 7 7 0 0021 12.79z" />
            </svg>
            <span>{{ isDarkMode ? 'Light mode' : 'Dark mode' }}</span>
          </button>
        </div>
      </div>
    </nav>

    <div class="flex flex-1 overflow-hidden relative">
      <div
        v-if="sidebarOpen"
        @click="closeSidebar"
        class="lg:hidden fixed inset-0 bg-black bg-opacity-50 z-40"
      ></div>

      <!-- Sidebar -->
      <nav
        :class="[
          'fixed lg:static inset-y-0 left-0 z-50 w-64 shadow-sm border-r flex flex-col transform transition-transform duration-300 ease-in-out',
          sidebarOpen ? 'translate-x-0' : '-translate-x-full lg:translate-x-0',
          isDarkMode ? 'bg-[#000000] border-[#0C94AB] shadow-lg' : 'bg-gray-50 border-gray-200'
        ]"
      >
        <div class="flex-1 py-4 overflow-y-auto">
          <nav class="space-y-px">
            <router-link
              to="/admin/dashboard"
              @click="closeSidebarOnMobile"
              :class="linkClass('AdminDashboard')"
            >
              <span class="mr-3 inline-flex items-center justify-center">
                <svg class="h-5 w-5" :class="isDarkMode ? 'text-[#0C94AB]' : 'text-cyan-500'" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 12l2-2 7-7 7 7 2 2v7a1 1 0 01-1 1h-4v-4a2 2 0 00-2-2h-4a2 2 0 00-2 2v4H4a1 1 0 01-1-1v-7z" />
                </svg>
              </span>
              Dashboard
            </router-link>

            <router-link
              to="/admin/gallery"
              @click="closeSidebarOnMobile"
              :class="linkClass('GalleryManagement', ['GalleryCategory'])"
            >
              <span class="mr-3 inline-flex items-center justify-center">
                <svg class="h-5 w-5" :class="isDarkMode ? 'text-[#0C94AB]' : 'text-cyan-500'" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 16l4.586-4.586a2 2 0 012.828 0L16 16l4-4m-8 4h.01" />
                </svg>
              </span>
              Galleries
            </router-link>
            <div class="ml-6 space-y-0.5">
              <router-link
                v-for="section in gallerySections"
                :key="section.value"
                :to="section.to"
                @click="closeSidebarOnMobile"
                :class="subLinkClass(section.value)"
              >
                <span class="mr-2 inline-flex items-center justify-center h-5 w-3">
                  <span class="h-1 w-1 rounded-full" :class="isDarkMode ? 'bg-[#0C94AB]' : 'bg-cyan-500'"></span>
                </span>
                {{ section.label }}
              </router-link>
            </div>

            <router-link
              to="/admin/events"
              @click="closeSidebarOnMobile"
              :class="linkClass('EventManagement')"
            >
              <span class="mr-3 inline-flex items-center justify-center">
                <svg class="h-5 w-5" :class="isDarkMode ? 'text-[#0C94AB]' : 'text-green-500'" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 7V3m8 4V3m-9 8h10M5 21h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v12a2 2 0 002 2z" />
                </svg>
              </span>
              Events
            </router-link>

            <router-link
              to="/admin/calendar"
              @click="closeSidebarOnMobile"
              :class="linkClass('AdminCalendar')"
            >
              <span class="mr-3 inline-flex items-center justify-center">
                <svg class="h-5 w-5" :class="isDarkMode ? 'text-[#0C94AB]' : 'text-amber-500'" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 7V3m8 4V3m-9 8h10M5 21h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v12a2 2 0 002 2z" />
                </svg>
              </span>
              Calendar
            </router-link>

            <router-link
              to="/admin/videos"
              @click="closeSidebarOnMobile"
              :class="linkClass('VideoManagement')"
            >
              <span class="mr-3 inline-flex items-center justify-center">
                <svg class="h-5 w-5" :class="isDarkMode ? 'text-[#0C94AB]' : 'text-purple-500'" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 10l4.553-2.276a1 1 0 010 1.788L15 12.222M9 10L4.447 7.724a1 1 0 000 1.788L9 12.222m6 0l4.553 2.276a1 1 0 010 1.788L15 18m-6-5.778l-4.553 2.276a1 1 0 000 1.788L9 18m0-8l6 3m-6 0l6 3" />
                </svg>
              </span>
              Videos
            </router-link>

            <router-link
              to="/admin/content"
              @click="closeSidebarOnMobile"
              :class="linkClass('ContentManagement')"
            >
              <span class="mr-3 inline-flex items-center justify-center">
                <svg class="h-5 w-5" :class="isDarkMode ? 'text-[#0C94AB]' : 'text-blue-500'" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11 5H6a2 2 0 00-2 2v11a2 2 0 002 2h11a2 2 0 002-2v-5m-1.414-9.414a2 2 0 112.828 2.828L11.828 15H9v-2.828l8.586-8.586z" />
                </svg>
              </span>
              Content
            </router-link>

            <router-link
              to="/admin/contacts"
              @click="closeSidebarOnMobile"
              :class="linkClass('ContactMessages', ['contact-requests', 'contact-request-detail'])"
            >
              <span class="mr-3 inline-flex items-center justify-center">
                <svg class="h-5 w-5" :class="isDarkMode ? 'text-[#0C94AB]' : 'text-cyan-500'" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 8l7.89 5.26a2 2 0 002.22 0L21 8m-4 8H7" />
                </svg>
              </span>
              Contact Requests
            </router-link>

            <router-link
              to="/admin/settings"
              @click="closeSidebarOnMobile"
              :class="linkClass('Settings')"
            >
              <span class="mr-3 inline-flex items-center justify-center">
                <svg class="h-5 w-5" :class="isDarkMode ? 'text-[#0C94AB]' : 'text-zinc-500'" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10.325 4.317c.426-1.756 2.924-1.756 3.35 0a1.724 1.724 0 002.573 1.066c1.543-.94 3.31.826 2.37 2.37a1.724 1.724 0 001.065 2.572c1.756.426 1.756 2.924 0 3.35a1.724 1.724 0 00-1.066 2.573c.94 1.543-.826 3.31-2.37 2.37a1.724 1.724 0 00-2.572 1.065c-.426 1.756-2.924 1.756-3.35 0a1.724 1.724 0 00-2.573-1.066c-1.543.94-3.31-.826-2.37-2.37a1.724 1.724 0 00-1.065-2.572c-1.756-.426-1.756-2.924 0-3.35a1.724 1.724 0 001.066-2.573c-.94-1.543.826-3.31 2.37-2.37.996.608 2.296.07 2.572-1.065z" />
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z" />
                </svg>
              </span>
              Settings
            </router-link>
          </nav>
        </div>

        <!-- User Section -->
        <div class="border-t" :class="isDarkMode ? 'border-[#0C94AB]' : 'border-gray-200'">
          <div class="p-4">
            <div class="flex items-center">
              <div class="flex-shrink-0">
                <div class="w-8 h-8 rounded-full flex items-center justify-center" :class="isDarkMode ? 'bg-[#0C94AB] text-black' : 'bg-cyan-100'">
                  <span class="text-sm font-medium" :class="isDarkMode ? 'text-black' : 'text-cyan-600'">{{ userInitial }}</span>
                </div>
              </div>
              <div class="ml-3 flex-1 min-w-0">
                <p class="text-sm font-medium truncate" :class="isDarkMode ? 'text-white' : 'text-zinc-700'">{{ userEmail }}</p>
                <button
                  @click="handleLogout"
                  class="text-xs"
                  :class="isDarkMode ? 'text-slate-300 hover:text-white' : 'text-zinc-500 hover:text-zinc-700'"
                >
                  Sign out
                </button>
              </div>
            </div>
          </div>
          <div class="flex items-center justify-center pt-3 pb-4 border-t px-4" :class="isDarkMode ? 'border-[#0C94AB]' : 'border-gray-200'">
            <div class="w-5 flex flex-col items-center justify-center mr-1 mb-0.5">
              <div class="flex gap-1 mb-0.5">
                <div class="w-1 h-1 bg-cyan-400 rounded-full"></div>
                <div class="w-1 h-1 bg-cyan-400 rounded-full"></div>
              </div>
              <div class="w-1 h-1 bg-cyan-400 rounded-full"></div>
            </div>
            <span class="text-xs font-light text-zinc-500 tracking-wider leading-tight" :class="isDarkMode ? 'text-slate-300' : ''">ANHELM</span>
          </div>
        </div>
      </nav>

      <!-- Main Content -->
      <div class="flex-1 flex flex-col overflow-hidden">
        <main :class="['flex-1 overflow-y-auto', isDarkMode ? 'bg-[#000000] text-white' : 'bg-white']">
          <div class="py-4 sm:py-6">
            <div class="px-4 sm:px-6">
              <slot />
            </div>
          </div>
        </main>
      </div>
    </div>
  </div>
</template>

<script>
import { ref, computed, onMounted } from 'vue'
import { useRouter, useRoute } from 'vue-router'

export default {
  name: 'AdminLayout',
  props: {
    user: {
      type: Object,
      default: null
    },
    isDark: {
      type: Boolean,
      default: false
    }
  },
  setup(props, { emit }) {
    const router = useRouter()
    const route = useRoute()
    const sidebarOpen = ref(false)
    const localUser = ref(null)
    const loadStoredTheme = () => {
      try {
        const stored = typeof window !== 'undefined' ? localStorage.getItem('adminDarkMode') : null
        if (stored === 'true') return true
        if (stored === 'false') return false
      } catch {
        /* ignore storage errors */
      }
      return null
    }

    const isDarkMode = ref(loadStoredTheme() ?? props.isDark)

    const currentUser = computed(() => props.user || localUser.value)
    const userEmail = computed(() => currentUser.value?.email || 'admin@ccpp.org')
    const userInitial = computed(() => (currentUser.value?.email || 'A').charAt(0).toUpperCase())
    const gallerySections = [
      { value: 'about', label: 'About', to: { name: 'GalleryCategory', params: { category: 'about' } } },
      { value: 'outreaches', label: 'Outreaches', to: { name: 'GalleryCategory', params: { category: 'outreaches' } } },
      { value: 'youth', label: 'Youth Ministry', to: { name: 'GalleryCategory', params: { category: 'youth' } } },
      { value: 'grace-church', label: 'Grace Church', to: { name: 'GalleryCategory', params: { category: 'grace-church' } } },
      { value: 'general', label: 'General', to: { name: 'GalleryCategory', params: { category: 'general' } } }
    ]

    const linkClass = (name, extraNames = []) => {
      const active = route.name === name || extraNames.includes(route.name)
      const base = 'group flex items-center px-4 sm:px-6 py-3 text-sm font-medium border-l-4 transition-all duration-200 ease-in-out'

      if (isDarkMode.value) {
        return active
          ? `${base} bg-slate-800 text-slate-100 border-cyan-400`
          : `${base} text-slate-300 border-transparent hover:bg-slate-800 hover:text-white hover:border-cyan-400`
      }

      return active
        ? `${base} bg-white text-zinc-700 border-cyan-500`
        : `${base} text-zinc-600 border-transparent hover:bg-white hover:text-zinc-700 hover:border-cyan-500`
    }

    const subLinkClass = (value) => {
      const active = route.name === 'GalleryCategory' && route.params.category === value
      const base = 'group flex items-center px-6 sm:px-8 py-2 text-sm font-medium transition-all duration-200 ease-in-out rounded-r-md'

      if (isDarkMode.value) {
        return active
          ? `${base} bg-slate-800 text-slate-100`
          : `${base} text-slate-300 hover:bg-slate-800 hover:text-white`
      }

      return active
        ? `${base} bg-white text-zinc-700`
        : `${base} text-zinc-600 hover:bg-white hover:text-zinc-700`
    }

    const toggleSidebar = () => {
      sidebarOpen.value = !sidebarOpen.value
    }

    const closeSidebar = () => {
      sidebarOpen.value = false
    }

    const closeSidebarOnMobile = () => {
      if (window.innerWidth < 1024) {
        closeSidebar()
      }
    }

    const handleLogout = () => {
      localStorage.removeItem('authToken')
      localStorage.removeItem('auth_token')
      localStorage.removeItem('authUser')
      emit('logout')
      router.push('/admin/login')
    }

    const toggleTheme = () => {
      isDarkMode.value = !isDarkMode.value
      try {
        localStorage.setItem('adminDarkMode', isDarkMode.value ? 'true' : 'false')
      } catch {
        /* ignore storage errors */
      }
      emit('toggle-theme', isDarkMode.value)
    }

    onMounted(() => {
      const saved = localStorage.getItem('authUser')
      if (saved) {
        try {
          localUser.value = JSON.parse(saved)
        } catch {
          localUser.value = null
        }
      }
    })

    return {
      sidebarOpen,
      isDarkMode,
      userEmail,
      userInitial,
      gallerySections,
      linkClass,
      subLinkClass,
      toggleSidebar,
      closeSidebar,
      closeSidebarOnMobile,
      handleLogout,
      toggleTheme
    }
  }
}
</script>

<style scoped>
.admin-dark-mode :deep(.bg-white) {
  background-color: #0c0f14 !important;
  border-color: #0c94ab40 !important;
}

.admin-dark-mode :deep(.border-gray-200) {
  border-color: #0c94ab40 !important;
}

.admin-dark-mode {
  color: #f5f7fb;
}

.admin-dark-mode :deep(.text-gray-500),
.admin-dark-mode :deep(.text-gray-600),
.admin-dark-mode :deep(.text-gray-700),
.admin-dark-mode :deep(.text-zinc-500),
.admin-dark-mode :deep(.text-zinc-600),
.admin-dark-mode :deep(.text-zinc-700) {
  color: #cfd7e3 !important;
}

.admin-dark-mode :deep(.text-gray-900),
.admin-dark-mode :deep(.text-zinc-900),
.admin-dark-mode :deep(.font-semibold) {
  color: #f8fbff !important;
}

.admin-dark-mode :deep(a),
.admin-dark-mode :deep(button),
.admin-dark-mode :deep(.router-link-active),
.admin-dark-mode :deep(.router-link-exact-active) {
  color: #f5f7fb;
}

.admin-dark-mode :deep(.router-link-active),
.admin-dark-mode :deep(.router-link-exact-active) {
  background-color: #0c94ab1a !important;
  border-color: #0c94ab !important;
}
</style>

