package models

import "time"

// User model
type User struct {
	ID           int64     `json:"id"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
	DeletedAt    time.Time `json:"deletedAt"`
	FullName     string    `json:"fullName"`
	PrimaryEmail string    `json:"primaryEmail"`
	Passw        string    `json:"-"`
}
