<template>
  <AdminLayout>
    <!-- Page Header -->
    <div class="mb-8">
      <h1 class="text-2xl font-semibold text-zinc-900">Contact Requests</h1>
      <p class="mt-1 text-sm text-zinc-600">View and manage contact form submissions</p>
    </div>


    <!-- Messages List -->
    <div v-if="!selectedRequest">
      <!-- Loading State -->
      <div v-if="loading" class="flex items-center justify-center py-12">
        <div class="animate-spin rounded-full h-8 w-8 border-b-2 border-[#23D3EE]"></div>
      </div>

      <!-- Error State -->
      <div v-else-if="error" class="bg-red-50 border border-red-200 rounded-lg p-4 mb-6">
        <p class="text-sm text-red-800">{{ error }}</p>
      </div>

      <!-- Contact Requests List -->
      <div v-else-if="filteredRequests.length > 0" class="bg-white border border-gray-200 rounded-lg overflow-hidden">
        <div class="overflow-x-auto">
          <table class="min-w-full divide-y divide-gray-200">
            <thead class="bg-gray-50">
              <tr>
                <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">Name</th>
                <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">Email</th>
                <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">Phone</th>
                <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">Submitted</th>
                <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">Actions</th>
              </tr>
            </thead>
            <tbody class="bg-white divide-y divide-gray-200">
              <tr 
                v-for="request in filteredRequests" 
                :key="request.id" 
                @click="viewRequest(request)"
                class="hover:bg-gray-50 cursor-pointer transition-colors"
              >
                <td class="px-6 py-4 whitespace-nowrap text-sm font-medium text-gray-900">
                  {{ request.name }}
                </td>
                <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-600">
                  <a 
                    :href="`mailto:${request.email}`" 
                    @click.stop
                    class="text-[#23D3EE] hover:text-[#1FC5D9]"
                  >
                    {{ request.email }}
                  </a>
                </td>
                <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-600">
                  <a 
                    v-if="request.phone"
                    :href="`tel:${request.phone}`" 
                    @click.stop
                    class="text-[#23D3EE] hover:text-[#1FC5D9]"
                  >
                    {{ request.phone }}
                  </a>
                  <span v-else class="text-gray-400">N/A</span>
                </td>
                <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
                  {{ formatDate(request.created_at || request.createdAt) }}
                </td>
                <td class="px-6 py-4 whitespace-nowrap text-right text-sm font-medium" @click.stop>
                  <button
                    class="text-red-600 hover:text-red-800 transition-colors font-medium"
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

      <!-- Empty State -->
      <div v-else class="bg-white p-6 border border-gray-200 rounded-lg">
        <div class="text-center py-12">
          <svg
            class="mx-auto h-16 w-16 text-zinc-300 mb-4"
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
          <h3 class="text-lg font-medium text-zinc-900 mb-2">No Contact Requests</h3>
          <p class="text-sm text-zinc-600">Contact form submissions will appear here</p>
        </div>
      </div>
    </div>

    <!-- View Contact Details Page -->
    <div v-if="selectedRequest" class="max-w-2xl mx-auto">
      <!-- Page Header -->
      <div class="flex justify-between items-start mb-4 pb-3 border-b border-gray-200">
        <h2 class="text-xl font-semibold text-zinc-900">
          Contact Message Details
        </h2>
        <button
          class="text-gray-500 hover:text-gray-700 transition-colors p-1 hover:bg-gray-100 rounded-full"
          @click="selectedRequest = null"
        >
          <svg
            class="h-5 w-5"
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
      <div class="space-y-3">
        <div class="grid grid-cols-2 gap-3">
          <div class="bg-gray-50 p-3 rounded-lg border border-gray-200">
            <label class="block text-xs font-semibold text-gray-600 uppercase tracking-wide mb-1"
              >Name</label
            >
            <p class="text-sm font-medium text-gray-900">{{ selectedRequest.name }}</p>
          </div>
          <div class="bg-gray-50 p-3 rounded-lg border border-gray-200">
            <label class="block text-xs font-semibold text-gray-600 uppercase tracking-wide mb-1"
              >Email</label
            >
            <p class="text-sm text-gray-900">
              <a
                :href="`mailto:${selectedRequest.email}`"
                class="text-[#23D3EE] hover:text-[#1FC5D9] font-medium hover:underline"
              >
                {{ selectedRequest.email }}
              </a>
            </p>
          </div>
        </div>

        <div class="grid grid-cols-2 gap-3">
          <div class="bg-gray-50 p-3 rounded-lg border border-gray-200">
            <label class="block text-xs font-semibold text-gray-600 uppercase tracking-wide mb-1"
              >Phone</label
            >
            <p class="text-sm text-gray-900">
              <a
                v-if="selectedRequest.phone"
                :href="`tel:${selectedRequest.phone}`"
                class="text-[#23D3EE] hover:text-[#1FC5D9] font-medium hover:underline"
              >
                {{ selectedRequest.phone }}
              </a>
              <span v-else class="text-gray-500">N/A</span>
            </p>
          </div>
          <div class="bg-gray-50 p-3 rounded-lg border border-gray-200">
            <label class="block text-xs font-semibold text-gray-600 uppercase tracking-wide mb-1"
              >Reason</label
            >
            <p class="text-sm font-medium text-gray-900">{{ selectedRequest.reason }}</p>
          </div>
        </div>

        <div>
          <label class="block text-xs font-semibold text-gray-600 uppercase tracking-wide mb-1"
            >Message</label
          >
          <div
            class="mt-1 p-3 bg-gray-50 rounded-lg border border-gray-200"
          >
            <p class="text-gray-900 whitespace-pre-wrap text-sm leading-relaxed">
              {{ selectedRequest.message }}
            </p>
          </div>
        </div>

        <div>
          <label class="block text-xs font-semibold text-gray-600 uppercase tracking-wide mb-1"
            >Submitted</label
          >
          <p class="text-xs font-medium text-gray-700">
            {{
              formatDate(
                selectedRequest.created_at || selectedRequest.createdAt,
              )
            }}
          </p>
        </div>
      </div>

      <!-- Footer Actions -->
      <div class="mt-4 pt-3 border-t border-gray-200 flex justify-end space-x-2">
        <button
          class="px-3 py-1.5 border border-gray-300 rounded-md text-sm text-gray-700 hover:bg-gray-50 transition-colors font-medium"
          @click="selectedRequest = null"
        >
          Close
        </button>
        <a
          :href="`mailto:${selectedRequest.email}?subject=Re: Contact Request`"
          class="px-3 py-1.5 bg-[#23D3EE] text-white rounded-md hover:bg-[#1FC5D9] transition-colors font-medium text-sm"
        >
          Reply via Email
        </a>
        <button
          class="px-3 py-1.5 text-sm text-red-600 hover:text-red-800 transition-colors font-medium"
          @click="deleteRequest(selectedRequest.id)"
        >
          Delete
        </button>
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
      if (!dateString) return 'N/A'
      const date = new Date(dateString);
      return date.toLocaleDateString("en-US", {
        year: "numeric",
        month: "short",
        day: "numeric",
        hour: "2-digit",
        minute: "2-digit",
      });
    };

    const truncateMessage = (message, maxLength) => {
      if (!message) return ''
      if (message.length <= maxLength) return message
      return message.substring(0, maxLength) + '...'
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
      truncateMessage,
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

<style>
</style>
