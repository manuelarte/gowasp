package handlers

import (
	"bufio"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ing-bank/ginerr/v3"
	"github.com/manuelarte/pagorminator"
	"gowasp/internal/models"
	"gowasp/internal/services"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"
)

type BlogsHandler struct {
	BlogService        services.BlogService
	BlogCommentService services.BlogCommentService
}

func (h *BlogsHandler) ViewBlogPage(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		code, response := ginerr.NewErrorResponse(c, err)
		c.JSON(code, response)
		return
	}
	blog, err := h.BlogService.GetById(c, id)
	if err != nil {
		code, response := ginerr.NewErrorResponse(c, err)
		c.JSON(code, response)
		return
	}
	commentsPageRequest, _ := pagorminator.PageRequest(0, 10)
	pageResponseBlogComments, err := h.BlogCommentService.GetAllForBlog(c, uint(id), commentsPageRequest)
	pageResponseBlogUserComments := models.Transform(pageResponseBlogComments, toBlogUserComment)

	c.HTML(http.StatusOK, "blogs/blog.tpl", gin.H{"blog": blog, "comments": pageResponseBlogUserComments})
}

func (h *BlogsHandler) GetStaticBlogFileByName(c *gin.Context) {
	name := c.Query("name")
	// CWE-918: Server-Side Request Forgery (SSRF) https://cwe.mitre.org/data/definitions/918.html
	file, err := os.Open(fmt.Sprintf("./resources/blogs/%s", name))
	if err != nil {
		code, response := ginerr.NewErrorResponse(c, err)
		c.JSON(code, response)
		return
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			code, response := ginerr.NewErrorResponse(c, err)
			c.JSON(code, response)
			return
		}
	}(file)
	// Get the file size
	stat, err := file.Stat()
	if err != nil {
		code, response := ginerr.NewErrorResponse(c, err)
		c.JSON(code, response)
		return
	}

	// Read the file into a byte slice
	bs := make([]byte, stat.Size())
	_, err = bufio.NewReader(file).Read(bs)
	if err != nil && err != io.EOF {
		code, response := ginerr.NewErrorResponse(c, err)
		c.JSON(code, response)
		return
	}

	c.Data(200, "text/plain", bs)
}

func (h *BlogsHandler) GetAll(c *gin.Context) {
	pageString := c.DefaultQuery("page", "0")
	page, _ := strconv.Atoi(pageString)
	sizeString := c.DefaultQuery("size", "10")
	size, _ := strconv.Atoi(sizeString)
	pageRequest, _ := pagorminator.PageRequest(page, size)
	pageBlogsResponse, err := h.BlogService.GetAll(c, pageRequest)
	if err != nil {
		code, response := ginerr.NewErrorResponse(c, err)
		c.JSON(code, response)
		return
	}
	c.JSON(200, pageBlogsResponse)
}

type BlogUserComment struct {
	ID        uint        `json:"id"`
	CreatedAt time.Time   `json:"createdAt"`
	UpdatedAt time.Time   `json:"updatedAt,omitempty"`
	PostedAt  time.Time   `json:"postedAt" gorm:"now"`
	BlogID    uint        `json:"blogId"`
	User      UserComment `json:"user"`
	Comment   string      `json:"comment"`
}

type UserComment struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
}

func toBlogUserComment(blogComment *models.BlogComment) BlogUserComment {
	return BlogUserComment{
		ID:        blogComment.ID,
		CreatedAt: blogComment.CreatedAt,
		UpdatedAt: blogComment.UpdatedAt,
		PostedAt:  blogComment.PostedAt,
		BlogID:    blogComment.BlogID,
		User: UserComment{
			ID:       blogComment.User.ID,
			Username: blogComment.User.Username,
		},
		Comment: blogComment.Comment,
	}
}
