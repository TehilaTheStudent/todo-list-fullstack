package handlers

import (
	"net/http"

	"github.com/TehilaTheStudent/todo-list-fullstack/backend/internal/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type TaskHandler struct {
	DB *gorm.DB
}

func NewTaskHandler(db *gorm.DB) *TaskHandler {
	return &TaskHandler{DB: db}
}

func (h *TaskHandler) GetTasks(c *gin.Context) {
	var tasks []models.Task
	h.DB.Find(&tasks)
	c.JSON(http.StatusOK, tasks)
}

func (h *TaskHandler) CreateTask(c *gin.Context) {
	var task models.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	h.DB.Create(&task)
	c.JSON(http.StatusCreated, task)
}

func (h *TaskHandler) UpdateTask(c *gin.Context) {
	var task models.Task
	if err := h.DB.First(&task, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	h.DB.Save(&task)
	c.JSON(http.StatusOK, task)
}

func (h *TaskHandler) DeleteTask(c *gin.Context) {
	var task models.Task
	if err := h.DB.First(&task, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}
	h.DB.Delete(&task)
	c.JSON(http.StatusOK, gin.H{"message": "Task deleted"})
}
