package request

import "time"

type CreateMapelRequest struct {
	Judul     string    `json:"judul"`
	CreatedAt time.Time `json:"created_at"`
}

type UpdateMapelRequest struct {
	Judul     string    `json:"judul"`
	UpdatedAt time.Time `json:"updated_at"`
}