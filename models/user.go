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
