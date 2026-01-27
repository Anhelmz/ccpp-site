package handlers

import (
	"ccpp-backend/internal/domain/models"
	"ccpp-backend/internal/domain/services"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type VideoHandler struct {
	videoService services.VideoService
}

func NewVideoHandler(videoService services.VideoService) *VideoHandler {
	return &VideoHandler{
		videoService: videoService,
	}
}

// extractYouTubeID extracts YouTube video ID from URL or returns the ID if it's already an ID
func extractYouTubeID(input string) string {
	input = strings.TrimSpace(input)
	if input == "" {
		return ""
	}
	
	// If it's already an 11-character ID, return it
	if len(input) == 11 {
		return input
	}
	
	// Try to extract from various YouTube URL formats
	patterns := []string{
		"youtube.com/watch?v=",
		"youtu.be/",
		"youtube.com/embed/",
		"youtube.com/v/",
	}
	
	for _, pattern := range patterns {
		idx := strings.Index(input, pattern)
		if idx != -1 {
			start := idx + len(pattern)
			id := input[start:]
			// Remove any query parameters
			if ampIdx := strings.Index(id, "&"); ampIdx != -1 {
				id = id[:ampIdx]
			}
			if len(id) == 11 {
				return id
			}
		}
	}
	
	return ""
}

// CreateVideo handles POST /api/videos
func (h *VideoHandler) CreateVideo(c *gin.Context) {
	var req struct {
		Title       string `json:"title" binding:"required"`
		Description string `json:"description"`
		YouTubeID   string `json:"youtubeId"`
		YouTubeURL  string `json:"youtubeUrl"`
		YouTubeInput string `json:"youtubeInput"` // For accepting URL or ID
	}
	
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Extract YouTube ID from input
	youtubeID := req.YouTubeID
	if youtubeID == "" && req.YouTubeInput != "" {
		youtubeID = extractYouTubeID(req.YouTubeInput)
	}
	
	if youtubeID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Valid YouTube URL or ID is required"})
		return
	}

	// Build YouTube URL if not provided
	youtubeURL := req.YouTubeURL
	if youtubeURL == "" {
		youtubeURL = "https://www.youtube.com/watch?v=" + youtubeID
	}

	video := &models.Video{
		Title:       strings.TrimSpace(req.Title),
		Description: strings.TrimSpace(req.Description),
		YouTubeID:   youtubeID,
		YouTubeURL:  youtubeURL,
	}

	if err := h.videoService.CreateVideo(video); err != nil {
		log.Printf("Error creating video: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create video", "details": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Video created successfully",
		"video":   video,
	})
}

// GetAllVideos handles GET /api/videos
func (h *VideoHandler) GetAllVideos(c *gin.Context) {
	videos, err := h.videoService.GetAllVideos()
	if err != nil {
		log.Printf("Error fetching videos: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch videos"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"videos": videos,
		"count":  len(videos),
	})
}

// GetVideoByID handles GET /api/videos/:id
func (h *VideoHandler) GetVideoByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid video ID"})
		return
	}

	video, err := h.videoService.GetVideoByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Video not found"})
		return
	}

	c.JSON(http.StatusOK, video)
}

// UpdateVideo handles PUT /api/videos/:id
func (h *VideoHandler) UpdateVideo(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid video ID"})
		return
	}

	var req struct {
		Title       string `json:"title" binding:"required"`
		Description string `json:"description"`
		YouTubeID   string `json:"youtubeId"`
		YouTubeURL  string `json:"youtubeUrl"`
		YouTubeInput string `json:"youtubeInput"` // For accepting URL or ID
	}
	
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Extract YouTube ID from input
	youtubeID := req.YouTubeID
	if youtubeID == "" && req.YouTubeInput != "" {
		youtubeID = extractYouTubeID(req.YouTubeInput)
	}
	
	if youtubeID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Valid YouTube URL or ID is required"})
		return
	}

	// Build YouTube URL if not provided
	youtubeURL := req.YouTubeURL
	if youtubeURL == "" {
		youtubeURL = "https://www.youtube.com/watch?v=" + youtubeID
	}

	video := &models.Video{
		ID:          uint(id),
		Title:       req.Title,
		Description: req.Description,
		YouTubeID:   youtubeID,
		YouTubeURL:  youtubeURL,
	}

	if err := h.videoService.UpdateVideo(video); err != nil {
		log.Printf("Error updating video: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update video"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Video updated successfully",
		"video":   video,
	})
}

// DeleteVideo handles DELETE /api/videos/:id
func (h *VideoHandler) DeleteVideo(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid video ID"})
		return
	}

	if err := h.videoService.DeleteVideo(uint(id)); err != nil {
		log.Printf("Error deleting video: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete video"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Video deleted successfully",
	})
}
