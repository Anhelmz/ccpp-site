import api from "./api";

export const eventService = {
  // Create new event
  async createEvent(eventData) {
    return await api.post("/events", eventData);
  },

  // Get all events
  async getEvents(upcoming = false) {
    const url = upcoming ? "/events?upcoming=true" : "/events";
    return await api.get(url);
  },

  // Get event by ID
  async getEvent(id) {
    return await api.get(`/events/${id}`);
  },

  // Update event
  async updateEvent(id, eventData) {
    return await api.put(`/events/${id}`, eventData);
  },

  // Delete event
  async deleteEvent(id) {
    return await api.delete(`/events/${id}`);
  },
};
