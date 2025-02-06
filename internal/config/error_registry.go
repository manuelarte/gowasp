package config

import (
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/ing-bank/ginerr/v3"
	"github.com/mattn/go-sqlite3"
	"gowasp/internal/models/errors"
	"net/http"
)

func validatorErrorHandler(_ context.Context, err *validator.ValidationErrors) (int, any) {
	return http.StatusBadRequest, errors.ErrorResponse{
		HtmlTemplate: "errors/500.tpl",
		Data:         gin.H{},
	}
}

func jsonSyntaxErrorHandler(_ context.Context, err *json.SyntaxError) (int, any) {
	return http.StatusBadRequest, errors.ErrorResponse{
		HtmlTemplate: "errors/400.tpl",
		Data:         gin.H{"message": err.Error()},
	}
}

func sqlite3ErrorHandler(_ context.Context, err sqlite3.Error) (int, any) {
	return http.StatusBadRequest, errors.ErrorResponse{
		HtmlTemplate: "errors/400.tpl",
		Data:         gin.H{"message": err.Error()},
	}
}

func RegisterErrorResponseHandlers() {
	ginerr.RegisterErrorHandler(&validator.ValidationErrors{}, validatorErrorHandler)
	ginerr.RegisterErrorHandler(&json.SyntaxError{}, jsonSyntaxErrorHandler)
	ginerr.RegisterErrorHandler(sqlite3.Error{}, sqlite3ErrorHandler)
	ginerr.RegisterErrorHandler(errors.PasswordNotValid{}, func(ctx context.Context, err errors.PasswordNotValid) (int, any) {
		return http.StatusBadRequest, errors.ErrorResponse{
			HtmlTemplate: "errors/400.tpl",
			Data:         gin.H{"message": err.Error()},
		}
	})

	ginerr.DefaultErrorRegistry.RegisterDefaultHandler(func(ctx context.Context, err error) (int, any) {
		return http.StatusBadRequest, errors.ErrorResponse{
			HtmlTemplate: "errors/500.tpl",
			Data:         gin.H{},
		}
	})
}
