package models

import (
	"time"

	"gorm.io/gorm"
)

// Story ...
type Story struct{
	gorm.Model
	Title string `json:"title"`
	Author string `json:"author"`
	CreatedAt time.Time `json:"createdAt"`
}