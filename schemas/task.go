package schemas

import (
	"time"

	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	Title       string
	Description string
	DueTo       string
}

type TaskResponse struct {
	ID          uint      `json:"id"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
	DeletedAt   time.Time `json:"deletedAt,omitempty"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	DueTo       string    `json:"dueTo"`
}
