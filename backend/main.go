package main

import (
	"github.com/TehilaTheStudent/todo-list-fullstack/backend/config"
	_ "github.com/TehilaTheStudent/todo-list-fullstack/backend/docs" // Swag docs
	"github.com/TehilaTheStudent/todo-list-fullstack/backend/internal/database"
	"github.com/TehilaTheStudent/todo-list-fullstack/backend/internal/handlers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
)

// @title Todo List API
// @version 1.0
// @description This is a simple todo list API.
func main() {
	// Load environment variables
	config.LoadEnv()

	// Connect to the database
	database.Connect()

	// Initialize Gin router
	r := gin.Default()

	// Enable CORS
	r.Use(cors.Default())

	// Setup Swagger documentation
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Setup application routes
	SetupRoutes(r)

	// Run the server on port 8080
	if err := r.Run(":8080"); err != nil {
		panic("Failed to start server: " + err.Error())
	}
}

// SetupRoutes sets up the API routes
func SetupRoutes(r *gin.Engine) {
	// @Summary Get all tasks
	// @Description Get a list of all tasks
	// @Tags tasks
	// @Produce json
	// @Success 200 {array} models.Task "List of tasks"
	// @Router /tasks [get]
	r.GET("/tasks", handlers.GetTasks)

	// @Summary Create a new task
	// @Description Create a new task with the input payload
	// @Tags tasks
	// @Accept json
	// @Produce json
	// @Param task body models.Task true "Task to create"
	// @Success 201 {object} models.Task "Created task"
	// @Router /tasks [post]
	r.POST("/tasks", handlers.CreateTask)

	// @Summary Update an existing task
	// @Description Update an existing task by ID
	// @Tags tasks
	// @Accept json
	// @Produce json
	// @Param id path int true "Task ID"
	// @Param task body models.Task true "Updated task"
	// @Success 200 {object} models.Task "Updated task"
	// @Router /tasks/{id} [put]
	r.PUT("/tasks/:id", handlers.UpdateTask)

	// @Summary Delete a task
	// @Description Delete a task by ID
	// @Tags tasks
	// @Produce json
	// @Param id path int true "Task ID"
	// @Success 204 {string} string "No Content"
	// @Router /tasks/{id} [delete]
	r.DELETE("/tasks/:id", handlers.DeleteTask)
}
