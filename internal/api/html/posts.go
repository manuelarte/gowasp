package html

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/ing-bank/ginerr/v3"
	"github.com/manuelarte/pagorminator"

	"github.com/manuelarte/gowasp/internal/api/dtos"
	"github.com/manuelarte/gowasp/internal/api/rest"
	"github.com/manuelarte/gowasp/internal/config"
	"github.com/manuelarte/gowasp/internal/models"
	"github.com/manuelarte/gowasp/internal/posts"
	"github.com/manuelarte/gowasp/internal/posts/postcomments"
	"github.com/manuelarte/gowasp/internal/sliceutils"
)

type Posts struct {
	service        posts.Service
	commentService postcomments.Service
}

func NewPosts(service posts.Service, commentService postcomments.Service) *Posts {
	return &Posts{
		service:        service,
		commentService: commentService,
	}
}

func (h *Posts) ViewPostPage(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, rest.ErrorResponse{
			Code:    http.StatusBadRequest,
			Details: err,
			Message: "Error parsing the post id",
		})

		return
	}
	post, err := h.service.GetByID(c, uint(id))
	if err != nil {
		c.JSON(http.StatusBadRequest, rest.ErrorResponse{
			Code:    http.StatusBadRequest,
			Details: err,
			Message: "Error getting the post",
		})

		return
	}
	defaultPageSize := 10
	commentsPageRequest, _ := pagorminator.PageRequest(0, defaultPageSize)
	pageResponsePostComments, errCmm := h.commentService.GetAllForPostID(c, uint(id), commentsPageRequest)
	if errCmm != nil {
		c.JSON(http.StatusBadRequest, rest.ErrorResponse{
			Code:    http.StatusBadRequest,
			Details: err,
			Message: "Error getting all post comments",
		})

		return
	}
	pageResponsePostUserComments := sliceutils.Transform(pageResponsePostComments, toPostUserCommentExtended)

	session := sessions.Default(c)
	var user dtos.UserSession
	sessionUserByte, ok := session.Get("user").([]byte)
	if !ok {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})

		return
	}
	_ = json.Unmarshal(sessionUserByte, &user)

	c.HTML(http.StatusOK, "posts/post.tpl", gin.H{"user": user, "post": post, "comments": pageResponsePostUserComments})
}

func (h *Posts) GetStaticPostFileByName(c *gin.Context) {
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

func RegisterPostsHandlers(r gin.IRouter, p *Posts) {
	r.GET("/static/posts", config.AuthMiddleware(), p.GetStaticPostFileByName)
	r.GET("/posts/:id/view", config.AuthMiddleware(), p.ViewPostPage)
}

func toPostUserCommentExtended(postComment *models.PostComment) PostCommentExtended {
	return PostCommentExtended{
		PostComment: rest.PostComment{
			ID:        postComment.ID,
			CreatedAt: postComment.CreatedAt,
			UpdatedAt: postComment.UpdatedAt,
			PostedAt:  postComment.PostedAt,
			PostID:    postComment.PostID,
			Comment:   postComment.Comment,
		},
		User: struct {
			ID       uint   `json:"id"`
			Username string `json:"username"`
		}{
			ID:       postComment.User.ID,
			Username: postComment.User.Username,
		},
	}
}

type PostCommentExtended struct {
	rest.PostComment

	User struct {
		ID       uint   `json:"id"`
		Username string `json:"username"`
	}
}
