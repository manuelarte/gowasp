package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ing-bank/ginerr/v3"

	"github.com/manuelarte/gowasp/internal/models"
	"github.com/manuelarte/gowasp/internal/posts/postcomments"
)

type PostCommentsHandler struct {
	PostCommentService postcomments.Service
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
