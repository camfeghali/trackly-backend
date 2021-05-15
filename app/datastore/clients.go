package datastore

import "gorm.io/gorm"

type Client struct {
	ID       uint       `json:"id"`
	Name     string     `json:"name"`
	Email    string     `json:"email"`
	UserID   uint       `json:"userId"`
	User     User       `json:"user"`
	Projects []*Project `gorm:"foreignKey:ClientID"`
	gorm.Model
}
