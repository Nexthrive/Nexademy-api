package entity

import "time"

type Mapel struct {
	Id_mapel  string    `json:"id_mapel"`
	Judul     string    `json:"judul"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
