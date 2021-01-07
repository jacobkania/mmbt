package models

import "time"

// UserEmail model
type UserEmail struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
	DeletedAt time.Time `json:"-"`
	Email     string    `json:"email"`
	User      *User     `json:"user" pg:"rel:has-one"`
}
