package config

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/manuelarte/pagorminator"

	"github.com/manuelarte/gowasp/internal/models"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		sessionUser := session.Get("user")
		if sessionUser == nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})

			return
		}

		var user models.User
		byteSession, ok := sessionUser.([]byte)
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})

			return
		}

		err := json.Unmarshal(byteSession, &user)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})

			return
		}

		c.Next()
	}
}

func PaginationMiddleware(defaultSize int) gin.HandlerFunc {
	return func(c *gin.Context) {
		pageString := c.DefaultQuery("page", "0")
		page, err := strconv.Atoi(pageString)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "error converting page to int"})

			return
		}
		sizeString := c.DefaultQuery("size", strconv.Itoa(defaultSize))
		size, err := strconv.Atoi(sizeString)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "error converting size to int"})

			return
		}

		pageRequest, err := pagorminator.PageRequest(page, size)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "page request error"})

			return
		}
		c.Set("pageRequest", pageRequest)

		c.Next()
	}
}
