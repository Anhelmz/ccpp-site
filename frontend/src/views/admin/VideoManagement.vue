<template>
  <AdminLayout>
    <div class="mb-6 flex justify-end">
      <button
        @click="openCreateModal"
        class="px-6 py-3 bg-main text-white rounded-lg hover:opacity-90 transition-colors shadow-lg font-medium flex items-center space-x-2"
      >
        <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4v16m8-8H4"></path>
        </svg>
        <span>New Video</span>
      </button>
    </div>

    <div class="bg-white rounded-lg shadow">
      <div class="p-6 border-b border-gray-200">
        <div class="flex flex-col sm:flex-row justify-between items-start sm:items-center gap-4">
          <h2 class="text-xl font-semibold text-gray-900">Videos</h2>
          <div class="flex items-center space-x-2 w-full sm:w-auto">
            <input
              v-model="searchQuery"
              type="text"
              placeholder="Search videos..."
              class="flex-1 sm:flex-none px-4 py-2 border border-gray-300 rounded-lg text-sm min-w-48"
            />
          </div>
        </div>
      </div>

      <div v-if="filteredVideos.length === 0" class="text-center py-12">
        <svg class="mx-auto h-12 w-12 text-gray-400 mb-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 10l4.553-2.276a1 1 0 010 1.788L15 12.222M9 10L4.447 7.724a1 1 0 000 1.788L9 12.222m6 0l4.553 2.276a1 1 0 010 1.788L15 18m-6-5.778l-4.553 2.276a1 1 0 000 1.788L9 18m0-8l6 3m-6 0l6 3" />
        </svg>
        <p class="text-gray-600 mb-2">{{ searchQuery ? 'No videos found' : 'No videos added yet' }}</p>
        <p class="text-sm text-gray-500 mb-4">{{ searchQuery ? 'Try a different search term' : 'Add your first video to get started' }}</p>
        <button
          v-if="!searchQuery"
          @click="openCreateModal"
          class="px-6 py-2 bg-main text-white rounded-lg hover:opacity-90 transition-colors"
        >
          Add Video
        </button>
      </div>

      <div v-else class="overflow-x-auto">
        <table class="min-w-full divide-y divide-gray-200">
          <thead class="bg-gray-50">
            <tr>
              <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">Title</th>
              <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">YouTube ID</th>
              <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">Actions</th>
            </tr>
          </thead>
          <tbody class="bg-white divide-y divide-gray-200">
            <tr 
              v-for="video in filteredVideos" 
              :key="video.id"
              class="hover:bg-gray-50 transition-colors"
            >
              <td class="px-6 py-4 whitespace-nowrap">
                <div class="text-sm font-medium text-gray-900">{{ video.title }}</div>
                <div v-if="video.description" class="text-sm text-gray-500 truncate max-w-xs">{{ video.description }}</div>
              </td>
              <td class="px-6 py-4 whitespace-nowrap">
                <div class="text-sm text-gray-900">{{ video.youtubeId }}</div>
                <a
                  :href="video.youtubeUrl || `https://www.youtube.com/watch?v=${video.youtubeId}`"
                  target="_blank"
                  rel="noopener noreferrer"
                  class="text-xs text-main hover:underline"
                >
                  Open link
                </a>
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm font-medium">
                <button 
                  @click="openEditModal(video)"
                  class="text-blue-600 hover:text-blue-800 mr-4"
                >
                  Edit
                </button>
                <button 
                  @click="deleteVideo(video.id)"
                  class="text-red-600 hover:text-red-800"
                >
                  Delete
                </button>
              </td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>

    <div
      v-if="showModal"
      class="fixed inset-0 bg-black bg-opacity-50 z-50 flex items-center justify-center p-4"
      @click="closeModal"
    >
      <div
        class="bg-white rounded-lg shadow-xl max-w-xl w-full max-h-[90vh] overflow-y-auto"
        @click.stop
      >
        <div class="p-6 border-b border-gray-200">
          <div class="flex justify-between items-center">
            <h3 class="text-xl font-semibold text-gray-900">
              {{ editingVideo ? 'Edit Video' : 'Add New Video' }}
            </h3>
            <button
              @click="closeModal"
              class="text-gray-400 hover:text-gray-600"
            >
              <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"></path>
              </svg>
            </button>
          </div>
        </div>
        
        <form @submit.prevent="saveVideo" class="p-6 space-y-6">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">
              Title *
            </label>
            <input
              v-model="videoForm.title"
              type="text"
              required
              class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-main focus:border-main"
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
              class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-main focus:border-main"
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
              class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-main focus:border-main"
              placeholder="https://www.youtube.com/watch?v=XXXXXXXXXXX"
            />
            <p class="text-xs text-gray-500 mt-1">Paste a full YouTube URL or the 11-character video ID.</p>
          </div>

          <div v-if="formError" class="p-4 bg-red-50 border border-red-200 rounded-lg">
            <p class="text-red-600 text-sm">{{ formError }}</p>
          </div>
          
          <div class="flex justify-end space-x-4 pt-4 border-t border-gray-200">
            <button
              type="button"
              @click="closeModal"
              class="px-6 py-2 border border-gray-300 text-gray-700 rounded-lg hover:bg-gray-50 transition-colors"
            >
              Cancel
            </button>
            <button
              type="submit"
              class="px-6 py-2 bg-main text-white rounded-lg hover:opacity-90 transition-colors"
            >
              {{ editingVideo ? 'Update Video' : 'Add Video' }}
            </button>
          </div>
        </form>
      </div>
    </div>
  </AdminLayout>
