<template>
  <div
    class="min-h-screen flex items-center justify-center bg-gray-50 py-12 px-4 sm:px-6 lg:px-8 relative"
  >
    <!-- Watermark -->
    <div class="absolute bottom-6 left-6 flex items-center">
      <div class="w-4 h-4 flex flex-col items-center justify-center mr-2">
        <div class="flex gap-1 mb-0.5">
          <div class="w-1 h-1 bg-[#0089AE] rounded-full"></div>
          <div class="w-1 h-1 bg-[#0089AE] rounded-full"></div>
        </div>
        <div class="w-1 h-1 bg-[#0089AE] rounded-full"></div>
      </div>
      <span class="text-xs font-light text-zinc-500 tracking-wider"
        >ANHELM</span
      >
    </div>

    <div class="max-w-md w-full space-y-6">
      <div class="text-center">
        <div class="mx-auto flex items-center justify-center mb-8">
          <div class="flex items-center">
            <div class="w-8 h-8 flex flex-col items-center justify-center mr-3">
              <div class="flex gap-1.5 mb-1">
                <div class="w-1.5 h-1.5 bg-[#0089AE] rounded-full"></div>
                <div class="w-1.5 h-1.5 bg-[#0089AE] rounded-full"></div>
              </div>
              <div class="w-1.5 h-1.5 bg-[#0089AE] rounded-full"></div>
            </div>
            <h1 class="text-2xl font-light text-zinc-700 tracking-wider">
              ANHELM
            </h1>
          </div>
        </div>
        <div class="mb-2">
          <h2 class="text-lg font-medium text-zinc-600">Welcome back</h2>
          <p class="text-sm text-zinc-500 mt-1">Sign in to your account</p>
        </div>
      </div>

      <form
        class="bg-white py-8 px-6 shadow-lg rounded-2xl border border-gray-100 space-y-6"
        @submit.prevent="handleLogin"
      >
        <div
          v-if="error"
          class="bg-red-50 border border-red-200 rounded-md p-4"
        >
          <div class="flex">
            <div class="flex-shrink-0">
              <svg
                class="h-5 w-5 text-red-400"
                viewBox="0 0 20 20"
                fill="currentColor"
              >
                <path
                  fill-rule="evenodd"
                  d="M10 18a8 8 0 100-16 8 8 0 000 16zM8.707 7.293a1 1 0 00-1.414 1.414L8.586 10l-1.293 1.293a1 1 0 101.414 1.414L10 11.414l1.293 1.293a1 1 0 001.414-1.414L11.414 10l1.293-1.293a1 1 0 00-1.414-1.414L10 8.586 8.707 7.293z"
                  clip-rule="evenodd"
                />
              </svg>
            </div>
            <div class="ml-3">
              <p class="text-sm text-red-800">{{ error }}</p>
            </div>
          </div>
        </div>

        <div>
          <label
            for="email"
            class="block text-sm font-medium text-zinc-700 mb-2"
          >
            Email address
          </label>
          <div class="relative">
            <div
              class="absolute inset-y-0 left-0 pl-3 flex items-center pointer-events-none"
            >
              <svg
                class="h-5 w-5 text-gray-400"
                fill="none"
                stroke="currentColor"
                viewBox="0 0 24 24"
              >
                <path
                  stroke-linecap="round"
                  stroke-linejoin="round"
                  stroke-width="2"
                  d="M16 12a4 4 0 10-8 0 4 4 0 008 0z"
                />
                <path
                  stroke-linecap="round"
                  stroke-linejoin="round"
                  stroke-width="2"
                  d="M12 14v7m0 0h-3m3 0h3"
                />
              </svg>
            </div>
            <input
              id="email"
              v-model="email"
              type="email"
              autocomplete="email"
              required
              class="block w-full pl-10 pr-3 py-3 border border-gray-300 rounded-md text-gray-800 placeholder-gray-500 focus:outline-none focus:ring-1 focus:ring-[#0089AE] focus:border-[#0089AE] transition duration-150 ease-in-out text-sm"
              placeholder="admin@example.com"
            />
          </div>
        </div>

        <div>
          <label
            for="password"
            class="block text-sm font-medium text-zinc-700 mb-2"
          >
            Password
          </label>
          <div class="relative">
            <div
              class="absolute inset-y-0 left-0 pl-3 flex items-center pointer-events-none"
            >
              <svg
                class="h-5 w-5 text-gray-400"
                fill="none"
                stroke="currentColor"
                viewBox="0 0 24 24"
              >
                <path
                  stroke-linecap="round"
                  stroke-linejoin="round"
                  stroke-width="2"
                  d="M12 11c1.657 0 3-1.343 3-3S13.657 5 12 5 9 6.343 9 8s1.343 3 3 3z"
                />
                <path
                  stroke-linecap="round"
                  stroke-linejoin="round"
                  stroke-width="2"
                  d="M5.121 17.804A9 9 0 1118.88 17.8L12 21l-6.879-3.196z"
                />
              </svg>
            </div>
            <input
              id="password"
              v-model="password"
              type="password"
              autocomplete="current-password"
              required
              class="block w-full pl-10 pr-3 py-3 border border-gray-300 rounded-md text-gray-800 placeholder-gray-500 focus:outline-none focus:ring-1 focus:ring-[#0089AE] focus:border-[#0089AE] transition duration-150 ease-in-out text-sm"
              placeholder="Enter your password"
            />
          </div>
        </div>

        <button
          type="submit"
          :disabled="loading"
          class="group relative w-full flex justify-center py-3 px-4 border border-transparent text-sm font-medium rounded-md text-white bg-[#0089AE] hover:bg-[#007A9D] focus:outline-none focus:ring-1 focus:ring-[#0089AE] disabled:opacity-50 disabled:cursor-not-allowed transition duration-150 ease-in-out"
        >
          <span
            v-if="loading"
            class="absolute left-0 inset-y-0 flex items-center pl-3"
          >
            <svg
              class="animate-spin h-5 w-5 text-[#0089AE]"
              fill="none"
              viewBox="0 0 24 24"
            >
              <circle
                class="opacity-25"
                cx="12"
                cy="12"
                r="10"
                stroke="currentColor"
                stroke-width="4"
              ></circle>
              <path
                class="opacity-75"
                fill="currentColor"
                d="M4 12a8 8 0 018-8v8z"
              ></path>
            </svg>
          </span>
          <span v-else class="absolute left-0 inset-y-0 flex items-center pl-3">
            <svg
              class="h-5 w-5 text-[#0089AE] group-hover:text-[#0089AE]/80"
              fill="none"
              stroke="currentColor"
              viewBox="0 0 24 24"
            >
              <path
                stroke-linecap="round"
                stroke-linejoin="round"
                stroke-width="2"
                d="M12 11c1.657 0 3-1.343 3-3S13.657 5 12 5 9 6.343 9 8s1.343 3 3 3z"
              />
              <path
                stroke-linecap="round"
                stroke-linejoin="round"
                stroke-width="2"
                d="M5.121 17.804A9 9 0 1118.88 17.8L12 21l-6.879-3.196z"
              />
            </svg>
          </span>
          {{ loading ? "Signing in..." : "Sign in" }}
        </button>
      </form>

      <div class="mt-6 text-center">
        <router-link to="/" class="text-sm text-[#0089AE] hover:text-[#007A9D] flex items-center justify-center gap-1.5">
          <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 12l2-2m0 0l7-7 7 7M5 10v10a1 1 0 001 1h3m10-11l2 2m-2-2v10a1 1 0 01-1 1h-3m-6 0a1 1 0 001-1v-4a1 1 0 011-1h2a1 1 0 011 1v4a1 1 0 001 1m-6 0h6"></path>
          </svg>
          Back to Website
        </router-link>
      </div>
    </div>
  </div>
</template>

<script>
import { ref } from "vue";
import { useRouter, useRoute } from "vue-router";
import { login } from "@/services/auth";

export default {
  name: "AdminLogin",
  setup() {
    const router = useRouter();
    const route = useRoute();
    const email = ref("");
    const password = ref("");
    const loading = ref(false);
    const error = ref("");

    const handleLogin = async () => {
      error.value = "";
      loading.value = true;
      try {
        const { token, user } = await login(email.value, password.value);
        localStorage.setItem("authToken", token);
        localStorage.setItem("auth_token", token);
        localStorage.setItem("authUser", JSON.stringify(user));

        const redirectTo = route.query.redirect || "/admin/dashboard";
        router.push(redirectTo);
      } catch (err) {
        error.value = err.message || "Unable to sign in";
      } finally {
        loading.value = false;
      }
    };

    return {
      email,
      password,
      loading,
      error,
      handleLogin,
    };
  },
};
</script>
