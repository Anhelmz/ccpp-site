<template>
  <AdminLayout>
    <div class="mb-6 flex flex-wrap gap-3 items-center justify-between">
      <div>
        <h1 class="text-2xl font-semibold text-gray-900 dark:text-gray-400">Calendar</h1>
        <p class="text-sm text-gray-600 dark:text-gray-500">Create events that show on the public calendar page.</p>
      </div>
      <div class="flex gap-3">
        <button
          @click="resetToToday"
          class="px-4 py-2 rounded-md border border-gray-300 dark:border-gray-700 text-gray-700 dark:text-gray-400 bg-white dark:bg-gray-800 hover:bg-gray-50 dark:hover:bg-gray-800 transition-colors font-medium shadow-sm"
        >
          Today
        </button>
        <button
          v-if="oldEventsCount > 0"
          @click="confirmDeleteOldEvents"
          class="px-4 py-2 rounded-md border border-orange-300 dark:border-orange-700 text-orange-700 dark:text-orange-400 bg-white dark:bg-gray-800 hover:bg-orange-50 dark:hover:bg-orange-900/20 transition-colors font-medium shadow-sm"
        >
          Delete Old Events ({{ oldEventsCount }})
        </button>
        <button
          v-if="events.length > 0"
          @click="confirmDeleteAll"
          class="px-4 py-2 rounded-md border border-red-300 dark:border-red-700 text-red-700 dark:text-red-400 bg-white dark:bg-gray-800 hover:bg-red-50 dark:hover:bg-red-900/20 transition-colors font-medium shadow-sm"
        >
          Delete All Events
        </button>
        <button
          @click="openCreateModal()"
          class="px-6 py-3 bg-main text-white rounded-lg hover:opacity-90 transition-colors shadow-lg font-medium flex items-center space-x-2"
        >
          <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4v16m8-8H4" />
          </svg>
          <span>New Event</span>
        </button>
      </div>
    </div>

    <div class="bg-white dark:bg-black rounded-xl shadow-lg border border-gray-200 dark:border-gray-700 p-4 md:p-6">
      <div class="flex items-center justify-between mb-4">
        <div class="flex items-center gap-3">
          <button
            @click="previousMonth"
            class="p-2 rounded-md border border-gray-300 dark:border-gray-700 bg-white dark:bg-gray-800 hover:bg-gray-50 dark:hover:bg-gray-800 text-gray-700 dark:text-gray-400 transition-colors shadow-sm"
            aria-label="Previous month"
          >
            <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 19l-7-7 7-7" />
            </svg>
          </button>
          <button
            @click="nextMonth"
            class="p-2 rounded-md border border-gray-300 dark:border-gray-700 bg-white dark:bg-gray-800 hover:bg-gray-50 dark:hover:bg-gray-800 text-gray-700 dark:text-gray-400 transition-colors shadow-sm"
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
            'min-h-28 border rounded-lg p-3 flex flex-col gap-2 cursor-pointer transition-all hover:shadow-md',
            day.isPast 
              ? (day.isCurrentMonth 
                  ? 'bg-white dark:bg-gray-800 hover:bg-gray-50 dark:hover:bg-gray-700 border-gray-200 dark:border-gray-600 text-gray-500 dark:text-gray-500' 
                  : 'bg-gray-50 dark:bg-gray-900 text-gray-400 dark:text-gray-600 border-gray-200 dark:border-gray-700')
              : (day.isCurrentMonth 
                  ? 'bg-white dark:bg-black hover:bg-gray-50 dark:hover:bg-gray-900 border-gray-300 dark:border-gray-700 text-gray-800 dark:text-gray-400' 
                  : 'bg-gray-50 dark:bg-black text-gray-400 dark:text-gray-600 border-gray-200 dark:border-gray-800'),
            day.isToday ? 'border-brand-blue dark:border-main ring-2 ring-brand-blue/20 dark:ring-main/30 bg-blue-50 dark:bg-main/10' : ''
          ]"
          @click="openCreateModal(day.date)"
        >
          <div class="flex items-center justify-between">
            <span :class="[
              'text-sm font-semibold', 
              day.isPast 
                ? (day.isCurrentMonth ? 'text-gray-500 dark:text-gray-500' : 'text-gray-400 dark:text-gray-600')
                : (day.isCurrentMonth ? 'text-gray-800 dark:text-gray-400' : 'text-gray-400 dark:text-gray-600'),
              day.isToday ? 'text-brand-blue dark:text-main font-bold' : ''
            ]">
              {{ day.day }}
            </span>
            <span 
              v-if="day.events.length" 
              :class="[
                'text-[11px] font-semibold px-2 py-1 rounded-full',
                day.isPast 
                  ? 'bg-gray-200 dark:bg-gray-700 text-gray-600 dark:text-gray-400' 
                  : 'bg-brand-blue/10 dark:bg-main/20 text-brand-blue dark:text-main'
              ]"
            >
              {{ day.events.length }} evt
            </span>
          </div>
          <div class="space-y-1">
            <div
              v-for="event in day.events.slice(0, 3)"
              :key="event.id"
              class="flex items-center gap-1 group"
            >
              <button
                type="button"
                :class="[
                  'flex-1 text-left text-xs rounded-md px-2 py-1 truncate transition-colors',
                  day.isPast 
                    ? 'bg-gray-200 dark:bg-gray-700 text-gray-600 dark:text-gray-400 hover:bg-gray-300 dark:hover:bg-gray-600' 
                    : 'bg-brand-blue/10 dark:bg-main/20 text-brand-blue dark:text-main hover:bg-brand-blue/20 dark:hover:bg-main/30'
                ]"
                @click.stop="openEditModal(event)"
              >
                {{ event.title }}
              </button>
              <button
                type="button"
                @click.stop="confirmDeleteEvent(event.id)"
                class="opacity-0 group-hover:opacity-100 transition-opacity text-red-600 hover:text-red-800 p-1"
                aria-label="Delete event"
              >
                <svg class="w-3 h-3" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12" />
                </svg>
              </button>
            </div>
            <div v-if="day.events.length > 3" :class="['text-xs', day.isPast ? 'text-gray-500 dark:text-gray-600' : 'text-gray-500 dark:text-gray-500']">+{{ day.events.length - 3 }} more</div>
          </div>
        </div>
      </div>
    </div>

    <div
      v-if="showModal"
      class="fixed inset-0 bg-black/50 backdrop-blur-sm z-50 flex items-center justify-center p-4"
      @click="closeModal"
    >
      <div
        class="bg-white rounded-xl shadow-2xl w-full max-w-xl max-h-[90vh] overflow-y-auto border-2 border-brand-blue"
        @click.stop
      >
        <div class="p-5 border-b border-gray-200 flex items-center justify-between">
          <h3 class="text-xl font-semibold text-black">{{ editingEvent ? 'Edit Event' : 'New Event' }}</h3>
          <button @click="closeModal" class="text-black hover:text-gray-800">
            <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12" />
            </svg>
          </button>
        </div>

        <form @submit.prevent="saveEvent" class="p-5 space-y-4">
          <div>
            <label class="block text-sm font-medium text-black mb-1">Title *</label>
            <input
              v-model="form.title"
              type="text"
              required
              class="w-full border border-gray-300 rounded-md px-3 py-2 text-black focus:ring-2 focus:ring-main focus:border-main"
              placeholder="Event title"
            />
          </div>

          <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
            <div>
              <label class="block text-sm font-medium text-black mb-1">Start *</label>
              <input
                v-model="form.start"
                type="datetime-local"
                required
                class="w-full border border-gray-300 rounded-md px-3 py-2 text-black focus:ring-2 focus:ring-main focus:border-main"
              />
            </div>
            <div>
              <label class="block text-sm font-medium text-black mb-1">End *</label>
              <input
                v-model="form.end"
                type="datetime-local"
                required
                class="w-full border border-gray-300 rounded-md px-3 py-2 text-black focus:ring-2 focus:ring-main focus:border-main"
              />
            </div>
          </div>

          <div>
            <label class="block text-sm font-medium text-black mb-1">Location</label>
            <input
              v-model="form.location"
              type="text"
              class="w-full border border-gray-300 rounded-md px-3 py-2 text-black focus:ring-2 focus:ring-main focus:border-main"
              placeholder="Main hall"
            />
          </div>

          <div>
            <label class="block text-sm font-medium text-black mb-1">Description</label>
            <textarea
              v-model="form.description"
              rows="4"
              class="w-full border border-gray-300 rounded-md px-3 py-2 text-black focus:ring-2 focus:ring-main focus:border-main"
              placeholder="Full details"
            ></textarea>
          </div>

          <div v-if="formError" class="p-3 border border-red-200 bg-red-50 text-red-700 text-sm rounded-md">
            {{ formError }}
          </div>

          <div class="flex justify-between gap-3 pt-3">
            <button
              v-if="editingEvent"
              type="button"
              @click="confirmDeleteEvent(editingEvent.id)"
              class="px-4 py-2 bg-red-600 text-white rounded-md hover:bg-red-700 transition-colors"
            >
              Delete
            </button>
            <div class="flex gap-3 ml-auto">
              <button type="button" @click="closeModal" class="px-4 py-2 border border-gray-300 rounded-md text-black hover:bg-gray-50">
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
          </div>
        </form>
      </div>
    </div>

    <!-- Confirmation Modals -->
    <ConfirmationModal
      v-if="showDeleteModal"
      :show="showDeleteModal"
      title="Delete Event"
      message="Are you sure you want to delete this event? This action cannot be undone."
      confirm-text="Delete Event"
      @confirm="deleteEvent"
      @cancel="cancelDelete"
    />

    <ConfirmationModal
      v-if="showDeleteAllModal"
      :show="showDeleteAllModal"
      title="Delete All Events"
      :message="`Are you sure you want to delete ALL ${events.length} events? This action cannot be undone.`"
      confirm-text="Delete All"
      @confirm="deleteAllEvents"
      @cancel="cancelDeleteAll"
    />

    <ConfirmationModal
      v-if="showDeleteOldModal"
      :show="showDeleteOldModal"
      title="Delete Old Events"
      :message="`Are you sure you want to permanently delete ${oldEventsCount} old (soft-deleted) event(s)? This action cannot be undone.`"
      confirm-text="Delete Old Events"
      @confirm="deleteOldEvents"
      @cancel="cancelDeleteOld"
    />
  </AdminLayout>
