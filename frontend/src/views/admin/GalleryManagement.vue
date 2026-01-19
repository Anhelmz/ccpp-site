<template>
  <AdminLayout>
    <!-- Upload Section -->
    <div class="bg-white dark:bg-[#0c0f14] rounded-xl shadow-lg p-6 mb-8 border border-gray-100 dark:border-[#0c94ab40]">
      <h2 class="text-xl font-bold text-gray-900 dark:text-white mb-6">Upload New Images</h2>

      <div
        :class="[
          'border-2 border-dashed rounded-xl p-12 text-center transition-all cursor-pointer',
          isDragging
            ? 'border-main bg-primary/30 dark:border-[#0C94AB] dark:bg-[#0C94AB33]'
            : 'border-gray-300 bg-gray-100 hover:border-main hover:bg-primary/20 dark:border-[#0c94ab40] dark:bg-[#0c0f14] dark:hover:border-[#0C94AB] dark:hover:bg-[#0C94AB26]',
        ]"
        @drop.prevent="handleDrop"
        @dragover.prevent="isDragging = true"
        @dragleave.prevent="isDragging = false"
        @click="triggerFileInput"
      >
        <input
          ref="fileInput"
          type="file"
          multiple
          accept="image/jpeg,image/jpg,image/png,image/gif"
          class="hidden"
          @change="handleFileSelect"
        />
        <div
          class="mx-auto w-16 h-16 bg-secondary rounded-full flex items-center justify-center mb-4"
        >
          <svg
            class="h-8 w-8 text-main"
            fill="none"
            stroke="currentColor"
            viewBox="0 0 24 24"
          >
            <path
              stroke-linecap="round"
              stroke-linejoin="round"
              stroke-width="2"
              d="M7 16a4 4 0 01-.88-7.903A5 5 0 1115.9 6L16 6a5 5 0 011 9.9M15 13l-3-3m0 0l-3 3m3-3v12"
            ></path>
          </svg>
        </div>
        <p class="text-gray-700 dark:text-gray-200 font-medium mb-2">
          Drag and drop images here, or click to select
        </p>
        <p class="text-sm text-gray-500 dark:text-gray-300 mb-6">
          Supports JPG, PNG, GIF (Max 10MB per image)
        </p>
        <button
          class="px-6 py-3 bg-main text-white rounded-lg hover:opacity-90 transition-colors font-medium shadow-md"
          @click.stop="triggerFileInput"
        >
          Select Images
        </button>
      </div>

    <!-- Category Display -->
    <div class="mt-6">
      <label class="block text-sm font-medium text-gray-700 dark:text-gray-200 mb-1">
        Category
      </label>
      <p class="text-base font-semibold text-gray-900 dark:text-white capitalize">
        {{ currentCategory.label }}
      </p>
    </div>

      <!-- Upload Progress -->
      <div v-if="uploading" class="mt-6">
        <div class="flex items-center space-x-2">
          <div
            class="animate-spin rounded-full h-5 w-5 border-2 border-main border-t-transparent"
          ></div>
          <span class="text-gray-700"
            >Uploading {{ uploadingCount }} image(s)...</span
          >
        </div>
      </div>

      <!-- Error Message -->
      <div
        v-if="uploadError"
        class="mt-4 p-4 bg-red-50 border border-red-200 rounded-lg"
      >
        <p class="text-red-600 text-sm">{{ uploadError }}</p>
      </div>

      <!-- Success Message -->
      <div
        v-if="uploadSuccess"
        class="mt-4 p-4 bg-green-50 border border-green-200 rounded-lg"
      >
        <p class="text-green-600 text-sm">{{ uploadSuccess }}</p>
      </div>
    </div>

    <!-- Gallery Images Grid -->
    <div class="bg-white dark:bg-[#0c0f14] rounded-xl shadow-lg p-6 border border-gray-100 dark:border-[#0c94ab40]">
      <div class="flex flex-col sm:flex-row justify-between items-start sm:items-center mb-6 gap-4">
        <div>
          <h2 class="text-xl font-bold text-gray-900 dark:text-white">Gallery Images</h2>
          <p class="text-sm text-gray-600 dark:text-gray-300">Section: {{ currentCategory.label }}</p>
        </div>
      </div>

      <!-- Loading State -->
      <div v-if="loading" class="text-center py-12">
        <div
          class="inline-block animate-spin rounded-full h-12 w-12 border-4 border-main border-t-transparent mb-4"
        ></div>
        <p class="text-gray-600">Loading gallery...</p>
      </div>

      <!-- Error State -->
      <div v-else-if="error" class="text-center py-12">
        <p class="text-red-600 mb-4">{{ error }}</p>
        <button
          class="px-4 py-2 bg-main text-white rounded-lg hover:opacity-90 transition-colors"
          @click="loadGalleries"
        >
          Retry
        </button>
      </div>

      <!-- Empty State -->
      <div v-else-if="galleries.length === 0" class="text-center py-12">
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
            d="M4 16l4.586-4.586a2 2 0 012.828 0L16 16m-2-2l1.586-1.586a2 2 0 012.828 0L20 14m-6-6h.01M6 20h12a2 2 0 002-2V6a2 2 0 00-2-2H6a2 2 0 00-2 2v12a2 2 0 002 2z"
          ></path>
        </svg>
        <p class="text-gray-600 mb-2">No images uploaded yet</p>
        <p class="text-sm text-gray-500">
          Upload your first image to get started
        </p>
      </div>

      <!-- Image Grid -->
      <div v-else>
        <div class="mb-8">
          <div class="flex items-center justify-between mb-3">
            <h3 class="text-lg font-semibold text-gray-900 dark:text-white capitalize">
              {{ currentCategory.label }}
            </h3>
            <span class="text-sm text-gray-500 dark:text-gray-300">{{ currentImages.length }} image(s)</span>
          </div>
          <div v-if="currentImages.length" class="grid grid-cols-2 md:grid-cols-4 lg:grid-cols-6 gap-4">
            <div
              v-for="gallery in currentImages"
              :key="gallery.id"
              class="relative group rounded-lg overflow-hidden border border-gray-200 dark:border-[#0c94ab40] bg-white dark:bg-[#0c0f14]"
            >
              <img
                :src="getImageUrl(gallery.path)"
                :alt="gallery.filename"
                class="w-full h-32 object-cover"
                @error="handleImageError"
              />
              <div
                class="absolute inset-0 bg-black bg-opacity-0 group-hover:bg-opacity-50 transition-opacity flex items-center justify-center"
              >
                <button
                  class="text-white opacity-0 group-hover:opacity-100 px-3 py-1 bg-red-600 rounded hover:bg-red-700 transition-colors text-sm font-medium"
                  @click="deleteGallery(gallery.id)"
                >
                  Delete
                </button>
              </div>
            </div>
          </div>
          <div v-else class="text-sm text-gray-500 dark:text-gray-300">No images in this category yet.</div>
        </div>
      </div>
    </div>
  </AdminLayout>
