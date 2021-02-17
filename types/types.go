package types

import "fibreApi/models"

// SSession ...
type SSession struct{
	User models.User
}

// Status ...
type Status struct{
	Success bool `json:"success"`
	Message string `json:"message"`
}