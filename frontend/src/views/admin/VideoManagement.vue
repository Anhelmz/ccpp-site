<template>
  <AdminLayout>
    <!-- Videos List View -->
    <div v-if="!selectedVideo && !showCreateForm">
      <!-- Page Header -->
      <div class="mb-8">
        <h1 class="text-2xl font-semibold text-zinc-900">Videos</h1>
        <p class="mt-1 text-sm text-zinc-600">Manage YouTube videos</p>
      </div>

      <div class="mb-4 flex flex-wrap gap-3 items-center justify-between">
        <!-- Search Bar -->
        <div class="flex-1 max-w-md">
          <div class="relative">
            <input
              v-model="searchQuery"
              type="text"
              placeholder="Search videos..."
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
        <button
          @click="openCreateForm"
          class="px-4 py-2 rounded-md text-sm font-medium transition-colors bg-[#0089AE] text-white hover:bg-[#007A9D] flex items-center gap-2"
        >
          <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4v16m8-8H4"></path>
          </svg>
          New Video
        </button>
      </div>

      <div class="bg-white border border-gray-200 rounded-lg overflow-hidden">
      <div v-if="filteredVideos.length === 0" class="text-center py-12">
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
            d="M15 10l4.553-2.276a1 1 0 010 1.788L15 12.222M9 10L4.447 7.724a1 1 0 000 1.788L9 12.222m6 0l4.553 2.276a1 1 0 010 1.788L15 18m-6-5.778l-4.553 2.276a1 1 0 000 1.788L9 18m0-8l6 3m-6 0l6 3"
          />
        </svg>
        <h3 class="text-lg font-medium text-zinc-900 mb-2">
          {{ searchQuery ? "No videos found" : "No videos added yet" }}
        </h3>
        <p class="text-sm text-zinc-600 mb-4">
          {{
            searchQuery
              ? "Try a different search term"
              : "Add your first video to get started"
          }}
        </p>
        <button
          v-if="!searchQuery"
          class="px-6 py-2 bg-[#0089AE] text-white rounded-lg hover:bg-[#007A9D] transition-colors font-medium shadow-sm"
          @click="openCreateForm"
        >
          Add Video
        </button>
      </div>

      <div v-else>
        <div class="overflow-x-auto">
          <table class="min-w-full divide-y divide-gray-200">
            <thead class="bg-gray-50">
              <tr>
                <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">Title</th>
                <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">Links</th>
                <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">Actions</th>
              </tr>
            </thead>
            <tbody class="bg-white divide-y divide-gray-200">
              <tr
                v-for="video in filteredVideos"
                :key="video.id"
                class="hover:bg-gray-50 transition-colors cursor-pointer"
                @click="viewVideo(video)"
              >
                <td class="px-6 py-4 whitespace-nowrap text-sm font-medium text-gray-900">
                  {{ video.title }}
                </td>
                <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-600">
                  <a
                    :href="
                      video.youtubeUrl ||
                      `https://www.youtube.com/watch?v=${video.youtubeId}`
                    "
                    target="_blank"
                    rel="noopener noreferrer"
                    class="text-xs text-[#0089AE] hover:text-[#007A9D] flex items-center gap-1.5"
                    @click.stop
                  >
                    <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 6H6a2 2 0 00-2 2v10a2 2 0 002 2h10a2 2 0 002-2v-4M14 4h6m0 0v6m0-6L10 14"></path>
                    </svg>
                    Open link
                  </a>
                </td>
                <td class="px-6 py-4 whitespace-nowrap text-right text-sm font-medium" @click.stop>
                  <button
                    class="text-red-600 hover:text-red-800 transition-colors font-medium"
                    @click="deleteVideo(video.id)"
                  >
                    Delete
                  </button>
                </td>
              </tr>
            </tbody>
          </table>
        </div>
      </div>
      </div>
    </div>

    <!-- View Video Detail Page -->
    <div v-if="selectedVideo && !showCreateForm">
      <!-- Breadcrumbs -->
      <nav class="mb-6" aria-label="Breadcrumb">
        <ol class="flex items-center space-x-2 text-sm">
          <li>
            <a href="#" @click.prevent="selectedVideo = null" class="text-gray-500 hover:text-gray-700">
              Videos
            </a>
          </li>
          <li>
            <span class="text-gray-400">/</span>
          </li>
          <li class="text-gray-900 font-medium">
            {{ selectedVideo.title || 'Video' }}
          </li>
        </ol>
      </nav>

      <!-- Video Details -->
      <div class="bg-white border border-gray-200 rounded-lg">
        <div class="flex flex-col lg:flex-row">
          <!-- Main Content Area -->
          <div class="flex-1 p-6 lg:pr-8">
            <h1 class="text-3xl font-bold text-gray-900 mb-6">{{ selectedVideo.title }}</h1>
            
            <!-- Video Embed -->
            <div class="mb-6">
              <div class="aspect-video bg-gray-100 rounded-lg overflow-hidden">
                <iframe
                  :src="`https://www.youtube.com/embed/${selectedVideo.youtubeId}`"
                  class="w-full h-full"
                  frameborder="0"
                  allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture"
                  allowfullscreen
                ></iframe>
              </div>
            </div>

            <!-- Description -->
            <div v-if="selectedVideo.description" class="prose max-w-none">
              <div class="bg-gray-50 rounded-lg border border-gray-200 p-6">
                <p class="text-gray-900 whitespace-pre-wrap text-base leading-relaxed">
                  {{ selectedVideo.description }}
                </p>
              </div>
            </div>
          </div>

          <!-- Sidebar -->
          <div class="lg:w-80 lg:border-l border-gray-200 p-6 space-y-6 bg-gray-50">
            <!-- Video Information -->
            <div>
              <label class="text-xs font-medium text-gray-500 uppercase tracking-wider">Title</label>
              <p class="mt-1 text-sm text-gray-900">{{ selectedVideo.title }}</p>
            </div>

            <div>
              <label class="text-xs font-medium text-gray-500 uppercase tracking-wider">YouTube ID</label>
              <p class="mt-1 text-sm text-gray-600 break-all">{{ selectedVideo.youtubeId }}</p>
            </div>

            <div>
              <label class="text-xs font-medium text-gray-500 uppercase tracking-wider">YouTube Link</label>
              <div class="mt-1">
                <a
                  :href="selectedVideo.youtubeUrl || `https://www.youtube.com/watch?v=${selectedVideo.youtubeId}`"
                  target="_blank"
                  rel="noopener noreferrer"
                  class="text-[#0089AE] hover:text-[#007A9D] text-sm flex items-center gap-1.5 break-all"
                >
                  <svg class="w-4 h-4 flex-shrink-0" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 6H6a2 2 0 00-2 2v10a2 2 0 002 2h10a2 2 0 002-2v-4M14 4h6m0 0v6m0-6L10 14"></path>
                  </svg>
                  Open on YouTube
                </a>
              </div>
            </div>

            <!-- Action Buttons -->
            <div class="pt-4 border-t border-gray-200 space-y-3">
              <button
                @click="editVideo"
                class="w-full px-4 py-2 bg-[#0089AE] text-white text-sm font-medium rounded-md hover:bg-[#007A9D] transition-colors"
              >
                Edit Video
              </button>
              <button
                @click="selectedVideo = null"
                class="w-full px-4 py-2 border border-gray-300 text-gray-700 text-sm font-medium rounded-md hover:bg-gray-50 transition-colors"
              >
                Back to List
              </button>
              <button
                @click="deleteVideo(selectedVideo.id)"
                class="w-full px-4 py-2 bg-red-600 text-white text-sm font-medium rounded-md hover:bg-red-700 transition-colors"
              >
                Delete Video
              </button>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- Create/Edit Video Form -->
    <div v-if="showCreateForm || editingVideo">
      <!-- Breadcrumbs -->
      <nav class="mb-6" aria-label="Breadcrumb">
        <ol class="flex items-center space-x-2 text-sm">
          <li>
            <a href="#" @click.prevent="closeForm" class="text-gray-500 hover:text-gray-700">
              Videos
            </a>
          </li>
          <li>
            <span class="text-gray-400">/</span>
          </li>
          <li class="text-gray-900 font-medium">
            {{ editingVideo ? 'Edit Video' : 'New Video' }}
          </li>
        </ol>
      </nav>

      <!-- Video Form -->
      <div class="bg-white border border-gray-200 rounded-lg p-6 space-y-6">
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">
            Title <span class="text-red-500">*</span>
          </label>
          <input
            v-model="videoForm.title"
            type="text"
            required
            class="w-full px-4 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-[#0089AE]"
            placeholder="Teaching title"
          />
        </div>

        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">
            Description
          </label>
          <textarea
            v-model="videoForm.description"
            rows="3"
            class="w-full px-4 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-[#0089AE]"
            placeholder="Brief description"
          ></textarea>
        </div>

        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">
            YouTube Link or ID <span class="text-red-500">*</span>
          </label>
          <input
            v-model="videoForm.youtubeInput"
            type="text"
            required
            class="w-full px-4 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-[#0089AE]"
            placeholder="https://www.youtube.com/watch?v=XXXXXXXXXXX"
          />
          <p class="text-xs text-gray-500 mt-1">
            Paste a full YouTube URL or the 11-character video ID.
          </p>
        </div>

        <div
          v-if="formError"
          class="p-4 bg-red-50 border border-red-200 rounded-lg"
        >
          <p class="text-sm text-red-800">{{ formError }}</p>
        </div>

        <div class="flex justify-end gap-3 pt-4 border-t border-gray-200">
          <button
            type="button"
            class="px-4 py-2 border border-gray-300 text-gray-700 rounded-md hover:bg-gray-50 transition-colors"
            @click="closeForm"
          >
            Cancel
          </button>
          <button
            type="button"
            @click="saveVideo"
            class="px-4 py-2 bg-[#0089AE] text-white rounded-md hover:bg-[#007A9D] transition-colors font-medium disabled:opacity-50 disabled:cursor-not-allowed"
            :disabled="saving"
          >
            {{ saving ? 'Saving...' : editingVideo ? 'Update Video' : 'Create Video' }}
          </button>
        </div>
      </div>
    </div>

    <!-- Confirmation Modal -->
    <ConfirmationModal
      :show="showConfirmModal"
      title="Delete Video"
      message="Are you sure you want to delete this video? This action cannot be undone."
      confirm-text="Delete Video"
      @confirm="confirmDelete"
      @cancel="cancelDelete"
    />
  </AdminLayout>
