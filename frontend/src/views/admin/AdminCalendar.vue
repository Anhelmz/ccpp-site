<template>
  <AdminLayout>
    <div class="mb-6 flex flex-wrap gap-3 items-center justify-between">
      <div>
        <h1 class="text-2xl font-semibold text-gray-900">Calendar</h1>
        <p class="text-sm text-gray-600">Create events that show on the public calendar page.</p>
      </div>
      <div class="flex gap-3">
        <button
          @click="resetToToday"
          class="px-4 py-2 rounded-md border border-gray-200 text-gray-700 hover:bg-gray-50"
        >
          Today
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

    <div class="bg-white rounded-xl shadow border border-gray-200 p-4 md:p-6">
      <div class="flex items-center justify-between mb-4">
        <div class="flex items-center gap-3">
          <button
            @click="previousMonth"
            class="p-2 rounded-md border border-gray-200 hover:bg-gray-50"
            aria-label="Previous month"
          >
            <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 19l-7-7 7-7" />
            </svg>
          </button>
          <button
            @click="nextMonth"
            class="p-2 rounded-md border border-gray-200 hover:bg-gray-50"
            aria-label="Next month"
          >
            <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7" />
            </svg>
          </button>
          <h2 class="text-xl font-semibold text-gray-900">{{ monthLabel }}</h2>
        </div>
      </div>

      <div class="grid grid-cols-7 gap-2 text-sm font-semibold text-gray-600 mb-2">
        <div v-for="day in weekdayLabels" :key="day" class="text-center py-2">
          {{ day }}
        </div>
      </div>

      <div class="grid grid-cols-7 gap-2">
        <div
          v-for="day in calendarDays"
          :key="day.date"
          :class="[
            'min-h-28 border rounded-lg p-3 flex flex-col gap-2 cursor-pointer transition-colors',
            day.isCurrentMonth ? 'bg-white hover:bg-gray-50 border-gray-200' : 'bg-gray-50 text-gray-400 border-gray-100',
            day.isToday ? 'border-main ring-1 ring-main/30' : ''
          ]"
          @click="openCreateModal(day.date)"
        >
          <div class="flex items-center justify-between">
            <span :class="['text-sm font-semibold', day.isCurrentMonth ? 'text-gray-900' : 'text-gray-400']">
              {{ day.day }}
            </span>
            <span v-if="day.events.length" class="text-[11px] font-semibold bg-main/10 text-main px-2 py-1 rounded-full">
              {{ day.events.length }} evt
            </span>
          </div>
          <div class="space-y-1">
            <button
              v-for="event in day.events.slice(0, 3)"
              :key="event.id"
              type="button"
              class="w-full text-left text-xs bg-main/10 text-main rounded-md px-2 py-1 truncate hover:bg-main/20"
              @click.stop="openEditModal(event)"
            >
              {{ event.title }}
            </button>
            <div v-if="day.events.length > 3" class="text-xs text-gray-500">+{{ day.events.length - 3 }} more</div>
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
        class="bg-white rounded-xl shadow-2xl w-full max-w-xl max-h-[90vh] overflow-y-auto"
        @click.stop
      >
        <div class="p-5 border-b border-gray-200 flex items-center justify-between">
          <h3 class="text-xl font-semibold text-gray-900">{{ editingEvent ? 'Edit Event' : 'New Event' }}</h3>
          <button @click="closeModal" class="text-gray-400 hover:text-gray-600">
            <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12" />
            </svg>
          </button>
        </div>

        <form @submit.prevent="saveEvent" class="p-5 space-y-4">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1">Title *</label>
            <input
              v-model="form.title"
              type="text"
              required
              class="w-full border border-gray-300 rounded-md px-3 py-2 focus:ring-2 focus:ring-main focus:border-main"
              placeholder="Event title"
            />
          </div>

          <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-1">Start *</label>
              <input
                v-model="form.start"
                type="datetime-local"
                required
                class="w-full border border-gray-300 rounded-md px-3 py-2 focus:ring-2 focus:ring-main focus:border-main"
              />
            </div>
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-1">End *</label>
              <input
                v-model="form.end"
                type="datetime-local"
                required
                class="w-full border border-gray-300 rounded-md px-3 py-2 focus:ring-2 focus:ring-main focus:border-main"
              />
            </div>
          </div>

          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1">Location</label>
            <input
              v-model="form.location"
              type="text"
              class="w-full border border-gray-300 rounded-md px-3 py-2 focus:ring-2 focus:ring-main focus:border-main"
              placeholder="Main hall"
            />
          </div>

          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1">Summary</label>
            <input
              v-model="form.summary"
              type="text"
              class="w-full border border-gray-300 rounded-md px-3 py-2 focus:ring-2 focus:ring-main focus:border-main"
              placeholder="Short blurb"
            />
          </div>

          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1">Description</label>
            <textarea
              v-model="form.description"
              rows="4"
              class="w-full border border-gray-300 rounded-md px-3 py-2 focus:ring-2 focus:ring-main focus:border-main"
              placeholder="Full details"
            ></textarea>
          </div>

          <div v-if="formError" class="p-3 border border-red-200 bg-red-50 text-red-700 text-sm rounded-md">
            {{ formError }}
          </div>

          <div class="flex justify-end gap-3 pt-3">
            <button type="button" @click="closeModal" class="px-4 py-2 border border-gray-300 rounded-md text-gray-700 hover:bg-gray-50">
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
  name: 'AdminCalendar',
  components: {
    AdminLayout
  },
  setup() {
    const events = ref([])
    const showModal = ref(false)
    const editingEvent = ref(null)
    const formError = ref('')
    const saving = ref(false)

    const currentMonth = ref(new Date().getMonth())
    const currentYear = ref(new Date().getFullYear())

    const form = ref({
      title: '',
      start: '',
      end: '',
      location: '',
      summary: '',
      description: ''
    })

    const weekdayLabels = ['Sun', 'Mon', 'Tue', 'Wed', 'Thu', 'Fri', 'Sat']
    const monthLabel = computed(() =>
      new Date(currentYear.value, currentMonth.value, 1).toLocaleString('default', { month: 'long', year: 'numeric' })
    )

    const loadEvents = async () => {
      try {
        const response = await eventService.getEvents(false)
        events.value = response.events || []
      } catch (err) {
        console.error('Failed to load events', err)
      }
    }

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
        days.push({
          date: dateKey,
          day: cursor.getDate(),
          isCurrentMonth: cursor.getMonth() === currentMonth.value,
          isToday: cursor.toDateString() === new Date().toDateString(),
          events: dayEvents
        })
        cursor.setDate(cursor.getDate() + 1)
      }
      return days
    })

    const resetForm = () => {
      form.value = {
        title: '',
        start: '',
        end: '',
        location: '',
        summary: '',
        description: ''
      }
      formError.value = ''
    }

    const openCreateModal = (dateKey) => {
      editingEvent.value = null
      resetForm()
      if (dateKey) {
        const start = new Date(dateKey)
        start.setHours(9, 0, 0, 0)
        const end = new Date(dateKey)
        end.setHours(10, 0, 0, 0)
        form.value.start = toDateTimeLocal(start.toISOString())
        form.value.end = toDateTimeLocal(end.toISOString())
      }
      showModal.value = true
    }

    const openEditModal = (event) => {
      editingEvent.value = event
      form.value = {
        title: event.title,
        start: toDateTimeLocal(event.startTime || event.start || event.date),
        end: toDateTimeLocal(event.endTime || event.end || event.startTime || event.date),
        location: event.location || '',
        summary: event.summary || '',
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

      const start = form.value.start ? new Date(form.value.start) : null
      const end = form.value.end ? new Date(form.value.end) : null

      if (!start || !end || end < start) {
        formError.value = 'Please provide valid start and end times.'
        saving.value = false
        return
      }

      const payload = {
        title: form.value.title.trim(),
        summary: form.value.summary.trim(),
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
      currentMonth.value = now.getMonth()
      currentYear.value = now.getFullYear()
    }

    onMounted(() => {
      loadEvents()
    })

    return {
      weekdayLabels,
      monthLabel,
      calendarDays,
      showModal,
      form,
      formError,
      saving,
      openCreateModal,
      openEditModal,
      closeModal,
      saveEvent,
      previousMonth,
      nextMonth,
      resetToToday
    }
  }
}
</script>
