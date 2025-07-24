package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/ing-bank/ginerr/v3"
	"github.com/sirupsen/logrus"

	"github.com/manuelarte/gowasp/internal/models"
	"github.com/manuelarte/gowasp/internal/posts"
	"github.com/manuelarte/gowasp/internal/users"
)

type UsersHandler struct {
	UserService users.Service
	PostService posts.Service
}

func (h *UsersHandler) Login(c *gin.Context) {
	user := models.User{}
	if err := c.BindJSON(&user); err != nil {
		code, response := ginerr.NewErrorResponse(c, err)
		c.JSON(code, response)

		return
	}
	user, err := h.UserService.Login(c, user.Username, user.Password)
	if err != nil {
		logrus.Infof("Login attempt failed for User '%s'", user.Username)
		code, response := ginerr.NewErrorResponse(c, err)
		c.JSON(code, response)

		return
	}
	session := sessions.Default(c)
	userBytes, _ := json.Marshal(user)
	session.Set("user", userBytes)
	err = session.Save()
	if err != nil {
		logrus.Infof("Login attempt failed for User '%s'", user.Username)
		code, response := ginerr.NewErrorResponse(c, err)
		c.JSON(code, response)

		return
	}
	c.JSON(http.StatusOK, user)
	logrus.Infof("User %s logged in", user.Username)
}

type UserSignup struct {
	Username string `binding:"required,max=20" json:"username" `
	Password string `binding:"required,max=20" json:"password"`
}

//nolint:unused // to be used later
func (u UserSignup) toUser() models.User {
	return models.User{
		ID:        0,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		Username:  u.Username,
		Password:  u.Password,
	}
}
