<template>
  <AdminLayout>
    <!-- Page Header -->
    <div class="mb-8">
      <h1 class="text-2xl font-semibold text-zinc-900">Dashboard</h1>
      <p class="mt-1 text-sm text-zinc-600">Welcome back, {{ userEmail }}</p>
    </div>
    
    <!-- Coming Soon Message -->
    <div class="bg-white border border-gray-200 rounded-lg p-12">
      <div class="text-center">
        <div class="flex justify-center mb-6">
          <div class="w-16 h-16 bg-[#0089AE]/10 rounded-full flex items-center justify-center">
            <svg
              class="w-8 h-8 text-[#0089AE]"
              fill="none"
              stroke="currentColor"
              viewBox="0 0 24 24"
            >
              <path
                stroke-linecap="round"
                stroke-linejoin="round"
                stroke-width="2"
                d="M9 19v-6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2a2 2 0 002-2zm0 0V9a2 2 0 012-2h2a2 2 0 012 2v10m-6 0a2 2 0 002 2h2a2 2 0 002-2m0 0V5a2 2 0 012-2h2a2 2 0 012 2v14a2 2 0 01-2 2h-2a2 2 0 01-2-2z"
              />
            </svg>
          </div>
        </div>
        <h2 class="text-xl font-semibold text-zinc-900 mb-3">Analytics Dashboard</h2>
        <p class="text-zinc-600 max-w-md mx-auto">
          A metrics dashboard is coming soon.
        </p>
      </div>
    </div>
  </AdminLayout>
</template>

<script>
import { ref, computed, onMounted } from "vue";
import AdminLayout from "@/components/admin/AdminLayout.vue";

export default {
  name: "AdminDashboard",
  components: {
    AdminLayout,
  },
  setup() {
    const user = ref(null);
    const userEmail = computed(() => user.value?.email || "admin@ccpp.org");

    onMounted(() => {
      const saved = localStorage.getItem("authUser");
      if (saved) {
        try {
          user.value = JSON.parse(saved);
        } catch (err) {
          console.error("Error parsing user data:", err);
        }
      }
    });

    return {
      user,
      userEmail,
    };
  },
};
</script>
