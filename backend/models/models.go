package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `json:"name"`
	Password string `json:"password"`
	Task     []Task `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

// gorm.Model definition
type Model struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type Task struct {
	gorm.Model
	Title    string `json:"title"`
	Todo     string `json:"todo"`
	Priority int    `json:"priority"`
	UserId   uint   `json:"userId"` // foreign key
}

type TaskPayload struct {
	Title    string `json:"title"`
	Todo     string `json:"todo"`
	Priority int    `json:"priority"`
	UserId   uint   `json:"userId"` // foreign key
}
