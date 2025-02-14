package models

import "time"

type Blog struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
	PostedAt  time.Time `json:"postedAt"`
	UserID    uint      `json:"userId"`
	Title     string    `json:"title"`
	Content   string    `json:"content" binding:"required"`
}
