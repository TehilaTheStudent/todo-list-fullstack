package main

import (
	"github.com/TehilaTheStudent/todo-list-fullstack/backend/config"
	"github.com/TehilaTheStudent/todo-list-fullstack/backend/internal/database"
	"github.com/TehilaTheStudent/todo-list-fullstack/backend/internal/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
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

	// Setup application routes
	routes.SetupRoutes(r)

	// Run the server on port 8080
	if err := r.Run(":8080"); err != nil {
		panic("Failed to start server: " + err.Error())
	}
}