<template>
  <AdminLayout>
    <!-- Page Selection -->
    <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
      <div
        v-for="page in pages"
        :key="page.id"
        class="bg-white rounded-lg shadow hover:shadow-lg transition-shadow p-6 cursor-pointer"
        @click="selectedPage = page"
      >
        <div class="flex items-center mb-4">
          <div class="p-3 bg-secondary rounded-lg mr-4">
            <svg
              class="w-6 h-6 text-main"
              fill="none"
              stroke="currentColor"
              viewBox="0 0 24 24"
            >
              <path
                stroke-linecap="round"
                stroke-linejoin="round"
                stroke-width="2"
                d="M9 12h6m-6 4h6m2 5H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z"
              ></path>
            </svg>
          </div>
          <h3 class="text-lg font-semibold text-gray-900">{{ page.name }}</h3>
        </div>
        <p class="text-sm text-gray-600">{{ page.description }}</p>
      </div>
    </div>

    <!-- Edit Modal -->
    <div
      v-if="selectedPage"
      class="fixed inset-0 bg-black bg-opacity-50 z-50 flex items-center justify-center p-4"
      @click="selectedPage = null"
    >
      <div
        class="bg-white rounded-lg shadow-xl max-w-4xl w-full max-h-[90vh] overflow-y-auto"
        @click.stop
      >
        <div class="p-6 border-b border-gray-200">
          <div class="flex justify-between items-center">
            <h3 class="text-xl font-semibold text-gray-900">
              Edit {{ selectedPage.name }}
            </h3>
            <button
              class="text-gray-400 hover:text-gray-600"
              @click="selectedPage = null"
            >
              <svg
                class="w-6 h-6"
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
        </div>

        <div class="p-6 space-y-6">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">
              Page Title
            </label>
            <input
              type="text"
              class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-[#23D3EE] focus:border-transparent"
              :placeholder="selectedPage.name"
            />
          </div>

          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">
              Content
            </label>
            <textarea
              rows="10"
              class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-[#23D3EE] focus:border-transparent"
              placeholder="Page content..."
            ></textarea>
          </div>

          <div class="flex justify-end space-x-4 pt-4 border-t border-gray-200">
            <button
              class="px-6 py-2 border border-gray-300 text-gray-700 rounded-lg hover:bg-gray-50 transition-colors"
              @click="selectedPage = null"
            >
              Cancel
            </button>
            <button
              class="px-6 py-2 bg-[#23D3EE] text-white rounded-lg hover:bg-[#1FC5D9] transition-colors"
            >
              Save Changes
            </button>
          </div>
        </div>
      </div>
    </div>
  </AdminLayout>
</template>

<script>
import { ref } from "vue";
import AdminLayout from "@/components/admin/AdminLayout.vue";

export default {
  name: "ContentManagement",
  components: {
    AdminLayout,
  },
  setup() {
    const selectedPage = ref(null);

    const pages = [
      { id: 1, name: "About Page", description: "Edit about us content" },
      {
        id: 2,
        name: "Leadership",
        description: "Update leadership team information",
      },
      {
        id: 3,
        name: "Ministries",
        description: "Edit ministries descriptions",
      },
      {
        id: 4,
        name: "What We Believe",
        description: "Update statement of faith",
      },
      { id: 5, name: "Cambodia", description: "Edit Cambodia page content" },
      { id: 6, name: "Home Page", description: "Edit home page content" },
    ];

    return {
      selectedPage,
      pages,
    };
  },
};
</script>
