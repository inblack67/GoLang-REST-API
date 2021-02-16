package types

// SSession ...
type SSession struct{
	Username string `json:"username"`
	User uint `json:"user"`
}

// Status ...
type Status struct{
	Success bool `json:"success"`
	Message string `json:"message"`
}