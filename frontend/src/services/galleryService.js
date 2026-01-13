import api from "./api";

export const galleryService = {
  // Upload gallery image
  async uploadImage(formData) {
    return await api.post("/gallery/upload", formData);
  },

  // Get all galleries
  async getGalleries(category = null) {
    const url = category ? `/gallery?category=${category}` : "/gallery";
    return await api.get(url);
  },

  // Get gallery by ID
  async getGallery(id) {
    return await api.get(`/gallery/${id}`);
  },

  // Delete gallery
  async deleteGallery(id) {
    return await api.delete(`/gallery/${id}`);
  },
};
