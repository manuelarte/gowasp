package rest

import (
	"encoding/json"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/manuelarte/ptrutils"
	"github.com/sirupsen/logrus"

	"github.com/manuelarte/gowasp/internal/models"
	"github.com/manuelarte/gowasp/internal/users"
)

type Users struct {
	service users.Service
}

func NewUsers(service users.Service) *Users {
	return &Users{service: service}
}

func (h *Users) UserSignup(c *gin.Context) {
	userSignup := User{}
	if err := c.BindJSON(&userSignup); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Code:    http.StatusBadRequest,
			Details: err,
			Message: "Failed to marshal user data",
		})

		return
	}
	user := userToDao(userSignup)
	if err := h.service.Create(c, &user); err != nil {
		logrus.Infof("Signup attempt failed for User %q", user.Username)
		// TODO(manuelarte): this error can be 400 or 500, depending on what's happening
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Code:    http.StatusBadRequest,
			Details: err, // TODO(manuelarte): we should check if we can send this information or not
			Message: "Failed to register user",
		})

		return
	}
	session := sessions.Default(c)
	userBytes, _ := json.Marshal(user)
	session.Set("user", userBytes)
	err := session.Save()
	if err != nil {
		logrus.Warnf("Failed to save the user's %q session: %s", user.Username, err)
		c.JSON(http.StatusInternalServerError, ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "Failed to save the user's session",
		})

		return
	}
	logrus.Infof("Signup for User %q completed", user.Username)
	c.JSON(http.StatusCreated, user)
}

func userToDao(u User) models.User {
	return models.User{
		Username: u.Username,
		Password: u.Password,
		IsAdmin:  ptrutils.DerefOr(u.IsAdmin, false),
	}
}
