package models

import "time"

type Blog struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"createdAt" gorm:"-"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"-"`
	PostedAt  time.Time `json:"postedAt" gorm:"-"`
	UserID    uint      `json:"userId"`
	User      User      `json:"user"`
	Title     string    `json:"title"`
	Content   string    `json:"content" binding:"required"`
}
