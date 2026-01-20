<template>
  <AdminLayout>
    <!-- Contact Requests Count -->
    <div class="mb-6">
      <div class="flex items-center justify-between mb-4">
        <div>
          <h1 class="text-2xl font-semibold text-gray-900">Contact Requests</h1>
          <p class="text-sm text-gray-600 mt-1">
            Total: <span class="font-semibold text-main">{{ requests.length }}</span> contact request{{ requests.length !== 1 ? 's' : '' }}
          </p>
        </div>
      </div>
    </div>

    <!-- Filters -->
    <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6 mb-6">
      <div class="flex flex-wrap items-center gap-4">
        <select
          v-model="statusFilter"
          class="px-4 py-2 border border-gray-300 rounded-lg text-sm focus:outline-none focus:ring-2 focus:ring-main focus:border-transparent"
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
          class="px-4 py-2 border border-gray-300 rounded-lg text-sm flex-1 min-w-64 focus:outline-none focus:ring-2 focus:ring-main focus:border-transparent"
          @input="filterContacts"
        />
        <button
          class="px-4 py-2 bg-main text-white rounded-lg hover:opacity-90 transition-colors text-sm font-medium shadow-sm"
          @click="loadRequests"
        >
          Refresh
        </button>
      </div>
    </div>

    <!-- Messages List -->
    <div>
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
          @click="loadRequests"
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
      <div v-else class="bg-white rounded-lg shadow-sm border border-gray-200 overflow-hidden">
        <div class="overflow-x-auto">
          <table class="min-w-full divide-y divide-gray-200">
            <thead class="bg-gray-50">
              <tr>
                <th
                  class="px-6 py-4 text-left text-xs font-semibold text-gray-700 uppercase tracking-wider"
                >
                  Name
                </th>
                <th
                  class="px-6 py-4 text-left text-xs font-semibold text-gray-700 uppercase tracking-wider"
                >
                  Email
                </th>
                <th
                  class="px-6 py-4 text-left text-xs font-semibold text-gray-700 uppercase tracking-wider"
                >
                  Phone
                </th>
                <th
                  class="px-6 py-4 text-left text-xs font-semibold text-gray-700 uppercase tracking-wider"
                >
                  Reason
                </th>
                <th
                  class="px-6 py-4 text-left text-xs font-semibold text-gray-700 uppercase tracking-wider"
                >
                  Date
                </th>
                <th
                  class="px-6 py-4 text-right text-xs font-semibold text-gray-700 uppercase tracking-wider"
                >
                  Actions
                </th>
              </tr>
            </thead>
            <tbody class="bg-white divide-y divide-gray-200">
              <tr
                v-for="request in filteredRequests"
                :key="request.id"
                class="hover:bg-gray-50 transition-colors cursor-pointer"
                @click="viewRequest(request)"
              >
                <td class="px-6 py-4 whitespace-nowrap">
                  <div class="text-sm font-semibold text-gray-900">
                    {{ request.name }}
                  </div>
                </td>
                <td class="px-6 py-4">
                  <div class="text-sm text-gray-900">
                    <a
                      :href="`mailto:${request.email}`"
                      class="text-blue-600 hover:text-blue-800 hover:underline"
                      @click.stop
                    >
                      {{ request.email }}
                    </a>
                  </div>
                </td>
                <td class="px-6 py-4 whitespace-nowrap">
                  <div class="text-sm text-gray-900">
                    <a
                      v-if="request.phone"
                      :href="`tel:${request.phone}`"
                      class="text-blue-600 hover:text-blue-800 hover:underline"
                      @click.stop
                    >
                      {{ request.phone }}
                    </a>
                    <span v-else class="text-gray-400">N/A</span>
                  </div>
                </td>
                <td class="px-6 py-4">
                  <div class="text-sm text-gray-900 max-w-xs">
                    {{ request.reason }}
                  </div>
                </td>
                <td class="px-6 py-4 whitespace-nowrap">
                  <div class="text-sm text-gray-600">
                    {{ formatDate(request.created_at || request.createdAt) }}
                  </div>
                </td>
                <td class="px-6 py-4 whitespace-nowrap text-right text-sm font-medium" @click.stop>
                  <button
                    class="px-4 py-2 bg-red-600 text-white rounded-lg hover:bg-red-700 transition-colors font-medium shadow-sm"
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
    </div>

    <!-- View Contact Modal -->
    <div
      v-if="selectedRequest"
      class="fixed inset-0 bg-black bg-opacity-75 backdrop-blur-sm flex items-center justify-center z-50 p-4"
      @click.self="selectedRequest = null"
    >
      <div
        class="bg-white rounded-xl shadow-2xl border-2 border-gray-200 max-w-2xl w-full max-h-[90vh] overflow-y-auto transform transition-all"
      >
        <div class="p-8">
          <!-- Modal Header -->
          <div class="flex justify-between items-start mb-6 pb-4 border-b border-gray-200">
            <h2 class="text-3xl font-bold text-gray-900">
              Contact Message Details
            </h2>
            <button
              class="text-gray-500 hover:text-gray-700 transition-colors p-1 hover:bg-gray-100 rounded-full"
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
          <div class="space-y-6">
            <div class="grid grid-cols-2 gap-6">
              <div class="bg-gray-50 p-4 rounded-lg border border-gray-200">
                <label class="block text-xs font-semibold text-gray-600 uppercase tracking-wide mb-2"
                  >Name</label
                >
                <p class="text-base font-medium text-gray-900">{{ selectedRequest.name }}</p>
              </div>
              <div class="bg-gray-50 p-4 rounded-lg border border-gray-200">
                <label class="block text-xs font-semibold text-gray-600 uppercase tracking-wide mb-2"
                  >Email</label
                >
                <p class="text-base text-gray-900">
                  <a
                    :href="`mailto:${selectedRequest.email}`"
                    class="text-blue-600 hover:text-blue-800 font-medium hover:underline"
                  >
                    {{ selectedRequest.email }}
                  </a>
                </p>
              </div>
            </div>

            <div class="grid grid-cols-2 gap-6">
              <div class="bg-gray-50 p-4 rounded-lg border border-gray-200">
                <label class="block text-xs font-semibold text-gray-600 uppercase tracking-wide mb-2"
                  >Phone</label
                >
                <p class="text-base text-gray-900">
                  <a
                    v-if="selectedRequest.phone"
                    :href="`tel:${selectedRequest.phone}`"
                    class="text-blue-600 hover:text-blue-800 font-medium hover:underline"
                  >
                    {{ selectedRequest.phone }}
                  </a>
                  <span v-else class="text-gray-500">N/A</span>
                </p>
              </div>
              <div class="bg-gray-50 p-4 rounded-lg border border-gray-200">
                <label class="block text-xs font-semibold text-gray-600 uppercase tracking-wide mb-2"
                  >Reason</label
                >
                <p class="text-base font-medium text-gray-900">{{ selectedRequest.reason }}</p>
              </div>
            </div>

            <div>
              <label class="block text-xs font-semibold text-gray-600 uppercase tracking-wide mb-2"
                >Message</label
              >
              <div
                class="mt-2 p-5 bg-gray-50 rounded-lg border-2 border-gray-300"
              >
                <p class="text-gray-900 whitespace-pre-wrap text-base leading-relaxed">
                  {{ selectedRequest.message }}
                </p>
              </div>
            </div>

            <div>
              <label class="block text-xs font-semibold text-gray-600 uppercase tracking-wide mb-2"
                >Submitted</label
              >
              <p class="text-sm font-medium text-gray-700">
                {{
                  formatDate(
                    selectedRequest.created_at || selectedRequest.createdAt,
                  )
                }}
              </p>
            </div>
          </div>

          <!-- Modal Footer -->
          <div class="mt-8 pt-6 border-t border-gray-200 flex justify-end space-x-3">
            <button
              class="px-6 py-2.5 border-2 border-gray-300 rounded-lg text-gray-700 hover:bg-gray-50 transition-colors font-semibold shadow-sm"
              @click="selectedRequest = null"
            >
              Close
            </button>
            <a
              :href="`mailto:${selectedRequest.email}?subject=Re: Contact Request`"
              class="px-6 py-2.5 bg-main text-white rounded-lg hover:opacity-90 transition-colors font-semibold shadow-md"
            >
              Reply via Email
            </a>
            <button
              class="px-6 py-2.5 bg-red-600 text-white rounded-lg hover:bg-red-700 transition-colors font-semibold shadow-md"
              @click="deleteRequest(selectedRequest.id)"
            >
              Delete
            </button>
          </div>
        </div>
      </div>
    </div>

    <!-- Confirmation Modal -->
    <ConfirmationModal
      :show="showConfirmModal"
      title="Delete Contact Request"
      message="Are you sure you want to delete this contact request? This action cannot be undone."
      confirm-text="Delete Request"
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
import { contactRequestService } from "@/services/contactRequestService";

