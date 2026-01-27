package handlers

import (
	"ccpp-backend/internal/domain/models"
	"ccpp-backend/internal/domain/services"
	"encoding/base64"
	"fmt"
	"io"
	"log"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type GalleryHandler struct {
	galleryService services.GalleryService
}

type CreateGalleryRequest struct {
	Path     string `json:"path"`
	Filename string `json:"filename"`
	Category string `json:"category" binding:"required"`
}

func NewGalleryHandler(galleryService services.GalleryService) *GalleryHandler {
	return &GalleryHandler{
		galleryService: galleryService,
	}
}

// CreateGallery handles POST /api/gallery (for URL/path)
func (h *GalleryHandler) CreateGallery(c *gin.Context) {
	var req CreateGalleryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate path is not empty
	if strings.TrimSpace(req.Path) == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Path is required"})
		return
	}

	// Validate filename is not empty
	if strings.TrimSpace(req.Filename) == "" {
		// If filename not provided, extract from path
		req.Filename = filepath.Base(req.Path)
	}

	// Validate category
	if strings.TrimSpace(req.Category) == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Category is required"})
		return
	}

	// Create gallery record
	gallery := &models.Gallery{
		Filename: req.Filename,
		Path:     req.Path,
		Category: req.Category,
	}

	if err := h.galleryService.CreateGallery(gallery); err != nil {
		log.Printf("Error creating gallery record: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save gallery record"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Gallery image added successfully",
		"gallery": gallery,
	})
}

// UploadGallery handles POST /api/gallery/upload (for file uploads)
func (h *GalleryHandler) UploadGallery(c *gin.Context) {
	// Get category from form
	category := c.PostForm("category")
	if category == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Category is required"})
		return
	}

	// Parse multipart form with memory limit to prevent file saving
	err := c.Request.ParseMultipartForm(10 << 20) // 10 MB max
	if err != nil {
		log.Printf("Error parsing multipart form: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to parse form"})
		return
	}

	// Get the file from form
	fileHeader := c.Request.MultipartForm.File["image"]
	if len(fileHeader) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Image file is required"})
		return
	}

	file := fileHeader[0]

	// Validate file type
	ext := strings.ToLower(filepath.Ext(file.Filename))
	allowedExts := []string{".jpg", ".jpeg", ".png", ".gif"}
	isValid := false
	var mimeType string
	for _, allowedExt := range allowedExts {
		if ext == allowedExt {
			isValid = true
			switch ext {
			case ".jpg", ".jpeg":
				mimeType = "image/jpeg"
			case ".png":
				mimeType = "image/png"
			case ".gif":
				mimeType = "image/gif"
			}
			break
		}
	}
	if !isValid {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid file type. Only JPG, PNG, and GIF are allowed"})
		return
	}

	// Validate file size (10MB max)
	if file.Size > 10*1024*1024 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "File size exceeds 10MB limit"})
		return
	}

	// Open the uploaded file (this reads from memory, not disk)
	src, err := file.Open()
	if err != nil {
		log.Printf("Error opening file: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read file"})
		return
	}
	defer src.Close()

	// Read file content directly from memory
	fileBytes, err := io.ReadAll(src)
	if err != nil {
		log.Printf("Error reading file: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read file"})
		return
	}

	// Encode to base64
	base64Data := base64.StdEncoding.EncodeToString(fileBytes)
	
	// Generate data URI for path field (for backward compatibility)
	dataURI := fmt.Sprintf("data:%s;base64,%s", mimeType, base64Data)

	// Generate unique filename
	timestamp := time.Now().Unix()
	filename := filepath.Base(file.Filename)
	nameWithoutExt := strings.TrimSuffix(filename, ext)
	uniqueFilename := fmt.Sprintf("%s_%d%s", nameWithoutExt, timestamp, ext)

	// Create gallery record with image data stored in database
	gallery := &models.Gallery{
		Filename:  uniqueFilename,
		Path:      dataURI, // Store data URI for backward compatibility
		ImageData: base64Data,
		MimeType:  mimeType,
		Category:  category,
	}

	if err := h.galleryService.CreateGallery(gallery); err != nil {
		log.Printf("Error creating gallery record: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save gallery record"})
		// Clean up multipart form even on error
		if c.Request.MultipartForm != nil {
			c.Request.MultipartForm.RemoveAll()
		}
		return
	}

	// Clean up multipart form to free memory and ensure no temp files remain
	if c.Request.MultipartForm != nil {
		c.Request.MultipartForm.RemoveAll()
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Image uploaded successfully",
		"gallery": gallery,
	})
}

