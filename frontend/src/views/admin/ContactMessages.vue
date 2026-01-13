<template>
  <AdminLayout>
    <!-- Filters -->
    <div class="bg-white rounded-lg shadow p-6 mb-6">
      <div class="flex flex-wrap items-center gap-4">
        <select
          v-model="statusFilter"
          class="px-4 py-2 border border-gray-300 rounded-lg text-sm"
          @change="loadRequests"
        >
          <option value="all">All Messages</option>
          <option value="unread">Unread</option>
          <option value="read">Read</option>
        </select>
        <input
          v-model="searchQuery"
          type="text"
          placeholder="Search messages..."
          class="px-4 py-2 border border-gray-300 rounded-lg text-sm flex-1 min-w-64"
          @input="filterContacts"
        />
        <button
          class="px-4 py-2 bg-main text-white rounded-lg hover:opacity-90 transition-colors text-sm"
          @click="loadRequests"
        >
          Refresh
        </button>
      </div>
    </div>

    <!-- Messages List -->
    <div class="bg-white rounded-lg shadow">
      <!-- Loading State -->
      <div v-if="loading" class="text-center py-12">
        <div
          class="inline-block animate-spin rounded-full h-12 w-12 border-4 border-main border-t-transparent mb-4"
        ></div>
        <p class="text-gray-600">Loading messages...</p>
      </div>

      <!-- Error State -->
      <div v-else-if="error" class="text-center py-12">
        <p class="text-red-600 mb-4">{{ error }}</p>
        <button
          class="px-4 py-2 bg-main text-white rounded-lg hover:opacity-90 transition-colors"
          @click="loadContacts"
        >
          Retry
        </button>
      </div>

      <!-- Empty State -->
      <div v-else-if="filteredRequests.length === 0" class="text-center py-12">
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
            d="M3 8l7.89 5.26a2 2 0 002.22 0L21 8M5 19h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v10a2 2 0 002 2z"
          ></path>
        </svg>
        <p class="text-gray-600 mb-2">No messages found</p>
        <p class="text-sm text-gray-500">
          {{
            searchQuery
              ? "Try a different search term"
              : "Contact form submissions will appear here"
          }}
        </p>
      </div>

      <!-- Messages Table -->
      <div v-else class="overflow-x-auto">
        <table class="min-w-full divide-y divide-gray-200">
          <thead class="bg-gray-50">
            <tr>
              <th
                class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider"
              >
                Name
              </th>
              <th
                class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider"
              >
                Email
              </th>
              <th
                class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider"
              >
                Phone
              </th>
              <th
                class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider"
              >
                Reason
              </th>
              <th
                class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider"
              >
                Date
              </th>
              <th
                class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider"
              >
                Actions
              </th>
            </tr>
          </thead>
          <tbody class="bg-white divide-y divide-gray-200">
            <tr
              v-for="request in filteredRequests"
              :key="request.id"
              class="hover:bg-gray-50 transition-colors"
            >
              <td class="px-6 py-4 whitespace-nowrap">
                <div class="text-sm font-medium text-gray-900">
                  {{ request.name }}
                </div>
              </td>
              <td class="px-6 py-4 whitespace-nowrap">
                <div class="text-sm text-gray-900">{{ request.email }}</div>
              </td>
              <td class="px-6 py-4 whitespace-nowrap">
                <div class="text-sm text-gray-900">
                  {{ request.phone || "N/A" }}
                </div>
              </td>
              <td class="px-6 py-4 whitespace-nowrap">
                <div
                  class="text-sm text-gray-900 max-w-xs truncate"
                  :title="request.reason"
                >
                  {{ request.reason }}
                </div>
              </td>
              <td class="px-6 py-4 whitespace-nowrap">
                <div class="text-sm text-gray-500">
                  {{ formatDate(request.created_at || request.createdAt) }}
                </div>
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm font-medium">
                <button
                  class="text-blue-600 hover:text-blue-800 mr-4"
                  @click="viewRequest(request)"
                >
                  View
                </button>
                <button
                  class="text-red-600 hover:text-red-800"
                  @click="deleteRequest(request.id)"
                >
                  Delete
                </button>
              </td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>

    <!-- View Contact Modal -->
    <div
      v-if="selectedRequest"
      class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50 p-4"
      @click.self="selectedRequest = null"
    >
      <div
        class="bg-white rounded-lg shadow-xl max-w-2xl w-full max-h-[90vh] overflow-y-auto"
      >
        <div class="p-6">
          <!-- Modal Header -->
          <div class="flex justify-between items-start mb-6">
            <h2 class="text-2xl font-bold text-gray-900">
              Contact Message Details
            </h2>
            <button
              class="text-gray-400 hover:text-gray-600 transition-colors"
              @click="selectedRequest = null"
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

          <!-- Contact Information -->
          <div class="space-y-4">
            <div class="grid grid-cols-2 gap-4">
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-1"
                  >Name</label
                >
                <p class="text-gray-900">{{ selectedRequest.name }}</p>
              </div>
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-1"
                  >Email</label
                >
                <p class="text-gray-900">
                  <a
                    :href="`mailto:${selectedRequest.email}`"
                    class="text-blue-600 hover:text-blue-800"
                  >
                    {{ selectedRequest.email }}
                  </a>
                </p>
              </div>
            </div>

            <div class="grid grid-cols-2 gap-4">
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-1"
                  >Phone</label
                >
                <p class="text-gray-900">
                  <a
                    v-if="selectedRequest.phone"
                    :href="`tel:${selectedRequest.phone}`"
                    class="text-blue-600 hover:text-blue-800"
                  >
                    {{ selectedRequest.phone }}
                  </a>
                  <span v-else class="text-gray-400">N/A</span>
                </p>
              </div>
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-1"
                  >Reason</label
                >
                <p class="text-gray-900">{{ selectedRequest.reason }}</p>
              </div>
            </div>

            <div>
              <label class="block text-sm font-medium text-gray-700 mb-1"
                >Message</label
              >
              <div
                class="mt-2 p-4 bg-gray-50 rounded-lg border border-gray-200"
              >
                <p class="text-gray-900 whitespace-pre-wrap">
                  {{ selectedRequest.message }}
                </p>
              </div>
            </div>

            <div>
              <label class="block text-sm font-medium text-gray-700 mb-1"
                >Submitted</label
              >
              <p class="text-gray-500 text-sm">
                {{
                  formatDate(
                    selectedRequest.created_at || selectedRequest.createdAt,
                  )
                }}
              </p>
            </div>
          </div>

          <!-- Modal Footer -->
          <div class="mt-6 flex justify-end space-x-4">
            <button
              class="px-4 py-2 border border-gray-300 rounded-lg text-gray-700 hover:bg-gray-50 transition-colors"
              @click="selectedRequest = null"
            >
              Close
            </button>
            <a
              :href="`mailto:${selectedRequest.email}?subject=Re: Contact Request`"
              class="px-4 py-2 bg-main text-white rounded-lg hover:opacity-90 transition-colors"
            >
              Reply via Email
            </a>
            <button
              class="px-4 py-2 bg-red-600 text-white rounded-lg hover:bg-red-700 transition-colors"
              @click="deleteRequest(selectedRequest.id)"
            >
              Delete
            </button>
          </div>
        </div>
      </div>
    </div>
  </AdminLayout>
