package config

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/ing-bank/ginerr/v3"
	"github.com/mattn/go-sqlite3"

	"github.com/manuelarte/gowasp/internal/models/gerrors"
	"github.com/manuelarte/gowasp/internal/users"
)

func RegisterErrorResponseHandlers() {
	ginerr.RegisterErrorHandler(&validator.ValidationErrors{}, validatorErrorHandler)
	ginerr.RegisterErrorHandler(&json.SyntaxError{}, jsonSyntaxErrorHandler)
	ginerr.RegisterErrorHandler(sqlite3.Error{}, sqlite3ErrorHandler)
	ginerr.RegisterErrorHandler(gerrors.PasswordNotValidError{},
		func(_ context.Context, err gerrors.PasswordNotValidError) (int, any) {
			return http.StatusBadRequest, gerrors.ErrorResponse{
				Data: gin.H{"message": err.Error()},
			}
		})
	ginerr.RegisterErrorHandler(users.ErrUserNotFound, func(_ context.Context, err error) (int, any) {
		return http.StatusBadRequest, gerrors.ErrorResponse{
			Data: gin.H{"message": err.Error()},
		}
	})

	ginerr.DefaultErrorRegistry.RegisterDefaultHandler(func(_ context.Context, err error) (int, any) {
		return http.StatusBadRequest, gerrors.ErrorResponse{
			Data: gin.H{
				"message": err.Error(),
			},
		}
	})
}

func jsonSyntaxErrorHandler(_ context.Context, err *json.SyntaxError) (int, any) {
	return http.StatusBadRequest, gerrors.ErrorResponse{
		Data: gin.H{"message": err.Error()},
	}
}

func sqlite3ErrorHandler(_ context.Context, err sqlite3.Error) (int, any) {
	return http.StatusBadRequest, gerrors.ErrorResponse{
		Data: gin.H{"message": err.Error()},
	}
}

func validatorErrorHandler(_ context.Context, err *validator.ValidationErrors) (int, any) {
	return http.StatusBadRequest, gerrors.ErrorResponse{
		Data: gin.H{
			"message": err.Error(),
		},
	}
}
