package rest

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/manuelarte/ptrutils"
	"github.com/sirupsen/logrus"

	"github.com/manuelarte/gowasp/internal/api/dtos"
	"github.com/manuelarte/gowasp/internal/models"
	"github.com/manuelarte/gowasp/internal/users"
)

type UsersHandler struct {
	service users.Service
}

func NewUsers(service users.Service) *UsersHandler {
	return &UsersHandler{service: service}
}

func (h *UsersHandler) UserLogin(c *gin.Context) {
	userLogin := UserCredential{}
	if err := c.BindJSON(&userLogin); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Code:    http.StatusBadRequest,
			Details: err,
			Message: "Failed to marshal user data",
		})

		return
	}
	userDao := userToDAO(userLogin)
	user, err := h.service.Login(c, userDao.Username, userDao.Password)
	if err != nil {
		logrus.Infof("Login attempt failed for User %q", userDao.Username)
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Code:    http.StatusBadRequest,
			Details: err,
			Message: "Login attempt failed",
		})

		return
	}
	session := sessions.Default(c)
	userSession := userToUserSession(user)
	userBytes, _ := json.Marshal(userSession)
	session.Set("user", userBytes)
	err = session.Save()
	if err != nil {
		logrus.Infof("Error saving session for User '%s'", user.Username)
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Code:    http.StatusBadRequest,
			Details: err,
			Message: "Error saving session",
		})

		return
	}
	c.JSON(http.StatusOK, userSession)
	logrus.Infof("User %q logged in", user.Username)
}

func (h *UsersHandler) UserLogout(c *gin.Context) {
	session := sessions.Default(c)
	sessionUserByte, ok := session.Get("user").([]byte)
	session.Clear()
	_ = session.Save()
	var user dtos.UserSession
	if !ok {
		return
	}
	_ = json.Unmarshal(sessionUserByte, &user)
	logrus.Infof("User %q logged out", user.Username)
}

func (h *UsersHandler) UserSignup(c *gin.Context) {
	userSignup := UserCredential{}
	if err := c.BindJSON(&userSignup); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Code:    http.StatusBadRequest,
			Details: err,
			Message: "Failed to marshal user data",
		})

		return
	}
	user := userToDAO(userSignup)
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
	userBytes, _ := json.Marshal(userToUserSession(user))
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
	c.JSON(http.StatusCreated, userToDTO(user))
}

func (h *UsersHandler) GetUserById(c *gin.Context, userID uint) {
	user, err := h.service.GetByID(c, userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Code:    http.StatusBadRequest,
			Details: err,
			Message: "Error getting the user",
		})

		return
	}
	c.JSON(http.StatusOK, userToDTO(user))
}

func userToDAO(u UserCredential) models.User {
	return models.User{
		Username: u.Username,
		Password: u.Password,
		IsAdmin:  ptrutils.DerefOr(u.IsAdmin, false),
	}
}

func userToDTO(u models.User) User {
	return User{
		CreatedAt: u.CreatedAt,
		//#nosec G115
		Self:      Paths{}.GetUserByIdEndpoint.Path(strconv.Itoa(int(u.ID))),
		ID:        u.ID,
		IsAdmin:   u.IsAdmin,
		Password:  u.Password,
		UpdatedAt: u.UpdatedAt,
		Username:  u.Username,
	}
}

func userToUserSession(u models.User) dtos.UserSession {
	return dtos.UserSession{
		ID:       u.ID,
		Username: u.Username,
		Password: u.Password,
		IsAdmin:  u.IsAdmin,
	}
}
