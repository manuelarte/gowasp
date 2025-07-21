package models

import (
	"time"
)

type PostComment struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
	PostedAt  time.Time `gorm:"now" json:"postedAt"`
	PostID    uint      `json:"postId"`
	UserID    uint      `json:"userId"`
	User      *User     `json:"user"`
	Comment   string    `json:"comment"`
}
