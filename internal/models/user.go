package models

import "time"

type User struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Username  string    `json:"username" binding:"required" form:"username"`
	Password  string    `json:"password" binding:"required" form:"password"`
}
