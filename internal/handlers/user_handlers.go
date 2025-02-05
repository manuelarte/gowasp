package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/ing-bank/ginerr/v3"
	"gowasp/internal/models"
	"gowasp/internal/services"
)

type UserHandler struct {
	UserService services.UserService
}

func (h *UserHandler) Create(c *gin.Context) {
	user := models.User{}
	if err := c.ShouldBindBodyWithJSON(&user); err != nil {
		code, response := ginerr.NewErrorResponse(c, err)
		c.JSON(code, response)
		return
	}
	if err := h.UserService.CreateUser(c, &user); err != nil {
		code, response := ginerr.NewErrorResponse(c, err)
		c.JSON(code, response)
		return
	}
	c.JSON(201, user)
	return
}
