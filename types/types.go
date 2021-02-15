package types

// SSession ...
type SSession struct{
	Username string `json:"username"`
}

// Status ...
type Status struct{
	Success bool `json:"success"`
	Message string `json:"message"`
}