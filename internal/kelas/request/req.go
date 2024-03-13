package request

import "time"

type CreateKelasRequest struct {
	ID_Kelas string `json:"id_kelas"`
	Walas string `json:"walas"`
	CreatedAt time.Time `json:"created_at"`
}