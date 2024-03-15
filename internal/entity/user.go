package entity

// User represents a user.
type User struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Nis        int    `json:"nis"`
	Passphrase string `json:"passphrase"`
	Email      string `json:"email"`
	No_telp    string `json:"no_telp"`
	Gender     string `json:"gender"`
	Religion   string `json:"religion"`
}

// GetID returns the user ID.
func (u User) GetID() string {
	return u.ID
}

// GetName returns the user name.
func (u User) GetName() string {
	return u.Name
}