</template>

<script>
import { ref, computed, onMounted } from "vue";
import AdminLayout from "@/components/admin/AdminLayout.vue";
import { contactRequestService } from "@/services/contactRequestService";

export default {
  name: "ContactMessages",
  components: {
    AdminLayout,
  },
  setup() {
    const requests = ref([]);
    const selectedRequest = ref(null);
    const loading = ref(false);
    const error = ref("");
    const statusFilter = ref("all");
    const searchQuery = ref("");

    const formatDate = (dateString) => {
      const date = new Date(dateString);
      return date.toLocaleDateString("en-US", {
        year: "numeric",
        month: "short",
        day: "numeric",
        hour: "2-digit",
        minute: "2-digit",
      });
    };

    const loadRequests = async () => {
      loading.value = true;
      error.value = "";
      try {
        const response = await contactRequestService.getContactRequests();
        // API returns { requests: [...] }
        requests.value = response.requests || [];
      } catch (err) {
        console.error("Error loading contact requests:", err);
        error.value = err.message || "Failed to load contact messages";
      } finally {
        loading.value = false;
      }
    };

    const filterContacts = () => {
      // Filtering is handled by computed property
    };

    const filteredRequests = computed(() => {
      let filtered = [...requests.value];

      // Apply status filter (currently just shows all since we don't have read/unread status)
      if (statusFilter.value !== "all") {
        // This can be implemented when read/unread status is added
      }

      // Apply search filter
      if (searchQuery.value) {
        const query = searchQuery.value.toLowerCase();
        filtered = filtered.filter(
          (request) =>
            request.name.toLowerCase().includes(query) ||
            request.email.toLowerCase().includes(query) ||
            (request.phone && request.phone.toLowerCase().includes(query)) ||
            request.reason.toLowerCase().includes(query) ||
            request.message.toLowerCase().includes(query),
        );
      }

      // Sort by date (newest first)
      return filtered.sort(
        (a, b) => new Date(b.createdAt) - new Date(a.createdAt),
      );
    });

    const viewRequest = (request) => {
      selectedRequest.value = request;
    };

    const deleteRequest = async (id) => {
      if (!confirm("Are you sure you want to delete this contact request?")) {
        return;
      }

      try {
        await contactRequestService.deleteContactRequest(id);
        await loadRequests();
        if (selectedRequest.value && selectedRequest.value.id === id) {
          selectedRequest.value = null;
        }
      } catch (err) {
        console.error("Error deleting contact request:", err);
        alert(`Failed to delete contact request: ${err.message}`);
      }
    };

    onMounted(() => {
      loadRequests();
    });

    return {
      requests,
      selectedRequest,
      loading,
      error,
      statusFilter,
      searchQuery,
      filteredRequests,
      formatDate,
      loadRequests,
      filterContacts,
      viewRequest,
      deleteRequest,
    };
  },
};
</script>
