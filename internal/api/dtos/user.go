package dtos

type UserSession struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"` // #nosec
	IsAdmin  bool   `json:"isAdmin"`
}
