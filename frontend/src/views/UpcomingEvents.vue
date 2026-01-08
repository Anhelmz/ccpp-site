<template>
  <div>
    <!-- Hero Section -->
    <section class="relative min-h-[55vh] flex items-center overflow-hidden pt-20 md:pt-24 pb-16 bg-gradient-to-br from-[#0a1f33] via-[#162c46] to-[#1e3d5d] text-white">
      <div class="absolute inset-0" aria-hidden="true">
        <div class="absolute inset-0 opacity-25 mix-blend-screen" style="background-image: radial-gradient(circle at 15% 25%, rgba(255,255,255,0.25), transparent 45%), radial-gradient(circle at 80% 12%, rgba(255,255,255,0.18), transparent 35%), radial-gradient(circle at 60% 80%, rgba(255,255,255,0.2), transparent 30%);"></div>
        <div class="absolute inset-y-[-30%] left-[-8%] w-1/2 bg-gradient-to-r from-brand-orange/25 via-transparent to-transparent blur-[150px] opacity-60"></div>
        <div class="absolute inset-y-[-40%] right-[-15%] w-2/3 bg-gradient-to-l from-brand-blue/35 via-transparent to-transparent blur-[200px] opacity-65"></div>
        <div class="absolute inset-x-0 bottom-0 h-40 bg-gradient-to-t from-black/60 via-black/30 to-transparent"></div>
      </div>
      <div class="relative z-10 w-full max-w-4xl mx-auto px-4 sm:px-6 lg:px-8 text-center space-y-4">
        <h1 class="text-4xl md:text-5xl lg:text-6xl font-bold leading-tight">Upcoming Events</h1>
        <p class="text-lg md:text-xl text-white/85 max-w-3xl mx-auto font-light">
          Prayer nights, serve days, retreats, and gatherings that keep our church present in Phnom Penh all year.
        </p>
        <div class="flex flex-wrap justify-center gap-3">
          <span class="inline-flex items-center px-4 py-2 rounded-full border border-white/25 text-white/80 text-xs uppercase tracking-[0.3em] backdrop-blur-sm">
            Worship moments
          </span>
          <span class="inline-flex items-center px-4 py-2 rounded-full border border-white/25 text-white/80 text-xs uppercase tracking-[0.3em] backdrop-blur-sm">
            Serve Phnom Penh
          </span>
          <span class="inline-flex items-center px-4 py-2 rounded-full border border-white/25 text-white/80 text-xs uppercase tracking-[0.3em] backdrop-blur-sm">
            Youth retreats
          </span>
          <span class="inline-flex items-center px-4 py-2 rounded-full border border-white/25 text-white/80 text-xs uppercase tracking-[0.3em] backdrop-blur-sm">
            Prayer & fasting
          </span>
        </div>
      </div>
    </section>

    <!-- Content Section -->
    <section class="bg-primary py-12 md:py-16">
      <div class="max-w-6xl mx-auto px-4 sm:px-6 lg:px-8 space-y-10">
        <!-- Loading State -->
        <div v-if="loading" class="py-12">
          <div class="text-center">
            <div class="inline-block animate-spin rounded-full h-12 w-12 border-b-2 border-main mb-4"></div>
            <p class="text-gray-600">Loading events...</p>
          </div>
        </div>

        <!-- Error State -->
        <div v-else-if="error" class="py-12 text-center">
          <p class="text-red-600">{{ error }}</p>
        </div>

        <!-- Events List -->
        <div v-else>
          <div v-if="events.length > 0" class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
            <div 
              v-for="event in events" 
              :key="event.id"
              class="relative rounded-[32px] border border-white/15 bg-white backdrop-blur-sm overflow-hidden shadow-2xl flex flex-col cursor-pointer group"
              @click="openModal(event)"
            >
              <div class="relative h-44 bg-gradient-to-br from-brand-blue/90 to-brand-blue flex items-center justify-center">
                <div class="text-center text-white space-y-1">
                  <p class="text-sm uppercase tracking-[0.35em] text-white/80">{{ formatMonth(event.startTime) }}</p>
                  <p class="text-5xl font-bold leading-none">{{ formatDay(event.startTime) }}</p>
                  <p class="text-xs uppercase tracking-[0.35em] text-white/70">{{ event.startTime.getFullYear() }}</p>
                </div>
                <div class="absolute inset-0 bg-gradient-to-t from-black/30 via-transparent to-transparent"></div>
              </div>
              <div class="p-6 flex flex-col gap-4 flex-1">
                <div>
                  <h2 class="text-xl font-semibold text-gray-900 mb-2 group-hover:text-brand-orange transition-colors">
                    {{ event.title }}
                  </h2>
                  <p class="text-gray-600 text-sm leading-relaxed line-clamp-3">{{ event.summary }}</p>
                </div>
                <div class="flex flex-wrap gap-2 text-xs text-gray-600">
                  <span class="inline-flex items-center gap-1 px-3 py-1 rounded-full bg-brand-blue/10 text-brand-blue font-medium">
                    <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 7V3m8 4V3m-9 8h10M5 21h14a2 2 0 002-2V7a 2 2 0 00-2-2H5a 2 2 0 00-2 2v12a2 2 0 002 2z" />
                    </svg>
                    {{ formatFullDate(event.startTime) }}
                  </span>
                  <span v-if="event.timeRange" class="inline-flex items-center gap-1 px-3 py-1 rounded-full bg-brand-orange/10 text-brand-orange font-medium">
                    <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z" />
                    </svg>
                    {{ event.timeRange }}
                  </span>
                  <span v-if="event.location" class="inline-flex items-center gap-1 px-3 py-1 rounded-full bg-gray-100 text-gray-700 font-medium">
                    <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17.657 16.657L13.414 20.9a1.998 1.998 0 01-2.827 0l-4.244-4.243a8 8 0 1111.314 0z" />
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 11a3 3 0 11-6 0 3 3 0 016 0z" />
                    </svg>
                    {{ event.location }}
                  </span>
                </div>
              </div>
            </div>
          </div>

          <!-- Empty State -->
          <div v-else class="card text-center py-16">
            <div class="text-6xl mb-6">📅</div>
            <h3 class="text-3xl font-bold text-gray-900 mb-3">No Upcoming Events</h3>
            <p class="text-gray-600 text-lg">Check back soon for upcoming church events and activities.</p>
          </div>
        </div>
      </div>
    </section>

    <!-- Event Modal -->
    <div
      v-if="selectedEvent"
      @click="closeModal"
      class="fixed inset-0 bg-black bg-opacity-60 z-50 flex items-center justify-center p-4 backdrop-blur-sm"
    >
      <div 
        @click.stop
        class="bg-white rounded-2xl max-w-4xl w-full max-h-[90vh] overflow-hidden shadow-2xl transform transition-all"
      >
        <!-- Modal Header -->
        <div class="sticky top-0 bg-gradient-to-r from-main to-secondary px-8 py-6 flex justify-between items-start z-10">
          <div class="flex-1">
            <h2 class="text-3xl font-bold text-white mb-4 pr-8">{{ selectedEvent.title }}</h2>
            <div class="flex flex-wrap items-center gap-4 text-sm">
              <div class="flex items-center text-white/90">
                <svg class="w-5 h-5 mr-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 7V3m8 4V3m-9 8h10M5 21h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v12a2 2 0 002 2z" />
                </svg>
                <span class="font-medium">{{ formatFullDate(selectedEvent.startTime) }}</span>
              </div>
              <div v-if="selectedEvent.timeRange" class="flex items-center text-white/90">
                <svg class="w-5 h-5 mr-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z" />
                </svg>
                <span class="font-medium">{{ selectedEvent.timeRange }}</span>
              </div>
              <div v-if="selectedEvent.location" class="flex items-center text-white/90">
                <svg class="w-5 h-5 mr-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17.657 16.657L13.414 20.9a1.998 1.998 0 01-2.827 0l-4.244-4.243a8 8 0 1111.314 0z" />
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 11a3 3 0 11-6 0 3 3 0 016 0z" />
                </svg>
                <span class="font-medium">{{ selectedEvent.location }}</span>
              </div>
            </div>
          </div>
          <button
            @click="closeModal"
            class="text-white hover:text-gray-200 transition-colors p-2 hover:bg-white/10 rounded-lg"
          >
            <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"></path>
            </svg>
          </button>
        </div>
        
        <!-- Modal Content -->
        <div class="px-8 py-8 overflow-y-auto max-h-[calc(90vh-180px)]">
          <div class="prose prose-lg max-w-none">
            <p class="text-gray-700 whitespace-pre-line leading-relaxed text-lg">{{ selectedEvent.content }}</p>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { ref, onMounted } from 'vue'
