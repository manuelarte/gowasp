package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ing-bank/ginerr/v3"
	"github.com/manuelarte/pagorminator"

	"github.com/manuelarte/gowasp/internal/posts"
	"github.com/manuelarte/gowasp/internal/posts/postcomments"
)

type PostsHandler struct {
	PostService        posts.Service
	PostCommentService postcomments.Service
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
