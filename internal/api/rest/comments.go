package rest

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/manuelarte/pagorminator"
	"github.com/manuelarte/ptrutils"

	"github.com/manuelarte/gowasp/internal/posts/postcomments"
)

const defaultPageRequestSize = 10

type Comments struct {
	service postcomments.Service
}

func NewComments(service postcomments.Service) *Comments {
	return &Comments{service: service}
}

func (h *Comments) GetPostComments(c *gin.Context, postID uint64, params GetPostCommentsParams) {
	pageRequest, err := pagorminator.PageRequest(
		ptrutils.DerefOr(params.Page, 0),
		ptrutils.DerefOr(params.Size, defaultPageRequestSize),
	)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Code:    http.StatusBadRequest,
			Details: err,
			Message: "Error creating the page request",
		})

		return
	}

	pageResponse, err := h.service.GetAllForPostID(c, postID, pageRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Code:    http.StatusBadRequest,
			Details: err,
			Message: "Error getting the post comments",
		})

		return
	}
	hourTime := time.Hour
	c.SetCookie("csrf", uuid.New().String(), int(hourTime),
		fmt.Sprintf("/posts/%d/comments", postID), "localhost", false, true)
	c.JSON(http.StatusOK, pageResponse)
}