export default {
  name: "ContactMessages",
  components: {
    AdminLayout,
    ConfirmationModal,
    ErrorModal,
  },
  setup() {
    const requests = ref([]);
    const selectedRequest = ref(null);
    const selectedItems = ref([]);
    const loading = ref(false);
    const error = ref("");
    const statusFilter = ref("all");
    const searchQuery = ref("");
    const showConfirmModal = ref(false);
    const showErrorModal = ref(false);
    const errorMessage = ref("");
    const requestToDelete = ref(null);

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
        (a, b) => 
          new Date(b.created_at || b.createdAt) - 
          new Date(a.created_at || a.createdAt),
      );
    });

    const viewRequest = (request) => {
      selectedRequest.value = request;
    };

    const deleteRequest = (id) => {
      requestToDelete.value = id;
      showConfirmModal.value = true;
    };

    const confirmDelete = async () => {
      if (!requestToDelete.value) return;

      const id = requestToDelete.value;
      showConfirmModal.value = false;

      try {
        await contactRequestService.deleteContactRequest(id);
        await loadRequests();
        if (selectedRequest.value && selectedRequest.value.id === id) {
          selectedRequest.value = null;
        }
        requestToDelete.value = null;
      } catch (err) {
        console.error("Error deleting contact request:", err);
        errorMessage.value = `Failed to delete contact request: ${err.message}`;
        showErrorModal.value = true;
        requestToDelete.value = null;
      }
    };

    const cancelDelete = () => {
      showConfirmModal.value = false;
      requestToDelete.value = null;
    };

    const closeErrorModal = () => {
      showErrorModal.value = false;
      errorMessage.value = "";
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
