package structs

// SLogin ...
type SLogin struct{
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}