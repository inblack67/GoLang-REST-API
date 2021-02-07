package models

import (
	"gorm.io/gorm"
)

// Story ...
type Story struct{
	gorm.Model
	Title string `json:"title"`
	Author string `json:"author"`
}