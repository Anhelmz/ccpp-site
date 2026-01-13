import api from "./api";

export const contactService = {
  // Create new contact submission
  async createContact(contactData) {
    return await api.post("/contacts", contactData);
  },

  // Get all contacts (admin use)
  async getContacts() {
    return await api.get("/contacts");
  },

  // Get contact by ID (admin use)
  async getContact(id) {
    return await api.get(`/contacts/${id}`);
  },

  // Delete contact (admin use)
  async deleteContact(id) {
    return await api.delete(`/contacts/${id}`);
  },
};
