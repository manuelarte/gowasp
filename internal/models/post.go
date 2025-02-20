package models

import "time"

type Post struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
	PostedAt  time.Time `json:"postedAt" gorm:"default:CURRENT_TIMESTAMP()"`
	UserID    uint      `json:"userId"`
	Title     string    `json:"title"`
	Content   string    `json:"content" binding:"required"`
}
