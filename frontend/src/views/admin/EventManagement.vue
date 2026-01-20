<template>
  <AdminLayout>
    <div class="mb-6 flex justify-end">
      <button
        class="px-6 py-3 bg-main text-white rounded-lg hover:opacity-90 transition-colors shadow-lg font-medium flex items-center space-x-2"
        @click="openCreateModal"
      >
        <svg
          class="w-5 h-5"
          fill="none"
          stroke="currentColor"
          viewBox="0 0 24 24"
        >
          <path
            stroke-linecap="round"
            stroke-linejoin="round"
            stroke-width="2"
            d="M12 4v16m8-8H4"
          ></path>
        </svg>
        <span>New Event</span>
      </button>
    </div>

    <!-- Events List -->
    <div>
      <div class="p-6 border-b border-gray-200">
        <div
          class="flex flex-col sm:flex-row justify-between items-start sm:items-center gap-4"
        >
          <h2 class="text-xl font-semibold text-gray-900">Events</h2>
          <div class="flex items-center space-x-2 w-full sm:w-auto">
            <input
              v-model="searchQuery"
              type="text"
              placeholder="Search events..."
              class="flex-1 sm:flex-none px-4 py-2 border border-gray-300 rounded-lg text-sm min-w-48"
              @input="filterEvents"
            />
            <select
              v-model="filterType"
              class="px-4 py-2 border border-gray-300 rounded-lg text-sm"
              @change="loadEvents"
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
        <div
          class="inline-block animate-spin rounded-full h-12 w-12 border-4 border-main border-t-transparent mb-4"
        ></div>
        <p class="text-gray-600">Loading events...</p>
      </div>

      <!-- Error State -->
      <div v-else-if="error" class="text-center py-12">
        <p class="text-red-600 mb-4">{{ error }}</p>
        <button
          class="px-4 py-2 bg-main text-white rounded-lg hover:opacity-90 transition-colors"
          @click="loadEvents"
        >
          Retry
        </button>
      </div>

      <!-- Empty State -->
      <div v-else-if="filteredEvents.length === 0" class="text-center py-12">
        <svg
          class="mx-auto h-12 w-12 text-gray-400 mb-4"
          fill="none"
          stroke="currentColor"
          viewBox="0 0 24 24"
        >
          <path
            stroke-linecap="round"
            stroke-linejoin="round"
            stroke-width="2"
            d="M8 7V3m8 4V3m-9 8h10M5 21h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v12a2 2 0 002 2z"
          ></path>
        </svg>
        <p class="text-gray-600 mb-2">
          {{ searchQuery ? "No events found" : "No events created yet" }}
        </p>
        <p class="text-sm text-gray-500 mb-4">
          {{
            searchQuery
              ? "Try a different search term"
              : "Create your first event to get started"
          }}
        </p>
        <button
          v-if="!searchQuery"
          class="px-6 py-2 bg-main text-white rounded-lg hover:opacity-90 transition-colors"
          @click="openCreateModal"
        >
          Create Event
        </button>
      </div>

      <!-- Events Table -->
      <div v-else class="overflow-x-auto">
        <table class="min-w-full divide-y divide-gray-200">
          <thead>
            <tr>
              <th
                class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider"
              >
                Title
              </th>
              <th
                class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider"
              >
                Start
              </th>
              <th
                class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider"
              >
                End
              </th>
              <th
                class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider"
              >
                Category
              </th>
              <th
                class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider"
              >
                Actions
              </th>
            </tr>
          </thead>
          <tbody class="divide-y divide-gray-200">
            <tr
              v-for="event in filteredEvents"
              :key="event.id"
              class="transition-colors hover:bg-gray-50 cursor-pointer"
              @click="openEditModal(event)"
            >
              <td class="px-6 py-4 whitespace-nowrap">
                <div class="text-sm font-medium text-gray-900">
                  {{ event.title }}
                </div>
                <div
                  v-if="event.summary"
                  class="text-sm text-gray-500 truncate max-w-xs"
                >
                  {{ event.summary }}
                </div>
              </td>
              <td class="px-6 py-4 whitespace-nowrap">
                <div class="text-sm text-gray-900">
                  {{ formatDateTime(event.startTime) }}
                </div>
              </td>
              <td class="px-6 py-4 whitespace-nowrap">
                <div class="text-sm text-gray-900">
                  {{ formatDateTime(event.endTime) }}
                </div>
              </td>
              <td class="px-6 py-4 whitespace-nowrap">
                <div class="text-sm text-gray-900">
                  {{ event.category || "—" }}
                </div>
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm font-medium" @click.stop>
                <button
                  class="text-red-600 hover:text-red-800"
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

    <!-- Create/Edit Event Modal -->
    <div
      v-if="showModal"
      class="fixed inset-0 bg-black bg-opacity-75 backdrop-blur-sm flex items-center justify-center z-50 p-4"
      @click.self="closeModal"
    >
      <div
        class="bg-white rounded-xl shadow-2xl border-2 border-gray-200 max-w-2xl w-full max-h-[90vh] overflow-y-auto transform transition-all"
        @click.stop
      >
        <div class="p-8">
          <!-- Modal Header -->
          <div class="flex justify-between items-start mb-6 pb-4 border-b border-gray-200">
            <h2 class="text-3xl font-bold text-gray-900">
              {{ editingEvent ? "Edit Event" : "Create New Event" }}
            </h2>
            <button
              class="text-gray-500 hover:text-gray-700 transition-colors p-1 hover:bg-gray-100 rounded-full"
              @click="closeModal"
            >
              <svg
                class="h-6 w-6"
                fill="none"
                stroke="currentColor"
                viewBox="0 0 24 24"
              >
                <path
                  stroke-linecap="round"
                  stroke-linejoin="round"
                  stroke-width="2"
                  d="M6 18L18 6M6 6l12 12"
                ></path>
              </svg>
            </button>
          </div>

          <form class="space-y-6" @submit.prevent="saveEvent">
            <!-- Event Title -->
            <div class="bg-gray-50 p-4 rounded-lg border border-gray-200">
              <label class="block text-xs font-semibold text-gray-600 uppercase tracking-wide mb-2">
                Event Title *
              </label>
              <input
                v-model="eventForm.title"
                type="text"
                required
                class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-main focus:border-transparent bg-white"
                placeholder="Sunday Worship Service"
              />
            </div>

            <!-- Summary -->
            <div class="bg-gray-50 p-4 rounded-lg border border-gray-200">
              <label class="block text-xs font-semibold text-gray-600 uppercase tracking-wide mb-2">
                Summary (Short Description)
              </label>
              <input
                v-model="eventForm.summary"
                type="text"
                class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-main focus:border-transparent bg-white"
                placeholder="Brief summary of the event"
              />
            </div>

            <!-- Date & Time -->
            <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
              <div class="bg-gray-50 p-4 rounded-lg border border-gray-200">
                <label class="block text-xs font-semibold text-gray-600 uppercase tracking-wide mb-2">
                  Start date & time *
                </label>
                <input
                  v-model="eventForm.start"
                  type="datetime-local"
                  required
                  class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-main focus:border-transparent bg-white"
                />
              </div>
              <div class="bg-gray-50 p-4 rounded-lg border border-gray-200">
                <label class="block text-xs font-semibold text-gray-600 uppercase tracking-wide mb-2">
                  End date & time *
                </label>
                <input
                  v-model="eventForm.end"
                  type="datetime-local"
                  required
                  class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-main focus:border-transparent bg-white"
                />
              </div>
            </div>

            <!-- Location & Category -->
            <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
              <div class="bg-gray-50 p-4 rounded-lg border border-gray-200">
                <label class="block text-xs font-semibold text-gray-600 uppercase tracking-wide mb-2">
                  Location
                </label>
                <input
                  v-model="eventForm.location"
                  type="text"
                  class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-main focus:border-transparent bg-white"
                  placeholder="Main hall"
                />
              </div>
              <div class="bg-gray-50 p-4 rounded-lg border border-gray-200">
                <label class="block text-xs font-semibold text-gray-600 uppercase tracking-wide mb-2">
                  Category
                </label>
                <input
                  v-model="eventForm.category"
                  type="text"
                  class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-main focus:border-transparent bg-white"
                  placeholder="Worship, Youth, Outreach..."
                />
              </div>
            </div>

            <!-- Color & Recurrence -->
            <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
              <div class="bg-gray-50 p-4 rounded-lg border border-gray-200">
                <label class="block text-xs font-semibold text-gray-600 uppercase tracking-wide mb-2">
                  Color (badge)
                </label>
                <input
                  v-model="eventForm.color"
                  type="color"
                  class="w-full h-11 px-2 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-main focus:border-transparent bg-white"
                />
              </div>
              <div class="bg-gray-50 p-4 rounded-lg border border-gray-200">
                <label class="block text-xs font-semibold text-gray-600 uppercase tracking-wide mb-2">
                  Recurrence
                </label>
                <select
                  v-model="eventForm.recurrence"
                  class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-main focus:border-transparent bg-white"
                >
                  <option value="none">Does not repeat</option>
                  <option value="daily">Daily</option>
                  <option value="weekly">Weekly</option>
                  <option value="monthly">Monthly</option>
                  <option value="yearly">Yearly</option>
                </select>
              </div>
            </div>

            <!-- Recurrence Ends -->
            <div class="bg-gray-50 p-4 rounded-lg border border-gray-200">
              <label class="block text-xs font-semibold text-gray-600 uppercase tracking-wide mb-2">
                Recurrence ends (optional)
              </label>
              <input
                v-model="eventForm.recurrenceEndsAt"
                type="datetime-local"
                class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-main focus:border-transparent bg-white"
              />
            </div>

            <!-- Description -->
            <div class="bg-gray-50 p-4 rounded-lg border border-gray-200">
              <label class="block text-xs font-semibold text-gray-600 uppercase tracking-wide mb-2">
                Description
              </label>
              <textarea
                v-model="eventForm.description"
                rows="4"
                class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-main focus:border-transparent bg-white"
                placeholder="Event description..."
              ></textarea>
            </div>

            <!-- Error Message -->
            <div
              v-if="formError"
              class="p-4 bg-red-50 border-2 border-red-200 rounded-lg"
            >
              <p class="text-red-600 text-sm font-medium">{{ formError }}</p>
            </div>

            <!-- Modal Footer -->
            <div class="mt-8 pt-6 border-t border-gray-200 flex justify-end space-x-3">
              <button
                type="button"
                class="px-6 py-2.5 border-2 border-gray-300 rounded-lg text-gray-700 hover:bg-gray-50 transition-colors font-semibold shadow-sm"
                @click="closeModal"
              >
                Cancel
              </button>
              <button
                type="submit"
                :disabled="saving"
                class="px-6 py-2.5 bg-main text-white rounded-lg hover:opacity-90 transition-colors font-semibold shadow-md disabled:opacity-50 disabled:cursor-not-allowed"
              >
                {{
                  saving
                    ? "Saving..."
                    : editingEvent
                      ? "Update Event"
                      : "Create Event"
                }}
              </button>
            </div>
          </form>
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

    <!-- Error Modal -->
    <ErrorModal
      :show="showErrorModal"
      :message="errorMessage"
      @close="closeErrorModal"
    />
  </AdminLayout>
