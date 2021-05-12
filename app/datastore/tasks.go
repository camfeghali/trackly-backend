package datastore

import (
	"time"

	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	ID          uint        `json:"id"`
	Title       string      `json:"title"`
	StartDate   time.Time   `json:"startDate"`
	EndtDate    time.Time   `json:"endDate"`
	TimeEntries []TimeEntry `gorm:"foreignKey:TaskID"`
	ProjectID   uint        `json:"projectId"`
	Project     Project     `json:"project"`
}
