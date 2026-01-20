<template>
  <AdminLayout>
    <!-- Upload Section -->
    <div class="bg-white dark:bg-[#0c0f14] rounded-xl shadow-lg p-6 mb-8 border border-gray-100 dark:border-[#0c94ab40]">
      <h2 class="text-xl font-bold text-gray-900 dark:text-white mb-6">Upload Images</h2>

      <!-- File Selection -->
      <div
        :class="[
          'border-2 border-dashed rounded-xl p-12 text-center transition-all cursor-pointer mb-6',
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
          Drag and drop images here, or click to browse folders
        </p>
        <p class="text-sm text-gray-500 dark:text-gray-300 mb-6">
          Supports JPG, PNG, GIF (Max 10MB per image)
        </p>
        <button
          type="button"
          class="px-6 py-3 bg-main text-white rounded-lg hover:opacity-90 transition-colors font-medium shadow-md"
          @click.stop="triggerFileInput"
        >
          Select Images from Folder
        </button>
      </div>

      <!-- Category Selection -->
      <div class="mb-4">
        <label class="block text-sm font-medium text-gray-700 dark:text-gray-200 mb-2">
          Category <span class="text-red-500">*</span>
        </label>
        <select
          v-model="formData.category"
          required
          class="w-full px-4 py-2 border border-gray-300 dark:border-[#0c94ab40] rounded-lg bg-white dark:bg-[#0c0f14] text-gray-900 dark:text-white focus:ring-2 focus:ring-main focus:border-transparent"
        >
          <option v-for="cat in categories" :key="cat.value" :value="cat.value">
            {{ cat.label }}
          </option>
        </select>
      </div>

      <!-- Selected Files Preview -->
      <div v-if="selectedFiles.length > 0" class="mb-4">
        <p class="text-sm font-medium text-gray-700 dark:text-gray-200 mb-2">
          Selected Files ({{ selectedFiles.length }})
        </p>
        <div class="max-h-40 overflow-y-auto space-y-1">
          <div
            v-for="(file, index) in selectedFiles"
            :key="index"
            class="flex items-center justify-between p-2 bg-gray-50 dark:bg-[#1a1f2e] rounded text-sm"
          >
            <span class="text-gray-700 dark:text-gray-300 truncate flex-1">{{ file.name }}</span>
            <button
              type="button"
              @click="removeFile(index)"
              class="ml-2 text-red-600 hover:text-red-800 dark:text-red-400 dark:hover:text-red-300"
            >
              Remove
            </button>
          </div>
        </div>
      </div>

      <!-- Upload Button -->
      <button
        type="button"
        :disabled="submitting || selectedFiles.length === 0"
        @click="uploadFiles"
        class="w-full px-6 py-3 bg-main text-white rounded-lg hover:opacity-90 transition-colors font-medium shadow-md disabled:opacity-50 disabled:cursor-not-allowed"
      >
        <span v-if="submitting" class="flex items-center justify-center space-x-2">
          <div class="animate-spin rounded-full h-5 w-5 border-2 border-white border-t-transparent"></div>
          <span>Uploading {{ uploadingCount }} image(s)...</span>
        </span>
        <span v-else>Upload {{ selectedFiles.length || 0 }} Image(s)</span>
      </button>

      <!-- Error Message -->
      <div
        v-if="error"
        class="mt-4 p-4 bg-red-50 dark:bg-red-900/20 border border-red-200 dark:border-red-800 rounded-lg"
      >
        <p class="text-red-600 dark:text-red-400 text-sm">{{ error }}</p>
      </div>

      <!-- Success Message -->
      <div
        v-if="success"
        class="mt-4 p-4 bg-green-50 dark:bg-green-900/20 border border-green-200 dark:border-green-800 rounded-lg"
      >
        <p class="text-green-600 dark:text-green-400 text-sm">{{ success }}</p>
      </div>
    </div>

    <!-- Category Selector -->
    <div class="bg-white dark:bg-[#0c0f14] rounded-xl shadow-lg p-6 mb-8 border border-gray-100 dark:border-[#0c94ab40]">
      <h2 class="text-xl font-bold text-gray-900 dark:text-white mb-4">Select Category</h2>
      <div class="flex flex-wrap gap-2">
        <button
          v-for="cat in categories"
          :key="cat.value"
          @click="selectCategory(cat.value)"
          :class="[
            'px-4 py-2 rounded-lg font-medium transition-colors',
            filterCategory === cat.value
              ? 'bg-main text-white shadow-md'
              : 'bg-gray-100 dark:bg-[#1a1f2e] text-gray-700 dark:text-gray-300 hover:bg-gray-200 dark:hover:bg-[#252a3a]'
          ]"
        >
          {{ cat.label }}
        </button>
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

    const filterCategory = ref(defaultCategory);
    const galleries = ref([]);
    const loading = ref(false);
    const error = ref("");
    const success = ref("");
    const submitting = ref(false);
    const route = useRoute();
    const router = useRouter();

    const fileInput = ref(null);
    const isDragging = ref(false);
    const selectedFiles = ref([]);
    const uploadingCount = ref(0);

    const formData = ref({
      category: defaultCategory,
    });

    const applyRouteCategory = () => {
      const routeCategory = route.params.category;
      const validValues = categories.map((c) => c.value);
      const isValid = routeCategory && validValues.includes(routeCategory);
      const nextCategory = isValid ? routeCategory : defaultCategory;

      filterCategory.value = nextCategory;
      formData.value.category = nextCategory; // Sync form category with selected category

      return isValid || !routeCategory;
    };

    const selectCategory = async (categoryValue) => {
      filterCategory.value = categoryValue;
      formData.value.category = categoryValue; // Update form category when switching
      
      // Update URL to reflect selected category
      if (categoryValue === defaultCategory) {
        await router.push({ name: "GalleryManagement" });
      } else {
        await router.push({ name: "GalleryCategory", params: { category: categoryValue } });
      }
      
      await loadGalleries();
    };

    const triggerFileInput = () => {
      fileInput.value?.click();
    };

    const handleFileSelect = (event) => {
      const files = Array.from(event.target.files);
      addFiles(files);
    };

    const handleDrop = (event) => {
      isDragging.value = false;
      const files = Array.from(event.dataTransfer.files).filter((file) =>
        file.type.startsWith("image/"),
      );
      if (files.length > 0) {
        addFiles(files);
      }
    };

    const addFiles = (files) => {
      const imageFiles = files.filter((file) => file.type.startsWith("image/"));
      selectedFiles.value = [...selectedFiles.value, ...imageFiles];
      
      // Clear file input to allow selecting the same files again
      if (fileInput.value) {
        fileInput.value.value = "";
      }
    };

    const removeFile = (index) => {
      selectedFiles.value.splice(index, 1);
    };

    const uploadFiles = async () => {
      if (selectedFiles.value.length === 0) {
        error.value = "Please select at least one image";
        return;
      }

      if (!formData.value.category) {
        error.value = "Please select a category";
        return;
      }

      submitting.value = true;
      uploadingCount.value = selectedFiles.value.length;
      error.value = "";
      success.value = "";

      let successCount = 0;
      let failCount = 0;

      for (const file of selectedFiles.value) {
        try {
          const formDataToSend = new FormData();
          formDataToSend.append("image", file);
          formDataToSend.append("category", formData.value.category);

          await galleryService.uploadImage(formDataToSend);
          successCount++;
        } catch (err) {
          console.error("Upload error:", err);
          failCount++;
          error.value = `Failed to upload ${failCount} image(s): ${err.response?.data?.error || err.message}`;
        }
      }

      submitting.value = false;
      uploadingCount.value = 0;

      if (successCount > 0) {
        success.value = `Successfully uploaded ${successCount} image(s)`;
        setTimeout(() => {
          success.value = "";
        }, 3000);
        
        // Clear selected files
        selectedFiles.value = [];
        
        await loadGalleries();
      }

      // Clear file input
      if (fileInput.value) {
        fileInput.value.value = "";
      }
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

    const loadGalleries = async () => {
      loading.value = true;
      error.value = "";
      try {
        // Load all galleries (no category filter) so we can filter on frontend
        const response = await galleryService.getGalleries(null);
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
      formData,
      filterCategory,
      galleries,
      loading,
      error,
      success,
      submitting,
      categories,
      groupedGalleries,
      currentCategory,
      currentImages,
      fileInput,
      isDragging,
      selectedFiles,
      uploadingCount,
      getImageUrl,
      handleImageError,
      triggerFileInput,
      handleFileSelect,
      handleDrop,
      addFiles,
      removeFile,
      uploadFiles,
      loadGalleries,
      deleteGallery,
      setFilter,
      selectCategory,
    };
  },
};
</script>
