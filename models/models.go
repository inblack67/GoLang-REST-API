package models

import "time"

// Story ...
type Story struct{
	// gorm.Model
	ID  uint `gorm:"primaryKey" json:"id,omitempty"`
	Title string `json:"title,omitempty" gorm:"unique;not null;default:null"`
	Author string `json:"author,omitempty" gorm:"not null;default:null"`
	CreatedAt time.Time `json:"createdAt,omitempty" gorm:"not null;default:null"`
	UpdatedAt time.Time `json:"updatedAt,omitempty" gorm:"not null;default:null"`
}

// User ...
type User struct{
	// gorm.Model
	ID  uint `gorm:"primaryKey" json:"id,omitempty"`
	Name string `json:"name,omitempty" gorm:"not null;default:null"`
	Username string `json:"username,omitempty" gorm:"unique;not null;default:null"`
	Email string `json:"email,omitempty" gorm:"unique;not null;default:null"`
	Password string `json:"password,omitempty" gorm:"not null;default:null"`
	CreatedAt time.Time `json:"createdAt,omitempty" gorm:"not null;default:null"`
	UpdatedAt time.Time `json:"updatedAt,omitempty" gorm:"not null;default:null"`
}