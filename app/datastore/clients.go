package datastore

import (
	"time"

	"gorm.io/gorm"
)

type Client struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	Name      string         `json:"name"`
	Email     string         `json:"email"`
	UserID    uint           `json:"userId"`
	User      User           `json:"user"`
	Projects  []*Project     `gorm:"foreignKey:ClientID"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"deletedAt" gorm:"index"`
}
