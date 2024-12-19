package database

import (
	"fmt"
	"time"

	"github.com/TehilaTheStudent/todo-list-fullstack/backend/config"
	"github.com/TehilaTheStudent/todo-list-fullstack/backend/internal/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Connect() (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.GetEnv("DB_USER"),
		config.GetEnv("DB_PASS"),
		config.GetEnv("DB_HOST"),
		config.GetEnv("DB_PORT"),
		config.GetEnv("DB_NAME"),
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info), // Verbose logging for development
	})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}
	// Configure connection pooling
	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to get generic DB object: %w", err)
	}

	// Match the database's maximum connection limit
	sqlDB.SetMaxOpenConns(5)                   // Maximum number of open connections
	sqlDB.SetMaxIdleConns(5)                   // Maximum number of idle connections
	sqlDB.SetConnMaxLifetime(15 * time.Minute) // Recycle connections every 15 minutes

	// Auto-migrate models
	db.AutoMigrate(&models.Task{})
	db.AutoMigrate(&models.User{})

	return db, nil
}
