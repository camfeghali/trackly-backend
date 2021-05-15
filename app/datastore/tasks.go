package datastore

import (
	"time"

	"gorm.io/gorm"
)

type Task struct {
	Id          uint           `json:"id" gorm:"primaryKey"`
	Title       string         `json:"title"`
	StartDate   time.Time      `json:"startDate"`
	EndtDate    time.Time      `json:"endDate"`
	TimeEntries []TimeEntry    `gorm:"foreignKey:TaskID"`
	ProjectID   uint           `json:"projectId"`
	Project     Project        `json:"project"`
	CreatedAt   time.Time      `json:"createdAt"`
	UpdatedAt   time.Time      `json:"updatedAt"`
	DeletedAt   gorm.DeletedAt `json:"deletedAt" gorm:"index"`
}
