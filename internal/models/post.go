package models

import "time"

type Post struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
	PostedAt  time.Time `gorm:"default:CURRENT_TIMESTAMP()" json:"postedAt"`
	UserID    uint      `json:"userId"`
	Title     string    `json:"title"`
	Content   string    `binding:"required" json:"content"`
}
