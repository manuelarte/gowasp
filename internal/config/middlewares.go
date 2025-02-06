package config

import (
	"encoding/json"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"gowasp/internal/models"
	"net/http"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		session := sessions.Default(c)
		sessionUser := session.Get("user")
		if sessionUser == nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{})
			return
		}

		var user models.User
		err := json.Unmarshal(sessionUser.([]byte), &user)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{})
			return
		}

		c.Next()
	}
}
