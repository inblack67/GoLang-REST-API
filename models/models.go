package models

// Story ...
type Story struct{
	// gorm.Model
	ID  uint `gorm:"primaryKey" json:"id,omitempty"`
	Title string `json:"title,omitempty" gorm:"unique;not null;default:null"`
	Author string `json:"author,omitempty" gorm:"not null;default:null"`
}