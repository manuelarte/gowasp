package errors

import "github.com/gin-gonic/gin"

type ErrorResponse struct {
	Data gin.H `json:"data"`
}

var _ error = new(PasswordNotValidError)

type PasswordNotValidError struct {
	Message string
}

func (p PasswordNotValidError) Error() string {
	return p.Message
}
