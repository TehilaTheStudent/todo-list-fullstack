package database

import (
	"fmt"
	"log"
	"time"

	"github.com/TehilaTheStudent/todo-list-fullstack/backend/config"
	"github.com/TehilaTheStudent/todo-list-fullstack/backend/internal/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func Connect() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.GetEnv("DB_USER"),
		config.GetEnv("DB_PASS"),
		config.GetEnv("DB_HOST"),
		config.GetEnv("DB_PORT"),
		config.GetEnv("DB_NAME"),
	)
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info), // Verbose logging for development
	})
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}
	// Configure connection pooling
	sqlDB, err := DB.DB()
	if err != nil {
		log.Fatal("Failed to get generic DB object: ", err)
	}

	// Match the database's maximum connection limit
	sqlDB.SetMaxOpenConns(5)                   // Maximum number of open connections
	sqlDB.SetMaxIdleConns(5)                   // Maximum number of idle connections
	sqlDB.SetConnMaxLifetime(15 * time.Minute) // Recycle connections every 15 minutes

	// Auto-migrate models
	DB.AutoMigrate(&models.Task{})
}
