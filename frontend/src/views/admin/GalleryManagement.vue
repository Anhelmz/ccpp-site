<template>
  <AdminLayout>
    <!-- Category List View -->
    <div v-if="!selectedCategory">
      <!-- Page Header -->
      <div class="mb-8">
        <h1 class="text-2xl font-semibold text-zinc-900">Gallery</h1>
        <p class="mt-1 text-sm text-zinc-600">Manage images by category</p>
      </div>

      <!-- Categories Table -->
      <div class="bg-white border border-gray-200 rounded-lg overflow-hidden">
        <div v-if="loading" class="flex items-center justify-center py-12">
          <div class="animate-spin rounded-full h-8 w-8 border-b-2 border-[#0089AE]"></div>
        </div>

        <div v-else-if="error" class="bg-red-50 border border-red-200 rounded-lg p-4 m-6">
          <p class="text-sm text-red-800">{{ error }}</p>
        </div>

        <div v-else-if="categoryList.length === 0" class="text-center py-12">
          <svg
            class="mx-auto h-16 w-16 text-zinc-300 mb-4"
            fill="none"
            stroke="currentColor"
            viewBox="0 0 24 24"
          >
            <path
              stroke-linecap="round"
              stroke-linejoin="round"
              stroke-width="2"
              d="M4 16l4.586-4.586a2 2 0 012.828 0L16 16m-2-2l1.586-1.586a2 2 0 012.828 0L20 14m-6-6h.01M6 20h12a2 2 0 002-2V6a2 2 0 00-2-2H6a2 2 0 00-2 2v12a2 2 0 002 2z"
            />
          </svg>
          <h3 class="text-lg font-medium text-zinc-900 mb-2">No Categories</h3>
          <p class="text-sm text-zinc-600">Categories will appear here</p>
      </div>

        <div v-else>
          <div class="overflow-x-auto">
            <table class="min-w-full divide-y divide-gray-200">
              <thead class="bg-gray-50">
                <tr>
                  <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">Category</th>
                  <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">Images</th>
                  <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">Actions</th>
                </tr>
              </thead>
              <tbody class="bg-white divide-y divide-gray-200">
                <tr
                  v-for="category in categoryList"
                  :key="category.value"
                  class="hover:bg-gray-50 transition-colors cursor-pointer"
                  @click="viewCategory(category.value)"
                >
                  <td class="px-6 py-4 whitespace-nowrap text-sm font-medium text-gray-900">
                    {{ category.label }}
                  </td>
                  <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-600">
                    {{ category.count }} image(s)
                  </td>
                  <td class="px-6 py-4 whitespace-nowrap text-right text-sm font-medium" @click.stop>
            <button
                      class="text-[#0089AE] hover:text-[#007A9D] transition-colors font-medium"
                      @click="viewCategory(category.value)"
            >
                      View
            </button>
                  </td>
                </tr>
              </tbody>
            </table>
          </div>
        </div>
      </div>
    </div>

    <!-- Category Detail View -->
    <div v-if="selectedCategory">
      <!-- Breadcrumbs -->
      <nav class="mb-6" aria-label="Breadcrumb">
        <ol class="flex items-center space-x-2 text-sm">
          <li>
            <a href="#" @click.prevent="selectedCategory = null" class="text-gray-500 hover:text-gray-700">
              Gallery
            </a>
          </li>
          <li>
            <span class="text-gray-400">/</span>
          </li>
          <li class="text-gray-900 font-medium">
            {{ currentCategoryData.label }}
          </li>
        </ol>
      </nav>

      <!-- Category Details -->
      <div class="bg-white border border-gray-200 rounded-lg">
        <div class="flex flex-col lg:flex-row">
          <!-- Main Content Area -->
          <div class="flex-1 p-6 lg:pr-8">
            <div class="flex items-center justify-between mb-6">
              <h1 class="text-3xl font-bold text-gray-900">{{ currentCategoryData.label }}</h1>
              <span class="text-sm text-gray-500">{{ currentImages.length }} image(s)</span>
            </div>

            <!-- Search Bar -->
      <div class="mb-6">
              <div class="relative max-w-md">
                <input
                  v-model="searchQuery"
                  type="text"
                  placeholder="Search images..."
                  class="w-full px-4 py-2 pl-10 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-[#0089AE] focus:border-transparent"
                />
                <svg
                  class="absolute left-3 top-1/2 transform -translate-y-1/2 w-5 h-5 text-gray-400"
                  fill="none"
                  stroke="currentColor"
                  viewBox="0 0 24 24"
                >
                  <path
                    stroke-linecap="round"
                    stroke-linejoin="round"
                    stroke-width="2"
                    d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z"
                  />
            </svg>
        </div>
      </div>

      <!-- Loading State -->
            <div v-if="loading" class="flex items-center justify-center py-12">
              <div class="animate-spin rounded-full h-8 w-8 border-b-2 border-[#0089AE]"></div>
      </div>

      <!-- Error State -->
            <div v-else-if="error" class="bg-red-50 border border-red-200 rounded-lg p-4 mb-6">
              <p class="text-sm text-red-800">{{ error }}</p>
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
                />
        </svg>
        <p class="text-gray-600 mb-2">No images in this category yet</p>
        <p class="text-sm text-gray-500">
                Upload images to the "{{ currentCategoryData.label }}" category to get started
        </p>
      </div>

      <!-- Image Grid -->
            <div v-else class="grid grid-cols-2 md:grid-cols-3 lg:grid-cols-4 gap-4">
            <div
              v-for="(gallery, index) in currentImages"
              :key="gallery.id"
                class="relative group rounded-lg overflow-hidden border border-gray-300 bg-white shadow-sm hover:shadow-lg transition-all cursor-pointer hover:border-[#0089AE]"
              @click="viewImage(index)"
            >
              <img
                :src="getImageUrl(gallery.path)"
                :alt="gallery.filename || 'Gallery image'"
                  class="w-full h-48 object-cover"
                @error="handleImageError"
              />
              <div
                class="absolute inset-0 bg-black bg-opacity-0 group-hover:bg-opacity-60 transition-opacity flex items-center justify-center gap-2"
              >
                <button
                    class="text-white opacity-0 group-hover:opacity-100 px-3 py-1.5 bg-[#0089AE] rounded-lg hover:bg-[#007A9D] transition-all text-xs font-medium shadow-lg"
                  @click.stop="viewImage(index)"
                >
                  View
                </button>
                <button
                    class="opacity-0 group-hover:opacity-100 px-3 py-1.5 text-white bg-red-600 rounded-lg hover:bg-red-700 transition-all text-xs font-medium"
                  @click.stop="openDeleteModal(gallery.id)"
                >
                  Delete
                </button>
                </div>
              </div>
            </div>
          </div>

          <!-- Sidebar -->
          <div class="lg:w-80 lg:border-l border-gray-200 p-6 space-y-6 bg-gray-50">
            <!-- Category Information -->
            <div>
              <label class="text-xs font-medium text-gray-500 uppercase tracking-wider">Category</label>
              <p class="mt-1 text-sm text-gray-900">{{ currentCategoryData.label }}</p>
            </div>

            <div>
              <label class="text-xs font-medium text-gray-500 uppercase tracking-wider">Total Images</label>
              <p class="mt-1 text-sm text-gray-600">{{ currentImages.length }} image(s)</p>
            </div>

            <!-- Upload Section -->
            <div class="pt-4 border-t border-gray-200">
              <h3 class="text-sm font-medium text-gray-900 mb-4">Upload Images</h3>
              
              <!-- File Selection -->
              <div
                :class="[
                  'border-2 border-dashed rounded-lg p-6 text-center transition-all cursor-pointer mb-4',
                  isDragging
                    ? 'border-[#0089AE] bg-blue-50'
                    : 'border-gray-300 bg-white hover:border-[#0089AE] hover:bg-gray-50',
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
                <svg
                  class="mx-auto h-8 w-8 text-[#0089AE] mb-2"
                  fill="none"
                  stroke="currentColor"
                  viewBox="0 0 24 24"
                >
                  <path
                    stroke-linecap="round"
                    stroke-linejoin="round"
                    stroke-width="2"
                    d="M7 16a4 4 0 01-.88-7.903A5 5 0 1115.9 6L16 6a5 5 0 011 9.9M15 13l-3-3m0 0l-3 3m3-3v12"
                  />
                </svg>
                <p class="text-xs text-gray-600 mb-1">Click or drag to upload</p>
                <p class="text-xs text-gray-500">JPG, PNG, GIF (Max 10MB)</p>
              </div>

              <!-- Selected Files Preview -->
              <div v-if="selectedFiles.length > 0" class="mb-3">
                <p class="text-xs font-medium text-gray-700 mb-2">
                  Selected ({{ selectedFiles.length }})
                </p>
                <div class="max-h-32 overflow-y-auto space-y-1 border border-gray-200 rounded-lg p-2 bg-white">
                  <div
                    v-for="(file, index) in selectedFiles"
                    :key="index"
                    class="flex items-center justify-between p-1.5 bg-white rounded text-xs border border-gray-200"
                  >
                    <span class="text-gray-700 truncate flex-1">{{ file.name }}</span>
                    <button
                      type="button"
                      @click.stop="removeFile(index)"
                      class="ml-2 text-red-600 hover:text-red-800 text-xs"
                    >
                      ×
                    </button>
                  </div>
                </div>
              </div>

              <!-- Upload Button -->
              <button
                type="button"
                :disabled="submitting || selectedFiles.length === 0"
                @click="uploadFiles"
                class="w-full px-4 py-2 text-sm bg-[#0089AE] text-white rounded-md hover:bg-[#007A9D] transition-colors font-medium disabled:opacity-50 disabled:cursor-not-allowed"
              >
                <span v-if="submitting" class="flex items-center justify-center space-x-2">
                  <div class="animate-spin rounded-full h-4 w-4 border-2 border-white border-t-transparent"></div>
                  <span>Uploading...</span>
                </span>
                <span v-else>Upload {{ selectedFiles.length || 0 }} Image(s)</span>
              </button>

              <!-- Error Message -->
              <div v-if="error" class="mt-3 p-3 bg-red-50 border border-red-200 rounded-lg">
                <p class="text-xs text-red-800">{{ error }}</p>
              </div>

              <!-- Success Message -->
              <div v-if="success" class="mt-3 p-3 bg-green-50 border border-green-200 rounded-lg">
                <p class="text-xs text-green-800">{{ success }}</p>
              </div>
            </div>

            <!-- Action Buttons -->
            <div class="pt-4 border-t border-gray-200 space-y-3">
              <button
                v-if="currentImages.length > 0"
                @click="deleteAllInCategory"
                class="w-full px-4 py-2 bg-red-600 text-white text-sm font-medium rounded-md hover:bg-red-700 transition-colors"
              >
                Delete All Images ({{ currentImages.length }})
              </button>
              <button
                @click="selectedCategory = null"
                class="w-full px-4 py-2 border border-gray-300 text-gray-700 text-sm font-medium rounded-md hover:bg-gray-50 transition-colors"
              >
                Back to Categories
              </button>
            </div>
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
          class="absolute top-4 right-4 z-10 text-white hover:text-gray-300 transition-colors bg-black/50 rounded-full p-1.5 hover:bg-black/70"
          aria-label="Close"
        >
          <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12" />
          </svg>
        </button>

        <!-- Delete Button -->
        <button
          @click="deleteCurrentImage"
          class="absolute top-4 right-20 z-10 text-red-600 hover:text-red-800 transition-colors px-3 py-1.5 text-xs font-medium bg-black/50 rounded-lg hover:bg-black/70"
          aria-label="Delete"
        >
          Delete
        </button>

        <!-- Previous Button -->
        <button
          v-if="currentImages.length > 1"
          @click.stop="previousImage"
          class="absolute left-4 z-10 text-white hover:text-gray-300 transition-colors bg-black/50 rounded-full p-2 hover:bg-black/70 disabled:opacity-50 disabled:cursor-not-allowed"
          :disabled="viewingImageIndex === 0"
          aria-label="Previous"
        >
          <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 19l-7-7 7-7" />
          </svg>
        </button>

        <!-- Next Button -->
        <button
          v-if="currentImages.length > 1"
          @click.stop="nextImage"
          class="absolute right-4 z-10 text-white hover:text-gray-300 transition-colors bg-black/50 rounded-full p-2 hover:bg-black/70 disabled:opacity-50 disabled:cursor-not-allowed"
          :disabled="viewingImageIndex === currentImages.length - 1"
          aria-label="Next"
        >
          <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
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
      class="fixed inset-0 z-[60] flex items-center justify-center bg-gray-900/20 p-4"
      @click.self="closeDeleteModal"
    >
      <div class="bg-white rounded-lg shadow-xl max-w-md w-full p-6 border border-gray-200">
        <div class="flex items-center justify-between mb-4">
          <h3 class="text-lg font-medium text-gray-900">Confirm Delete</h3>
          <button
            @click="closeDeleteModal"
            class="text-gray-400 hover:text-gray-600 transition-colors"
            aria-label="Close"
          >
            <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12" />
            </svg>
          </button>
        </div>
        
        <p class="text-gray-700 mb-6">
          Are you sure you want to delete this image? This action cannot be undone.
        </p>
        
        <div class="flex gap-3 justify-end">
          <button
            @click="closeDeleteModal"
            class="px-4 py-2 border border-gray-300 text-gray-700 rounded-md hover:bg-gray-50 transition-colors font-medium"
          >
            Cancel
          </button>
          <button
            @click="confirmDelete"
            class="px-4 py-2 bg-red-600 text-white rounded-md hover:bg-red-700 transition-colors font-medium"
          >
            Delete Image
          </button>
        </div>
      </div>
    </div>

    <!-- Delete All Confirmation Modal -->
    <div
      v-if="showDeleteAllModal"
      class="fixed inset-0 z-[60] flex items-center justify-center bg-gray-900/20 p-4"
      @click.self="cancelDeleteAll"
    >
      <div class="bg-white rounded-lg shadow-xl max-w-md w-full p-6 border border-gray-200">
        <div class="flex items-center justify-between mb-4">
          <h3 class="text-lg font-medium text-gray-900">Confirm Delete All</h3>
          <button
            @click="cancelDeleteAll"
            class="text-gray-400 hover:text-gray-600 transition-colors"
            aria-label="Close"
          >
            <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12" />
            </svg>
          </button>
        </div>
        
        <p class="text-gray-700 mb-6">
          Are you sure you want to delete all {{ currentImages.length }} image(s) from the "{{ currentCategoryData.label }}" category? This action cannot be undone.
        </p>
        
        <div class="flex gap-3 justify-end">
          <button
            @click="cancelDeleteAll"
            class="px-4 py-2 border border-gray-300 text-gray-700 rounded-md hover:bg-gray-50 transition-colors font-medium"
          >
            Cancel
          </button>
          <button
            @click="confirmDeleteAll"
            class="px-4 py-2 bg-red-600 text-white rounded-md hover:bg-red-700 transition-colors font-medium"
          >
            Delete All Images
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

    const selectedCategory = ref(null);
    const searchQuery = ref("");
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
    const showDeleteAllModal = ref(false);

    const applyRouteCategory = () => {
      const routeCategory = route.params.category;
      const validValues = categories.map((c) => c.value);
      const isValid = routeCategory && validValues.includes(routeCategory);
      
      if (isValid) {
        selectedCategory.value = routeCategory;
      } else if (routeCategory) {
        return false;
      }
      
      return true;
    };

    const viewCategory = (categoryValue) => {
      selectedCategory.value = categoryValue;
      searchQuery.value = "";
      
      // Update URL
      if (categoryValue === defaultCategory) {
        router.push({ name: "GalleryManagement" });
      } else {
        router.push({ name: "GalleryCategory", params: { category: categoryValue } });
      }
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

      if (!selectedCategory.value) {
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
          formDataToSend.append("category", selectedCategory.value);

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
        
        selectedFiles.value = [];
        await loadGalleries();
      }

      if (fileInput.value) {
        fileInput.value.value = "";
      }
    };

    const getImageUrl = (path) => {
      // If path is a data URI (starts with "data:"), return it directly
      if (path && path.startsWith("data:")) {
        return path;
      }
      // If path is a full URL, return it
      if (path && path.startsWith("http")) {
        return path;
      }
      // Otherwise, treat as relative path
      return path && path.startsWith("/") ? path : `/${path}`;
    };

    const handleImageError = (event) => {
      event.target.src =
        'data:image/svg+xml,%3Csvg xmlns="http://www.w3.org/2000/svg" width="200" height="200"%3E%3Crect fill="%23ddd" width="200" height="200"/%3E%3Ctext fill="%23999" font-family="sans-serif" font-size="14" dy="10.5" font-weight="bold" x="50%25" y="50%25" text-anchor="middle"%3EImage not found%3C/text%3E%3C/svg%3E';
    };

    const loadGalleries = async () => {
      loading.value = true;
      error.value = "";
      try {
        const response = await galleryService.getGalleries(null);
        galleries.value = response.galleries || response.data?.galleries || [];
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

    const deleteAllInCategory = () => {
      if (currentImages.value.length === 0) return;
      showDeleteAllModal.value = true;
    };

    const confirmDeleteAll = async () => {
      if (!selectedCategory.value) return;

      showDeleteAllModal.value = false;

      try {
        await galleryService.deleteGalleriesByCategory(selectedCategory.value);
        success.value = `Successfully deleted all images from ${currentCategoryData.value.label} category`;
        setTimeout(() => {
          success.value = "";
        }, 3000);
        await loadGalleries();
        
        if (viewingImageIndex.value !== null) {
          closeImageViewer();
        }
      } catch (err) {
        console.error("Error deleting galleries by category:", err);
        error.value = `Failed to delete images: ${err.message || err}`;
        setTimeout(() => {
          error.value = "";
        }, 5000);
      }
    };

    const cancelDeleteAll = () => {
      showDeleteAllModal.value = false;
    };

    const groupedGalleries = computed(() => {
      const map = {};
      galleries.value.forEach((g) => {
        const cat = (g.category || defaultCategory).trim();
        if (!map[cat]) map[cat] = [];
        map[cat].push(g);
      });
      return map;
    });

    const categoryList = computed(() => {
      return categories.map((cat) => ({
        ...cat,
        count: (groupedGalleries.value[cat.value] || []).length,
      }));
    });

    const currentCategoryData = computed(() => {
      const found = categories.find((c) => c.value === selectedCategory.value);
      return found || categories[0];
    });

    const currentImages = computed(() => {
      if (!selectedCategory.value) return [];
      
      const categoryKey = selectedCategory.value || defaultCategory;
      let images = groupedGalleries.value[categoryKey] || [];
      
      if (searchQuery.value) {
        const query = searchQuery.value.toLowerCase();
        images = images.filter((image) =>
          (image.filename && image.filename.toLowerCase().includes(query)) ||
          (image.path && image.path.toLowerCase().includes(query)) ||
          (image.category && image.category.toLowerCase().includes(query))
        );
      }
      
      return images;
    });

    const currentViewingImage = computed(() => {
      if (viewingImageIndex.value === null) return null;
      return currentImages.value[viewingImageIndex.value] || null;
    });

    const viewImage = (index) => {
      viewingImageIndex.value = index;
      document.body.style.overflow = "hidden";
    };

    const closeImageViewer = () => {
      viewingImageIndex.value = null;
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
      
      window.addEventListener("keydown", handleKeyDown);
    });

    onUnmounted(() => {
      window.removeEventListener("keydown", handleKeyDown);
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
        if (valid) {
        await loadGalleries();
        }
      },
    );

    return {
      selectedCategory,
      searchQuery,
      galleries,
      loading,
      error,
      success,
      submitting,
      categories,
      categoryList,
      currentCategoryData,
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
      viewCategory,
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
      showDeleteAllModal,
      deleteAllInCategory,
      confirmDeleteAll,
      cancelDeleteAll,
    };
  },
};
</script>