</template>

<script>
import { ref, computed, onMounted } from 'vue'
import AdminLayout from '@/components/admin/AdminLayout.vue'
import { videoStore } from '@/services/videoStore'

const extractYouTubeId = (input) => {
  const trimmed = (input || '').trim()
  if (!trimmed) return ''
  if (trimmed.length === 11) return trimmed
  const match = trimmed.match(/(?:v=|\/)([0-9A-Za-z_-]{11})/)
  return match ? match[1] : ''
}

export default {
  name: 'VideoManagement',
  components: {
    AdminLayout
  },
  setup() {
    const showModal = ref(false)
    const editingVideo = ref(null)
    const videos = ref([])
    const formError = ref('')
    const searchQuery = ref('')

    const videoForm = ref({
      title: '',
      description: '',
      youtubeInput: ''
    })

    const loadVideos = () => {
      videos.value = videoStore.getAll()
    }

    const filteredVideos = computed(() => {
      let list = [...videos.value]
      if (searchQuery.value) {
        const q = searchQuery.value.toLowerCase()
        list = list.filter(
          (video) =>
            video.title.toLowerCase().includes(q) ||
            (video.description && video.description.toLowerCase().includes(q)) ||
            (video.youtubeId && video.youtubeId.toLowerCase().includes(q))
        )
      }
      return list
    })

    const openCreateModal = () => {
      editingVideo.value = null
      videoForm.value = {
        title: '',
        description: '',
        youtubeInput: ''
      }
      formError.value = ''
      showModal.value = true
    }

    const openEditModal = (video) => {
      editingVideo.value = video
      videoForm.value = {
        title: video.title,
        description: video.description || '',
        youtubeInput: video.youtubeUrl || video.youtubeId || ''
      }
      formError.value = ''
      showModal.value = true
    }

    const closeModal = () => {
      showModal.value = false
      editingVideo.value = null
      formError.value = ''
    }

    const saveVideo = () => {
      formError.value = ''
      const id = extractYouTubeId(videoForm.value.youtubeInput)
      if (!videoForm.value.title.trim() || !id) {
        formError.value = 'Title and a valid YouTube link/ID are required.'
        return
      }

      const payload = {
        title: videoForm.value.title.trim(),
        description: videoForm.value.description.trim(),
        youtubeId: id,
        youtubeUrl: `https://www.youtube.com/watch?v=${id}`
      }

      if (editingVideo.value) {
        videoStore.update(editingVideo.value.id, payload)
      } else {
        videoStore.add(payload)
      }

      loadVideos()
      closeModal()
    }

    const deleteVideo = (id) => {
      if (!confirm('Are you sure you want to delete this video?')) {
        return
      }
      videoStore.remove(id)
      loadVideos()
    }

    onMounted(() => {
      loadVideos()
    })

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
      deleteVideo
    }
  }
}
</script>
