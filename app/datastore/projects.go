package datastore

import (
	"time"

	"gorm.io/gorm"
)

type Project struct {
	Id        uint           `json:"id" gorm:"primaryKey"`
	Title     string         `json:"title"`
	StartDate time.Time      `json:"startDate"`
	EndtDate  time.Time      `json:"endDate"`
	Tasks     []*Task        `gorm:"foreignKey:ProjectID"`
	ClientID  uint           `json:"clientId"`
	Client    Client         `json:"client"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"deletedAt" gorm:"index"`
}
