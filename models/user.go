package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `json:"name"`
	Password string `json:"password"`
	Task []Task     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
