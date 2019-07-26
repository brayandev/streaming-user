package user

import "time"

// User structure to register user.
type User struct {
	ID       int64     `json:"id"`
	Name     string    `json:"name,omitempty"`
	Email    string    `json:"email,omitempty"`
	Creation time.Time `json:"creation"`
}

// Version user version.
func (User) Version() string {
	return "strm-user.user.v1"
}
