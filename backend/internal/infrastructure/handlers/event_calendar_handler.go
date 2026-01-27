package handlers

import (
	"ccpp-backend/internal/domain/models"
	"ccpp-backend/internal/domain/services"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type EventCalendarHandler struct {
	eventService services.EventCalendarService
}

func NewEventCalendarHandler(eventService services.EventCalendarService) *EventCalendarHandler {
	return &EventCalendarHandler{
		eventService: eventService,
	}
}

// CreateEvent handles POST /api/events
func (h *EventCalendarHandler) CreateEvent(c *gin.Context) {
	var event models.EventCalendar
	if err := c.ShouldBindJSON(&event); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.eventService.CreateEvent(&event); err != nil {
		log.Printf("Error creating event: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create event"})
		return
	}

	// Convert to response format with UTC dates
	response := models.EventCalendarResponse{
		ID:        event.ID,
		Title:     event.Title,
		Details:   event.Details,
		StartDate: event.StartDate.UTC().Format("2006-01-02T15:04:05Z"),
		EndDate:   event.EndDate.UTC().Format("2006-01-02T15:04:05Z"),
		Location:  event.Location,
		CreatedAt: event.CreatedAt,
		UpdatedAt: event.UpdatedAt,
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Event created successfully",
		"event":   response,
	})
}

// GetAllEvents handles GET /api/events
func (h *EventCalendarHandler) GetAllEvents(c *gin.Context) {
	upcoming := c.Query("upcoming") == "true"
	
	var events []models.EventCalendar
	var err error
	
	if upcoming {
		events, err = h.eventService.GetUpcomingEvents()
	} else {
		events, err = h.eventService.GetAllEvents()
	}
	
	if err != nil {
		log.Printf("Error fetching events: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch events"})
		return
	}

	// Convert to response format with UTC dates
	responseEvents := make([]models.EventCalendarResponse, len(events))
	for i, event := range events {
		responseEvents[i] = models.EventCalendarResponse{
			ID:        event.ID,
			Title:     event.Title,
			Details:   event.Details,
			StartDate: event.StartDate.UTC().Format("2006-01-02T15:04:05Z"),
			EndDate:   event.EndDate.UTC().Format("2006-01-02T15:04:05Z"),
			Location:  event.Location,
			CreatedAt: event.CreatedAt,
			UpdatedAt: event.UpdatedAt,
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"events": responseEvents,
		"count":  len(responseEvents),
	})
}

// GetEventByID handles GET /api/events/:id
func (h *EventCalendarHandler) GetEventByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid event ID"})
		return
	}

	event, err := h.eventService.GetEventByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Event not found"})
		return
	}

	// Convert to response format with UTC dates
	response := models.EventCalendarResponse{
		ID:        event.ID,
		Title:     event.Title,
		Details:   event.Details,
		StartDate: event.StartDate.UTC().Format("2006-01-02T15:04:05Z"),
		EndDate:   event.EndDate.UTC().Format("2006-01-02T15:04:05Z"),
		Location:  event.Location,
		CreatedAt: event.CreatedAt,
		UpdatedAt: event.UpdatedAt,
	}

	c.JSON(http.StatusOK, response)
}

// UpdateEvent handles PUT /api/events/:id
func (h *EventCalendarHandler) UpdateEvent(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid event ID"})
		return
	}

	var event models.EventCalendar
	if err := c.ShouldBindJSON(&event); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	event.ID = uint(id)
	if err := h.eventService.UpdateEvent(&event); err != nil {
		log.Printf("Error updating event: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update event"})
		return
	}

	// Convert to response format with UTC dates
	response := models.EventCalendarResponse{
		ID:        event.ID,
		Title:     event.Title,
		Details:   event.Details,
		StartDate: event.StartDate.UTC().Format("2006-01-02T15:04:05Z"),
		EndDate:   event.EndDate.UTC().Format("2006-01-02T15:04:05Z"),
		Location:  event.Location,
		CreatedAt: event.CreatedAt,
		UpdatedAt: event.UpdatedAt,
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Event updated successfully",
		"event":   response,
	})
}

// DeleteEvent handles DELETE /api/events/:id
func (h *EventCalendarHandler) DeleteEvent(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid event ID"})
		return
	}

	if err := h.eventService.DeleteEvent(uint(id)); err != nil {
		log.Printf("Error deleting event: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete event"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Event deleted successfully",
	})
}
