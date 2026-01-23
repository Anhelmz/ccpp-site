<template>
  <AdminLayout>
    <!-- Page Header -->
    <div class="mb-8">
      <h1 class="text-2xl font-semibold text-zinc-900">Events & Calendar</h1>
      <p class="mt-1 text-sm text-zinc-600">Manage events and view them as a calendar or list</p>
    </div>

    <div class="mb-6 flex flex-wrap gap-3 items-center justify-end">
      <div class="flex gap-3 items-center">
        <!-- View Toggle -->
        <button
          @click="viewMode = 'list'"
          class="px-4 py-2 rounded-md text-sm font-medium transition-colors bg-[#23D3EE] text-white hover:bg-[#1FC5D9]"
        >
          <span class="flex items-center gap-2">
            <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 6h16M4 10h16M4 14h16M4 18h16"></path>
            </svg>
            List
          </span>
        </button>
        <button
          @click="viewMode = 'calendar'"
          class="px-4 py-2 rounded-md text-sm font-medium transition-colors bg-[#23D3EE] text-white hover:bg-[#1FC5D9]"
        >
          <span class="flex items-center gap-2">
            <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 7V3m8 4V3m-9 8h10M5 21h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v12a2 2 0 002 2z"></path>
            </svg>
            Calendar
          </span>
        </button>
        <button
          @click="openCreateModal()"
          type="button"
          class="px-4 py-2 rounded-md text-sm font-medium transition-colors bg-[#23D3EE] text-white hover:bg-[#1FC5D9] flex items-center gap-2"
        >
          <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4v16m8-8H4"></path>
          </svg>
          New Event
        </button>
      </div>
    </div>
    
    <!-- Calendar View -->
    <div v-if="viewMode === 'calendar'" class="calendar-container bg-white rounded-xl shadow-lg border border-gray-200 p-4 md:p-6">
      <div class="flex items-center justify-between mb-4">
        <div class="flex items-center gap-3">
          <button
            @click="previousMonth"
            class="p-2 rounded-md border border-gray-300 bg-white hover:bg-gray-50 text-gray-700 transition-colors shadow-sm"
            aria-label="Previous month"
          >
            <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 19l-7-7 7-7" />
            </svg>
          </button>
          <button
            @click="nextMonth"
            class="p-2 rounded-md border border-gray-300 bg-white hover:bg-gray-50 text-gray-700 transition-colors shadow-sm"
            aria-label="Next month"
          >
            <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7" />
            </svg>
          </button>
          <h2 class="text-xl font-semibold text-gray-900 dark:text-gray-400">{{ monthLabel }}</h2>
        </div>
      </div>

      <!-- Calendar with horizontal scroll on mobile -->
      <div class="overflow-x-auto -mx-4 md:mx-0 px-4 md:px-0">
        <div class="min-w-[700px] md:min-w-0">
          <div class="grid grid-cols-7 gap-2 text-sm font-semibold text-gray-500 dark:text-gray-500 mb-2">
            <div v-for="day in weekdayLabels" :key="day" class="text-center py-2">
              {{ day }}
            </div>
          </div>

          <div class="grid grid-cols-7 gap-2">
            <div
              v-for="day in calendarDays"
              :key="day.date"
              :class="[
                'calendar-day-cell min-h-32 md:min-h-28 border rounded-lg p-3 flex flex-col gap-2 cursor-pointer transition-all hover:shadow-md',
                day.isPast 
                  ? (day.isCurrentMonth 
                      ? 'bg-white hover:bg-gray-50 border-gray-200 text-gray-500' 
                      : 'bg-white text-gray-400 border-gray-200')
                  : (day.isCurrentMonth 
                      ? 'bg-white hover:bg-gray-50 border-gray-300 text-gray-800' 
                      : 'bg-white text-gray-400 border-gray-200'),
                day.isToday ? 'border-brand-blue ring-2 ring-brand-blue/20 bg-white' : ''
              ]"
              @click="openCreateModal(day.date)"
            >
              <div class="flex items-center justify-between">
                <span :class="[
                  'text-sm font-semibold', 
                  day.isPast 
                    ? (day.isCurrentMonth ? 'text-gray-500' : 'text-gray-400')
                    : (day.isCurrentMonth ? 'text-gray-800' : 'text-gray-400'),
                  day.isToday ? 'text-brand-blue font-bold' : ''
                ]">
                  {{ day.day }}
                </span>
              </div>
              <div class="space-y-1">
                <button
                  v-for="event in day.events.slice(0, 3)"
                  :key="event.id"
                  type="button"
                  :class="[
                    'w-full text-left text-sm md:text-base font-semibold rounded-md px-2 py-1.5 truncate transition-colors text-[#23D3EE]',
                    day.isPast 
                      ? 'bg-gray-200 dark:bg-transparent hover:bg-gray-300 dark:hover:bg-transparent opacity-70' 
                      : 'bg-brand-blue/10 dark:bg-transparent hover:bg-brand-blue/20 dark:hover:bg-transparent'
                  ]"
                  @click.stop="openEditModal(event)"
                >
                  {{ event.title }}
                </button>
                <div v-if="day.events.length > 3" class="text-xs text-black dark:text-white">+{{ day.events.length - 3 }} more</div>
                <button
                  v-if="day.events.length > 0"
                  type="button"
                  class="add-event-button w-full text-center text-xs font-medium rounded-md px-2 py-1.5 mt-1 transition-colors bg-[#23D3EE] text-white hover:bg-[#1FC5D9] hover:text-white"
                  @click.stop="openCreateModal(day.date)"
                >
                  + Add Event
                </button>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- Events List -->
    <div v-else>
      <!-- Loading State -->
      <div v-if="loading" class="flex items-center justify-center py-12">
        <div class="animate-spin rounded-full h-8 w-8 border-b-2 border-[#23D3EE]"></div>
      </div>

      <!-- Error State -->
      <div v-else-if="error" class="bg-red-50 border border-red-200 rounded-lg p-4 mb-6">
        <p class="text-sm text-red-800">{{ error }}</p>
      </div>

      <!-- Events Table -->
      <div v-else-if="filteredEvents.length > 0">
        <!-- Mobile Card View -->
        <div class="md:hidden space-y-3">
          <div
            v-for="event in filteredEvents"
            :key="event.id"
            class="bg-white border border-gray-200 rounded-lg p-4 shadow-sm hover:shadow-md transition-shadow cursor-pointer"
            @click="openEditModal(event)"
          >
            <div class="flex items-start justify-between mb-3">
              <h3 class="text-base font-semibold text-gray-900 pr-2">{{ event.title }}</h3>
              <button
                class="text-red-600 hover:text-red-800 transition-colors font-medium text-sm flex-shrink-0"
                @click.stop="deleteEvent(event.id)"
                aria-label="Delete event"
              >
                <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16"></path>
                </svg>
              </button>
            </div>
            <div class="space-y-2 text-sm">
              <div class="flex items-center gap-2 text-gray-600">
                <svg class="w-4 h-4 text-gray-400 flex-shrink-0" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 7V3m8 4V3m-9 8h10M5 21h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v12a2 2 0 002 2z"></path>
                </svg>
                <span class="font-medium text-gray-700">Date:</span>
                <span>{{ formatDate(event.date || event.startTime || event.start) }}</span>
              </div>
              <div class="flex items-center gap-2 text-gray-600">
                <svg class="w-4 h-4 text-gray-400 flex-shrink-0" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z"></path>
                </svg>
                <span class="font-medium text-gray-700">Time:</span>
                <span v-if="event.time">{{ event.time }}</span>
                <span v-else-if="event.startTime || event.start">
                  {{ formatTime(event.startTime || event.start) }}
                  <span v-if="event.endTime || event.end"> - {{ formatTime(event.endTime || event.end) }}</span>
                </span>
                <span v-else class="text-gray-400">N/A</span>
              </div>
              <div class="flex items-center gap-2 text-gray-600">
                <svg class="w-4 h-4 text-gray-400 flex-shrink-0" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17.657 16.657L13.414 20.9a1.998 1.998 0 01-2.827 0l-4.244-4.243a8 8 0 1111.314 0z"></path>
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 11a3 3 0 11-6 0 3 3 0 016 0z"></path>
                </svg>
                <span class="font-medium text-gray-700">Location:</span>
                <span>{{ event.location || 'N/A' }}</span>
              </div>
            </div>
          </div>
        </div>

        <!-- Desktop Table View -->
        <div class="hidden md:block bg-white border border-gray-200 rounded-lg overflow-hidden">
          <div class="overflow-x-auto">
            <table class="min-w-full divide-y divide-gray-200">
              <thead class="bg-gray-50">
                <tr>
                  <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">Title</th>
                  <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">Date</th>
                  <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">Time</th>
                  <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">Location</th>
                  <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">Actions</th>
                </tr>
              </thead>
              <tbody class="bg-white divide-y divide-gray-200">
                <tr 
                  v-for="event in filteredEvents" 
                  :key="event.id"
                  @click="openEditModal(event)"
                  class="hover:bg-gray-50 cursor-pointer transition-colors"
                >
                  <td class="px-6 py-4 whitespace-nowrap text-sm font-medium text-gray-900">
                    {{ event.title }}
                  </td>
                  <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-600">
                    {{ formatDate(event.date || event.startTime || event.start) }}
                  </td>
                  <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-600">
                    <span v-if="event.time">{{ event.time }}</span>
                    <span v-else-if="event.startTime || event.start">
                      {{ formatTime(event.startTime || event.start) }}
                      <span v-if="event.endTime || event.end"> - {{ formatTime(event.endTime || event.end) }}</span>
                    </span>
                    <span v-else class="text-gray-400">N/A</span>
                  </td>
                  <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-600">
                    {{ event.location || 'N/A' }}
                  </td>
                  <td class="px-6 py-4 whitespace-nowrap text-right text-sm font-medium" @click.stop>
                    <button
                      class="text-red-600 hover:text-red-800 transition-colors font-medium"
                      @click="deleteEvent(event.id)"
                    >
                      Delete
                    </button>
                  </td>
                </tr>
              </tbody>
            </table>
          </div>
        </div>
      </div>

      <!-- Empty State -->
      <div v-else class="bg-white p-6 border border-gray-200 rounded-lg">
        <div class="text-center py-12">
          <svg class="mx-auto h-16 w-16 text-zinc-300 mb-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 7V3m8 4V3m-9 8h10M5 21h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v12a2 2 0 002 2z"></path>
          </svg>
          <h3 class="text-lg font-medium text-zinc-900 mb-2">{{ searchQuery ? 'No events found' : 'No Events' }}</h3>
          <p class="text-sm text-zinc-600">{{ searchQuery ? 'Try a different search term' : 'Create your first event to get started' }}</p>
        </div>
      </div>
    </div>

    <!-- Confirmation Modal -->
    <ConfirmationModal
      :show="showConfirmModal"
      title="Delete Event"
      message="Are you sure you want to delete this event? This action cannot be undone."
      confirm-text="Delete Event"
      @confirm="confirmDelete"
      @cancel="cancelDelete"
    />

    <!-- Create/Edit Event Modal -->
    <div
      v-if="showModal"
      class="fixed inset-0 bg-black/50 backdrop-blur-sm z-[60] overflow-y-auto"
      @click="closeModal"
    >
      <div
        class="event-form-modal bg-white rounded-xl shadow-2xl w-full max-w-xl mx-auto my-8 border border-gray-200"
        @click.stop
      >
        <div class="p-4 border-b border-gray-200">
          <div class="flex justify-between items-center">
            <h3 class="text-xl font-semibold text-gray-900">{{ editingEvent ? 'Edit Event' : 'New Event' }}</h3>
            <button @click="closeModal" class="text-gray-500 hover:text-gray-700 transition-colors p-1 hover:bg-gray-100 rounded-full">
              <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12" />
              </svg>
            </button>
          </div>
        </div>

        <form @submit.prevent="saveEvent" class="p-6 space-y-4">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1">Title *</label>
            <input
              v-model="eventForm.title"
              type="text"
              required
              class="w-full border border-gray-300 rounded-md px-3 py-2 text-gray-900 bg-white focus:ring-2 focus:ring-[#23D3EE] focus:border-[#23D3EE]"
              placeholder="Event title"
            />
          </div>

          <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-1">Start *</label>
              <input
                v-model="eventForm.start"
                type="datetime-local"
                required
                class="w-full border border-gray-300 rounded-md px-3 py-2 text-gray-900 bg-white focus:ring-2 focus:ring-[#23D3EE] focus:border-[#23D3EE]"
              />
            </div>
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-1">End *</label>
              <input
                v-model="eventForm.end"
                type="datetime-local"
                required
                class="w-full border border-gray-300 rounded-md px-3 py-2 text-gray-900 bg-white focus:ring-2 focus:ring-[#23D3EE] focus:border-[#23D3EE]"
              />
            </div>
          </div>

          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1">Location</label>
            <input
              v-model="eventForm.location"
              type="text"
              class="w-full border border-gray-300 rounded-md px-3 py-2 text-gray-900 bg-white focus:ring-2 focus:ring-[#23D3EE] focus:border-[#23D3EE]"
              placeholder="Main hall"
            />
          </div>

          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1">Summary</label>
            <input
              v-model="eventForm.summary"
              type="text"
              class="w-full border border-gray-300 rounded-md px-3 py-2 text-gray-900 bg-white focus:ring-2 focus:ring-[#23D3EE] focus:border-[#23D3EE]"
              placeholder="Short blurb"
            />
          </div>

          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1">Description</label>
            <textarea
              v-model="eventForm.description"
              rows="3"
              class="w-full border border-gray-300 rounded-md px-3 py-2 text-gray-900 bg-white focus:ring-2 focus:ring-[#23D3EE] focus:border-[#23D3EE]"
              placeholder="Full details"
            ></textarea>
          </div>

          <div v-if="formError" class="p-3 border border-red-200 bg-red-50 text-red-700 text-sm rounded-md">
            {{ formError }}
          </div>

          <div class="flex justify-end gap-3 pt-2">
            <button type="button" @click="closeModal" class="px-4 py-2 border border-gray-300 rounded-md text-gray-700 hover:bg-gray-50 transition-colors">
              Cancel
            </button>
            <button
              type="submit"
              class="px-5 py-2 text-white rounded-md transition-colors disabled:opacity-50"
              :class="editingEvent ? 'bg-[#23D3EE] hover:bg-[#1FC5D9]' : 'bg-[#23D3EE] hover:bg-[#1FC5D9]'"
              :disabled="saving"
            >
              {{ saving ? 'Saving...' : editingEvent ? 'Update Event' : 'Create Event' }}
            </button>
          </div>
        </form>
      </div>
    </div>
  </AdminLayout>
