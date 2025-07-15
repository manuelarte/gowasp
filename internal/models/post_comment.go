package models

import (
	"time"
)

type PostComment struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
	PostedAt  time.Time `gorm:"now" json:"postedAt"`
	PostID    uint      `json:"postID"`
	UserID    uint      `json:"userID"`
	User      *User     `json:"user"`
	Comment   string    `json:"comment"`
}
