package models

import (
	"time"
)

type Task struct {
	ID          uint      `json:"id" gorm:"primary_key"`
	Title       string    `json:"title" gorm:"not null"`
	Description string    `json:"description"`
	Status      string    `json:"status" gorm:"default:pending"`
	CreatedAt   time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

type AddTaskModel struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description"`
}

type UpdateTaskModel struct {
	ID          uint   `json:"id" binding:"required"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
}
