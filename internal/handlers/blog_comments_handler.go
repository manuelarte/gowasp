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

type BlogCommentsHandler struct {
	BlogCommentService services.BlogCommentService
}

func (h *BlogCommentsHandler) GetBlogComments(c *gin.Context) {
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
	pageResponse, err := h.BlogCommentService.GetAllForBlog(c, uint(id), pageRequest)
	if err != nil {
		code, response := ginerr.NewErrorResponse(c, err)
		c.JSON(code, response)
		return
	}
	c.SetCookie("csrf", uuid.New().String(), 3600*24, fmt.Sprintf("/blogs/%d/comments", id), "localhost", false, true)
	c.JSON(200, pageResponse)
}

func (h *BlogCommentsHandler) CreateBlogComment(c *gin.Context) {
	//id, err := strconv.Atoi(c.Param("id"))
	//if err != nil {
	//	code, response := ginerr.NewErrorResponse(c, err)
	//	c.JSON(code, response)
	//	return
	//}
	newBlogComment := NewBlogComment{}
	newBlogComment.PostedAt = time.Now()
	if err := c.BindJSON(&newBlogComment); err != nil {
		code, response := ginerr.NewErrorResponse(c, err)
		c.JSON(code, response)
		return
	}
	// create
	// vulnerabilities, csrf, template injection, not validating that the user who created comment is the one authenticated
	blogComment := newBlogComment.toBlogComment()
	err := h.BlogCommentService.Create(c, &blogComment)
	if err != nil {
		code, response := ginerr.NewErrorResponse(c, err)
		c.JSON(code, response)
		return
	}
	c.JSON(200, blogComment)
}

type NewBlogComment struct {
	PostedAt time.Time `json:"postedAt" binding:"required"`
	BlogID   uint      `json:"blogId" binding:"required"`
	UserID   uint      `json:"userId" binding:"required"`
	Comment  string    `json:"comment" binding:"required,max=1000"`
}

func (b *NewBlogComment) toBlogComment() models.BlogComment {
	return models.BlogComment{
		ID:        0,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		PostedAt:  b.PostedAt,
		BlogID:    b.BlogID,
		UserID:    b.UserID,
		Comment:   b.Comment,
	}
}
