package routes

import (
	"os"

	"ccpp-backend/internal/domain/services"
	"ccpp-backend/internal/infrastructure/handlers"
	"ccpp-backend/internal/infrastructure/middleware"

	"github.com/gin-gonic/gin"
)

// registerAPIRoutes mounts all API endpoints under the provided group (e.g., /api or /api/v1)
func registerAPIRoutes(
	api *gin.RouterGroup,
	authHandler *handlers.AuthHandler,
	userHandler *handlers.UserHandler,
	contactHandler *handlers.ContactHandler,
	contactRequestHandler *handlers.ContactRequestHandler,
	galleryHandler *handlers.GalleryHandler,
	eventHandler *handlers.EventHandler,
) {
	// Auth routes
	auth := api.Group("/auth")
	{
		auth.POST("/login", authHandler.Login)
		auth.GET("/logout", authHandler.Logout)
		auth.GET("/profile", middleware.Auth0Middleware(), authHandler.Profile)
	}

	// Contact requests (public create, protected list/detail/delete)
	contactRequests := api.Group("/contact-requests")
	{
		contactRequests.POST("", contactRequestHandler.CreateContactRequest)

		contactRequests.GET("", contactRequestHandler.GetContactRequests)
		contactRequests.GET("/:id", contactRequestHandler.GetContactRequest)
		contactRequests.DELETE("/:id", contactRequestHandler.DeleteContactRequest)
	}

	// Protected routes
	protected := api.Group("")
	protected.Use(middleware.Auth0Middleware())

	users := protected.Group("/users")
	{
		users.POST("", userHandler.CreateUser)
		users.GET("", userHandler.GetAllUsers)
		users.GET("/:id", userHandler.GetUser)
		users.PUT("/:id", userHandler.UpdateUser)
		users.DELETE("/:id", userHandler.DeleteUser)
	}

	contacts := protected.Group("/contacts")
	{
		contacts.POST("", contactHandler.CreateContact)
		contacts.GET("", contactHandler.GetAllContacts)
		contacts.GET("/:id", contactHandler.GetContactByID)
		contacts.DELETE("/:id", contactHandler.DeleteContact)
	}

	gallery := api.Group("/gallery")
	{
		gallery.POST("", galleryHandler.CreateGallery)
		gallery.POST("/upload", galleryHandler.UploadGallery)
		gallery.GET("", galleryHandler.GetAllGalleries)
		gallery.GET("/:id", galleryHandler.GetGalleryByID)
		gallery.DELETE("/:id", galleryHandler.DeleteGallery)
	}

	events := api.Group("/events")
	{
		// Auth removed to allow admin calendar without token
		events.POST("", eventHandler.CreateEvent)
		events.GET("", eventHandler.GetAllEvents)
		events.GET("/:id", eventHandler.GetEventByID)
		events.PUT("/:id", eventHandler.UpdateEvent)
		events.DELETE("/:id", eventHandler.DeleteEvent)
	}
}

func SetupRoutes(
	router *gin.Engine,
	userService services.UserService,
	contactService services.ContactService,
	contactRequestService services.ContactRequestService,
	galleryService services.GalleryService,
	eventService services.EventService,
) {
	userHandler := handlers.NewUserHandler(userService)
	contactHandler := handlers.NewContactHandler(contactService)
	contactRequestHandler := handlers.NewContactRequestHandler(contactRequestService)
	galleryHandler := handlers.NewGalleryHandler(galleryService)
	eventHandler := handlers.NewEventHandler(eventService)
	authHandler := handlers.NewAuthHandler()

	// Versioned API (existing)
	apiV1 := router.Group("/api/v1")
	registerAPIRoutes(apiV1, authHandler, userHandler, contactHandler, contactRequestHandler, galleryHandler, eventHandler)

	// Unversioned API (go-vue-base compatibility)
	api := router.Group("/api")
	registerAPIRoutes(api, authHandler, userHandler, contactHandler, contactRequestHandler, galleryHandler, eventHandler)

	// Health check endpoint
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  "ok",
			"message": "API is running",
		})
	})

	// Serve static files for uploads
	uploadPath := os.Getenv("UPLOAD_PATH")
	if uploadPath == "" {
		uploadPath = "./uploads"
	}
	router.Static("/uploads", uploadPath)
}