</template>

<script>
import { ref, computed, onMounted } from "vue";
import AdminLayout from "@/components/admin/AdminLayout.vue";
import ConfirmationModal from "@/components/admin/ConfirmationModal.vue";
import ErrorModal from "@/components/admin/ErrorModal.vue";
import { eventService } from "@/services/eventService";

export default {
  name: "EventManagement",
  components: {
    AdminLayout,
    ConfirmationModal,
    ErrorModal,
  },
  setup() {
    const showModal = ref(false);
    const editingEvent = ref(null);
    const events = ref([]);
    const loading = ref(false);
    const error = ref("");
    const saving = ref(false);
    const formError = ref("");
    const filterType = ref("all");
    const searchQuery = ref("");
    const showConfirmModal = ref(false);
    const showErrorModal = ref(false);
    const errorMessage = ref("");
    const eventToDelete = ref(null);

    const eventForm = ref({
      title: "",
      summary: "",
      description: "",
      start: "",
      end: "",
      location: "",
      category: "",
      color: "#f97316",
      recurrence: "none",
      recurrenceEndsAt: "",
    });

    const toDateTimeLocal = (value) => {
      if (!value) return "";
      const date = new Date(value);
      const offset = date.getTimezoneOffset();
      const local = new Date(date.getTime() - offset * 60000);
      return local.toISOString().slice(0, 16);
    };

    const formatDateTime = (dateString) => {
      if (!dateString) return "—";
      const date = new Date(dateString);
      return date.toLocaleString("en-US", {
        year: "numeric",
        month: "short",
        day: "numeric",
        hour: "2-digit",
        minute: "2-digit",
      });
    };

    const loadEvents = async () => {
      loading.value = true;
      error.value = "";
      try {
        const upcoming = filterType.value === "upcoming";
        const response = await eventService.getEvents(upcoming);
        const apiEvents = response.events || [];

        events.value = apiEvents.map((event) => ({
          ...event,
          startTime: event.startTime || event.start || event.date,
          endTime: event.endTime || event.end || event.startTime || event.date,
          recurrence: event.recurrence || "none",
          color: event.color || "#f97316",
        }));
      } catch (err) {
        console.error("Error loading events:", err);
        error.value = err.message || "Failed to load events";
      } finally {
        loading.value = false;
      }
    };

    const filterEvents = () => {
      // Filtering is handled by computed property
    };

    const filteredEvents = computed(() => {
      let filtered = [...events.value];

      // Apply search filter
      if (searchQuery.value) {
        const query = searchQuery.value.toLowerCase();
        filtered = filtered.filter(
          (event) =>
            event.title.toLowerCase().includes(query) ||
            (event.summary && event.summary.toLowerCase().includes(query)) ||
            (event.description &&
              event.description.toLowerCase().includes(query)) ||
            (event.location && event.location.toLowerCase().includes(query)),
        );
      }

      // Apply past filter
      if (filterType.value === "past") {
        const now = new Date();
        filtered = filtered.filter((event) => new Date(event.startTime) < now);
      }

      return filtered.sort(
        (a, b) => new Date(a.startTime) - new Date(b.startTime),
      );
    });

    const openCreateModal = () => {
      editingEvent.value = null;
      eventForm.value = {
        title: "",
        summary: "",
        description: "",
        start: "",
        end: "",
        location: "",
        category: "",
        color: "#f97316",
        recurrence: "none",
        recurrenceEndsAt: "",
      };
      formError.value = "";
      showModal.value = true;
    };

    const openEditModal = (event) => {
      editingEvent.value = event;
      eventForm.value = {
        title: event.title,
        summary: event.summary || "",
        description: event.description || "",
        start: toDateTimeLocal(event.startTime || event.start || event.date),
        end: toDateTimeLocal(
          event.endTime || event.end || event.startTime || event.date,
        ),
        location: event.location || "",
        category: event.category || "",
        color: event.color || "#f97316",
        recurrence: event.recurrence || "none",
        recurrenceEndsAt: event.recurrenceEndsAt
          ? toDateTimeLocal(event.recurrenceEndsAt)
          : "",
      };
      formError.value = "";
      showModal.value = true;
    };

    const closeModal = () => {
      showModal.value = false;
      editingEvent.value = null;
      formError.value = "";
    };

    const saveEvent = async () => {
      saving.value = true;
      formError.value = "";

      const parseDateTime = (value) => (value ? new Date(value) : null);

      try {
        const start = parseDateTime(eventForm.value.start);
        const end = parseDateTime(eventForm.value.end) || start;
        const recurrenceEndsAt = parseDateTime(
          eventForm.value.recurrenceEndsAt,
        );

        if (!start || !end) {
          formError.value = "Start and end date/time are required.";
          saving.value = false;
          return;
        }

        if (end < start) {
          formError.value = "End time must be after the start time.";
          saving.value = false;
          return;
        }

        if (recurrenceEndsAt && recurrenceEndsAt < start) {
          formError.value = "Recurrence end must be after the start time.";
          saving.value = false;
          return;
        }

        const eventData = {
          title: eventForm.value.title,
          summary: eventForm.value.summary || "",
          description: eventForm.value.description || "",
          startTime: start.toISOString(),
          endTime: end.toISOString(),
          location: eventForm.value.location || "",
          category: eventForm.value.category || "",
          color: eventForm.value.color || "",
          recurrence: eventForm.value.recurrence || "none",
          recurrenceEndsAt: recurrenceEndsAt
            ? recurrenceEndsAt.toISOString()
            : null,
        };

        if (editingEvent.value) {
          await eventService.updateEvent(editingEvent.value.id, eventData);
        } else {
          await eventService.createEvent(eventData);
        }

        await loadEvents();
        closeModal();
      } catch (err) {
        console.error("Error saving event:", err);
        formError.value = err.message || "Failed to save event";
      } finally {
        saving.value = false;
      }
    };

    const deleteEvent = (id) => {
      eventToDelete.value = id;
      showConfirmModal.value = true;
    };

    const confirmDelete = async () => {
      if (!eventToDelete.value) return;

      const id = eventToDelete.value;
      showConfirmModal.value = false;

      try {
        await eventService.deleteEvent(id);
        await loadEvents();
        eventToDelete.value = null;
      } catch (err) {
        console.error("Error deleting event:", err);
        errorMessage.value = `Failed to delete event: ${err.message}`;
        showErrorModal.value = true;
        eventToDelete.value = null;
      }
    };

    const cancelDelete = () => {
      showConfirmModal.value = false;
      eventToDelete.value = null;
    };

    const closeErrorModal = () => {
      showErrorModal.value = false;
      errorMessage.value = "";
    };

    onMounted(() => {
      loadEvents();
    });

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
      formatDateTime,
      loadEvents,
      filterEvents,
      openCreateModal,
      openEditModal,
      closeModal,
      saveEvent,
      deleteEvent,
      showConfirmModal,
      showErrorModal,
      errorMessage,
      confirmDelete,
      cancelDelete,
      closeErrorModal,
    };
  },
};
</script>
