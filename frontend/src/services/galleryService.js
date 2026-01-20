import api from "./api";

export const galleryService = {
  // Create gallery image (using URL/path)
  async createGallery(data) {
    return await api.post("/gallery", data);
  },

  // Upload gallery image (file upload)
  async uploadImage(formData) {
    return await api.post("/gallery/upload", formData, {
      headers: {
        "Content-Type": "multipart/form-data",
      },
    });
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
