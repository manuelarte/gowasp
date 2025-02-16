package models

import "time"

type BlogComment struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
	PostedAt  time.Time `json:"postedAt"`
	BlogID    uint      `json:"blogId"`
	UserID    uint      `json:"userId"`
	Comment   string    `json:"comment"`
}