</template>

<script>
import { ref, computed, onMounted } from 'vue'
import AdminLayout from '@/components/admin/AdminLayout.vue'
import ConfirmationModal from '@/components/admin/ConfirmationModal.vue'
import { eventService } from '@/services/eventService'

export default {
  name: 'AdminCalendar',
  components: {
    AdminLayout,
    ConfirmationModal
  },
  setup() {
    const events = ref([])
    const showModal = ref(false)
    const editingEvent = ref(null)
    const formError = ref('')
    const saving = ref(false)
    const showDeleteModal = ref(false)
    const showDeleteAllModal = ref(false)
    const showDeleteOldModal = ref(false)
    const eventToDelete = ref(null)
    const deleting = ref(false)
    const oldEventsCount = ref(0)

    // Use UTC for month/year to ensure consistency
    const currentMonth = ref(new Date().getUTCMonth())
    const currentYear = ref(new Date().getUTCFullYear())

    const form = ref({
      title: '',
      start: '',
      end: '',
      location: '',
      description: ''
    })

    const weekdayLabels = ['Sun', 'Mon', 'Tue', 'Wed', 'Thu', 'Fri', 'Sat']
    const monthLabel = computed(() => {
      // Use UTC date for month label to ensure consistency
      const months = ['January', 'February', 'March', 'April', 'May', 'June', 'July', 'August', 'September', 'October', 'November', 'December']
      return `${months[currentMonth.value]} ${currentYear.value}`
    })

    const normalizeEvents = (apiEvents) => {
      // Normalize events - CRITICAL: Extract date keys from raw API strings BEFORE any Date parsing
      // This ensures no timezone conversion happens
      return apiEvents.map((event) => {
        // Get the raw string values from API response (these are ISO 8601 UTC strings)
        // IMPORTANT: Use the raw string value, not a converted Date object
        const startTimeStr = typeof event.startTime === 'string' 
          ? event.startTime 
          : (event.startTime ? event.startTime.toISOString() : (event.date || ''))
        const endTimeStr = typeof event.endTime === 'string'
          ? event.endTime
          : (event.endTime ? event.endTime.toISOString() : startTimeStr)
        
        // Parse as Date objects ONLY for time operations (display, comparison)
        // But for date key extraction, we ALWAYS use the raw string
        const start = new Date(startTimeStr)
        const end = endTimeStr ? new Date(endTimeStr) : start
        
        return {
          ...event,
          startTime: start,
          endTime: end,
          // CRITICAL: Keep original raw strings for date key extraction
          // These are the UTC ISO strings from the API, e.g. "2026-01-30T09:00:00Z"
          _startTimeStr: startTimeStr,
          _endTimeStr: endTimeStr
        }
      })
    }

    const loadEvents = async () => {
      try {
        const response = await eventService.getEvents(false)
        events.value = normalizeEvents(response.events || [])
      } catch (err) {
        console.error('Failed to load events', err)
      }
    }

    const loadOldEventsCount = async () => {
      try {
        const response = await eventService.getOldEvents()
        oldEventsCount.value = response.count || 0
      } catch (err) {
        console.error('Failed to load old events count', err)
        oldEventsCount.value = 0
      }
    }

    const toDateTimeLocal = (dateString) => {
      if (!dateString) return ''
      
      // If it's already in datetime-local format (YYYY-MM-DDTHH:mm), return as-is
      if (/^\d{4}-\d{2}-\d{2}T\d{2}:\d{2}$/.test(dateString)) {
        return dateString
      }
      
      // Parse the ISO string (e.g., "2026-01-30T09:00:00Z")
      // Extract date and time parts directly from the string to avoid timezone conversion
      if (dateString.includes('T')) {
        const [datePart, timePart] = dateString.split('T')
        // If it has timezone info (Z or +/-), extract just the time part before that
        const timeOnly = timePart.split(/[Z+-]/)[0] // Gets "09:00:00" or "09:00"
        const time = timeOnly.substring(0, 5) // Gets "09:00" (HH:mm)
        return `${datePart}T${time}`
      }
      
      // Fallback: parse as Date and extract UTC components
      const date = new Date(dateString)
      const year = date.getUTCFullYear()
      const month = String(date.getUTCMonth() + 1).padStart(2, '0')
      const day = String(date.getUTCDate()).padStart(2, '0')
      const hours = String(date.getUTCHours()).padStart(2, '0')
      const minutes = String(date.getUTCMinutes()).padStart(2, '0')
      return `${year}-${month}-${day}T${hours}:${minutes}`
    }

    const calendarRange = computed(() => {
      // Build calendar range using pure UTC date math - no Date object manipulation
      // Calculate first and last day of month
      const year = currentYear.value
      const month = currentMonth.value
      
      // First day of month: year, month, 1
      // Last day of month: year, month+1, 0 (day 0 of next month = last day of current month)
      const firstDayOfMonth = { year, month, day: 1 }
      const lastDayOfMonth = { year, month, day: new Date(Date.UTC(year, month + 1, 0)).getUTCDate() }
      
      // Calculate what day of week the first day falls on (0=Sunday, 6=Saturday)
      const firstDayDate = new Date(Date.UTC(year, month, 1))
      const firstDayOfWeek = firstDayDate.getUTCDay()
      
      // Start date is first day of month minus days to get to Sunday
      let startYear = year
      let startMonth = month
      let startDay = 1 - firstDayOfWeek
      
      // Handle case where startDay goes negative (previous month)
      if (startDay <= 0) {
        startMonth--
        if (startMonth < 0) {
          startMonth = 11
          startYear--
        }
        const daysInPrevMonth = new Date(Date.UTC(startYear, startMonth + 1, 0)).getUTCDate()
        startDay = daysInPrevMonth + startDay
      }
      
      // Calculate what day of week the last day falls on
      const lastDayDate = new Date(Date.UTC(year, month, lastDayOfMonth.day))
      const lastDayOfWeek = lastDayDate.getUTCDay()
      
      // End date is last day of month plus days to get to Saturday
      let endYear = year
      let endMonth = month
      let endDay = lastDayOfMonth.day + (6 - lastDayOfWeek)
      
      // Handle case where endDay exceeds days in month (next month)
      const daysInEndMonth = new Date(Date.UTC(endYear, endMonth + 1, 0)).getUTCDate()
      if (endDay > daysInEndMonth) {
        endDay = endDay - daysInEndMonth
        endMonth++
        if (endMonth > 11) {
          endMonth = 0
          endYear++
        }
      }
      
      return { 
        start: { year: startYear, month: startMonth, day: startDay },
        end: { year: endYear, month: endMonth, day: endDay }
      }
    })

    const calendarDays = computed(() => {
      const days = []
      const { start, end } = calendarRange.value
      
      // Work with pure year/month/day numbers - no Date object manipulation
      let year = start.year
      let month = start.month
      let day = start.day
      
      const endYear = end.year
      const endMonth = end.month
      const endDay = end.day
      
      // Loop through days using pure date math
      while (true) {
        // Build UTC date key directly from year/month/day - exactly like event date keys
        const dateKey = `${year}-${String(month + 1).padStart(2, '0')}-${String(day).padStart(2, '0')}`
        
        // Check if we've reached the end
        if (year > endYear || 
            (year === endYear && month > endMonth) ||
            (year === endYear && month === endMonth && day > endDay)) {
          break
        }
        
        const dayEvents = events.value
          .map((ev) => {
            // CRITICAL: Extract date key ONLY from the raw ISO string
            // This is the ONLY way to guarantee no timezone conversion
            if (!ev._startTimeStr) {
              console.warn('Event missing _startTimeStr:', ev)
              return null
            }
            
            // Extract date directly from ISO string: "2026-01-30T09:00:00Z" -> "2026-01-30"
            // Split on 'T' and take first part - this is the UTC date as stored
            // IMPORTANT: Do NOT use Date object methods - only string manipulation
            let eventDateKey = ev._startTimeStr.split('T')[0]
            
            // Handle case where there might be a space instead of T (some API formats)
            if (!eventDateKey && ev._startTimeStr.includes(' ')) {
              eventDateKey = ev._startTimeStr.split(' ')[0]
            }
            
            // Validate the date key format (should be YYYY-MM-DD)
            if (!/^\d{4}-\d{2}-\d{2}$/.test(eventDateKey)) {
              console.warn('Invalid date key format:', eventDateKey, 'from:', ev._startTimeStr, 'event:', ev)
              return null
            }
            
            // Debug: Log if date keys don't match
            if (eventDateKey !== dateKey) {
              // Only log once per mismatch to avoid spam
              if (!window._dateKeyMismatchLogged) {
                console.log('Date key mismatch - Calendar:', dateKey, 'Event:', eventDateKey, 'Event ISO:', ev._startTimeStr)
                window._dateKeyMismatchLogged = true
              }
            }
            
            return {
              ...ev,
              dateKey: eventDateKey
            }
          })
          .filter((ev) => ev && ev.dateKey === dateKey)
        // Compare dates for today/past using pure date comparison
        const today = new Date()
        const todayYear = today.getUTCFullYear()
        const todayMonth = today.getUTCMonth()
        const todayDay = today.getUTCDate()
        
        const isPast = year < todayYear ||
          (year === todayYear && month < todayMonth) ||
          (year === todayYear && month === todayMonth && day < todayDay)
        
        const isToday = year === todayYear && month === todayMonth && day === todayDay
        
        // CRITICAL: Ensure the displayed day number matches the date key
        // Extract day from dateKey to guarantee they match
        const dateKeyDay = parseInt(dateKey.split('-')[2], 10)
        
        days.push({
          date: dateKey,
          day: dateKeyDay, // Use day from dateKey to ensure it matches
          isCurrentMonth: month === currentMonth.value,
          isToday: isToday,
          isPast: isPast,
          events: dayEvents
        })
        
        // Debug: Log if day number doesn't match date key
        if (dateKeyDay !== day && !window._dayMismatchLogged) {
          console.warn('⚠️ Day number mismatch!', {
            dateKey: dateKey,
            dateKeyDay: dateKeyDay,
            trackedDay: day,
            month: month,
            year: year
          })
          window._dayMismatchLogged = true
        }
        
        // Increment to next day using pure date math
        const daysInMonth = new Date(Date.UTC(year, month + 1, 0)).getUTCDate()
        day++
        if (day > daysInMonth) {
          day = 1
          month++
          if (month > 11) {
            month = 0
            year++
          }
        }
      }
      return days
    })

    const resetForm = () => {
      form.value = {
        title: '',
        start: '',
        end: '',
        location: '',
        description: ''
      }
      formError.value = ''
    }

    const openCreateModal = (dateKey) => {
      editingEvent.value = null
      resetForm()
      if (dateKey) {
        // dateKey is already in format "YYYY-MM-DD" (UTC date from calendar)
        // datetime-local inputs need the date/time in YYYY-MM-DDTHH:mm format
        // Since dateKey is already correctly formatted, just append the time
        // This represents the UTC date/time that will be stored
        form.value.start = `${dateKey}T09:00`
        form.value.end = `${dateKey}T10:00`
      }
      showModal.value = true
    }

    const openEditModal = (event) => {
      editingEvent.value = event
      // CRITICAL: Use the raw ISO strings (_startTimeStr, _endTimeStr) to avoid any timezone conversion
      // These are the exact UTC times as stored in the database
      form.value = {
        title: event.title,
        start: toDateTimeLocal(event._startTimeStr || event.startTime || event.start || event.date),
        end: toDateTimeLocal(event._endTimeStr || event.endTime || event.end || event.startTime || event.date),
        location: event.location || '',
        description: event.description || ''
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
      formError.value = ''
      saving.value = true

      // Parse datetime-local strings (which are in local timezone)
      // and convert them to UTC for storage
      // datetime-local format: "YYYY-MM-DDTHH:mm" (no timezone, interpreted as local)
      let start = null
      let end = null

      if (form.value.start) {
        // Parse datetime-local string directly - it's already in UTC format (YYYY-MM-DDTHH:mm)
        // We treat it as UTC to preserve the exact date/time the user entered
        const [datePart, timePart] = form.value.start.split('T')
        const [year, month, day] = datePart.split('-').map(Number)
        const [hours, minutes] = timePart.split(':').map(Number)
        
        // CRITICAL: Create UTC date with the exact values from the form
        // This ensures the date stored matches exactly what the user entered
        start = new Date(Date.UTC(year, month - 1, day, hours, minutes || 0, 0, 0))
        
        // Debug: Verify the date key will match
        const storedDateKey = `${year}-${String(month).padStart(2, '0')}-${String(day).padStart(2, '0')}`
        const isoString = start.toISOString()
        const extractedDateKey = isoString.split('T')[0]
        if (storedDateKey !== extractedDateKey) {
          console.error('❌ Date key mismatch on save!', {
            formDateKey: storedDateKey,
            isoDateKey: extractedDateKey,
            isoString: isoString,
            formInput: form.value.start,
            year, month, day, hours, minutes
          })
        } else {
          console.log('✅ Date key matches on save:', storedDateKey, '→', isoString)
        }
      }

      if (form.value.end) {
        const [datePart, timePart] = form.value.end.split('T')
        const [year, month, day] = datePart.split('-').map(Number)
        const [hours, minutes] = timePart.split(':').map(Number)
        end = new Date(Date.UTC(year, month - 1, day, hours, minutes || 0, 0, 0))
      }

      if (!start || !end || end < start) {
        formError.value = 'Please provide valid start and end times.'
        saving.value = false
        return
      }

      const payload = {
        title: form.value.title.trim(),
        description: form.value.description.trim(),
        startTime: start.toISOString(),
        endTime: end.toISOString(),
        location: form.value.location.trim(),
        recurrence: 'none'
      }

      try {
        if (editingEvent.value) {
          await eventService.updateEvent(editingEvent.value.id, payload)
        } else {
          await eventService.createEvent(payload)
        }
        await loadEvents()
        closeModal()
      } catch (err) {
        console.error('Failed to save event', err)
        formError.value = err.message || 'Failed to save event'
      } finally {
        saving.value = false
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
      currentMonth.value = now.getUTCMonth()
      currentYear.value = now.getUTCFullYear()
    }

    const confirmDeleteEvent = (id) => {
      eventToDelete.value = id
      showDeleteModal.value = true
    }

    const deleteEvent = async () => {
      if (!eventToDelete.value) return
      
      deleting.value = true
      const id = eventToDelete.value
      showDeleteModal.value = false
      
      try {
        await eventService.deleteEvent(id)
        await loadEvents()
        eventToDelete.value = null
        // Close edit modal if we were editing the deleted event
        if (editingEvent.value && editingEvent.value.id === id) {
          closeModal()
        }
      } catch (err) {
        console.error('Error deleting event:', err)
        formError.value = err.message || 'Failed to delete event'
        eventToDelete.value = null
      } finally {
        deleting.value = false
      }
    }

    const cancelDelete = () => {
      showDeleteModal.value = false
      eventToDelete.value = null
    }

    const confirmDeleteAll = () => {
      showDeleteAllModal.value = true
    }

    const deleteAllEvents = async () => {
      deleting.value = true
      showDeleteAllModal.value = false
      
      try {
        await eventService.deleteAllEvents()
        await loadEvents()
        // Close any open modals
        closeModal()
      } catch (err) {
        console.error('Error deleting all events:', err)
        formError.value = err.message || 'Failed to delete all events'
      } finally {
        deleting.value = false
      }
    }

    const cancelDeleteAll = () => {
      showDeleteAllModal.value = false
    }

    const confirmDeleteOldEvents = () => {
      showDeleteOldModal.value = true
    }

    const deleteOldEvents = async () => {
      deleting.value = true
      showDeleteOldModal.value = false
      
      try {
        await eventService.deleteOldEvents()
        await loadEvents()
        await loadOldEventsCount()
        // Close any open modals
        closeModal()
      } catch (err) {
        console.error('Error deleting old events:', err)
        formError.value = err.message || 'Failed to delete old events'
      } finally {
        deleting.value = false
      }
    }

    const cancelDeleteOld = () => {
      showDeleteOldModal.value = false
    }

    onMounted(() => {
      loadEvents()
      loadOldEventsCount()
    })

    return {
      weekdayLabels,
      monthLabel,
      calendarDays,
      showModal,
      form,
      formError,
      saving,
      deleting,
      showDeleteModal,
      showDeleteAllModal,
      events,
      editingEvent,
      openCreateModal,
      openEditModal,
      closeModal,
      saveEvent,
      previousMonth,
      nextMonth,
      resetToToday,
      confirmDeleteEvent,
      deleteEvent,
      cancelDelete,
      confirmDeleteAll,
      deleteAllEvents,
      cancelDeleteAll,
      oldEventsCount,
      showDeleteOldModal,
      confirmDeleteOldEvents,
      deleteOldEvents,
      cancelDeleteOld
    }
  }
}
</script>
