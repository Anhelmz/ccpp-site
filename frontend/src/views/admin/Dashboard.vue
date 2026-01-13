<template>
  <AdminLayout>
    <!-- Quick Stats -->
    <div class="grid grid-cols-1 md:grid-cols-3 gap-6 mb-8">
      <div class="bg-white rounded-xl shadow-lg p-6 border-l-4 border-main">
        <div class="flex items-center justify-between">
          <div>
            <p
              class="text-sm font-medium text-gray-600 uppercase tracking-wide"
            >
              Total Images
            </p>
            <p class="text-3xl font-bold text-gray-900 mt-2">0</p>
            <p class="text-xs text-gray-500 mt-1">Across all galleries</p>
          </div>
          <div class="p-4 bg-main rounded-xl shadow-lg">
            <svg
              class="w-8 h-8 text-white"
              fill="none"
              stroke="currentColor"
              viewBox="0 0 24 24"
            >
              <path
                stroke-linecap="round"
                stroke-linejoin="round"
                stroke-width="2"
                d="M4 16l4.586-4.586a2 2 0 012.828 0L16 16m-2-2l1.586-1.586a2 2 0 012.828 0L20 14m-6-6h.01M6 20h12a2 2 0 002-2V6a2 2 0 00-2-2H6a2 2 0 00-2 2v12a2 2 0 002 2z"
              ></path>
            </svg>
          </div>
        </div>
      </div>

      <div
        class="bg-white rounded-xl shadow-lg p-6 border-l-4 border-green-600"
      >
        <div class="flex items-center justify-between">
          <div>
            <p
              class="text-sm font-medium text-gray-600 uppercase tracking-wide"
            >
              Upcoming Events
            </p>
            <p class="text-3xl font-bold text-gray-900 mt-2">0</p>
            <p class="text-xs text-gray-500 mt-1">Scheduled events</p>
          </div>
          <div
            class="p-4 bg-gradient-to-br from-green-500 to-green-600 rounded-xl shadow-lg"
          >
            <svg
              class="w-8 h-8 text-white"
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
          </div>
        </div>
      </div>

      <div class="bg-white rounded-xl shadow-lg p-6 border-l-4 border-blue-600">
        <div class="flex items-center justify-between">
          <div>
            <p
              class="text-sm font-medium text-gray-600 uppercase tracking-wide"
            >
              Contact Messages
            </p>
            <p class="text-3xl font-bold text-gray-900 mt-2">0</p>
            <p class="text-xs text-gray-500 mt-1">Unread messages</p>
          </div>
          <div
            class="p-4 bg-gradient-to-br from-blue-500 to-blue-600 rounded-xl shadow-lg"
          >
            <svg
              class="w-8 h-8 text-white"
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
          </div>
        </div>
      </div>
    </div>

    <!-- Management Cards -->
    <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
      <!-- Gallery Management -->
      <div
        class="bg-white rounded-xl shadow-lg hover:shadow-xl transition-all duration-300 p-6 border border-gray-100 hover:border-secondary group"
      >
        <div class="flex items-center mb-4">
          <div
            class="p-3 bg-main rounded-xl mr-4 group-hover:scale-110 transition-transform"
          >
            <svg
              class="w-6 h-6 text-white"
              fill="none"
              stroke="currentColor"
              viewBox="0 0 24 24"
            >
              <path
                stroke-linecap="round"
                stroke-linejoin="round"
                stroke-width="2"
                d="M4 16l4.586-4.586a2 2 0 012.828 0L16 16m-2-2l1.586-1.586a2 2 0 012.828 0L20 14m-6-6h.01M6 20h12a2 2 0 002-2V6a2 2 0 00-2-2H6a2 2 0 00-2 2v12a2 2 0 002 2z"
              ></path>
            </svg>
          </div>
          <h2 class="text-xl font-bold text-gray-900">Gallery</h2>
        </div>
        <p class="text-gray-600 mb-4 text-sm">
          Upload and manage images for Youth Ministry, Grace Church, and
          Outreaches galleries.
        </p>
        <router-link
          to="/admin/gallery"
          class="inline-flex items-center px-4 py-2 bg-main text-white rounded-lg hover:opacity-90 transition-colors font-medium text-sm shadow-md"
        >
          Manage Gallery
          <svg
            class="ml-2 w-4 h-4"
            fill="none"
            stroke="currentColor"
            viewBox="0 0 24 24"
          >
            <path
              stroke-linecap="round"
              stroke-linejoin="round"
              stroke-width="2"
              d="M9 5l7 7-7 7"
            ></path>
          </svg>
        </router-link>
      </div>

      <!-- Events Management -->
      <div
        class="bg-white rounded-xl shadow-lg hover:shadow-xl transition-all duration-300 p-6 border border-gray-100 hover:border-green-300 group"
      >
        <div class="flex items-center mb-4">
          <div
            class="p-3 bg-gradient-to-br from-green-500 to-green-600 rounded-xl mr-4 group-hover:scale-110 transition-transform"
          >
            <svg
              class="w-6 h-6 text-white"
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
          </div>
          <h2 class="text-xl font-bold text-gray-900">Events</h2>
        </div>
        <p class="text-gray-600 mb-4 text-sm">
          Create, edit, and manage church events and calendar entries.
        </p>
        <router-link
          to="/admin/events"
          class="inline-flex items-center px-4 py-2 bg-green-600 text-white rounded-lg hover:bg-green-700 transition-colors font-medium text-sm shadow-md"
        >
          Manage Events
          <svg
            class="ml-2 w-4 h-4"
            fill="none"
            stroke="currentColor"
            viewBox="0 0 24 24"
          >
            <path
              stroke-linecap="round"
              stroke-linejoin="round"
              stroke-width="2"
              d="M9 5l7 7-7 7"
            ></path>
          </svg>
        </router-link>
      </div>

      <!-- Content Management -->
      <div
        class="bg-white rounded-xl shadow-lg hover:shadow-xl transition-all duration-300 p-6 border border-gray-100 hover:border-blue-300 group"
      >
        <div class="flex items-center mb-4">
          <div
            class="p-3 bg-gradient-to-br from-blue-500 to-blue-600 rounded-xl mr-4 group-hover:scale-110 transition-transform"
          >
            <svg
              class="w-6 h-6 text-white"
              fill="none"
              stroke="currentColor"
              viewBox="0 0 24 24"
            >
              <path
                stroke-linecap="round"
                stroke-linejoin="round"
                stroke-width="2"
                d="M11 5H6a2 2 0 00-2 2v11a2 2 0 002 2h11a2 2 0 002-2v-5m-1.414-9.414a2 2 0 112.828 2.828L11.828 15H9v-2.828l8.586-8.586z"
              ></path>
            </svg>
          </div>
          <h2 class="text-xl font-bold text-gray-900">Content</h2>
        </div>
        <p class="text-gray-600 mb-4 text-sm">
          Edit page content, leadership information, and other site content.
        </p>
        <router-link
          to="/admin/content"
          class="inline-flex items-center px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700 transition-colors font-medium text-sm shadow-md"
        >
          Manage Content
          <svg
            class="ml-2 w-4 h-4"
            fill="none"
            stroke="currentColor"
            viewBox="0 0 24 24"
          >
            <path
              stroke-linecap="round"
              stroke-linejoin="round"
              stroke-width="2"
              d="M9 5l7 7-7 7"
            ></path>
          </svg>
        </router-link>
      </div>

      <!-- Contact Messages -->
      <div
        class="bg-white rounded-xl shadow-lg hover:shadow-xl transition-all duration-300 p-6 border border-gray-100 hover:border-purple-300 group"
      >
        <div class="flex items-center mb-4">
          <div
            class="p-3 bg-gradient-to-br from-purple-500 to-purple-600 rounded-xl mr-4 group-hover:scale-110 transition-transform"
          >
            <svg
              class="w-6 h-6 text-white"
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
          </div>
          <h2 class="text-xl font-bold text-gray-900">Messages</h2>
        </div>
        <p class="text-gray-600 mb-4 text-sm">
          View and respond to contact form submissions.
        </p>
        <router-link
          to="/admin/contacts"
          class="inline-flex items-center px-4 py-2 bg-purple-600 text-white rounded-lg hover:bg-purple-700 transition-colors font-medium text-sm shadow-md"
        >
          View Messages
          <svg
            class="ml-2 w-4 h-4"
            fill="none"
            stroke="currentColor"
            viewBox="0 0 24 24"
          >
            <path
              stroke-linecap="round"
              stroke-linejoin="round"
              stroke-width="2"
              d="M9 5l7 7-7 7"
            ></path>
          </svg>
        </router-link>
      </div>

      <!-- Settings -->
      <div
        class="bg-white rounded-xl shadow-lg hover:shadow-xl transition-all duration-300 p-6 border border-gray-100 hover:border-gray-300 group"
      >
        <div class="flex items-center mb-4">
          <div
            class="p-3 bg-gradient-to-br from-gray-500 to-gray-600 rounded-xl mr-4 group-hover:scale-110 transition-transform"
          >
            <svg
              class="w-6 h-6 text-white"
              fill="none"
              stroke="currentColor"
              viewBox="0 0 24 24"
            >
              <path
                stroke-linecap="round"
                stroke-linejoin="round"
                stroke-width="2"
                d="M10.325 4.317c.426-1.756 2.924-1.756 3.35 0a1.724 1.724 0 002.573 1.066c1.543-.94 3.31.826 2.37 2.37a1.724 1.724 0 001.065 2.572c1.756.426 1.756 2.924 0 3.35a1.724 1.724 0 00-1.066 2.573c.94 1.543-.826 3.31-2.37 2.37a1.724 1.724 0 00-2.572 1.065c-.426 1.756-2.924 1.756-3.35 0a1.724 1.724 0 00-2.573-1.066c-1.543.94-3.31-.826-2.37-2.37a1.724 1.724 0 00-1.065-2.572c-1.756-.426-1.756-2.924 0-3.35a1.724 1.724 0 001.066-2.573c-.94-1.543.826-3.31 2.37-2.37.996.608 2.296.07 2.572-1.065z"
              ></path>
              <path
                stroke-linecap="round"
                stroke-linejoin="round"
                stroke-width="2"
                d="M15 12a3 3 0 11-6 0 3 3 0 016 0z"
              ></path>
            </svg>
          </div>
          <h2 class="text-xl font-bold text-gray-900">Settings</h2>
        </div>
        <p class="text-gray-600 mb-4 text-sm">
          Configure admin settings and preferences.
        </p>
        <router-link
          to="/admin/settings"
          class="inline-flex items-center px-4 py-2 bg-gray-600 text-white rounded-lg hover:bg-gray-700 transition-colors font-medium text-sm shadow-md"
        >
          Open Settings
          <svg
            class="ml-2 w-4 h-4"
            fill="none"
            stroke="currentColor"
            viewBox="0 0 24 24"
          >
            <path
              stroke-linecap="round"
              stroke-linejoin="round"
              stroke-width="2"
              d="M9 5l7 7-7 7"
            ></path>
          </svg>
        </router-link>
      </div>
    </div>
  </AdminLayout>
</template>

<script>
import AdminLayout from "@/components/admin/AdminLayout.vue";

export default {
  name: "AdminDashboard",
  components: {
    AdminLayout,
  },
};
</script>
