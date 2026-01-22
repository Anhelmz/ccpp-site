<template>
  <AdminLayout>
    <!-- Upload Section -->
    <div class="gallery-upload-section bg-white rounded-xl shadow-lg p-6 mb-8 border border-gray-200">
      <h2 class="text-xl font-bold text-gray-900 mb-6">Upload Images</h2>

      <!-- File Selection -->
      <div
        :class="[
          'gallery-upload-area border-2 border-dashed rounded-xl p-12 text-center transition-all cursor-pointer mb-6',
          isDragging
            ? 'border-main bg-blue-50'
            : 'border-gray-300 bg-white hover:border-main hover:bg-gray-50',
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
        <p class="text-gray-700 font-medium mb-2">
          Drag and drop images here, or click to browse folders
        </p>
        <p class="text-sm text-gray-500 mb-6">
          Supports JPG, PNG, GIF (Max 10MB per image)
        </p>
        <button
          type="button"
          class="select-images-button px-6 py-3 bg-white border border-gray-300 text-green-600 rounded-lg hover:bg-gray-50 transition-colors font-medium shadow-md"
          @click.stop="triggerFileInput"
        >
          Select Images from Folder
        </button>
      </div>

      <!-- Category Selection for Upload -->
      <div class="mb-6">
        <label class="block text-base font-semibold text-gray-900 mb-3">
          Upload Category <span class="text-red-500">*</span>
        </label>
        <div class="grid grid-cols-2 sm:grid-cols-3 lg:grid-cols-5 gap-3">
          <button
            v-for="cat in categories"
            :key="cat.value"
            type="button"
            @click="formData.category = cat.value"
            :class="[
              'px-4 py-4 rounded-lg font-medium text-sm transition-all duration-200 border-2 shadow-sm',
              formData.category === cat.value
                ? 'bg-brand-blue text-white border-brand-blue shadow-md shadow-brand-blue/20'
                : 'bg-white text-gray-700 border-gray-300 hover:border-brand-blue/50 hover:bg-gray-50 hover:shadow-md'
            ]"
          >
            {{ cat.label }}
          </button>
        </div>
        <p class="text-xs text-gray-600 mt-2">
          Currently selected: <span class="font-semibold text-gray-900">{{ categories.find(c => c.value === formData.category)?.label }}</span>
        </p>
      </div>

      <!-- Selected Files Preview -->
      <div v-if="selectedFiles.length > 0" class="mb-4">
        <p class="text-sm font-medium text-gray-700 mb-2">
          Selected Files ({{ selectedFiles.length }})
        </p>
        <div class="max-h-40 overflow-y-auto space-y-1 border border-gray-200 rounded-lg p-2 bg-white">
          <div
            v-for="(file, index) in selectedFiles"
            :key="index"
            class="flex items-center justify-between p-2 bg-white rounded text-sm border border-gray-200 hover:shadow-sm transition-shadow"
          >
            <span class="text-gray-700 truncate flex-1">{{ file.name }}</span>
            <button
              type="button"
              @click="removeFile(index)"
              class="ml-2 text-red-600 hover:text-red-800 font-medium"
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
        class="w-full px-6 py-3 bg-brand-blue dark:bg-main text-white rounded-lg hover:bg-brand-blue/90 dark:hover:opacity-90 transition-colors font-medium shadow-md disabled:opacity-50 disabled:cursor-not-allowed"
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
    <div class="gallery-category-section bg-white rounded-xl shadow-lg p-8 mb-8 border border-gray-200">
      <div class="mb-6">
        <h2 class="text-2xl font-bold text-gray-900 mb-2">Gallery Category</h2>
        <p class="text-sm text-gray-600">Select a category to view and manage images</p>
      </div>
      <div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 xl:grid-cols-5 gap-4">
        <button
          v-for="cat in categories"
          :key="cat.value"
          @click="selectCategory(cat.value)"
          :class="[
            'relative px-6 py-8 rounded-xl font-semibold text-lg transition-all duration-200 transform hover:scale-105',
            'border-2 flex flex-col items-center justify-center min-h-[120px] shadow-sm',
            filterCategory === cat.value
              ? 'bg-brand-blue text-white border-brand-blue shadow-lg shadow-brand-blue/30 scale-105'
              : 'bg-white text-gray-700 border-gray-300 hover:border-brand-blue/50 hover:bg-gray-50 hover:shadow-md'
          ]"
        >
          <div v-if="filterCategory === cat.value" class="absolute top-2 right-2">
            <svg class="w-5 h-5 text-white" fill="currentColor" viewBox="0 0 20 20">
              <path fill-rule="evenodd" d="M10 18a8 8 0 100-16 8 8 0 000 16zm3.707-9.293a1 1 0 00-1.414-1.414L9 10.586 7.707 9.293a1 1 0 00-1.414 1.414l2 2a1 1 0 001.414 0l4-4z" clip-rule="evenodd"/>
            </svg>
          </div>
          <span class="text-center">{{ cat.label }}</span>
          <span 
            v-if="groupedGalleries[cat.value] && groupedGalleries[cat.value].length > 0"
            :class="[
              'text-sm mt-2 font-normal',
              filterCategory === cat.value ? 'text-white/80' : 'text-gray-500'
            ]"
          >
            {{ groupedGalleries[cat.value].length }} image(s)
          </span>
        </button>
      </div>
    </div>

    <!-- Gallery Images Grid -->
    <div class="gallery-images-section bg-white rounded-xl shadow-lg p-6 border border-gray-200">
      <div class="flex flex-col sm:flex-row justify-between items-start sm:items-center mb-6 gap-4">
        <div>
          <h2 class="text-xl font-bold text-gray-900">Gallery Images</h2>
          <p class="text-sm text-gray-600">Section: {{ currentCategory.label }}</p>
        </div>
      </div>

      <!-- Loading State -->
      <div v-if="loading" class="text-center py-12">
        <div
          class="inline-block animate-spin rounded-full h-12 w-12 border-4 border-brand-blue dark:border-main border-t-transparent mb-4"
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
      <div v-else-if="currentImages.length === 0" class="text-center py-12">
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
        <p class="text-gray-600 mb-2">No images in this category yet</p>
        <p class="text-sm text-gray-500">
          Upload images to the "{{ currentCategory.label }}" category to get started
        </p>
      </div>

      <!-- Image Grid -->
      <div v-else>
        <div class="mb-8">
          <div class="flex items-center justify-between mb-3">
            <h3 class="text-lg font-semibold text-gray-900 capitalize">
              {{ currentCategory.label }}
            </h3>
            <span class="text-sm text-gray-500">{{ currentImages.length }} image(s)</span>
          </div>
          <div v-if="currentImages.length > 0" class="grid grid-cols-2 md:grid-cols-4 lg:grid-cols-6 gap-4">
            <div
              v-for="(gallery, index) in currentImages"
              :key="gallery.id"
              class="gallery-image-item relative group rounded-lg overflow-hidden border border-gray-300 bg-white shadow-sm hover:shadow-lg transition-all cursor-pointer hover:border-brand-blue"
              @click="viewImage(index)"
            >
              <img
                :src="getImageUrl(gallery.path)"
                :alt="gallery.filename || 'Gallery image'"
                class="w-full h-32 object-cover"
                @error="handleImageError"
              />
              <div
                class="absolute inset-0 bg-black bg-opacity-0 group-hover:bg-opacity-60 transition-opacity flex items-center justify-center gap-2"
              >
                <button
                  class="text-white opacity-0 group-hover:opacity-100 px-4 py-2 bg-brand-blue dark:bg-blue-600 rounded-lg hover:bg-brand-blue/90 dark:hover:bg-blue-700 transition-all text-sm font-semibold shadow-lg transform hover:scale-105"
                  @click.stop="viewImage(index)"
                >
                  View
                </button>
                <button
                  class="text-white opacity-0 group-hover:opacity-100 px-4 py-2 bg-red-600 rounded-lg hover:bg-red-700 transition-all text-sm font-semibold shadow-lg transform hover:scale-105"
                  @click.stop="openDeleteModal(gallery.id)"
                >
                  Delete
                </button>
              </div>
              <div class="absolute top-2 right-2 opacity-0 group-hover:opacity-100 transition-opacity">
                <div class="bg-black/50 text-white text-xs px-2 py-1 rounded">
                  {{ gallery.category || 'uncategorized' }}
                </div>
              </div>
            </div>
          </div>
          <div v-else class="text-center py-8">
            <p class="text-sm text-gray-500">No images in this category yet.</p>
          </div>
        </div>
      </div>
    </div>

    <!-- Image Viewer Modal -->
    <div
      v-if="viewingImageIndex !== null"
      class="fixed inset-0 z-50 flex items-center justify-center bg-black bg-opacity-90 p-4"
      @click.self="closeImageViewer"
    >
      <div class="relative max-w-7xl max-h-full w-full h-full flex items-center justify-center">
        <!-- Close Button -->
        <button
          @click="closeImageViewer"
          class="absolute top-4 right-4 z-10 text-white hover:text-gray-300 transition-colors bg-black/50 rounded-full p-2 hover:bg-black/70"
          aria-label="Close"
        >
          <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12" />
          </svg>
        </button>

        <!-- Delete Button -->
        <button
          @click="deleteCurrentImage"
          class="absolute top-4 right-20 z-10 text-white hover:text-red-300 transition-colors bg-red-600/80 hover:bg-red-600 rounded-full px-4 py-2 text-sm font-semibold"
          aria-label="Delete"
        >
          Delete Image
        </button>

        <!-- Previous Button -->
        <button
          v-if="currentImages.length > 1"
          @click.stop="previousImage"
          class="absolute left-4 z-10 text-white hover:text-gray-300 transition-colors bg-black/50 rounded-full p-3 hover:bg-black/70 disabled:opacity-50 disabled:cursor-not-allowed"
          :disabled="viewingImageIndex === 0"
          aria-label="Previous"
        >
          <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 19l-7-7 7-7" />
          </svg>
        </button>

        <!-- Next Button -->
        <button
          v-if="currentImages.length > 1"
          @click.stop="nextImage"
          class="absolute right-4 z-10 text-white hover:text-gray-300 transition-colors bg-black/50 rounded-full p-3 hover:bg-black/70 disabled:opacity-50 disabled:cursor-not-allowed"
          :disabled="viewingImageIndex === currentImages.length - 1"
          aria-label="Next"
        >
          <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7" />
          </svg>
        </button>

        <!-- Image -->
        <div class="flex flex-col items-center justify-center max-w-full max-h-full">
          <img
            :src="getImageUrl(currentViewingImage?.path)"
            :alt="currentViewingImage?.filename || 'Gallery image'"
            class="max-w-full max-h-[90vh] object-contain rounded-lg shadow-2xl"
            @error="handleImageError"
          />
          <!-- Image Info -->
          <div class="mt-4 text-white text-center">
            <p class="text-sm opacity-80">
              {{ viewingImageIndex + 1 }} of {{ currentImages.length }}
            </p>
            <p v-if="currentViewingImage?.filename" class="text-sm opacity-60 mt-1">
              {{ currentViewingImage.filename }}
            </p>
          </div>
        </div>
      </div>
    </div>

    <!-- Delete Confirmation Modal -->
    <div
      v-if="showDeleteModal"
      class="fixed inset-0 z-[60] flex items-center justify-center bg-black bg-opacity-50 p-4"
      @click.self="closeDeleteModal"
    >
      <div class="bg-white dark:bg-[#0c0f14] rounded-xl shadow-2xl max-w-md w-full p-6 border border-gray-200 dark:border-[#0c94ab40]">
        <div class="flex items-center justify-between mb-4">
          <h3 class="text-xl font-bold text-gray-900 dark:text-white">Confirm Delete</h3>
          <button
            @click="closeDeleteModal"
            class="text-gray-400 hover:text-gray-600 dark:hover:text-gray-200 transition-colors"
            aria-label="Close"
          >
            <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12" />
            </svg>
          </button>
        </div>
        
        <p class="text-gray-700 dark:text-gray-300 mb-6">
          Are you sure you want to delete this image? This action cannot be undone.
        </p>
        
        <div class="flex gap-3 justify-end">
          <button
            @click="closeDeleteModal"
            class="px-4 py-2 bg-gray-200 dark:bg-[#1a1f2e] text-gray-700 dark:text-gray-300 rounded-lg hover:bg-gray-300 dark:hover:bg-[#252a3a] transition-colors font-medium"
          >
            Cancel
          </button>
          <button
            @click="confirmDelete"
            class="px-4 py-2 bg-red-600 text-white rounded-lg hover:bg-red-700 transition-colors font-medium"
          >
            Delete Image
          </button>
        </div>
      </div>
    </div>
  </AdminLayout>
