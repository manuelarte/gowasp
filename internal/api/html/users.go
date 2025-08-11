package html

import (
	"encoding/json"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/manuelarte/pagorminator"

	"github.com/manuelarte/gowasp/internal/api/dtos"
	"github.com/manuelarte/gowasp/internal/config"
	"github.com/manuelarte/gowasp/internal/posts"
)

type Users struct {
	PostService posts.Service
}

func NewUsers(postService posts.Service) *Users {
	return &Users{postService}
}

func (h *Users) LoginPage(c *gin.Context) {
	c.HTML(http.StatusOK, "users/login.tpl", gin.H{})
}

func (h *Users) Logout(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	_ = session.Save()
}

func (h *Users) SignupPage(c *gin.Context) {
	c.HTML(http.StatusOK, "users/signup.tpl", gin.H{})
}

func (h *Users) WelcomePage(c *gin.Context) {
	session := sessions.Default(c)
	var user dtos.UserSession
	sessionUserByte, ok := session.Get("user").([]byte)
	if !ok {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})

		return
	}
	err := json.Unmarshal(sessionUserByte, &user)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})

		return
	}

	defaultPageSize := 5
	postPageRequest, _ := pagorminator.PageRequest(0, defaultPageSize)
	latestPosts, err := h.PostService.GetAll(c, postPageRequest)
	if err != nil {
		_ = c.AbortWithError(http.StatusInternalServerError, err)

		return
	}
	c.HTML(http.StatusOK, "users/welcome.tpl", gin.H{"user": user, "latestPosts": latestPosts})
}

func RegisterUsersHandlers(r gin.IRouter, u *Users) {
	r.GET("/users/signup", u.SignupPage)
	r.GET("/users/login", u.LoginPage)
	r.DELETE("/users/logout", u.Logout)
	r.GET("/users/welcome", config.AuthMiddleware(), u.WelcomePage)
}
