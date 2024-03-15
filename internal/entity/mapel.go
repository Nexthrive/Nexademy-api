package entity

import "time"

type Mapel struct {
	ID        string    `json:"id"`
	Judul     string    `json:"judul"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
