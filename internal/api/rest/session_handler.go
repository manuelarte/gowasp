package rest

import (
	"encoding/json"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"

	"github.com/manuelarte/gowasp/internal/api/dtos"
	"github.com/manuelarte/gowasp/internal/users"
)

type SessionHandler struct {
	service users.Service
}

func NewSession(service users.Service) SessionHandler {
	return SessionHandler{service: service}
}

func (s SessionHandler) GetSession(c *gin.Context) {
	session := sessions.Default(c)
	userJSON, ok := session.Get("user").([]byte)
	if !ok {
		c.JSON(http.StatusUnauthorized, ErrorResponse{
			Code:    http.StatusNotFound,
			Message: "User not logged in",
		})

		return
	}

	var user dtos.UserSession
	if err := json.Unmarshal(userJSON, &user); err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{
			Code:    http.StatusNotFound,
			Details: err.Error(),
			Message: "Error parsing the user session",
		})
	}

	c.JSON(http.StatusOK, user)
}
