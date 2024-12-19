package routes

import (
	"github.com/TehilaTheStudent/todo-list-fullstack/backend/internal/handlers"
	"github.com/TehilaTheStudent/todo-list-fullstack/backend/internal/middlewares"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// SetupRoutes sets up the API routes
func SetupRoutes(r *gin.Engine, db *gorm.DB) {
	taskHandler := handlers.NewTaskHandler(db)
	authHandler := handlers.NewAuthHandler(db)
	userHandler := handlers.NewUserHandler(db) // Add user handler
	renderHandler := handlers.NewRenderHandler() // Add render handler

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "healthy",
		})
	})

	// Public routes
	r.POST("/register", authHandler.Register)
	r.POST("/login", authHandler.Login)
	r.GET("/users", userHandler.GetAllUsers) // Add route for getting all users
	r.GET("/render", renderHandler.GetServices) // Add route for render handler

	// Protected routes
	protected := r.Group("/")
	protected.Use(middlewares.JWTMiddleware())
	protected.GET("/tasks", taskHandler.GetTasks)
	protected.POST("/tasks", taskHandler.CreateTask)
	protected.PUT("/tasks/:id", taskHandler.UpdateTask)
	protected.DELETE("/tasks/:id", taskHandler.DeleteTask)

	// User routes
}
