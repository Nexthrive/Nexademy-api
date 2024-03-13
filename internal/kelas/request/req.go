package request

import "time"

type CreateKelasRequest struct {
	ID string `json:"id"`
	Walas string `json:"walas"`
	CreatedAt time.Time `json:"created_at"`
}