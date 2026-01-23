<template>
  <div class="admin-root flex flex-col h-screen bg-gray-50">
    <!-- Top Bar -->
    <nav class="border-b shadow-sm bg-gray-50 border-gray-200">
      <div class="px-4 sm:px-6 flex items-center justify-between h-12">
        <div class="flex items-center">
          <button
            @click="toggleSidebar"
            class="lg:hidden mr-3 p-2 rounded-md transition-colors text-zinc-600 hover:text-zinc-900 hover:bg-gray-100"
            aria-label="Toggle menu"
          >
            <svg class="h-5 w-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 6h16M4 12h16M4 18h16" />
            </svg>
          </button>
          <div class="flex items-center">
            <div class="w-8 sm:w-10 flex flex-col items-center justify-center mr-1 mb-0.5">
              <div class="flex gap-1 sm:gap-1.5 mb-0.5">
                <div class="w-1 sm:w-1.5 h-1 sm:h-1.5 bg-[#23D3EE] rounded-full"></div>
                <div class="w-1 sm:w-1.5 h-1 sm:h-1.5 bg-[#23D3EE] rounded-full"></div>
              </div>
              <div class="w-1 sm:w-1.5 h-1 sm:h-1.5 bg-[#23D3EE] rounded-full"></div>
            </div>
            <h1 class="text-xl sm:text-2xl font-light tracking-wider leading-tight text-zinc-700">
              ANHELM
            </h1>
          </div>
        </div>
        <a
          href="/"
          target="_blank"
          class="px-3 py-1.5 rounded-md text-xs font-medium transition-colors bg-[#23D3EE] text-white hover:bg-[#1FC5D9] flex items-center gap-1.5"
        >
          <svg class="w-3.5 h-3.5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 6H6a2 2 0 00-2 2v10a2 2 0 002 2h10a2 2 0 002-2v-4M14 4h6m0 0v6m0-6L10 14"></path>
          </svg>
          View Site
        </a>
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
          'fixed lg:static inset-y-0 left-0 z-50 w-64 shadow-sm border-r flex flex-col transform transition-transform duration-300 ease-in-out bg-gray-50 border-gray-200',
          sidebarOpen ? 'translate-x-0' : '-translate-x-full lg:translate-x-0'
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
                <svg class="h-5 w-5 text-[#23D3EE]" fill="none" stroke="currentColor" viewBox="0 0 24 24">
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
                <svg class="h-5 w-5 text-[#23D3EE]" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 16l4.586-4.586a2 2 0 012.828 0L16 16l4-4m-8 4h.01" />
                </svg>
              </span>
              Galleries
            </router-link>

            <router-link
              to="/admin/events"
              @click="closeSidebarOnMobile"
              :class="linkClass('EventManagement')"
            >
              <span class="mr-3 inline-flex items-center justify-center">
                <svg class="h-5 w-5 text-green-500" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 7V3m8 4V3m-9 8h10M5 21h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v12a2 2 0 002 2z" />
                </svg>
              </span>
              Events & Calendar
            </router-link>

            <router-link
              to="/admin/videos"
              @click="closeSidebarOnMobile"
              :class="linkClass('VideoManagement')"
            >
              <span class="mr-3 inline-flex items-center justify-center">
                <svg class="h-5 w-5 text-purple-500" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 10l4.553-2.276a1 1 0 010 1.788L15 12.222M9 10L4.447 7.724a1 1 0 000 1.788L9 12.222m6 0l4.553 2.276a1 1 0 010 1.788L15 18m-6-5.778l-4.553 2.276a1 1 0 000 1.788L9 18m0-8l6 3m-6 0l6 3" />
                </svg>
              </span>
              Videos
            </router-link>

            <router-link
              to="/admin/contacts"
              @click="closeSidebarOnMobile"
              :class="linkClass('ContactMessages', ['contact-requests', 'contact-request-detail'])"
            >
              <span class="mr-3 inline-flex items-center justify-center">
                <svg class="h-5 w-5 text-[#23D3EE]" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 8l7.89 5.26a2 2 0 002.22 0L21 8m-4 8H7" />
                </svg>
              </span>
              <span class="flex-1">Contact Requests</span>
              <span
                v-if="contactRequestCount > 0"
                class="ml-2 inline-flex items-center justify-center w-6 h-6 rounded-full bg-green-500 text-white text-xs font-semibold"
              >
                {{ contactRequestCount > 99 ? '99+' : contactRequestCount }}
              </span>
            </router-link>
          </nav>
        </div>

        <!-- User Section -->
        <div class="border-t border-gray-200">
          <div class="p-4">
            <div class="flex items-center">
              <div class="flex-shrink-0">
                <div class="w-8 h-8 rounded-full flex items-center justify-center bg-[#23D3EE]/20">
                  <span class="text-sm font-medium text-[#23D3EE]">{{ userInitial }}</span>
                </div>
              </div>
              <div class="ml-3 flex-1 min-w-0">
                <p class="text-sm font-medium truncate text-zinc-700">{{ userEmail }}</p>
                <button
                  @click="handleLogout"
                  class="text-xs text-zinc-500 hover:text-zinc-700"
                >
                  Sign out
                </button>
              </div>
            </div>
          </div>
          <div class="flex items-center justify-center pt-3 pb-4 border-t px-4 border-gray-200">
            <div class="w-5 flex flex-col items-center justify-center mr-1 mb-0.5">
              <div class="flex gap-1 mb-0.5">
                <div class="w-1 h-1 bg-[#23D3EE] rounded-full"></div>
                <div class="w-1 h-1 bg-[#23D3EE] rounded-full"></div>
              </div>
              <div class="w-1 h-1 bg-[#23D3EE] rounded-full"></div>
            </div>
            <span class="text-xs font-light text-zinc-500 tracking-wider leading-tight">ANHELM</span>
          </div>
        </div>
      </nav>

      <!-- Main Content -->
      <div class="flex-1 flex flex-col overflow-hidden">
        <main class="flex-1 overflow-y-auto bg-white text-black">
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
import { ref, computed, onMounted, onUnmounted } from 'vue'
import { useRouter, useRoute } from 'vue-router'
import { contactRequestService } from '@/services/contactRequestService'

