package models

import "time"

type Post struct {
	ID        uint
	CreatedAt time.Time
	UpdatedAt time.Time
	PostedAt  time.Time `gorm:"default:CURRENT_TIMESTAMP()"`
	UserID    uint
	Title     string
	Content   string
}
