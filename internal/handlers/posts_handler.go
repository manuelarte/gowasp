package handlers

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/gin-contrib/sessions"
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

type PostsHandler struct {
	PostService        services.PostService
	PostCommentService services.PostCommentService
}

func (h *PostsHandler) ViewPostPage(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		code, response := ginerr.NewErrorResponse(c, err)
		c.JSON(code, response)
		return
	}
	post, err := h.PostService.GetById(c, id)
	if err != nil {
		code, response := ginerr.NewErrorResponse(c, err)
		c.JSON(code, response)
		return
	}
	commentsPageRequest, _ := pagorminator.PageRequest(0, 10)
	pageResponsePostComments, err := h.PostCommentService.GetAllForPostID(c, uint(id), commentsPageRequest)
	pageResponsePostUserComments := models.Transform(pageResponsePostComments, toPostUserComment)

	session := sessions.Default(c)
	var user models.User
	_ = json.Unmarshal(session.Get("user").([]byte), &user)

	c.HTML(http.StatusOK, "posts/post.tpl", gin.H{"user": user, "post": post, "comments": pageResponsePostUserComments})
}

func (h *PostsHandler) GetStaticPostFileByName(c *gin.Context) {
	name := c.Query("name")
	// CWE-918: Server-Side Request Forgery (SSRF) https://cwe.mitre.org/data/definitions/918.html
	file, err := os.Open(fmt.Sprintf("./resources/posts/%s", name))
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

func (h *PostsHandler) GetAll(c *gin.Context) {
	pageString := c.DefaultQuery("page", "0")
	page, _ := strconv.Atoi(pageString)
	sizeString := c.DefaultQuery("size", "10")
	size, _ := strconv.Atoi(sizeString)
	pageRequest, _ := pagorminator.PageRequest(page, size)
	pagePostsResponse, err := h.PostService.GetAll(c, pageRequest)
	if err != nil {
		code, response := ginerr.NewErrorResponse(c, err)
		c.JSON(code, response)
		return
	}
	c.JSON(200, pagePostsResponse)
}

type PostUserComment struct {
	ID        uint        `json:"id"`
	CreatedAt time.Time   `json:"createdAt"`
	UpdatedAt time.Time   `json:"updatedAt,omitempty"`
	PostedAt  time.Time   `json:"postedAt" gorm:"now"`
	PostID    uint        `json:"postID"`
	User      UserComment `json:"user"`
	Comment   string      `json:"comment"`
}

type UserComment struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
}

func toPostUserComment(postComment *models.PostComment) PostUserComment {
	return PostUserComment{
		ID:        postComment.ID,
		CreatedAt: postComment.CreatedAt,
		UpdatedAt: postComment.UpdatedAt,
		PostedAt:  postComment.PostedAt,
		PostID:    postComment.PostID,
		User: UserComment{
			ID:       postComment.User.ID,
			Username: postComment.User.Username,
		},
		Comment: postComment.Comment,
	}
}
