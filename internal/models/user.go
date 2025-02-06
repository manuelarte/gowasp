package models

import "time"

type User struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Email     string    `json:"email" binding:"required" form:"email"`
	Password  string    `json:"password" binding:"required" form:"password"`
}
