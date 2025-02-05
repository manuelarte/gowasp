package config

import (
	"context"
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"github.com/ing-bank/ginerr/v3"
	"github.com/mattn/go-sqlite3"
	"gowasp/internal/models/errors"
	"net/http"
)

func validatorErrorHandler(_ context.Context, err *validator.ValidationErrors) (int, any) {
	return http.StatusBadRequest, errors.ErrorResponse{
		Status:  http.StatusBadRequest,
		Message: err.Error(),
	}
}

func jsonSyntaxErrorHandler(_ context.Context, err *json.SyntaxError) (int, any) {
	return http.StatusBadRequest, errors.ErrorResponse{
		Status:  http.StatusBadRequest,
		Message: err.Error(),
	}
}

func sqlite3ErrorHandler(_ context.Context, err sqlite3.Error) (int, any) {
	return http.StatusBadRequest, errors.ErrorResponse{
		Status:  http.StatusBadRequest,
		Message: err.Error(),
	}
}

func RegisterErrorResponseHandlers() {
	ginerr.RegisterErrorHandler(&validator.ValidationErrors{}, validatorErrorHandler)
	ginerr.RegisterErrorHandler(&json.SyntaxError{}, jsonSyntaxErrorHandler)
	ginerr.RegisterErrorHandler(sqlite3.Error{}, sqlite3ErrorHandler)
	ginerr.RegisterErrorHandler(errors.PasswordNotValid{}, func(ctx context.Context, err errors.PasswordNotValid) (int, any) {
		return http.StatusBadRequest, errors.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		}
	})
}
