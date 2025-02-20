package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/ing-bank/ginerr/v3"
	"github.com/manuelarte/pagorminator"
	"gowasp/internal/models"
	"gowasp/internal/services"
	"strconv"
	"time"
)

type PostCommentsHandler struct {
	PostCommentService services.PostCommentService
}

func (h *PostCommentsHandler) GetPostComments(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
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
	pageResponse, err := h.PostCommentService.GetAllForPostID(c, uint(id), pageRequest)
	if err != nil {
		code, response := ginerr.NewErrorResponse(c, err)
		c.JSON(code, response)
		return
	}
	c.SetCookie("csrf", uuid.New().String(), 3600*24, fmt.Sprintf("/posts/%d/comments", id), "localhost", false, true)
	c.JSON(200, pageResponse)
}

func (h *PostCommentsHandler) CreatePostComment(c *gin.Context) {
	//id, err := strconv.Atoi(c.Param("id"))
	//if err != nil {
	//	code, response := ginerr.NewErrorResponse(c, err)
	//	c.JSON(code, response)
	//	return
	//}
	newPostComment := NewPostComment{}
	newPostComment.PostedAt = time.Now()
	if err := c.BindJSON(&newPostComment); err != nil {
		code, response := ginerr.NewErrorResponse(c, err)
		c.JSON(code, response)
		return
	}
	// create
	// vulnerabilities, csrf, template injection, not validating that the user who created comment is the one authenticated
	postComment := newPostComment.toPostComment()
	err := h.PostCommentService.Create(c, &postComment)
	if err != nil {
		code, response := ginerr.NewErrorResponse(c, err)
		c.JSON(code, response)
		return
	}
	c.JSON(200, postComment)
}

type NewPostComment struct {
	PostedAt time.Time `json:"postedAt" binding:"required"`
	PostID   uint      `json:"postID" binding:"required"`
	UserID   uint      `json:"userId" binding:"required"`
	Comment  string    `json:"comment" binding:"required,max=1000"`
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
