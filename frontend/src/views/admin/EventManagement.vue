<template>
  <AdminLayout>
    <div class="mb-6 flex flex-wrap gap-3 items-center justify-between">
      <div>
        <h1 class="text-2xl font-semibold text-gray-900 dark:text-gray-400">Events & Calendar</h1>
        <p class="text-sm text-gray-600 dark:text-gray-500">Manage events and view them as a calendar or list.</p>
      </div>
      <div class="flex-1 flex justify-center">
        <button
          v-if="viewMode === 'list'"
          @click="openCreateModal"
          class="new-event-button px-6 py-3 bg-white border border-gray-300 text-green-600 rounded-lg hover:bg-gray-50 transition-colors shadow-lg font-medium flex items-center space-x-2"
        >
          <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4v16m8-8H4"></path>
          </svg>
          <span>New Event</span>
        </button>
      </div>
      <div class="flex gap-3 items-center">
        <!-- View Toggle -->
        <div class="view-toggle-container flex items-center gap-2 px-2 py-1 bg-white rounded-lg border border-gray-300">
          <button
            @click="viewMode = 'list'"
            :class="[
              'px-4 py-2 rounded-md text-sm font-medium transition-colors',
              viewMode === 'list'
                ? 'bg-main text-green-600 shadow-sm'
                : 'bg-white text-gray-700 hover:bg-gray-50 hover:text-gray-900'
            ]"
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
            :class="[
              'px-4 py-2 rounded-md text-sm font-medium transition-colors',
              viewMode === 'calendar'
                ? 'bg-main text-green-600 shadow-sm'
                : 'bg-white text-gray-700 hover:bg-gray-50 hover:text-gray-900'
            ]"
          >
            <span class="flex items-center gap-2">
              <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 7V3m8 4V3m-9 8h10M5 21h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v12a2 2 0 002 2z"></path>
              </svg>
              Calendar
            </span>
          </button>
        </div>
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
            'calendar-day-cell min-h-28 border rounded-lg p-3 flex flex-col gap-2 cursor-pointer transition-all hover:shadow-md',
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
            <span 
              v-if="day.events.length" 
              :class="[
                'text-[11px] font-semibold px-2 py-1 rounded-full',
                day.isPast 
                  ? 'bg-gray-200 dark:bg-transparent text-black dark:text-white' 
                  : 'bg-brand-blue/10 dark:bg-transparent text-black dark:text-white'
              ]"
            >
              {{ day.events.length }} evt
            </span>
          </div>
          <div class="space-y-1">
            <button
              v-for="event in day.events.slice(0, 3)"
              :key="event.id"
              type="button"
              :class="[
                'w-full text-left text-xs rounded-md px-2 py-1 truncate transition-colors',
                day.isPast 
                  ? 'bg-gray-200 dark:bg-transparent text-black dark:text-white hover:bg-gray-300 dark:hover:bg-transparent' 
                  : 'bg-brand-blue/10 dark:bg-transparent text-black dark:text-white hover:bg-brand-blue/20 dark:hover:bg-transparent'
              ]"
              @click.stop="openEditModal(event)"
            >
              {{ event.title }}
            </button>
            <div v-if="day.events.length > 3" class="text-xs text-black dark:text-white">+{{ day.events.length - 3 }} more</div>
          </div>
        </div>
      </div>
    </div>

    <!-- Events List -->
    <div v-else class="events-list-container bg-white rounded-lg shadow">
      <div class="p-6 border-b border-gray-200">
        <div class="flex flex-col sm:flex-row justify-between items-start sm:items-center gap-4">
          <h2 class="text-xl font-semibold text-gray-900">Events</h2>
          <div class="flex items-center space-x-2 w-full sm:w-auto">
            <input
              v-model="searchQuery"
              @input="filterEvents"
              type="text"
              placeholder="Search events..."
              class="flex-1 sm:flex-none px-4 py-2 border border-gray-300 rounded-lg text-sm min-w-48"
            />
            <select 
              v-model="filterType"
              @change="loadEvents"
              class="px-4 py-2 border border-gray-300 rounded-lg text-sm"
            >
              <option value="all">All Events</option>
              <option value="upcoming">Upcoming</option>
              <option value="past">Past</option>
            </select>
          </div>
        </div>
      </div>

      <!-- Loading State -->
      <div v-if="loading" class="text-center py-12">
        <div class="inline-block animate-spin rounded-full h-12 w-12 border-4 border-main border-t-transparent mb-4"></div>
        <p class="text-gray-600">Loading events...</p>
      </div>

      <!-- Error State -->
      <div v-else-if="error" class="text-center py-12">
        <p class="text-red-600 mb-4">{{ error }}</p>
        <button 
          @click="loadEvents"
          class="px-4 py-2 bg-main text-white rounded-lg hover:opacity-90 transition-colors"
        >
          Retry
        </button>
      </div>

      <!-- Empty State -->
      <div v-else-if="filteredEvents.length === 0" class="text-center py-12">
        <svg class="mx-auto h-12 w-12 text-gray-400 mb-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 7V3m8 4V3m-9 8h10M5 21h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v12a2 2 0 002 2z"></path>
        </svg>
        <p class="text-gray-600 mb-2">{{ searchQuery ? 'No events found' : 'No events created yet' }}</p>
        <p class="text-sm text-gray-500 mb-4">{{ searchQuery ? 'Try a different search term' : 'Create your first event to get started' }}</p>
        <button
          v-if="!searchQuery"
          @click="openCreateModal"
          class="px-6 py-2 bg-main text-white rounded-lg hover:opacity-90 transition-colors"
        >
          Create Event
        </button>
      </div>

      <!-- Events Table -->
      <div v-else class="overflow-x-auto events-table-container">
        <table class="min-w-full divide-y divide-gray-200">
          <thead class="bg-gray-50">
            <tr>
              <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">Title</th>
              <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">Date</th>
              <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">Time</th>
              <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">Location</th>
              <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">Actions</th>
            </tr>
          </thead>
          <tbody class="bg-white divide-y divide-gray-200">
            <tr 
              v-for="event in filteredEvents" 
              :key="event.id"
              @click="openEditModal(event)"
              class="hover:bg-gray-50 transition-colors cursor-pointer"
            >
              <td class="px-6 py-4 whitespace-nowrap">
                <div class="text-sm font-medium text-gray-900">{{ event.title }}</div>
                <div v-if="event.summary" class="text-sm text-gray-500 truncate max-w-xs">{{ event.summary }}</div>
              </td>
              <td class="px-6 py-4 whitespace-nowrap">
                <div class="text-sm text-gray-900">{{ formatDate(event.date || event.startTime || event.start) }}</div>
              </td>
              <td class="px-6 py-4 whitespace-nowrap">
                <div class="text-sm text-gray-900">
                  <span v-if="event.time">{{ event.time }}</span>
                  <span v-else-if="event.startTime || event.start">
                    {{ formatTime(event.startTime || event.start) }}
                    <span v-if="event.endTime || event.end"> - {{ formatTime(event.endTime || event.end) }}</span>
                  </span>
                  <span v-else>N/A</span>
                </div>
              </td>
              <td class="px-6 py-4 whitespace-nowrap">
                <div class="text-sm text-gray-900">{{ event.location || 'N/A' }}</div>
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm font-medium" @click.stop>
                <button 
                  @click="deleteEvent(event.id)"
                  class="text-red-600 hover:text-red-800"
                >
                  Delete
                </button>
              </td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>

    <!-- Create/Edit Event Modal -->
    <div
      v-if="showModal"
      class="fixed inset-0 bg-black/50 backdrop-blur-sm z-50 flex items-center justify-center p-4"
      @click="closeModal"
    >
      <div
        class="event-form-modal bg-white dark:bg-gray-800 rounded-xl shadow-2xl w-full max-w-xl max-h-[90vh] overflow-y-auto border-2 border-gray-200 dark:border-gray-700"
        @click.stop
      >
        <div class="p-8 border-b border-gray-200 dark:border-gray-700">
          <div class="flex justify-between items-start mb-6">
            <h3 class="text-3xl font-bold text-gray-900 dark:text-white">{{ editingEvent ? 'Edit Event' : 'New Event' }}</h3>
            <button @click="closeModal" class="text-gray-500 dark:text-white hover:text-gray-700 dark:hover:text-gray-200 transition-colors p-1 hover:bg-gray-100 dark:hover:bg-gray-700 rounded-full">
              <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12" />
              </svg>
            </button>
          </div>
        </div>

        <form @submit.prevent="saveEvent" class="p-8 space-y-6">
          <div>
            <label class="block text-sm font-medium text-gray-700 dark:text-white mb-2">Title *</label>
            <input
              v-model="eventForm.title"
              type="text"
              required
              class="w-full border border-gray-300 dark:border-gray-600 rounded-md px-3 py-2 text-gray-900 dark:text-white bg-white dark:bg-gray-900 focus:ring-2 focus:ring-main focus:border-main"
              placeholder="Event title"
            />
          </div>

          <div v-if="viewMode === 'calendar'" class="grid grid-cols-1 md:grid-cols-2 gap-6">
            <div>
              <label class="block text-sm font-medium text-gray-700 dark:text-white mb-2">Start *</label>
              <input
                v-model="eventForm.start"
                type="datetime-local"
                required
                class="w-full border border-gray-300 dark:border-gray-600 rounded-md px-3 py-2 text-gray-900 dark:text-white bg-white dark:bg-gray-900 focus:ring-2 focus:ring-main focus:border-main"
              />
            </div>
            <div>
              <label class="block text-sm font-medium text-gray-700 dark:text-white mb-2">End *</label>
              <input
                v-model="eventForm.end"
                type="datetime-local"
                required
                class="w-full border border-gray-300 dark:border-gray-600 rounded-md px-3 py-2 text-gray-900 dark:text-white bg-white dark:bg-gray-900 focus:ring-2 focus:ring-main focus:border-main"
              />
            </div>
          </div>

          <div v-else class="grid grid-cols-1 md:grid-cols-2 gap-6">
            <div>
              <label class="block text-sm font-medium text-gray-700 dark:text-white mb-2">Date *</label>
              <input
                v-model="eventForm.date"
                type="date"
                required
                class="w-full border border-gray-300 dark:border-gray-600 rounded-md px-3 py-2 text-gray-900 dark:text-white bg-white dark:bg-gray-900 focus:ring-2 focus:ring-main focus:border-main"
              />
            </div>
            <div>
              <label class="block text-sm font-medium text-gray-700 dark:text-white mb-2">Time</label>
              <input
                v-model="eventForm.time"
                type="time"
                class="w-full border border-gray-300 dark:border-gray-600 rounded-md px-3 py-2 text-gray-900 dark:text-white bg-white dark:bg-gray-900 focus:ring-2 focus:ring-main focus:border-main"
              />
            </div>
          </div>

          <div>
            <label class="block text-sm font-medium text-gray-700 dark:text-white mb-2">Location</label>
            <input
              v-model="eventForm.location"
              type="text"
              class="w-full border border-gray-300 dark:border-gray-600 rounded-md px-3 py-2 text-gray-900 dark:text-white bg-white dark:bg-gray-900 focus:ring-2 focus:ring-main focus:border-main"
              placeholder="Main hall"
            />
          </div>

          <div>
            <label class="block text-sm font-medium text-gray-700 dark:text-white mb-2">Summary</label>
            <input
              v-model="eventForm.summary"
              type="text"
              class="w-full border border-gray-300 dark:border-gray-600 rounded-md px-3 py-2 text-gray-900 dark:text-white bg-white dark:bg-gray-900 focus:ring-2 focus:ring-main focus:border-main"
              placeholder="Short blurb"
            />
          </div>

          <div>
            <label class="block text-sm font-medium text-gray-700 dark:text-white mb-2">Description</label>
            <textarea
              v-model="eventForm.description"
              rows="4"
              class="w-full border border-gray-300 dark:border-gray-600 rounded-md px-3 py-2 text-gray-900 dark:text-white bg-white dark:bg-gray-900 focus:ring-2 focus:ring-main focus:border-main"
              placeholder="Full details"
            ></textarea>
          </div>

          <div v-if="formError" class="p-3 border border-red-200 dark:border-red-800 bg-red-50 dark:bg-red-900/20 text-red-700 dark:text-red-400 text-sm rounded-md">
            {{ formError }}
          </div>

          <div class="flex justify-end gap-3 pt-4">
            <button type="button" @click="closeModal" class="px-4 py-2 border border-gray-300 dark:border-gray-600 rounded-md text-gray-700 dark:text-white hover:bg-gray-50 dark:hover:bg-gray-700">
              Cancel
            </button>
            <button
              type="submit"
              class="px-5 py-2 bg-main text-white rounded-md hover:opacity-90 transition-colors disabled:opacity-50"
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
import { eventService } from '@/services/eventService'

