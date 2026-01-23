<template>
  <AdminLayout>
    <div class="max-w-4xl mx-auto">
      <div class="space-y-6">
        <!-- Account Settings -->
        <div class="bg-white rounded-lg shadow p-6">
          <h2 class="text-xl font-semibold text-gray-900 mb-6">
            Account Settings
          </h2>

          <div class="space-y-4">
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">
                Email
              </label>
              <input
                type="email"
                class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-gray-500 focus:border-transparent"
                placeholder="admin@example.com"
              />
            </div>

            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">
                Change Password
              </label>
              <input
                type="password"
                class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-gray-500 focus:border-transparent"
                placeholder="Enter new password"
              />
            </div>

            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">
                Confirm Password
              </label>
              <input
                type="password"
                class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-gray-500 focus:border-transparent"
                placeholder="Confirm new password"
              />
            </div>

            <button
              class="px-6 py-2 bg-gray-600 text-white rounded-lg hover:bg-gray-700 transition-colors"
            >
              Update Account
            </button>
          </div>
        </div>

        <!-- Site Settings -->
        <div class="bg-white rounded-lg shadow p-6">
          <h2 class="text-xl font-semibold text-gray-900 mb-6">
            Site Settings
          </h2>

          <div class="space-y-4">
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">
                Site Name
              </label>
              <input
                type="text"
                class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-gray-500 focus:border-transparent"
                value="Calvary Chapel Phnom Penh"
              />
            </div>

            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">
                Contact Email
              </label>
              <input
                type="email"
                class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-gray-500 focus:border-transparent"
                placeholder="contact@example.com"
              />
            </div>

            <button
              class="px-6 py-2 bg-gray-600 text-white rounded-lg hover:bg-gray-700 transition-colors"
            >
              Save Settings
            </button>
          </div>
        </div>

        <!-- Danger Zone -->
        <div class="bg-white rounded-lg shadow p-6 border-2 border-red-200">
          <h2 class="text-xl font-semibold text-red-900 mb-6">Danger Zone</h2>

          <div class="space-y-4">
            <div
              class="flex items-center justify-between p-4 bg-red-50 rounded-lg"
            >
              <div>
                <h3 class="font-medium text-red-900">
                  Clear All Gallery Images
                </h3>
                <p class="text-sm text-red-700">
                  This will permanently delete all gallery images
                </p>
              </div>
              <button
                class="px-4 py-2 text-red-600 hover:text-red-800 transition-colors"
              >
                Clear Gallery
              </button>
            </div>

            <div
              class="flex items-center justify-between p-4 bg-red-50 rounded-lg"
            >
              <div>
                <h3 class="font-medium text-red-900">Reset All Settings</h3>
                <p class="text-sm text-red-700">
                  Reset all settings to default values
                </p>
              </div>
              <button
                class="px-4 py-2 text-red-600 hover:text-red-800 transition-colors"
              >
                Reset Settings
              </button>
            </div>
          </div>
        </div>
      </div>
    </div>
  </AdminLayout>
</template>

<script>
import AdminLayout from "@/components/admin/AdminLayout.vue";

export default {
  name: "Settings",
  components: {
    AdminLayout,
  },
};
</script>
