package config

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/manuelarte/gowasp/internal/models/errors"
	"github.com/manuelarte/gowasp/internal/repositories"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/ing-bank/ginerr/v3"
	"github.com/mattn/go-sqlite3"
)

func validatorErrorHandler(_ context.Context, err *validator.ValidationErrors) (int, any) {
	return http.StatusBadRequest, errors.ErrorResponse{
		Data: gin.H{
			"message": err.Error(),
		},
	}
}

func jsonSyntaxErrorHandler(_ context.Context, err *json.SyntaxError) (int, any) {
	return http.StatusBadRequest, errors.ErrorResponse{
		Data: gin.H{"message": err.Error()},
	}
}

func sqlite3ErrorHandler(_ context.Context, err sqlite3.Error) (int, any) {
	return http.StatusBadRequest, errors.ErrorResponse{
		Data: gin.H{"message": err.Error()},
	}
}

func RegisterErrorResponseHandlers() {
	ginerr.RegisterErrorHandler(&validator.ValidationErrors{}, validatorErrorHandler)
	ginerr.RegisterErrorHandler(&json.SyntaxError{}, jsonSyntaxErrorHandler)
	ginerr.RegisterErrorHandler(sqlite3.Error{}, sqlite3ErrorHandler)
	ginerr.RegisterErrorHandler(errors.PasswordNotValidError{},
		func(_ context.Context, err errors.PasswordNotValidError) (int, any) {
			return http.StatusBadRequest, errors.ErrorResponse{
				Data: gin.H{"message": err.Error()},
			}
		})
	ginerr.RegisterErrorHandler(repositories.ErrUserNotFound, func(_ context.Context, err error) (int, any) {
		return http.StatusBadRequest, errors.ErrorResponse{
			Data: gin.H{"message": err.Error()},
		}
	})

	ginerr.DefaultErrorRegistry.RegisterDefaultHandler(func(_ context.Context, err error) (int, any) {
		return http.StatusBadRequest, errors.ErrorResponse{
			Data: gin.H{
				"message": err.Error(),
			},
		}
	})
}
