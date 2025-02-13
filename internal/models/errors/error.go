package errors

import "github.com/gin-gonic/gin"

type ErrorResponse struct {
	Data gin.H `json:"data"`
}

var _ error = new(PasswordNotValid)

type PasswordNotValid struct {
	Message string
}

func (p PasswordNotValid) Error() string {
	return p.Message
}
