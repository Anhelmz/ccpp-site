<template>
  <AdminLayout>
    <!-- Page Header -->
    <div class="mb-8">
      <h1 class="text-2xl font-semibold text-zinc-900">Videos</h1>
      <p class="mt-1 text-sm text-zinc-600">Manage YouTube videos</p>
    </div>

    <div class="mb-4 flex flex-wrap gap-3 items-center justify-end">
      <button
        @click="openCreateModal"
        class="px-4 py-2 rounded-md text-sm font-medium transition-colors bg-[#23D3EE] text-white hover:bg-[#1FC5D9] flex items-center gap-2"
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
          class="px-6 py-2 bg-[#23D3EE] text-white rounded-lg hover:bg-[#1FC5D9] transition-colors font-medium shadow-sm"
          @click="openCreateModal"
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
                @click="openEditModal(video)"
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
                    class="text-xs text-[#23D3EE] hover:text-[#1FC5D9]"
                    @click.stop
                  >
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

    <div
      v-if="showModal"
      class="fixed inset-0 bg-black/50 backdrop-blur-sm z-50 flex items-center justify-center p-4"
      @click="closeModal"
    >
      <div
        class="bg-white rounded-xl border border-gray-200 shadow-2xl max-w-xl w-full max-h-[90vh] overflow-y-auto"
        @click.stop
      >
        <div class="p-8 border-b border-gray-200">
          <div class="flex justify-between items-start mb-6">
            <h3 class="text-3xl font-bold text-gray-900">
              {{ editingVideo ? "Edit Video" : "New Video" }}
            </h3>
            <button
              class="text-gray-400 hover:text-gray-600 transition-colors"
              @click="closeModal"
            >
              <svg
                class="w-6 h-6"
                fill="none"
                stroke="currentColor"
                viewBox="0 0 24 24"
              >
                <path
                  stroke-linecap="round"
                  stroke-linejoin="round"
                  stroke-width="2"
                  d="M6 18L18 6M6 6l12 12"
                ></path>
              </svg>
            </button>
          </div>
        </div>

        <form class="p-8 space-y-6" @submit.prevent="saveVideo">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">
              Title *
            </label>
            <input
              v-model="videoForm.title"
              type="text"
              required
              class="w-full border border-gray-300 rounded-md px-3 py-2 text-gray-900 bg-white focus:ring-2 focus:ring-[#23D3EE] focus:border-[#23D3EE]"
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
              class="w-full border border-gray-300 rounded-md px-3 py-2 text-gray-900 bg-white focus:ring-2 focus:ring-[#23D3EE] focus:border-[#23D3EE]"
              placeholder="Brief description"
            ></textarea>
          </div>

          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">
              YouTube Link or ID *
            </label>
            <input
              v-model="videoForm.youtubeInput"
              type="text"
              required
              class="w-full border border-gray-300 rounded-md px-3 py-2 text-gray-900 bg-white focus:ring-2 focus:ring-[#23D3EE] focus:border-[#23D3EE]"
              placeholder="https://www.youtube.com/watch?v=XXXXXXXXXXX"
            />
            <p class="text-xs text-gray-500 mt-1">
              Paste a full YouTube URL or the 11-character video ID.
            </p>
          </div>

          <div
            v-if="formError"
            class="p-4 bg-red-50 border border-red-200 rounded-2xl"
          >
            <p class="text-red-600 text-sm">{{ formError }}</p>
          </div>

          <div class="flex justify-end gap-3 pt-4">
            <button
              type="button"
              class="px-4 py-2 border border-gray-300 rounded-md text-gray-700 hover:bg-gray-50 transition-colors"
              @click="closeModal"
            >
              Cancel
            </button>
            <button
              type="submit"
              class="px-5 py-2 bg-[#23D3EE] text-white rounded-md hover:bg-[#1FC5D9] transition-colors disabled:opacity-50"
              :disabled="saving"
            >
              {{ saving ? 'Saving...' : editingVideo ? 'Update Video' : 'Create Video' }}
            </button>
          </div>
        </form>
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
    const showModal = ref(false);
    const editingVideo = ref(null);
    const videos = ref([]);
    const formError = ref("");
    const searchQuery = ref("");
    const showConfirmModal = ref(false);
    const videoToDelete = ref(null);

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

    const openCreateModal = () => {
      editingVideo.value = null;
      videoForm.value = {
        title: "",
        description: "",
        youtubeInput: "",
      };
      formError.value = "";
      showModal.value = true;
    };

    const openEditModal = (video) => {
      editingVideo.value = video;
      videoForm.value = {
        title: video.title,
        description: video.description || "",
        youtubeInput: video.youtubeUrl || video.youtubeId || "",
      };
      formError.value = "";
      showModal.value = true;
    };

    const closeModal = () => {
      showModal.value = false;
      editingVideo.value = null;
      formError.value = "";
    };

    const saveVideo = () => {
      formError.value = "";
      const id = extractYouTubeId(videoForm.value.youtubeInput);
      if (!videoForm.value.title.trim() || !id) {
        formError.value = "Title and a valid YouTube link/ID are required.";
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
      closeModal();
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
      showModal,
      editingVideo,
      videos,
      formError,
      searchQuery,
      videoForm,
      filteredVideos,
      loadVideos,
      openCreateModal,
      openEditModal,
      closeModal,
      saveVideo,
      deleteVideo,
      showConfirmModal,
      confirmDelete,
      cancelDelete,
    };
  },
};
</script>

