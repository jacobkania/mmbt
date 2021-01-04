package models

import "time"

// User model
type User struct {
	ID           int64
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    time.Time
	FullName     string
	PrimaryEmail string
	Passw        string
}

// SetAllowedParams only allows publicly accessible params to be set
func (u *User) SetAllowedParams(req &interface{}) {
	u.FullName = req.FullName
	u.PrimaryEmail = req.PrimaryEmail
	u.Passw = req.Passw
}