import api from './api'

export const contactRequestService = {
  async createContactRequest(payload) {
    return await api.post('/contact-requests', payload)
  },

  async getContactRequests() {
    return await api.get('/contact-requests')
  },

  async getContactRequest(id) {
    return await api.get(`/contact-requests/${id}`)
  },

  async deleteContactRequest(id) {
    return await api.delete(`/contact-requests/${id}`)
  },
}




