package routes

import (
	"github.com/TehilaTheStudent/todo-list-fullstack/backend/internal/handlers"
	"github.com/gin-gonic/gin"
)

// SetupRoutes sets up the API routes
func SetupRoutes(r *gin.Engine) {

	r.GET("/tasks", handlers.GetTasks)

	r.POST("/tasks", handlers.CreateTask)

	r.PUT("/tasks/:id", handlers.UpdateTask)

	r.DELETE("/tasks/:id", handlers.DeleteTask)
}