</template>

<script>
import { ref, computed, onMounted } from 'vue'
import AdminLayout from '@/components/admin/AdminLayout.vue'
import ConfirmationModal from '@/components/admin/ConfirmationModal.vue'
import { eventService } from '@/services/eventService'

export default {
  name: 'EventManagement',
  components: {
    AdminLayout,
    ConfirmationModal
  },
  setup() {
    const showModal = ref(false)
    const editingEvent = ref(null)
    const events = ref([])
    const loading = ref(false)
    const error = ref('')
    const saving = ref(false)
    const formError = ref('')
    const filterType = ref('all')
    const searchQuery = ref('')
    const viewMode = ref('list') // 'list' or 'calendar'
    const showConfirmModal = ref(false)
    const eventToDelete = ref(null)

    // Calendar state
    const currentMonth = ref(new Date().getMonth())
    const currentYear = ref(new Date().getFullYear())
    const weekdayLabels = ['Sun', 'Mon', 'Tue', 'Wed', 'Thu', 'Fri', 'Sat']

    const eventForm = ref({
      title: '',
      summary: '',
      description: '',
      date: '',
      time: '',
      location: '',
      start: '',
      end: ''
    })

    const formatDate = (dateString) => {
      if (!dateString) return 'N/A'
      const date = new Date(dateString)
      return date.toLocaleDateString('en-US', {
        year: 'numeric',
        month: 'short',
        day: 'numeric'
      })
    }

    const formatTime = (dateString) => {
      if (!dateString) return ''
      const date = new Date(dateString)
      return date.toLocaleTimeString('en-US', {
        hour: 'numeric',
        minute: '2-digit',
        hour12: true
      })
    }

    const monthLabel = computed(() =>
      new Date(currentYear.value, currentMonth.value, 1).toLocaleString('default', { month: 'long', year: 'numeric' })
    )

    const toDateTimeLocal = (dateString) => {
      if (!dateString) return ''
      const date = new Date(dateString)
      const offset = date.getTimezoneOffset()
      const local = new Date(date.getTime() - offset * 60000)
      return local.toISOString().slice(0, 16)
    }

    const calendarRange = computed(() => {
      const firstDay = new Date(currentYear.value, currentMonth.value, 1)
      const lastDay = new Date(currentYear.value, currentMonth.value + 1, 0)

      const start = new Date(firstDay)
      start.setDate(start.getDate() - start.getDay())

      const end = new Date(lastDay)
      end.setDate(end.getDate() + (6 - lastDay.getDay()))

      return { start, end }
    })

    const calendarDays = computed(() => {
      const days = []
      const { start, end } = calendarRange.value
      const cursor = new Date(start)
      while (cursor <= end) {
        const dateKey = cursor.toISOString().split('T')[0]
        const dayEvents = events.value
          .map((ev) => ({
            ...ev,
            startTime: ev.startTime || ev.start || ev.date,
            endTime: ev.endTime || ev.end || ev.startTime || ev.date,
            dateKey: (ev.startTime || ev.start || ev.date || '').split('T')[0]
          }))
          .filter((ev) => ev.dateKey === dateKey)
        const today = new Date()
        today.setHours(0, 0, 0, 0)
        const dayDate = new Date(cursor)
        dayDate.setHours(0, 0, 0, 0)
        const isPast = dayDate < today
        
        days.push({
          date: dateKey,
          day: cursor.getDate(),
          isCurrentMonth: cursor.getMonth() === currentMonth.value,
          isToday: cursor.toDateString() === new Date().toDateString(),
          isPast: isPast,
          events: dayEvents
        })
        cursor.setDate(cursor.getDate() + 1)
      }
      return days
    })

    const loadEvents = async () => {
      loading.value = true
      error.value = ''
      try {
        const upcoming = filterType.value === 'upcoming'
        const response = await eventService.getEvents(upcoming)
        events.value = response.events || []
      } catch (err) {
        console.error('Error loading events:', err)
        error.value = err.message || 'Failed to load events'
      } finally {
        loading.value = false
      }
    }

    const filterEvents = () => {
      // Filtering is handled by computed property
    }

    const filteredEvents = computed(() => {
      let filtered = [...events.value]

      // Apply search filter
      if (searchQuery.value) {
        const query = searchQuery.value.toLowerCase()
        filtered = filtered.filter(event => 
          event.title.toLowerCase().includes(query) ||
          (event.summary && event.summary.toLowerCase().includes(query)) ||
          (event.description && event.description.toLowerCase().includes(query)) ||
          (event.location && event.location.toLowerCase().includes(query))
        )
      }

      // Apply past filter
      if (filterType.value === 'past') {
        const now = new Date()
        filtered = filtered.filter(event => {
          const eventDate = event.date || event.startTime || event.start
          return eventDate && new Date(eventDate) < now
        })
      }

      return filtered.sort((a, b) => {
        const dateA = a.date || a.startTime || a.start || ''
        const dateB = b.date || b.startTime || b.start || ''
        return new Date(dateA) - new Date(dateB)
      })
    })

    const openCreateModal = (dateKey) => {
      editingEvent.value = null
      eventForm.value = {
        title: '',
        summary: '',
        description: '',
        date: '',
        time: '',
        location: '',
        start: '',
        end: ''
      }
      
      if (dateKey) {
        const start = new Date(dateKey)
        start.setHours(9, 0, 0, 0)
        const end = new Date(dateKey)
        end.setHours(10, 0, 0, 0)
        eventForm.value.start = toDateTimeLocal(start.toISOString())
        eventForm.value.end = toDateTimeLocal(end.toISOString())
      } else {
        // If no dateKey provided, use today's date
        const today = new Date()
        today.setHours(9, 0, 0, 0)
        const end = new Date(today)
        end.setHours(10, 0, 0, 0)
        eventForm.value.start = toDateTimeLocal(today.toISOString())
        eventForm.value.end = toDateTimeLocal(end.toISOString())
      }
      
      formError.value = ''
      showModal.value = true
    }

    const openEditModal = (event) => {
      editingEvent.value = event
      
      eventForm.value = {
        title: event.title,
        start: toDateTimeLocal(event.startTime || event.start || event.date),
        end: toDateTimeLocal(event.endTime || event.end || event.startTime || event.date),
        location: event.location || '',
        summary: event.summary || '',
        description: event.description || '',
        date: '',
        time: ''
      }
      
      formError.value = ''
      showModal.value = true
    }

    const closeModal = () => {
      showModal.value = false
      editingEvent.value = null
      formError.value = ''
    }

    const saveEvent = async () => {
      saving.value = true
      formError.value = ''

      try {
        const start = eventForm.value.start ? new Date(eventForm.value.start) : null
        const end = eventForm.value.end ? new Date(eventForm.value.end) : null

        if (!start || !end || end < start) {
          formError.value = 'Please provide valid start and end times.'
          saving.value = false
          return
        }

        const eventData = {
          title: eventForm.value.title.trim(),
          summary: eventForm.value.summary.trim(),
          description: eventForm.value.description.trim(),
          startTime: start.toISOString(),
          endTime: end.toISOString(),
          location: eventForm.value.location.trim(),
          recurrence: 'none'
        }

        if (editingEvent.value) {
          await eventService.updateEvent(editingEvent.value.id, eventData)
        } else {
          await eventService.createEvent(eventData)
        }

        await loadEvents()
        closeModal()
      } catch (err) {
        console.error('Error saving event:', err)
        formError.value = err.message || 'Failed to save event'
      } finally {
        saving.value = false
      }
    }

    const deleteEvent = (id) => {
      eventToDelete.value = id
      showConfirmModal.value = true
    }

    const confirmDelete = async () => {
      if (!eventToDelete.value) return

      const id = eventToDelete.value
      showConfirmModal.value = false

      try {
        await eventService.deleteEvent(id)
        await loadEvents()
        eventToDelete.value = null
      } catch (err) {
        console.error('Error deleting event:', err)
        alert(`Failed to delete event: ${err.message}`)
        eventToDelete.value = null
      }
    }

    const cancelDelete = () => {
      showConfirmModal.value = false
      eventToDelete.value = null
    }

    const previousMonth = () => {
      if (currentMonth.value === 0) {
        currentMonth.value = 11
        currentYear.value--
      } else {
        currentMonth.value--
      }
    }

    const nextMonth = () => {
      if (currentMonth.value === 11) {
        currentMonth.value = 0
        currentYear.value++
      } else {
        currentMonth.value++
      }
    }

    const resetToToday = () => {
      const now = new Date()
      currentMonth.value = now.getMonth()
      currentYear.value = now.getFullYear()
    }

    onMounted(() => {
      loadEvents()
    })

    return {
      showModal,
      editingEvent,
      events,
      loading,
      error,
      saving,
      formError,
      filterType,
      searchQuery,
      eventForm,
      filteredEvents,
      formatDate,
      formatTime,
      loadEvents,
      filterEvents,
      openCreateModal,
      openEditModal,
      closeModal,
      saveEvent,
      deleteEvent,
      showConfirmModal,
      eventToDelete,
      confirmDelete,
      cancelDelete,
      viewMode,
      weekdayLabels,
      monthLabel,
      calendarDays,
      previousMonth,
      nextMonth,
      resetToToday
    }
  }
}
</script>

