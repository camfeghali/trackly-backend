package datastore

import (
	"time"

	"gorm.io/gorm"
)

type TimeEntry struct {
	gorm.Model
	ID          uint      `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	StartDate   time.Time `json:"startDate"`
	EndtDate    time.Time `json:"endDate"`
	TaskID      uint      `json:"taskId"`
	Task        Task      `json:"task"`
}