</template>

<script>
import { ref, onMounted, onUnmounted, computed, watch } from "vue";
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
    const viewingImageIndex = ref(null);
    const showDeleteModal = ref(false);
    const imageToDelete = ref(null);

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
        // API interceptor already unwraps response.data, so response is the data object
        galleries.value = response.galleries || response.data?.galleries || [];
        console.log("Loaded galleries:", galleries.value.length, "Total");
        console.log("All galleries:", galleries.value);
        console.log("Grouped by category:", groupedGalleries.value);
        console.log("Current category filter:", filterCategory.value);
        console.log("Current images for category:", currentImages.value.length, currentImages.value);
      } catch (err) {
        console.error("Error loading galleries:", err);
        error.value = err.message || "Failed to load galleries";
      } finally {
        loading.value = false;
      }
    };

    const openDeleteModal = (id) => {
      imageToDelete.value = id;
      showDeleteModal.value = true;
    };

    const closeDeleteModal = () => {
      showDeleteModal.value = false;
      imageToDelete.value = null;
    };

    const confirmDelete = async () => {
      if (!imageToDelete.value) return;

      const id = imageToDelete.value;
      closeDeleteModal();

      try {
        await galleryService.deleteGallery(id);
        success.value = "Image deleted successfully";
        setTimeout(() => {
          success.value = "";
        }, 3000);
        await loadGalleries();
        
        // If we're viewing an image and it was deleted, close viewer or adjust index
        if (viewingImageIndex.value !== null) {
          if (currentImages.value.length === 0) {
            closeImageViewer();
          } else if (viewingImageIndex.value >= currentImages.value.length) {
            viewingImageIndex.value = currentImages.value.length - 1;
          }
        }
      } catch (err) {
        console.error("Error deleting gallery:", err);
        error.value = `Failed to delete image: ${err.message || err}`;
        setTimeout(() => {
          error.value = "";
        }, 5000);
      }
    };

    const deleteGallery = async (id) => {
      openDeleteModal(id);
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
        // Use the category as-is, but normalize to handle null/undefined
        // Default to "about" if category is missing
        const cat = (g.category || defaultCategory).trim();
        if (!map[cat]) map[cat] = [];
        map[cat].push(g);
      });
      return map;
    });

    const currentCategory = computed(() => {
      const found = categories.find((c) => c.value === filterCategory.value);
      return found || categories[0];
    });

    const currentImages = computed(() => {
      const categoryKey = filterCategory.value || defaultCategory;
      return groupedGalleries.value[categoryKey] || [];
    });

    const currentViewingImage = computed(() => {
      if (viewingImageIndex.value === null) return null;
      return currentImages.value[viewingImageIndex.value] || null;
    });

    const viewImage = (index) => {
      viewingImageIndex.value = index;
      // Prevent body scroll when modal is open
      document.body.style.overflow = "hidden";
    };

    const closeImageViewer = () => {
      viewingImageIndex.value = null;
      // Restore body scroll
      document.body.style.overflow = "";
    };

    const nextImage = () => {
      if (viewingImageIndex.value < currentImages.value.length - 1) {
        viewingImageIndex.value++;
      }
    };

    const previousImage = () => {
      if (viewingImageIndex.value > 0) {
        viewingImageIndex.value--;
      }
    };

    const deleteCurrentImage = () => {
      if (!currentViewingImage.value) return;
      const imageId = currentViewingImage.value.id;
      openDeleteModal(imageId);
    };

    const handleKeyDown = (event) => {
      if (viewingImageIndex.value === null) return;
      
      if (event.key === "Escape") {
        closeImageViewer();
      } else if (event.key === "ArrowLeft") {
        previousImage();
      } else if (event.key === "ArrowRight") {
        nextImage();
      }
    };

    onMounted(async () => {
      const valid = applyRouteCategory();
      if (!valid && route.params.category) {
        await router.replace({ name: "GalleryManagement" });
        return;
      }
      await loadGalleries();
      
      // Add keyboard event listener
      window.addEventListener("keydown", handleKeyDown);
    });

    onUnmounted(() => {
      // Remove keyboard event listener
      window.removeEventListener("keydown", handleKeyDown);
      // Ensure body scroll is restored
      document.body.style.overflow = "";
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
      viewingImageIndex,
      currentViewingImage,
      viewImage,
      closeImageViewer,
      nextImage,
      previousImage,
      deleteCurrentImage,
      showDeleteModal,
      openDeleteModal,
      closeDeleteModal,
      confirmDelete,
    };
  },
};
</script>

<style>
/* Force white backgrounds for gallery sections in light mode only */
.gallery-upload-section {
  background-color: white !important;
  border-color: #e5e7eb !important;
}

.gallery-upload-area {
  background-color: white !important;
  border-color: #d1d5db !important;
}

.gallery-upload-area:hover {
  background-color: #f9fafb !important;
}

.gallery-category-section {
  background-color: white !important;
  border-color: #e5e7eb !important;
}

.gallery-category-section button:not(.bg-brand-blue) {
  background-color: white !important;
  border-color: #d1d5db !important;
  color: #374151 !important;
}

.gallery-category-section button:not(.bg-brand-blue):hover {
  background-color: #f9fafb !important;
}

.gallery-images-section {
  background-color: white !important;
  border-color: #e5e7eb !important;
}

.gallery-image-item {
  background-color: white !important;
  border-color: #d1d5db !important;
}

.gallery-upload-section h2,
.gallery-category-section h2,
.gallery-images-section h2 {
  color: #111827 !important;
}

/* Force green text for Select Images button */
.select-images-button {
  background-color: white !important;
  border-color: #d1d5db !important;
  color: #16a34a !important;
}

.select-images-button:hover {
  background-color: #f9fafb !important;
  color: #16a34a !important;
}
</style>