</template>

<script>
import { ref, computed, onMounted } from "vue";
import AdminLayout from "@/components/admin/AdminLayout.vue";
import ConfirmationModal from "@/components/admin/ConfirmationModal.vue";
import { videoStore } from "@/services/videoStore";

const extractYouTubeId = (input) => {
  const trimmed = (input || "").trim();
  if (!trimmed) return "";
  if (trimmed.length === 11) return trimmed;
  const match = trimmed.match(/(?:v=|\/)([0-9A-Za-z_-]{11})/);
  return match ? match[1] : "";
};

export default {
  name: "VideoManagement",
  components: {
    AdminLayout,
    ConfirmationModal,
  },
  setup() {
    const selectedVideo = ref(null);
    const showCreateForm = ref(false);
    const editingVideo = ref(null);
    const videos = ref([]);
    const formError = ref("");
    const searchQuery = ref("");
    const showConfirmModal = ref(false);
    const videoToDelete = ref(null);
    const saving = ref(false);

    const videoForm = ref({
      title: "",
      description: "",
      youtubeInput: "",
    });

    const loadVideos = () => {
      videos.value = videoStore.getAll();
    };

    const filteredVideos = computed(() => {
      let list = [...videos.value];
      if (searchQuery.value) {
        const q = searchQuery.value.toLowerCase();
        list = list.filter(
          (video) =>
            video.title.toLowerCase().includes(q) ||
            (video.description &&
              video.description.toLowerCase().includes(q)) ||
            (video.youtubeId && video.youtubeId.toLowerCase().includes(q)),
        );
      }
      return list;
    });

    const viewVideo = (video) => {
      selectedVideo.value = video;
      showCreateForm.value = false;
      editingVideo.value = null;
    };

    const openCreateForm = () => {
      selectedVideo.value = null;
      editingVideo.value = null;
      showCreateForm.value = true;
      videoForm.value = {
        title: "",
        description: "",
        youtubeInput: "",
      };
      formError.value = "";
    };

    const editVideo = () => {
      if (!selectedVideo.value) return;
      editingVideo.value = selectedVideo.value;
      showCreateForm.value = true;
      selectedVideo.value = null;
      videoForm.value = {
        title: editingVideo.value.title,
        description: editingVideo.value.description || "",
        youtubeInput: editingVideo.value.youtubeUrl || editingVideo.value.youtubeId || "",
      };
      formError.value = "";
    };

    const closeForm = () => {
      showCreateForm.value = false;
      editingVideo.value = null;
      selectedVideo.value = null;
      formError.value = "";
    };

    const saveVideo = () => {
      formError.value = "";
      saving.value = true;
      const id = extractYouTubeId(videoForm.value.youtubeInput);
      if (!videoForm.value.title.trim() || !id) {
        formError.value = "Title and a valid YouTube link/ID are required.";
        saving.value = false;
        return;
      }

      const payload = {
        title: videoForm.value.title.trim(),
        description: videoForm.value.description.trim(),
        youtubeId: id,
        youtubeUrl: `https://www.youtube.com/watch?v=${id}`,
      };

      if (editingVideo.value) {
        videoStore.update(editingVideo.value.id, payload);
      } else {
        videoStore.add(payload);
      }

      loadVideos();
      saving.value = false;
      
      // After saving, view the video if editing, or go back to list if creating
      if (editingVideo.value) {
        const updatedVideo = videoStore.getAll().find(v => v.id === editingVideo.value.id);
        if (updatedVideo) {
          selectedVideo.value = updatedVideo;
          showCreateForm.value = false;
          editingVideo.value = null;
        } else {
          closeForm();
        }
      } else {
        closeForm();
      }
    };

    const deleteVideo = (id) => {
      videoToDelete.value = id;
      showConfirmModal.value = true;
    };

    const confirmDelete = () => {
      if (!videoToDelete.value) return;
      videoStore.remove(videoToDelete.value);
      loadVideos();
      showConfirmModal.value = false;
      videoToDelete.value = null;
    };

    const cancelDelete = () => {
      showConfirmModal.value = false;
      videoToDelete.value = null;
    };

    onMounted(() => {
      loadVideos();
    });

    return {
      selectedVideo,
      showCreateForm,
      editingVideo,
      videos,
      formError,
      searchQuery,
      videoForm,
      filteredVideos,
      saving,
      loadVideos,
      viewVideo,
      openCreateForm,
      editVideo,
      closeForm,
      saveVideo,
      deleteVideo,
      showConfirmModal,
      confirmDelete,
      cancelDelete,
    };
  },
};
</script>

