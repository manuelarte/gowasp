package models

import "time"

type User struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Username  string    `binding:"required" form:"username" json:"username"`
	Password  string    `binding:"required" form:"password" json:"password"`
	IsAdmin   bool      `json:"isAdmin"`
}