// GetAllGalleries handles GET /api/gallery
func (h *GalleryHandler) GetAllGalleries(c *gin.Context) {
	category := c.Query("category")
	var galleries []models.Gallery
	var err error

	if category != "" {
		galleries, err = h.galleryService.GetGalleriesByCategory(category)
	} else {
		galleries, err = h.galleryService.GetAllGalleries()
	}

	if err != nil {
		log.Printf("Error getting galleries: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve galleries"})
		return
	}

	// Ensure Path field contains data URI for images stored in database
	for i := range galleries {
		if galleries[i].ImageData != "" && !strings.HasPrefix(galleries[i].Path, "data:") {
			// Convert ImageData to data URI if Path doesn't already have it
			galleries[i].Path = fmt.Sprintf("data:%s;base64,%s", galleries[i].MimeType, galleries[i].ImageData)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"galleries": galleries,
		"count":     len(galleries),
	})
}

// GetGalleryByID handles GET /api/gallery/:id
func (h *GalleryHandler) GetGalleryByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid gallery ID"})
		return
	}

	gallery, err := h.galleryService.GetGalleryByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Gallery not found"})
		return
	}

	// Ensure Path field contains data URI for images stored in database
	if gallery.ImageData != "" && !strings.HasPrefix(gallery.Path, "data:") {
		gallery.Path = fmt.Sprintf("data:%s;base64,%s", gallery.MimeType, gallery.ImageData)
	}

	c.JSON(http.StatusOK, gallery)
}

// DeleteGallery handles DELETE /api/gallery/:id
func (h *GalleryHandler) DeleteGallery(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid gallery ID"})
		return
	}

	// Delete from database (hard delete - permanently remove from database)
	if err := h.galleryService.DeleteGallery(uint(id)); err != nil {
		log.Printf("Error deleting gallery ID %d: %v", id, err)
		if err.Error() == "record not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": "Gallery not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete gallery"})
		return
	}

	log.Printf("Successfully deleted gallery ID %d from database", id)
	c.JSON(http.StatusOK, gin.H{"message": "Gallery deleted successfully"})
}

// GetGalleryImage handles GET /api/gallery/:id/image (serves image from database)
func (h *GalleryHandler) GetGalleryImage(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid gallery ID"})
		return
	}

	gallery, err := h.galleryService.GetGalleryByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Gallery not found"})
		return
	}

	// Decode base64 image data
	imageData, err := base64.StdEncoding.DecodeString(gallery.ImageData)
	if err != nil {
		log.Printf("Error decoding image data: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decode image"})
		return
	}

	// Set content type and serve image
	c.Data(http.StatusOK, gallery.MimeType, imageData)
}

// DeleteGalleriesByCategory handles DELETE /api/gallery/category/:category
func (h *GalleryHandler) DeleteGalleriesByCategory(c *gin.Context) {
	category := c.Param("category")
	if category == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Category is required"})
		return
	}

	// Get count before deletion for response
	galleries, err := h.galleryService.GetGalleriesByCategory(category)
	if err != nil {
		log.Printf("Error getting galleries by category: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve galleries"})
		return
	}

	count := len(galleries)

	// Delete all galleries in this category
	if err := h.galleryService.DeleteGalleriesByCategory(category); err != nil {
		log.Printf("Error deleting galleries by category: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete galleries"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Successfully deleted %d gallery image(s) from category '%s'", count, category),
		"count":   count,
	})
}

// GetOldPhotos handles GET /api/gallery/old - returns old photos (soft-deleted or without ImageData)
func (h *GalleryHandler) GetOldPhotos(c *gin.Context) {
	oldPhotos, err := h.galleryService.GetOldPhotos()
	if err != nil {
		log.Printf("Error getting old photos: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve old photos"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"oldPhotos": oldPhotos,
		"count":     len(oldPhotos),
	})
}

// DeleteOldPhotos handles DELETE /api/gallery/old - permanently deletes old photos
func (h *GalleryHandler) DeleteOldPhotos(c *gin.Context) {
	deletedCount, err := h.galleryService.DeleteOldPhotos()
	if err != nil {
		log.Printf("Error deleting old photos: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete old photos"})
		return
	}

	log.Printf("Successfully deleted %d old photos from database", deletedCount)
	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Successfully deleted %d old photo(s) from database", deletedCount),
		"count":   deletedCount,
	})
}
