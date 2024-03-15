package request

type Login struct {
	Nis        int    `json:"nis"`
	Passphrase string `json:"passphrase"`
}
type CreateUser struct {
	ID         string `json:"id"`
	Nis        int    `json:"nis"`
	Name       string `json:"name"`
	Passphrase string `json:"passphrase"`
	Email      string `json:"email"`
	No_telp    string `json:"no_telp"`
	Gender     string `json:"gender"`
	Religion   string `json:"religion"`
}
type UpdateUser struct {
	Name       string `json:"name"`
	No_telp    string `json:"no_telp"`
	Religion   string `json:"religion"`
	Passphrase string `json:"passphrase"`
}
