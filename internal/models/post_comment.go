package models

import (
	"time"
)

type PostComment struct {
	ID        uint
	CreatedAt time.Time
	UpdatedAt time.Time
	PostedAt  time.Time `gorm:"default:CURRENT_TIMESTAMP()"`
	PostID    uint
	UserID    uint
	User      *User
	Comment   string
}