export default {
  name: 'AdminLayout',
  props: {
    user: {
      type: Object,
      default: null
    }
  },
  setup(props, { emit }) {
    const router = useRouter()
    const route = useRoute()
    const sidebarOpen = ref(false)
    const localUser = ref(null)
    const contactRequestCount = ref(0)
    let countRefreshInterval = null

    const currentUser = computed(() => props.user || localUser.value)
    const userEmail = computed(() => currentUser.value?.email || 'admin@ccpp.org')
    const userInitial = computed(() => (currentUser.value?.email || 'A').charAt(0).toUpperCase())

    const linkClass = (name, extraNames = []) => {
      const active = route.name === name || extraNames.includes(route.name)
      const base = 'group flex items-center px-4 sm:px-6 py-3 text-sm font-medium border-l-4 transition-all duration-200 ease-in-out'

      return active
        ? `${base} bg-white text-zinc-700 border-[#23D3EE]`
        : `${base} text-zinc-600 border-transparent hover:bg-white hover:text-zinc-700 hover:border-[#23D3EE]`
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

    const loadContactRequestCount = async () => {
      try {
        const response = await contactRequestService.getContactRequests()
        contactRequestCount.value = response.requests?.length || 0
      } catch (err) {
        console.error('Error loading contact request count:', err)
        contactRequestCount.value = 0
      }
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
      
      // Load contact request count and refresh every 30 seconds
      loadContactRequestCount()
      countRefreshInterval = setInterval(loadContactRequestCount, 30000)
    })

    onUnmounted(() => {
      if (countRefreshInterval) {
        clearInterval(countRefreshInterval)
      }
    })

    return {
      sidebarOpen,
      userEmail,
      userInitial,
      contactRequestCount,
      linkClass,
      toggleSidebar,
      closeSidebar,
      closeSidebarOnMobile,
      handleLogout
    }
  }
}
</script>

<style scoped>
</style>