<style>
/* Force white background for modal in light mode */
.event-form-modal {
  background-color: white !important;
}

/* Force black text for all modal text */
.event-form-modal label,
.event-form-modal h3,
.event-form-modal p,
.event-form-modal span,
.event-form-modal div,
.event-form-modal button:not(.bg-main):not(.text-white) {
  color: #000000 !important;
}

/* Ensure form inputs are white */
.event-form-modal input[type="text"],
.event-form-modal input[type="date"],
.event-form-modal input[type="time"],
.event-form-modal input[type="datetime-local"],
.event-form-modal textarea,
.event-form-modal select {
  background-color: white !important;
  color: #111827 !important;
  border-color: #d1d5db !important;
}

/* Force black text for calendar event buttons and text */
.calendar-container .calendar-day-cell button[type="button"] {
  color: #000000 !important;
}

.calendar-container .calendar-day-cell .space-y-1 button {
  color: #000000 !important;
}

.calendar-container .calendar-day-cell .space-y-1 .add-event-button {
  color: #ffffff !important;
}

.calendar-container .calendar-day-cell .space-y-1 div {
  color: #000000 !important;
}

.calendar-container .calendar-day-cell span[class*="rounded-full"] {
  color: #000000 !important;
}

/* Force white background for view toggle container in light mode only */
.view-toggle-container {
  background-color: white !important;
  border-color: #d1d5db !important;
}

.view-toggle-container button:not(.bg-main) {
  background-color: white !important;
  color: #374151 !important;
}

.view-toggle-container button:not(.bg-main):hover {
  background-color: #f9fafb !important;
  color: #111827 !important;
}

/* Ensure selected toggle button text is green */
.view-toggle-container button.bg-main {
  color: #16a34a !important;
}

.view-toggle-container button.bg-main svg {
  color: #16a34a !important;
}

/* Force white background for New Event button in light mode only */
.new-event-button {
  background-color: white !important;
  border-color: #d1d5db !important;
  color: #16a34a !important;
}

.new-event-button:hover {
  background-color: #f9fafb !important;
  color: #16a34a !important;
}

.new-event-button svg {
  color: #16a34a !important;
}
</style>
