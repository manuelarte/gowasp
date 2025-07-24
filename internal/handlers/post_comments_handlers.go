package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/ing-bank/ginerr/v3"
	"github.com/manuelarte/pagorminator"

	"github.com/manuelarte/gowasp/internal/models"
	"github.com/manuelarte/gowasp/internal/posts/postcomments"
)

type PostCommentsHandler struct {
	PostCommentService postcomments.Service
}

func (h *PostCommentsHandler) GetPostComments(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		code, response := ginerr.NewErrorResponse(c, err)
		c.JSON(code, response)

		return
	}
	pageRequest, err := getPageRequest(c)
	if err != nil {
		code, response := ginerr.NewErrorResponse(c, err)
		c.JSON(code, response)

		return
	}

	pageResponse, err := h.PostCommentService.GetAllForPostID(c, id, pageRequest)
	if err != nil {
		code, response := ginerr.NewErrorResponse(c, err)
		c.JSON(code, response)

		return
	}
	hourTime := time.Hour
	c.SetCookie("csrf", uuid.New().String(), int(hourTime),
		fmt.Sprintf("/posts/%d/comments", id), "localhost", false, true)
	c.JSON(http.StatusOK, pageResponse)
}

func (h *PostCommentsHandler) CreatePostComment(c *gin.Context) {
	newPostComment := NewPostComment{}
	if err := c.BindJSON(&newPostComment); err != nil {
		code, response := ginerr.NewErrorResponse(c, err)
		c.JSON(code, response)

		return
	}
	postComment := newPostComment.toPostComment(time.Now())
	err := h.PostCommentService.Create(c, &postComment)
	if err != nil {
		code, response := ginerr.NewErrorResponse(c, err)
		c.JSON(code, response)

		return
	}
	c.JSON(http.StatusOK, postComment)
}

func getPageRequest(c *gin.Context) (*pagorminator.Pagination, error) {
	pageRequestI, ok := c.Get("pageRequest")
	if !ok {
		return nil, errors.New("missing 'pageRequest' field")
	}
	pageRequest, ok := pageRequestI.(*pagorminator.Pagination)
	if !ok {
		return nil, errors.New("invalid 'pageRequest' field")
	}

	return pageRequest, nil
}

type NewPostComment struct {
	PostID  uint   `binding:"required" json:"postId" `
	UserID  uint   `binding:"required" json:"userId"`
	Comment string `binding:"required,min=1,max=1000" json:"comment"`
}

func (dto *NewPostComment) toPostComment(postedAt time.Time) models.PostComment {
	return models.PostComment{
		PostedAt: postedAt,
		PostID:   dto.PostID,
		UserID:   dto.UserID,
		Comment:  dto.Comment,
	}
}
