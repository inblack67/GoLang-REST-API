package models

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

// Story ...
type Story struct{
	ID  uint `gorm:"primaryKey" json:"id,omitempty"`
	UUID  uint `gorm:"unique" json:"uuid,omitempty"`
	Title string `json:"title,omitempty" gorm:"unique;not null;default:null"`
	Author string `json:"author,omitempty" gorm:"not null;default:null"`
	CreatedAt time.Time `json:"createdAt,omitempty" gorm:"not null;default:null"`
	UpdatedAt time.Time `json:"updatedAt,omitempty" gorm:"not null;default:null"`
}

// ValidateStory ...
func (newStory Story) ValidateStory() error{
	return validation.ValidateStruct(&newStory, 

		validation.Field(&newStory.Title, validation.Required, validation.Length(5,20)),

		validation.Field(&newStory.Author, validation.Required),
	)
}

// User ...
type User struct{
	ID  uint `gorm:"primaryKey" json:"id,omitempty"`
	UUID  uint `gorm:"unique" json:"uuid,omitempty"`
	Name string `json:"name,omitempty" gorm:"not null;default:null"`
	Username string `json:"username,omitempty" gorm:"unique;not null;default:null"`
	Email string `json:"email,omitempty" gorm:"unique;not null;default:null"`
	Password string `json:"password,omitempty" gorm:"not null;default:null"`
	CreatedAt time.Time `json:"createdAt,omitempty" gorm:"not null;default:null"`
	UpdatedAt time.Time `json:"updatedAt,omitempty" gorm:"not null;default:null"`
}


// ValidateMe ...
func (newUser User) ValidateMe() error{
	return validation.ValidateStruct(&newUser, 

		validation.Field(&newUser.Name, validation.Required, validation.Length(5,20)),

		validation.Field(&newUser.Username, validation.Required, validation.Length(5,20)),

		validation.Field(&newUser.Email, validation.Required, is.Email),

		validation.Field(&newUser.Password, validation.Required, validation.Length(8,20)),
	)
}