export default {
  name: 'EventManagement',
  components: {
    AdminLayout
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
      
      if (viewMode.value === 'calendar' && dateKey) {
        const start = new Date(dateKey)
        start.setHours(9, 0, 0, 0)
        const end = new Date(dateKey)
        end.setHours(10, 0, 0, 0)
        eventForm.value.start = toDateTimeLocal(start.toISOString())
        eventForm.value.end = toDateTimeLocal(end.toISOString())
      }
      
      formError.value = ''
      showModal.value = true
    }

    const openEditModal = (event) => {
      editingEvent.value = event
      
      if (viewMode.value === 'calendar') {
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
      } else {
        const eventDate = new Date(event.date || event.startTime || event.start)
        eventForm.value = {
          title: event.title,
          summary: event.summary || '',
          description: event.description || '',
          date: eventDate.toISOString().split('T')[0],
          time: event.time || '',
          location: event.location || '',
          start: '',
          end: ''
        }
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
        let eventData

        if (viewMode.value === 'calendar') {
          const start = eventForm.value.start ? new Date(eventForm.value.start) : null
          const end = eventForm.value.end ? new Date(eventForm.value.end) : null

          if (!start || !end || end < start) {
            formError.value = 'Please provide valid start and end times.'
            saving.value = false
            return
          }

          eventData = {
            title: eventForm.value.title.trim(),
            summary: eventForm.value.summary.trim(),
            description: eventForm.value.description.trim(),
            startTime: start.toISOString(),
            endTime: end.toISOString(),
            location: eventForm.value.location.trim(),
            recurrence: 'none'
          }
        } else {
          eventData = {
            title: eventForm.value.title.trim(),
            summary: eventForm.value.summary.trim(),
            description: eventForm.value.description.trim(),
            date: new Date(eventForm.value.date).toISOString(),
            time: eventForm.value.time.trim(),
            location: eventForm.value.location.trim()
          }
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

    const deleteEvent = async (id) => {
      if (!confirm('Are you sure you want to delete this event?')) {
        return
      }

      try {
        await eventService.deleteEvent(id)
        await loadEvents()
      } catch (err) {
        console.error('Error deleting event:', err)
        alert(`Failed to delete event: ${err.message}`)
      }
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

/* Allow dark mode background */
.admin-dark-mode .event-form-modal {
  background-color: #1f2937 !important;
}

/* Force black text for all modal text in light mode */
.event-form-modal label,
.event-form-modal h3,
.event-form-modal p,
.event-form-modal span,
.event-form-modal div,
.event-form-modal button:not(.bg-main) {
  color: #000000 !important;
}

/* Ensure form inputs are white in light mode (when NOT in admin-dark-mode) */
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

/* Override for dark mode only - make inputs dark */
.admin-dark-mode .event-form-modal input[type="text"],
.admin-dark-mode .event-form-modal input[type="date"],
.admin-dark-mode .event-form-modal input[type="time"],
.admin-dark-mode .event-form-modal input[type="datetime-local"],
.admin-dark-mode .event-form-modal textarea,
.admin-dark-mode .event-form-modal select {
  background-color: #000000 !important;
  color: #ffffff !important;
  border-color: #0c94ab40 !important;
}

/* Ensure all text is white in dark mode only */
.admin-dark-mode .event-form-modal label,
.admin-dark-mode .event-form-modal h3,
.admin-dark-mode .event-form-modal p,
.admin-dark-mode .event-form-modal span,
.admin-dark-mode .event-form-modal button:not(.bg-main) {
  color: #ffffff !important;
}

.admin-dark-mode .event-form-modal input::placeholder,
.admin-dark-mode .event-form-modal textarea::placeholder {
  color: #94a3b8 !important;
}

/* Force black text for calendar event buttons and text in light mode - override AdminLayout's :deep() selectors */
.calendar-container .calendar-day-cell button[type="button"] {
  color: #000000 !important;
}

.calendar-container .calendar-day-cell .space-y-1 button {
  color: #000000 !important;
}

.calendar-container .calendar-day-cell .space-y-1 div {
  color: #000000 !important;
}

.calendar-container .calendar-day-cell span[class*="rounded-full"] {
  color: #000000 !important;
}

/* Override for dark mode - make event text white */
.admin-dark-mode .calendar-container .calendar-day-cell button[type="button"] {
  color: #ffffff !important;
}

.admin-dark-mode .calendar-container .calendar-day-cell .space-y-1 button {
  color: #ffffff !important;
}

.admin-dark-mode .calendar-container .calendar-day-cell .space-y-1 div {
  color: #ffffff !important;
}

.admin-dark-mode .calendar-container .calendar-day-cell span[class*="rounded-full"] {
  color: #ffffff !important;
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
