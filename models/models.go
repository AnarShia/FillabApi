package models

import (
	"gorm.io/gorm"
)

type Fact struct {
	gorm.Model
	Content string `json:"content" gorm:"not null"`
}
