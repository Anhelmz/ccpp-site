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
                <div class="w-1 sm:w-1.5 h-1 sm:h-1.5 bg-[#0089AE] rounded-full"></div>
                <div class="w-1 sm:w-1.5 h-1 sm:h-1.5 bg-[#0089AE] rounded-full"></div>
              </div>
              <div class="w-1 sm:w-1.5 h-1 sm:h-1.5 bg-[#0089AE] rounded-full"></div>
            </div>
            <h1 class="text-xl sm:text-2xl font-light tracking-wider leading-tight text-zinc-700">
              ANHELM
            </h1>
          </div>
        </div>
        <a
          href="/"
          target="_blank"
          class="px-3 py-1.5 rounded-md text-xs font-medium transition-colors bg-[#0089AE] text-white hover:bg-[#007A9D] flex items-center gap-1.5"
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
                <svg class="h-5 w-5 text-[#0089AE]" fill="currentColor" viewBox="0 0 576 512">
                  <path d="M575.8 255.5c0 18-15 32.1-32 32.1h-32l.7 160.2c0 2.7-.2 5.4-.5 8.1V472c0 22.1-17.9 40-40 40H456c-1.1 0-2.2 0-3.3-.1c-1.4 .1-2.8 .1-4.2 .1H416 392c-22.1 0-40-17.9-40-40V448 384c0-17.7-14.3-32-32-32H256c-17.7 0-32 14.3-32 32v64 24c0 22.1-17.9 40-40 40H160 128.1c-1.5 0-3-.1-4.5-.2c-1.2 .1-2.4 .2-3.6 .2H104c-22.1 0-40-17.9-40-40V360c0-.9 0-1.9 .1-2.8V287.6H32c-17 0-32-14-32-32.1c0-9 3-17 10-24L266.4 8c7-7 15-8 22-8s15 2 21 7L564.8 231.5c8 7 12 15 11 24z"/>
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
                <svg class="h-5 w-5 text-[#0089AE]" fill="currentColor" viewBox="0 0 576 512">
                  <path d="M160 80H512c8.8 0 16 7.2 16 16V320c0 8.8-7.2 16-16 16H490.8L388.1 178.9c-4.5-6.6-12-10.9-20.1-10.9s-15.6 4.3-20.1 10.9l-52.2 76.8-12.8-19.2c-4.5-6.6-12-10.9-20.1-10.9s-15.6 4.3-20.1 10.9l-64.2 96.4L160 336V96c0-8.8 7.2-16 16-16zM96 96V416c0 17.7 14.3 32 32 32H512c17.7 0 32-14.3 32-32V96c0-17.7-14.3-32-32-32H128c-17.7 0-32 14.3-32 32zM48 120c0-13.3-10.7-24-24-24S0 106.7 0 120V392c0 13.3 10.7 24 24 24s24-10.7 24-24V120zM272 232a24 24 0 1 0 0-48 24 24 0 1 0 0 48z"/>
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
                <svg class="h-5 w-5 text-[#0089AE]" fill="currentColor" viewBox="0 0 448 512">
                  <path d="M128 0c17.7 0 32 14.3 32 32V64H288V32c0-17.7 14.3-32 32-32s32 14.3 32 32V64h48c26.5 0 48 21.5 48 48v48H0V112C0 85.5 21.5 64 48 64H96V32c0-17.7 14.3-32 32-32zM0 192H448V464c0 26.5-21.5 48-48 48H48c-26.5 0-48-21.5-48-48V192zm64 80v32c0 8.8 7.2 16 16 16h32c8.8 0 16-7.2 16-16V272c0-8.8-7.2-16-16-16H80c-8.8 0-16 7.2-16 16zm128 0v32c0 8.8 7.2 16 16 16h32c8.8 0 16-7.2 16-16V272c0-8.8-7.2-16-16-16H208c-8.8 0-16 7.2-16 16zm144-16c-8.8 0-16 7.2-16 16v32c0 8.8 7.2 16 16 16h32c8.8 0 16-7.2 16-16V272c0-8.8-7.2-16-16-16H336zM64 400v32c0 8.8 7.2 16 16 16h32c8.8 0 16-7.2 16-16V400c0-8.8-7.2-16-16-16H80c-8.8 0-16 7.2-16 16zm128 0v32c0 8.8 7.2 16 16 16h32c8.8 0 16-7.2 16-16V400c0-8.8-7.2-16-16-16H208c-8.8 0-16 7.2-16 16zm112-16c-8.8 0-16 7.2-16 16v32c0 8.8 7.2 16 16 16h32c8.8 0 16-7.2 16-16V400c0-8.8-7.2-16-16-16H336z"/>
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
                <svg class="h-5 w-5 text-[#0089AE]" fill="currentColor" viewBox="0 0 576 512">
                  <path d="M549.7 124.1c-6.3-23.7-24.8-42.3-48.3-48.6C458.8 64 288 64 288 64S117.2 64 74.6 75.5c-23.5 6.3-42 24.9-48.3 48.6-11.4 42.9-11.4 132.3-11.4 132.3s0 89.4 11.4 132.3c6.3 23.7 24.8 41.5 48.3 47.8C117.2 448 288 448 288 448s170.8 0 213.4-11.5c23.5-6.3 42-24.2 48.3-47.8 11.4-42.9 11.4-132.3 11.4-132.3s0-89.4-11.4-132.3zm-317.5 213.5V175.2l142.7 81.2-142.7 81.2z"/>
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
                <svg class="h-5 w-5 text-[#0089AE]" fill="currentColor" viewBox="0 0 512 512">
                  <path d="M48 64C21.5 64 0 85.5 0 112c0 15.1 7.1 29.3 19.2 38.4L236.8 313.6c11.4 8.5 27 8.5 38.4 0L492.8 150.4c12.1-9.1 19.2-23.3 19.2-38.4c0-26.5-21.5-48-48-48H48zM0 176V384c0 35.3 28.7 64 64 64H448c35.3 0 64-28.7 64-64V176L294.4 339.2c-22.8 17.1-54 17.1-76.8 0L0 176z"/>
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
                <div class="w-8 h-8 rounded-full flex items-center justify-center bg-[#0089AE]/20">
                  <span class="text-sm font-medium text-[#0089AE]">{{ userInitial }}</span>
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
                <div class="w-1 h-1 bg-[#0089AE] rounded-full"></div>
                <div class="w-1 h-1 bg-[#0089AE] rounded-full"></div>
              </div>
              <div class="w-1 h-1 bg-[#0089AE] rounded-full"></div>
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
        ? `${base} bg-white text-zinc-700 border-[#0089AE]`
        : `${base} text-zinc-600 border-transparent hover:bg-white hover:text-zinc-700 hover:border-[#0089AE]`
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

