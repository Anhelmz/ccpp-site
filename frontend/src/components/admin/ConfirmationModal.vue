<template>
  <div
    v-if="show"
    class="fixed inset-0 z-[60] flex items-center justify-center bg-black bg-opacity-50 backdrop-blur-sm p-4"
    @click.self="handleCancel"
  >
    <div
      class="bg-white dark:bg-[#0c0f14] rounded-xl shadow-2xl max-w-md w-full p-6 border-2 border-brand-blue"
    >
      <div class="flex items-center justify-between mb-4">
        <h3 class="text-xl font-bold text-gray-900 dark:text-white">
          {{ title }}
        </h3>
        <button
          @click="handleCancel"
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

      <div class="flex gap-3 justify-end">
        <button
          @click="handleCancel"
          class="px-6 py-2 bg-gray-200 dark:bg-[#1a1f2e] text-gray-700 dark:text-gray-300 rounded-lg hover:bg-gray-300 dark:hover:bg-[#252a3a] transition-colors font-medium"
        >
          Cancel
        </button>
        <button
          @click="handleConfirm"
          :class="[
            'px-6 py-2 rounded-lg transition-colors font-medium',
            confirmButtonClass || 'bg-red-600 text-white hover:bg-red-700'
          ]"
        >
          {{ confirmText }}
        </button>
      </div>
    </div>
  </div>
</template>

<script>
import { watch, onUnmounted } from "vue";

export default {
  name: "ConfirmationModal",
  props: {
    show: {
      type: Boolean,
      default: false,
    },
    title: {
      type: String,
      default: "Confirm Action",
    },
    message: {
      type: String,
      required: true,
    },
    confirmText: {
      type: String,
      default: "Confirm",
    },
    confirmButtonClass: {
      type: String,
      default: null,
    },
  },
  emits: ["confirm", "cancel"],
  setup(props, { emit }) {
    const handleConfirm = () => {
      emit("confirm");
    };

    const handleCancel = () => {
      emit("cancel");
    };

    let escapeHandler = null;

    // Close on Escape key
    watch(
      () => props.show,
      (isShowing) => {
        if (isShowing) {
          escapeHandler = (e) => {
            if (e.key === "Escape") {
              handleCancel();
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
      handleConfirm,
      handleCancel,
    };
  },
};
</script>
