package entity

import "time"

type Kelas struct {
	ID_Kelas  string    `json:"id_kelas"`
	Walas     string    `json:"walas"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time
}
