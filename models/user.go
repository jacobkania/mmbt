package models

import "time"

// User model
type User struct {
	ID              int64             `json:"id"`
	CreatedAt       time.Time         `json:"createdAt"`
	UpdatedAt       time.Time         `json:"updatedAt"`
	DeletedAt       time.Time         `json:"deletedAt"`
	FullName        string            `json:"fullName"`
	Username        string            `json:"Username"`
	Passw           string            `json:"-"`
	Emails          []*UserEmail      `json:"emails" pg:"rel:has-many"`
	UserLoginTokens []*UserLoginToken `json:"-" pg:"rel:has-many"`
}
