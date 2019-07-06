package models

import "time"

// User represents the auth for the api
type User struct {
	ID        int       `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	LastName  string    `json:"lastName,omitempty"`
	Email     string    `json:"email,omitempty"`
	Password  string    `json:"password,omitempty"`
	AvatarURL string    `json:"avatarurl,omitempty"`
	UUID      string    `json:"uuiduser,omitempty"`
	Cpf       string    `json:"cpf,omitempty"`
	DataStart time.Time `json:"datastart,omitempty"`
}
