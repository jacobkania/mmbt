package models

import "time"

// UserLoginToken model
type UserLoginToken struct {
	ID        int64     `json:"-"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
	Token     string    `json:"token"`
	UserID    int64     `json:"userID"`
}
