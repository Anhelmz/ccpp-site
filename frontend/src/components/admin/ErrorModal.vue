<template>
  <div
    v-if="show"
    class="fixed inset-0 z-[60] flex items-center justify-center bg-black bg-opacity-50 backdrop-blur-sm p-4"
    @click.self="handleClose"
  >
    <div
      class="bg-white dark:bg-[#0c0f14] rounded-xl shadow-2xl max-w-md w-full p-6 border-2 border-red-500"
    >
      <div class="flex items-center justify-between mb-4">
        <div class="flex items-center gap-3">
          <div class="flex-shrink-0 w-10 h-10 bg-red-100 dark:bg-red-900/30 rounded-full flex items-center justify-center">
            <svg
              class="w-6 h-6 text-red-600 dark:text-red-400"
              fill="none"
              stroke="currentColor"
              viewBox="0 0 24 24"
            >
              <path
                stroke-linecap="round"
                stroke-linejoin="round"
                stroke-width="2"
                d="M12 8v4m0 4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z"
              />
            </svg>
          </div>
          <h3 class="text-xl font-bold text-gray-900 dark:text-white">
            {{ title }}
          </h3>
        </div>
        <button
          @click="handleClose"
          class="text-gray-400 hover:text-gray-600 dark:hover:text-gray-200 transition-colors"
          aria-label="Close"
        >
          <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path
              stroke-linecap="round"
              stroke-linejoin="round"
              stroke-width="2"
              d="M6 18L18 6M6 6l12 12"
            />
          </svg>
        </button>
      </div>

      <p class="text-gray-700 dark:text-gray-300 mb-6">{{ message }}</p>

      <div class="flex justify-end">
        <button
          @click="handleClose"
          class="px-6 py-2 bg-red-600 text-white rounded-lg hover:bg-red-700 transition-colors font-medium"
        >
          OK
        </button>
      </div>
    </div>
  </div>
</template>

<script>
import { watch, onUnmounted } from "vue";

export default {
  name: "ErrorModal",
  props: {
    show: {
      type: Boolean,
      default: false,
    },
    title: {
      type: String,
      default: "Error",
    },
    message: {
      type: String,
      required: true,
    },
  },
  emits: ["close"],
  setup(props, { emit }) {
    const handleClose = () => {
      emit("close");
    };

    let escapeHandler = null;

    // Close on Escape key
    watch(
      () => props.show,
      (isShowing) => {
        if (isShowing) {
          escapeHandler = (e) => {
            if (e.key === "Escape") {
              handleClose();
            }
          };
          window.addEventListener("keydown", escapeHandler);
        } else {
          if (escapeHandler) {
            window.removeEventListener("keydown", escapeHandler);
            escapeHandler = null;
          }
        }
      },
    );

    onUnmounted(() => {
      if (escapeHandler) {
        window.removeEventListener("keydown", escapeHandler);
      }
    });

    return {
      handleClose,
    };
  },
};
</script>
