import api from "./api";

export const userService = {
  // Get all users
  async getUsers() {
    return await api.get("/users");
  },

  // Get user by ID
  async getUser(id) {
    return await api.get(`/users/${id}`);
  },

  // Create new user
  async createUser(userData) {
    return await api.post("/users", userData);
  },

  // Update user
  async updateUser(id, userData) {
    return await api.put(`/users/${id}`, userData);
  },

  // Delete user
  async deleteUser(id) {
    return await api.delete(`/users/${id}`);
  },
};
