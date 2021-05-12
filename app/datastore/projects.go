package datastore

import (
	"time"

	"gorm.io/gorm"
)

type Project struct {
	gorm.Model
	ID        uint      `json:"id"`
	Title     string    `json:"title"`
	StartDate time.Time `json:"startDate"`
	EndtDate  time.Time `json:"endDate"`
	Tasks     []*Task   `gorm:"foreignKey:ProjectID"`
	ClientID  uint      `json:"clientId"`
	Client    Client    `json:"client"`
}
