package main

import (
	"log"
	"os"

	"ccpp-backend/internal/domain/models"
	"ccpp-backend/internal/infrastructure/database"
	infraRepos "ccpp-backend/internal/infrastructure/repositories"
	"ccpp-backend/internal/infrastructure/routes"
	infraServices "ccpp-backend/internal/infrastructure/services"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	// Initialize database
	database.InitDB()
	db := database.GetDB()

	// Handle mime_type migration BEFORE AutoMigrate
	// Check if mime_type column exists and handle it
	if db.Migrator().HasColumn(&models.Gallery{}, "mime_type") {
		// Column exists - check if it's nullable and has NULL values
		var isNullable string
		var nullCount int64
		db.Raw(`
			SELECT is_nullable 
			FROM information_schema.columns 
			WHERE table_name = 'galleries' AND column_name = 'mime_type'
		`).Scan(&isNullable)
		
		db.Model(&models.Gallery{}).Where("mime_type IS NULL").Count(&nullCount)
		
		// If there are NULL values, backfill them
		if nullCount > 0 {
			log.Printf("Found %d galleries with NULL mime_type, backfilling...", nullCount)
			
			// Backfill mime_type based on filename extension
			result := db.Exec(`
				UPDATE galleries
				SET mime_type = CASE
					WHEN LOWER(filename) LIKE '%.jpg' OR LOWER(filename) LIKE '%.jpeg' THEN 'image/jpeg'
					WHEN LOWER(filename) LIKE '%.png' THEN 'image/png'
					WHEN LOWER(filename) LIKE '%.gif' THEN 'image/gif'
					WHEN LOWER(filename) LIKE '%.webp' THEN 'image/webp'
					ELSE 'image/jpeg'
				END
				WHERE mime_type IS NULL
			`)
			if result.Error != nil {
				log.Printf("Warning: Error backfilling mime_type from filename: %v", result.Error)
			}
			
			// Also check path field for data URIs
			result = db.Exec(`
				UPDATE galleries
				SET mime_type = CASE
					WHEN path LIKE 'data:image/jpeg%' THEN 'image/jpeg'
					WHEN path LIKE 'data:image/png%' THEN 'image/png'
					WHEN path LIKE 'data:image/gif%' THEN 'image/gif'
					WHEN path LIKE 'data:image/webp%' THEN 'image/webp'
					ELSE mime_type
				END
				WHERE mime_type IS NULL OR (mime_type = 'image/jpeg' AND path LIKE 'data:image/%')
			`)
			if result.Error != nil {
				log.Printf("Warning: Error backfilling mime_type from path: %v", result.Error)
			}
			
			// Set default for any remaining NULL values
			result = db.Exec(`UPDATE galleries SET mime_type = 'image/jpeg' WHERE mime_type IS NULL`)
			if result.Error != nil {
				log.Printf("Warning: Error setting default mime_type: %v", result.Error)
			}
		}
		
		// Make column NOT NULL if it's currently nullable
		if isNullable == "YES" {
			result := db.Exec(`ALTER TABLE galleries ALTER COLUMN mime_type SET NOT NULL`)
			if result.Error != nil {
				log.Printf("Warning: Error setting mime_type to NOT NULL: %v", result.Error)
			} else {
				log.Println("Successfully set mime_type column to NOT NULL")
			}
		}
	}

	// Auto migrate the schema (mime_type will be added as nullable if it doesn't exist)
	if err := db.AutoMigrate(&models.User{}, &models.Contact{}, &models.ContactRequest{}, &models.Gallery{}, &models.EventCalendar{}); err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	// After AutoMigrate, if mime_type was just added, backfill it and make it NOT NULL
	if db.Migrator().HasColumn(&models.Gallery{}, "mime_type") {
		var isNullable string
		db.Raw(`
			SELECT is_nullable 
			FROM information_schema.columns 
			WHERE table_name = 'galleries' AND column_name = 'mime_type'
		`).Scan(&isNullable)
		
		// Backfill any NULL values that might have been created
		db.Exec(`
			UPDATE galleries
			SET mime_type = CASE
				WHEN LOWER(filename) LIKE '%.jpg' OR LOWER(filename) LIKE '%.jpeg' THEN 'image/jpeg'
				WHEN LOWER(filename) LIKE '%.png' THEN 'image/png'
				WHEN LOWER(filename) LIKE '%.gif' THEN 'image/gif'
				WHEN LOWER(filename) LIKE '%.webp' THEN 'image/webp'
				ELSE 'image/jpeg'
			END
			WHERE mime_type IS NULL
		`)
		
		// Make column NOT NULL if it's currently nullable
		if isNullable == "YES" {
			db.Exec(`ALTER TABLE galleries ALTER COLUMN mime_type SET NOT NULL`)
		}
	}

	// Initialize repositories
	userRepo := infraRepos.NewUserRepository(db)
	contactRepo := infraRepos.NewContactRepository(db)
	contactRequestRepo := infraRepos.NewContactRequestRepository(db)
	galleryRepo := infraRepos.NewGalleryRepository(db)
	eventCalendarRepo := infraRepos.NewEventCalendarRepository(db)

	// Initialize services
	userService := infraServices.NewUserService(userRepo)
	contactService := infraServices.NewContactService(contactRepo)
	contactRequestService := infraServices.NewContactRequestService(contactRequestRepo, infraServices.NewEmailService())
	galleryService := infraServices.NewGalleryService(galleryRepo)
	eventCalendarService := infraServices.NewEventCalendarService(eventCalendarRepo)

	// Setup Gin router
	router := gin.Default()
	
	// Configure multipart form to not save files to disk
	router.MaxMultipartMemory = 10 << 20 // 10 MB

	// Add CORS middleware
	router.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	})

	// Setup routes
	routes.SetupRoutes(router, userService, contactService, contactRequestService, galleryService, eventCalendarService)

	// Get port from environment or use default
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server starting on port %s", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
