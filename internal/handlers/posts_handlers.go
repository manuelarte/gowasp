package handlers

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/ing-bank/ginerr/v3"
	"github.com/manuelarte/pagorminator"

	"github.com/manuelarte/gowasp/internal/models"
	"github.com/manuelarte/gowasp/internal/posts"
	"github.com/manuelarte/gowasp/internal/posts/postcomments"
)

type PostsHandler struct {
	PostService        posts.Service
	PostCommentService postcomments.Service
}

func (h *PostsHandler) ViewPostPage(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		code, response := ginerr.NewErrorResponse(c, err)
		c.JSON(code, response)

		return
	}
	post, err := h.PostService.GetByID(c, id)
	if err != nil {
		code, response := ginerr.NewErrorResponse(c, err)
		c.JSON(code, response)

		return
	}
	defaultPageSize := 10
	commentsPageRequest, _ := pagorminator.PageRequest(0, defaultPageSize)
	pageResponsePostComments, errCmm := h.PostCommentService.GetAllForPostID(c, uint(id), commentsPageRequest)
	if errCmm != nil {
		code, response := ginerr.NewErrorResponse(c, errCmm)
		c.JSON(code, response)

		return
	}
	pageResponsePostUserComments := models.Transform(pageResponsePostComments, toPostUserComment)

	session := sessions.Default(c)
	var user models.User
	sessionUserByte, ok := session.Get("user").([]byte)
	if !ok {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})

		return
	}
	_ = json.Unmarshal(sessionUserByte, &user)

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
		errCls := file.Close()
		if errCls != nil {
			code, response := ginerr.NewErrorResponse(c, errCls)
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
	if err != nil && !errors.Is(err, io.EOF) {
		code, response := ginerr.NewErrorResponse(c, err)
		c.JSON(code, response)

		return
	}

	c.Data(http.StatusOK, "text/plain", bs)
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
	c.JSON(http.StatusOK, pagePostsResponse)
}

type PostUserComment struct {
	ID        uint        `json:"id"`
	CreatedAt time.Time   `json:"createdAt"`
	UpdatedAt time.Time   `json:"updatedAt,omitempty"`
	PostedAt  time.Time   `gorm:"now" json:"postedAt"`
	PostID    uint        `json:"postId"`
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