import { eventService } from '@/services/eventService'

export default {
  name: 'UpcomingEvents',
  setup() {
    const loading = ref(true)
    const error = ref('')
    const events = ref([])
    const selectedEvent = ref(null)


    const formatDay = (date) => {
      return date.getDate()
    }

    const formatMonth = (date) => {
      const months = ['Jan', 'Feb', 'Mar', 'Apr', 'May', 'Jun', 'Jul', 'Aug', 'Sep', 'Oct', 'Nov', 'Dec']
      return months[date.getMonth()]
    }

    const formatFullDate = (date) => {
      const months = ['January', 'February', 'March', 'April', 'May', 'June', 'July', 'August', 'September', 'October', 'November', 'December']
      return `${months[date.getMonth()]} ${date.getDate()}, ${date.getFullYear()}`
    }

    const formatTimeRange = (start, end) => {
      if (!start) return ''
      const opts = { hour: '2-digit', minute: '2-digit' }
      const startStr = start.toLocaleTimeString([], opts)
      const endStr = end ? end.toLocaleTimeString([], opts) : ''
      return endStr && endStr !== startStr ? `${startStr} - ${endStr}` : startStr
    }

    const fetchEvents = async () => {
      loading.value = true
      error.value = ''
      
      try {
        const response = await eventService.getEvents(true) // Get upcoming events
        const apiEvents = response.events || []
        
        // Convert API events to the format expected by the component
        events.value = apiEvents.map(event => {
          const start = new Date(event.startTime || event.date)
          const end = event.endTime ? new Date(event.endTime) : start
          return {
            id: event.id,
            title: event.title,
            summary: event.summary || event.description || '',
            content: event.description || event.summary || '',
            description: event.description || '',
            startTime: start,
            endTime: end,
            timeRange: formatTimeRange(start, end),
            location: event.location || ''
          }
        }).sort((a, b) => a.startTime - b.startTime)
        
        loading.value = false
      } catch (err) {
        console.error('Error loading events:', err)
        error.value = err.message || 'Failed to load events'
        loading.value = false
      }
    }

    const openModal = (event) => {
      selectedEvent.value = event
    }

    const closeModal = () => {
      selectedEvent.value = null
    }

    onMounted(() => {
      fetchEvents()
    })

    return {
      loading,
      error,
      events,
      selectedEvent,
      formatDay,
      formatMonth,
      formatFullDate,
      openModal,
      closeModal
    }
  }
}
</script>

