package models


type Task struct {
	ID          uint   `gorm:"primaryKey"`
	Name        string `gorm:"size:255;not null"`
	IsCompleted bool   `gorm:"default:false"`
}
