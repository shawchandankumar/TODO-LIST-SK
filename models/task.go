package models

import (
	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	Title    string `json:"title"`
	Todo     string `json:"todo"`
	Priority int    `json:"priority"`
	UserId uint     `json:"userId"`   // foreign key
}
