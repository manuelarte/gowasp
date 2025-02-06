package errors

import "github.com/gin-gonic/gin"

type ErrorResponse struct {
	HtmlTemplate string
	Data         gin.H
}

var _ error = new(PasswordNotValid)

type PasswordNotValid struct {
	Message string
}

func (p PasswordNotValid) Error() string {
	return p.Message
}
