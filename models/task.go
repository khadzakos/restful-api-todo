package models

import "time"

type Task struct {
	ID          uint      `json:"id" gorm:"primary_key"`
	Title       string    `json:"title" gorm:"not null"`
	Description string    `json:"description"`
	Status      string    `json:"status" gorm:"default:pending"`
	CreatedAt   time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

func NewTask(title, description string) *Task {
	return &Task{
		Title:       title,
		Description: description,
		Status:      "pending", // default status
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
}

func (t *Task) Update(title, description, status string) {
	t.Title = title
	t.Description = description
	t.Status = status
	t.UpdatedAt = time.Now()
}

func (t *Task) Complete() {
	t.Status = "completed"
	t.UpdatedAt = time.Now()
}
