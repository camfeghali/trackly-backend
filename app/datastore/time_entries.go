package datastore

import (
	"time"

	"gorm.io/gorm"
)

type TimeEntry struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	Title       string         `json:"title"`
	Description string         `json:"description"`
	StartDate   time.Time      `json:"startDate"`
	EndtDate    time.Time      `json:"endDate"`
	TaskID      uint           `json:"taskId"`
	Task        Task           `json:"task"`
	CreatedAt   time.Time      `json:"createdAt"`
	UpdatedAt   time.Time      `json:"updatedAt"`
	DeletedAt   gorm.DeletedAt `json:"deletedAt" gorm:"index"`
}
