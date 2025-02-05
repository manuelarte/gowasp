package errors

import (
	"context"
	"encoding/json"
	"github.com/ing-bank/ginerr/v3"
	"github.com/mattn/go-sqlite3"
	"net/http"
)

type ErrorResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

var _ error = new(PasswordNotValid)

type PasswordNotValid struct {
	Message string
}

func (p PasswordNotValid) Error() string {
	return p.Message
}

func jsonSyntaxErrorHandler(_ context.Context, err *json.SyntaxError) (int, any) {
	return http.StatusBadRequest, ErrorResponse{
		Status:  http.StatusBadRequest,
		Message: err.Error(),
	}
}

func sqlite3ErrorHandler(_ context.Context, err sqlite3.Error) (int, any) {
	return http.StatusBadRequest, ErrorResponse{
		Status:  http.StatusBadRequest,
		Message: err.Error(),
	}
}

func RegisterErrorResponseHandlers() {
	ginerr.RegisterErrorHandler(&json.SyntaxError{}, jsonSyntaxErrorHandler)
	ginerr.RegisterErrorHandler(sqlite3.Error{}, sqlite3ErrorHandler)
	ginerr.RegisterErrorHandler(PasswordNotValid{}, func(ctx context.Context, err PasswordNotValid) (int, any) {
		return http.StatusBadRequest, ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		}
	})
}
