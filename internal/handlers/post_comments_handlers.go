package handlers

import (
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
	pageString := c.DefaultQuery("page", "0")
	page, _ := strconv.Atoi(pageString)
	sizeString := c.DefaultQuery("size", "10")
	size, _ := strconv.Atoi(sizeString)
	pageRequest, _ := pagorminator.PageRequest(page, size)
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
	newPostComment.PostedAt = time.Now()
	if err := c.BindJSON(&newPostComment); err != nil {
		code, response := ginerr.NewErrorResponse(c, err)
		c.JSON(code, response)

		return
	}
	postComment := newPostComment.toPostComment()
	err := h.PostCommentService.Create(c, &postComment)
	if err != nil {
		code, response := ginerr.NewErrorResponse(c, err)
		c.JSON(code, response)

		return
	}
	c.JSON(http.StatusOK, postComment)
}

type NewPostComment struct {
	PostedAt time.Time `binding:"required" json:"postedAt"`
	PostID   uint      `binding:"required" json:"postID" `
	UserID   uint      `binding:"required" json:"userID"`
	Comment  string    `binding:"required,min=1,max=1000" json:"comment"`
}

func (b *NewPostComment) toPostComment() models.PostComment {
	return models.PostComment{
		ID:        0,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		PostedAt:  b.PostedAt,
		PostID:    b.PostID,
		UserID:    b.UserID,
		Comment:   b.Comment,
	}
}
