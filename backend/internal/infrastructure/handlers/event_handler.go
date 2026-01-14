package handlers

import (
	"errors"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"ccpp-backend/internal/domain/models"
	"ccpp-backend/internal/domain/services"

	"github.com/gin-gonic/gin"
)

type EventHandler struct {
	eventService services.EventService
}

func NewEventHandler(eventService services.EventService) *EventHandler {
	return &EventHandler{eventService: eventService}
}

func normalizeEvent(event *models.Event) error {
	if event.Title == "" {
		return errors.New("title is required")
	}

	if event.StartTime.IsZero() {
		return errors.New("startTime is required")
	}

	if event.EndTime.IsZero() {
		event.EndTime = event.StartTime
	}

	// Maintain legacy date column for compatibility
	if event.Date.IsZero() {
		event.Date = event.StartTime
	}

	if event.EndTime.Before(event.StartTime) {
		return errors.New("endTime must be after startTime")
	}

	allowedRecurrences := map[string]bool{
		"none":    true,
		"daily":   true,
		"weekly":  true,
		"monthly": true,
		"yearly":  true,
		"":        true,
	}

	event.Recurrence = strings.ToLower(strings.TrimSpace(event.Recurrence))
	if event.Recurrence == "" {
		event.Recurrence = "none"
	}

	if !allowedRecurrences[event.Recurrence] {
		return errors.New("recurrence must be one of: none, daily, weekly, monthly, yearly")
	}

	if event.RecurrenceEndsAt != nil && event.RecurrenceEndsAt.Before(event.StartTime) {
		return errors.New("recurrenceEndsAt must be after startTime")
	}

	if event.Timezone == "" {
		if tz := os.Getenv("DEFAULT_TIMEZONE"); tz != "" {
			event.Timezone = tz
		} else {
			event.Timezone = "UTC"
		}
	}

	// Normalize times to UTC for storage while preserving provided instants
	event.StartTime = event.StartTime.UTC()
	event.EndTime = event.EndTime.UTC()
	if event.RecurrenceEndsAt != nil {
		t := event.RecurrenceEndsAt.UTC()
		event.RecurrenceEndsAt = &t
	}

	return nil
}

// CreateEvent handles POST /api/events
func (h *EventHandler) CreateEvent(c *gin.Context) {
	var event models.Event

	if err := c.ShouldBindJSON(&event); err != nil {
		log.Printf("Error binding JSON: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := normalizeEvent(&event); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.eventService.CreateEvent(&event); err != nil {
		log.Printf("Error creating event: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create event"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Event created successfully",
		"event":   event,
	})
}

// GetAllEvents handles GET /api/events
func (h *EventHandler) GetAllEvents(c *gin.Context) {
	upcoming := c.Query("upcoming")
	var events []models.Event
	var err error

	if upcoming == "true" {
		events, err = h.eventService.GetUpcomingEvents()
	} else {
		events, err = h.eventService.GetAllEvents()
	}

	if err != nil {
		log.Printf("Error getting events: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve events"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"events": events,
		"count":  len(events),
	})
}

// GetEventByID handles GET /api/events/:id
func (h *EventHandler) GetEventByID(c *gin.Context) {
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

	c.JSON(http.StatusOK, event)
}

// UpdateEvent handles PUT /api/events/:id
func (h *EventHandler) UpdateEvent(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid event ID"})
		return
	}

	var event models.Event
	if err := c.ShouldBindJSON(&event); err != nil {
		log.Printf("Error binding JSON: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	event.ID = uint(id)

	if err := normalizeEvent(&event); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.eventService.UpdateEvent(&event); err != nil {
		log.Printf("Error updating event: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update event"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Event updated successfully",
		"event":   event,
	})
}

// DeleteEvent handles DELETE /api/events/:id
func (h *EventHandler) DeleteEvent(c *gin.Context) {
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

	c.JSON(http.StatusOK, gin.H{"message": "Event deleted successfully"})
}
