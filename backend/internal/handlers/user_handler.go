package handlers

import (
	"net/http"

	"github.com/TehilaTheStudent/todo-list-fullstack/backend/internal/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// UserHandler handles user-related requests
type UserHandler struct {
	DB *gorm.DB
}

// NewUserHandler creates a new UserHandler
func NewUserHandler(db *gorm.DB) *UserHandler {
	return &UserHandler{DB: db}
}

// GetAllUsers handles GET requests to fetch all users
func (h *UserHandler) GetAllUsers(c *gin.Context) {
	var users []models.User
	if err := h.DB.Find(&users).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, users)
}

