import api from "./api";

export const videoService = {
  // Create new video
  async createVideo(videoData) {
    return await api.post("/videos", videoData);
  },

  // Get all videos
  async getVideos() {
    return await api.get("/videos");
  },

  // Get video by ID
  async getVideo(id) {
    return await api.get(`/videos/${id}`);
  },

  // Update video
  async updateVideo(id, videoData) {
    return await api.put(`/videos/${id}`, videoData);
  },

  // Delete video
  async deleteVideo(id) {
    return await api.delete(`/videos/${id}`);
  },
};
