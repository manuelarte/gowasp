package handlers

import (
	"encoding/json"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/ing-bank/ginerr/v3"
	"github.com/sirupsen/logrus"
	"gowasp/internal/models"
	"gowasp/internal/services"
	"net/http"
)

type UsersHandler struct {
	UserService services.UserService
}

func (h *UsersHandler) SignupPage(c *gin.Context) {
	c.HTML(http.StatusOK, "users/signup.tpl", gin.H{})
}

func (h *UsersHandler) LoginPage(c *gin.Context) {
	c.HTML(http.StatusOK, "users/login.tpl", gin.H{})
}

func (h *UsersHandler) WelcomePage(c *gin.Context) {
	session := sessions.Default(c)
	var user models.User
	_ = json.Unmarshal(session.Get("user").([]byte), &user)

	c.HTML(http.StatusOK, "users/welcome.tpl", gin.H{"user": user})
}

func (h *UsersHandler) Signup(c *gin.Context) {
	user := models.User{}
	if err := c.BindJSON(&user); err != nil {
		code, response := ginerr.NewErrorResponse(c, err)
		c.JSON(code, response)
		return
	}
	if err := h.UserService.CreateUser(c, &user); err != nil {
		logrus.Infof("Signup attempt failed for User '%s'", user.Username)
		code, response := ginerr.NewErrorResponse(c, err)
		c.JSON(code, response)
		return
	}
	session := sessions.Default(c)
	userBytes, _ := json.Marshal(user)
	session.Set("user", userBytes)
	err := session.Save()
	if err != nil {
		code, response := ginerr.NewErrorResponse(c, err)
		c.JSON(code, response)
		return
	}
	logrus.Infof("Signup for User '%s'", user.Username)
	c.JSON(http.StatusCreated, user)
}

func (h *UsersHandler) Login(c *gin.Context) {
	user := models.User{}
	if err := c.BindJSON(&user); err != nil {
		code, response := ginerr.NewErrorResponse(c, err)
		c.JSON(code, response)
		return
	}
	user, err := h.UserService.LoginUser(c, user.Username, user.Password)
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
	logrus.Infof("User %s logged in", user.Username)
}

func (h *UsersHandler) Logout(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	_ = session.Save()
}
