package handlers

import (
	"encoding/json"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/ing-bank/ginerr/v3"
	"github.com/sirupsen/logrus"
	"gowasp/internal/models"
	"gowasp/internal/models/errors"
	"gowasp/internal/services"
	"net/http"
)

type UserHandler struct {
	UserService services.UserService
}

func (h *UserHandler) SignupPage(c *gin.Context) {
	c.HTML(http.StatusOK, "users/signup.tpl", gin.H{})
}

func (h *UserHandler) LoginPage(c *gin.Context) {
	c.HTML(http.StatusOK, "users/login.tpl", gin.H{})
}

func (h *UserHandler) WelcomePage(c *gin.Context) {
	session := sessions.Default(c)
	var user models.User
	_ = json.Unmarshal(session.Get("user").([]byte), &user)

	c.HTML(http.StatusOK, "users/welcome.tpl", gin.H{"user": user})
}

func (h *UserHandler) Signup(c *gin.Context) {
	user := models.User{}
	if err := c.Bind(&user); err != nil {
		code, response := ginerr.NewErrorResponse(c, err)
		c.HTML(code, response.(errors.ErrorResponse).HtmlTemplate, response.(errors.ErrorResponse).Data)
		return
	}
	if err := h.UserService.CreateUser(c, &user); err != nil {
		code, response := ginerr.NewErrorResponse(c, err)
		c.HTML(code, response.(errors.ErrorResponse).HtmlTemplate, response.(errors.ErrorResponse).Data)
		return
	}
	session := sessions.Default(c)
	userBytes, _ := json.Marshal(user)
	session.Set("user", userBytes)
	err := session.Save()
	if err != nil {
		code, response := ginerr.NewErrorResponse(c, err)
		c.HTML(code, response.(errors.ErrorResponse).HtmlTemplate, response.(errors.ErrorResponse).Data)
		return
	}
	c.Redirect(http.StatusFound, "/users/welcome")
}

func (h *UserHandler) Login(c *gin.Context) {
	user := models.User{}
	if err := c.Bind(&user); err != nil {
		code, response := ginerr.NewErrorResponse(c, err)
		c.HTML(code, response.(errors.ErrorResponse).HtmlTemplate, response.(errors.ErrorResponse).Data)
		return
	}
	user, err := h.UserService.LoginUser(c, user.Username, user.Password)
	if err != nil {
		code, response := ginerr.NewErrorResponse(c, err)
		c.HTML(code, response.(errors.ErrorResponse).HtmlTemplate, response.(errors.ErrorResponse).Data)
		return
	}
	session := sessions.Default(c)
	userBytes, _ := json.Marshal(user)
	session.Set("user", userBytes)
	err = session.Save()
	if err != nil {
		code, response := ginerr.NewErrorResponse(c, err)
		c.HTML(code, response.(errors.ErrorResponse).HtmlTemplate, response.(errors.ErrorResponse).Data)
		return
	}
	logrus.Infof("User %s logged in", user.Username)
	c.Redirect(http.StatusFound, "/users/welcome")
}

func (h *UserHandler) Logout(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	_ = session.Save()
}