</template>

<script>
import { ref, onMounted, computed, watch } from "vue";
import { useRoute, useRouter } from "vue-router";
import AdminLayout from "@/components/admin/AdminLayout.vue";
import { galleryService } from "@/services/galleryService";

export default {
  name: "GalleryManagement",
  components: {
    AdminLayout,
  },
  setup() {
    const categories = [
      { value: "about", label: "About" },
      { value: "outreaches", label: "Outreaches" },
      { value: "youth", label: "Youth Ministry" },
      { value: "grace-church", label: "Grace Church" },
      { value: "general", label: "General" },
    ];
    const defaultCategory = categories[0].value;

    const fileInput = ref(null);
    const isDragging = ref(false);
    const selectedCategory = ref(defaultCategory);
    const filterCategory = ref(defaultCategory);
    const galleries = ref([]);
    const loading = ref(false);
    const error = ref("");
    const uploading = ref(false);
    const uploadingCount = ref(0);
    const uploadError = ref("");
    const uploadSuccess = ref("");
    const route = useRoute();
    const router = useRouter();

    const applyRouteCategory = () => {
      const routeCategory = route.params.category;
      const validValues = categories.map((c) => c.value);
      const isValid = routeCategory && validValues.includes(routeCategory);
      const nextCategory = isValid ? routeCategory : defaultCategory;

      selectedCategory.value = nextCategory;
      filterCategory.value = nextCategory;

      return isValid || !routeCategory;
    };

    const getImageUrl = (path) => {
      // If path already starts with http, return as is
      if (path.startsWith("http")) {
        return path;
      }
      // Ensure path starts with / for proper routing
      return path.startsWith("/") ? path : `/${path}`;
    };

    const handleImageError = (event) => {
      event.target.src =
        'data:image/svg+xml,%3Csvg xmlns="http://www.w3.org/2000/svg" width="200" height="200"%3E%3Crect fill="%23ddd" width="200" height="200"/%3E%3Ctext fill="%23999" font-family="sans-serif" font-size="14" dy="10.5" font-weight="bold" x="50%25" y="50%25" text-anchor="middle"%3EImage not found%3C/text%3E%3C/svg%3E';
    };

    const triggerFileInput = () => {
      fileInput.value?.click();
    };

    const handleFileSelect = (event) => {
      const files = Array.from(event.target.files);
      uploadFiles(files);
    };

    const handleDrop = (event) => {
      isDragging.value = false;
      const files = Array.from(event.dataTransfer.files).filter((file) =>
        file.type.startsWith("image/"),
      );
      if (files.length > 0) {
        uploadFiles(files);
      }
    };

    const uploadFiles = async (files) => {
      if (!selectedCategory.value) {
        uploadError.value = "Please select a category";
        return;
      }

      uploading.value = true;
      uploadingCount.value = files.length;
      uploadError.value = "";
      uploadSuccess.value = "";

      let successCount = 0;
      let failCount = 0;

      for (const file of files) {
        try {
          const formData = new FormData();
          formData.append("image", file);
          formData.append("category", selectedCategory.value);

          await galleryService.uploadImage(formData);
          successCount++;
        } catch (err) {
          console.error("Upload error:", err);
          failCount++;
          uploadError.value = `Failed to upload ${failCount} image(s): ${err.message}`;
        }
      }

      uploading.value = false;
      uploadingCount.value = 0;

      if (successCount > 0) {
        uploadSuccess.value = `Successfully uploaded ${successCount} image(s)`;
        setTimeout(() => {
          uploadSuccess.value = "";
        }, 3000);
        await loadGalleries();
      }

      // Clear file input
      if (fileInput.value) {
        fileInput.value.value = "";
      }
    };

    const loadGalleries = async () => {
      loading.value = true;
      error.value = "";
      try {
        const response = await galleryService.getGalleries(
          filterCategory.value || null,
        );
        galleries.value = response.galleries || [];
      } catch (err) {
        console.error("Error loading galleries:", err);
        error.value = err.message || "Failed to load galleries";
      } finally {
        loading.value = false;
      }
    };

    const deleteGallery = async (id) => {
      if (!confirm("Are you sure you want to delete this image?")) {
        return;
      }

      try {
        await galleryService.deleteGallery(id);
        await loadGalleries();
      } catch (err) {
        console.error("Error deleting gallery:", err);
        alert(`Failed to delete image: ${err.message}`);
      }
    };

    const setFilter = async (value) => {
      const target = value
        ? { name: "GalleryCategory", params: { category: value } }
        : { name: "GalleryManagement" };
      await router.push(target);
    };

    const groupedGalleries = computed(() => {
      const map = {};
      galleries.value.forEach((g) => {
        const cat = g.category || "general";
        if (!map[cat]) map[cat] = [];
        map[cat].push(g);
      });
      return map;
    });

    const currentCategory = computed(() => {
      const found = categories.find((c) => c.value === filterCategory.value);
      return found || categories[0];
    });

    const currentImages = computed(() => groupedGalleries.value[currentCategory.value] || []);

    onMounted(async () => {
      const valid = applyRouteCategory();
      if (!valid && route.params.category) {
        await router.replace({ name: "GalleryManagement" });
        return;
      }
      await loadGalleries();
    });

    watch(
      () => route.params.category,
      async (newCategory) => {
        const valid = applyRouteCategory();
        if (!valid && newCategory) {
          await router.replace({ name: "GalleryManagement" });
          return;
        }
        await loadGalleries();
      },
    );

    return {
      fileInput,
      isDragging,
      selectedCategory,
      filterCategory,
      galleries,
      loading,
      error,
      uploading,
      uploadingCount,
      uploadError,
      uploadSuccess,
      categories,
      groupedGalleries,
      currentCategory,
      currentImages,
      getImageUrl,
      handleImageError,
      triggerFileInput,
      handleFileSelect,
      handleDrop,
      loadGalleries,
      deleteGallery,
      setFilter,
    };
  },
};
</script>
