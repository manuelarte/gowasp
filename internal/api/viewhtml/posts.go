package viewhtml

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"

	"github.com/manuelarte/gowasp/internal/api/rest"
	"github.com/manuelarte/gowasp/internal/config"
	"github.com/manuelarte/gowasp/internal/posts"
	"github.com/manuelarte/gowasp/internal/posts/postcomments"
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

func (h *Posts) GetStaticPostFileByName(c *gin.Context) {
	name := c.Query("name")
	// CWE-918: Server-Side Request Forgery (SSRF) https://cwe.mitre.org/data/definitions/918.html
	file, err := os.Open(fmt.Sprintf("./resources/posts/%s", name))
	if err != nil {
		c.JSON(http.StatusBadRequest, rest.ErrorResponse{
			Code:    http.StatusBadRequest,
			Details: nil,
			Message: err.Error(),
		})

		return
	}
	defer func(file *os.File) {
		errCls := file.Close()
		if errCls != nil {
			c.JSON(http.StatusInternalServerError, rest.ErrorResponse{
				Code:    http.StatusInternalServerError,
				Details: nil,
				Message: errCls.Error(),
			})

			return
		}
	}(file)
	// Get the file size
	stat, err := file.Stat()
	if err != nil {
		c.JSON(http.StatusInternalServerError, rest.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Details: nil,
			Message: err.Error(),
		})

		return
	}

	// Read the file into a byte slice
	bs := make([]byte, stat.Size())
	_, err = bufio.NewReader(file).Read(bs)
	if err != nil && !errors.Is(err, io.EOF) {
		c.JSON(http.StatusInternalServerError, rest.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Details: nil,
			Message: err.Error(),
		})

		return
	}

	c.Data(http.StatusOK, "text/plain", bs)
}

func RegisterPostsHandlers(r gin.IRouter, p *Posts) {
	r.GET("/static/posts", config.AuthMiddleware(), p.GetStaticPostFileByName)
}
